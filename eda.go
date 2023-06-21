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
	"errors"
	"io"
)

// Parse returns entries until io.EOF is reached or an error
// occurs. When done nil error is returned.
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
