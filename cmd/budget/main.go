package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/eda"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		shared = cli.Option("-p, --people",
			"number of people sharing the expenses",
		).Int(2)
		file = cli.Option("-f, --filename").String("")
	)
	cli.Parse()
	log.SetFlags(0)

	// select input
	var input io.Reader = os.Stdin
	var err error
	if file != "" {
		input, err = os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	// parse entries
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
	var totalLoans int
	for _, t := range entries {
		monthly += t.Monthly()
		if l, ok := t.(*eda.Loan); ok {
			totalLoans += l.Left
		}
	}

	// write result
	p := message.NewPrinter(language.Swedish)
	write := func(v int, txt string) {
		p.Printf("%10d %s\n", v, txt)
	}
	write(totalLoans, "loans left")
	fmt.Println("---------- --------------------")
	for k, t := range tagged {
		if t.Count == 1 {
			continue
		}
		write(t.Amount, k)
	}
	fmt.Println("+ -------- --------------------")
	write(monthly, "sum")
	fmt.Printf("%10v people\n", shared)
	fmt.Println("/ -------- --------------------")
	write(monthly/int(shared), "")
}

type Tag struct {
	Count  int // number of tags
	Amount int
}
