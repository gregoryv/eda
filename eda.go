/*

Expenses data format

# comments, empty lines are ignored
# special comments, [expense|loan] signals the subsequent line formats

# expense
int/[y|m] tag1 ... tagn

# loan
# date left original interest installment tags
20220228 686453 720000     3.34        599 lån Bolån_1 3994 15 72194 gemensamt 3år
20220228 668800 720000     3.39        700 lån Bolån_2 3994 15 72208 gemensamt 4år

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

func Parse(r io.Reader) ([]*Entry, error) {
	scanner := NewScanner(r)
	entries := make([]*Entry, 0)
	for {
		e, err := scanner.Scan()
		if errors.Is(err, io.EOF) {
			break
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

func (s *Scanner) Scan() (*Entry, error) {
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

		// todo
		var e Entry
		return &e, nil
	} else {
		// parse amount
		amount, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid amount on line %v: %s", s.lineno, line)
		}

		var e Entry
		e.amount = amount
		e.Period = parts[1]
		e.Tags = parts[2:]
		return &e, nil
	}
}

// ----------------------------------------

type Entry struct {
	amount int
	Period string
	Tags   []string
}

type Loan struct {
	Date        string
	Left        int
	Orig        int
	Interest    float64
	installment int
	Tags        []string
}

func (l *Loan) Monthly() int{
	// todo
	return 0
}


func (e *Entry) String() string {
	return fmt.Sprintf("%v/%s %s", e.amount, string(e.Period), strings.Join(e.Tags, " "))
}

func (e *Entry) Monthly() int {
	switch e.Period {
	case "y":
		return e.amount / 12
	default:
		return e.amount
	}
}
