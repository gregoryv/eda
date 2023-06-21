package eda

import (
	"fmt"
	"strings"
)

// Entry represents data lines
type Entry interface {
	// Monthly returns the monthly amount.
	Monthly() int
	Tags() []string
}

// ----------------------------------------

// Expense represents expense lines
type Expense struct {
	amount int
	Period string
	tags   []string
}

func (e *Expense) String() string {
	return fmt.Sprintf("%v/%s %s", e.amount, string(e.Period), strings.Join(e.tags, " "))
}

func (e *Expense) Monthly() int {
	switch e.Period {
	case "y":
		return e.amount / 12
	default:
		return e.amount
	}
}

func (e *Expense) Tags() []string { return e.tags }

// ----------------------------------------

// Loan represents loan lines
type Loan struct {
	Left        int
	Interest    float64
	installment int
	tags        []string
}

func (l *Loan) Monthly() int {
	return l.installment + int(((l.Interest / 100.0) * float64(l.Left) / 12.0))
}

func (l *Loan) Tags() []string { return l.tags }
