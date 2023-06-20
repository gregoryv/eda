package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/eda"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		shared = cli.Option("-p, --people").Int(2)
		file   = cli.Option("-f, --filename").String("")
	)
	cli.Parse()

	log.SetFlags(0)

	var input io.Reader = os.Stdin
	var err error
	if file != "" {
		input, err = os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	entries, err := eda.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	// group by tags
	tagged := map[string]*Tag{}
	for _, e := range entries {
		for _, t := range e.Tags() {
			if _, found := tagged[t]; !found {
				tagged[t] = &Tag{}
			}
			tagged[t].Amount += e.Monthly()
			tagged[t].Count++
		}
	}

	// summarize
	var monthly int
	for _, t := range entries {
		monthly += t.Monthly()
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
