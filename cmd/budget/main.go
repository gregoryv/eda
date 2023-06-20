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
	write := func(v int, txt string) {
		fmt.Printf("%8s %s\n", formatAmount(v), txt)
	}
	for k, t := range tagged {
		if t.Count == 1 {
			continue
		}
		write(t.Amount, k)
	}
	write(other, "other")
	fmt.Println("+ ------ --------------------")
	write(monthly, "sum")
	fmt.Printf("%8v people\n", shared)
	fmt.Println("/ ------ --------------------")
	write(monthly/int(shared), "")
}

func formatAmount(v int) string {
	switch {
	case v < 1_000:
		return fmt.Sprintf("%v", v)

	case v < 1_000_000:
		return fmt.Sprintf("%v %03v", v/1000, v%1000)
	default:
		return fmt.Sprintf("%v", v)
	}
}

type Tag struct {
	Count  int // number of tags
	Amount int
}
