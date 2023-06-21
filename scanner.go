package eda

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		s: bufio.NewScanner(r),
	}
}

type Scanner struct {
	lineno int
	s      *bufio.Scanner
	loan   bool
}

// multiple spaces regexp
var re = regexp.MustCompile(`\s+`)

func (s *Scanner) Scan() (Entry, error) {
next:
	s.lineno++
	more := s.s.Scan()
	if !more {
		return nil, io.EOF
	}
	line := strings.TrimSpace(s.s.Text())
	// remove multiple spaces
	line = re.ReplaceAllString(line, " ")
	// toggle loan|expense parsing
	if strings.HasPrefix(line, "# loan") {
		s.loan = true
	}
	if strings.HasPrefix(line, "# expense") {
		s.loan = false
	}
	if strings.HasPrefix(line, "#") || len(line) == 0 {
		goto next
	}

	// start parsing the line
	line = strings.Replace(line, "/", " ", 1) // in expense entries
	parts := strings.Split(line, " ")

	if s.loan {
		var err error
		var l Loan
		l.Left, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid left on line %v: %s", s.lineno, line)
		}

		l.Interest, err = strconv.ParseFloat(parts[1], 32)
		if err != nil {
			return nil, fmt.Errorf("invalid interest on line %v: %s", s.lineno, line)
		}

		l.Installment, err = strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid installment on line %v: %s", s.lineno, line)
		}
		l.tags = parts[3:]
		return &l, nil
	} else {
		// parse amount
		Amount, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid amount on line %v: %s", s.lineno, line)
		}

		var e Expense
		e.Amount = Amount
		e.Period = parts[1]
		e.tags = parts[2:]
		return &e, nil
	}
}
