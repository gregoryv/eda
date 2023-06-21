package eda

import (
	"fmt"
	"log"
	"strings"
)

func ExampleParse() {
	budget := `
# home
6000/year electricity
1000/m food

# clothes and stuff
500/m clothes linnen

# loan
100000  5.0 0 car

# expense
100/m internet
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
	// 500 [clothes linnen]
	// 416 [car]
	// 100 [internet]
}
