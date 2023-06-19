package eda

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func Example() {
	budget := `
# home
6000 y electricity
1000 m food

# clothes and stuff
500 m clothes, linnen
`
	scanner := NewScanner(strings.NewReader(budget))
	for {
		e, err := scanner.Scan()
		if errors.Is(err, io.EOF) {
			break
		}
		fmt.Println(e)
	}
	// output:
	// 6000 y electricity
	// 1000 m food
	// 500 m clothes, linnen
}

func TestScanner(t *testing.T) {
	scanner := NewScanner(strings.NewReader("y food"))
	_, err := scanner.Scan()
	if err == nil {
		t.Error("expected error on missing amount")
	}
}
