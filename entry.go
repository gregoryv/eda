package eda

// Entry represents data lines
type Entry interface {
	// Monthly returns the monthly amount.
	Monthly() int
	Tags() []string
}

// ----------------------------------------

// Expense represents expense lines
type Expense struct {
	Amount int
	Period string
	tags   []string
}

func (e *Expense) Monthly() int {
	switch e.Period {
	case "y":
		return e.Amount / 12
	default:
		return e.Amount
	}
}

func (e *Expense) Tags() []string { return e.tags }

// ----------------------------------------

// Loan represents loan lines
type Loan struct {
	Left        int
	Interest    float64
	Installment int
	tags        []string
}

func (l *Loan) Monthly() int {
	return l.Installment + int(((l.Interest / 100.0) * float64(l.Left) / 12.0))
}

func (l *Loan) Tags() []string { return l.tags }
