package eda

import "testing"

func TestEntry(t *testing.T) {
	// each entry should result in monthly amount of 15
	entries := []Entry{
		&Loan{
			Left:        1200,
			Interest:    5.0,
			Installment: 10,
		},
		&Expense{
			Amount: 180,
			Period: "y",
		},
		&Expense{
			Amount: 15,
			Period: "m",
		},
	}
	for _, e := range entries {
		if v := e.Monthly(); v != 15 {
			t.Error(v)
		}
	}
}
