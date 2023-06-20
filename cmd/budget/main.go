package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/eda"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		shared = cli.Option("-p, --people").Int(2)
	)
	cli.Parse()

	log.SetFlags(0)
	entries, err := eda.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// group by tags
	tagged := make(map[string]*Tag)
	for _, e := range entries {
		for _, t := range e.Tags {
			if _, found := tagged[t]; !found {
				tagged[t] = &Tag{}
			}
			tagged[t].Amount += e.Monthly()
			tagged[t].Count++
		}
	}

	// summarize
	var other int // with only one tag
	var monthly int
	for _, t := range tagged {
		monthly += t.Amount
		if t.Count == 1 {
			other += t.Amount
			continue
		}
	}

	// write result
	for k, t := range tagged {
		if t.Count == 1 {
			continue
		}
		fmt.Println(t.Amount, k)
	}
	fmt.Println(other, "other")
	fmt.Println(monthly, "sum")
	fmt.Println(monthly/shared, "per person")
}

type Tag struct {
	Count  int // number of tags
	Amount int
}
