package eda

import "testing"

func TestLoan(t *testing.T) {
	l := Loan{
		Left:        1200,
		Interest:    5.0,
		Installment: 10,
	}
	if v := l.Monthly(); v != 15 {
		t.Error(v)
	}
}
