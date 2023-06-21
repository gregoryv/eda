package eda

import (
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{"missing amount", "y food"},
		{"missing left", "# loan\n3.9 100 loan\n"},
		{"bad interest", "# loan\n1000 bad 100 loan"},
		{"bad installment", "# loan\n1000 2.0 bad loan"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(c.input))
			_, err := scanner.Scan()
			if err == nil {
				t.Error("expected error")
			}

		})
	}
}
