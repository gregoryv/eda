/*
Expenses data format

  # Comments start with a '#' and empty lines are ignored

  # Special comments
  #
  # expense   signals following lines are expenses (default)
  # loan      signals following lines are loan entries
  #
  # expense
  # amount/(y|m) tags
  1000/m electricity
  100/m mobile
  40/m github
  ...

  # loan
  # left interest installment tags
  686453     3.34        599  loan house
   68800     5.39        700  loan car
  ...
*/
package eda

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func Parse(r io.Reader) ([]Entry, error) {
	scanner := NewScanner(r)
	entries := make([]Entry, 0)
	for {
		e, err := scanner.Scan()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

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
		amount, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid amount on line %v: %s", s.lineno, line)
		}

		var e Expense
		e.amount = amount
		e.Period = parts[1]
		e.tags = parts[2:]
		return &e, nil
	}
}
