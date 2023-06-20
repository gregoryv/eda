/*

Expenses data format

# comments, empty lines are ignored
int/[y|m] tag1 ... tagn

*/
package eda

import (
	"bufio"
	"fmt"
	"io"
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
}

func (s *Scanner) Scan() (*Entry, error) {
next:
	s.lineno++
	more := s.s.Scan()
	if !more {
		return nil, io.EOF
	}
	line := strings.TrimSpace(s.s.Text())
	if strings.HasPrefix(line, "#") || len(line) == 0 {
		goto next
	}
	line = strings.Replace(line, "/", " ", 1) // optional '/'

	parts := strings.Split(line, " ")
	// parse amount
	amount, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid amount on line %v: %s", s.lineno, line)
	}

	var e Entry
	e.Amount = amount
	e.Period = parts[1]
	e.Tags = parts[2:]
	return &e, nil
}

// ----------------------------------------

type Entry struct {
	Amount int
	Period string
	Tags   []string
}

func (e *Entry) String() string {
	return fmt.Sprintf("%v/%s %s", e.Amount, string(e.Period), strings.Join(e.Tags, " "))
}
