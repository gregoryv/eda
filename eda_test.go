package eda

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func ExampleParse() {
	budget := `
# home
6000/year electricity
1000/m food

# clothes and stuff
500/m clothes, linnen

# loan
100000  5.0 0 car
`
	entries, err := Parse(strings.NewReader(budget))
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Monthly(), e.Tags())
	}
	// output:
	// 6000 [electricity]
	// 1000 [food]
	// 500 [clothes, linnen]
	// 416 [car]
}

func TestScanner(t *testing.T) {
	scanner := NewScanner(strings.NewReader("y food"))
	_, err := scanner.Scan()
	if err == nil {
		t.Error("expected error on missing amount")
	}
}

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
