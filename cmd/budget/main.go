// Command budget parses EDA files and prints out a summary
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

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

		groupByTag = cli.Option("-g, --group-by-tag", "group entries by tags").Bool()
		files      = cli.NamedArg("FILES...").Strings()
	)
	cli.Parse()
	log.SetFlags(0)

	// select input
	var input io.Reader = os.Stdin

	entries := make([]eda.Entry, 0)
	if len(files) > 0 {
		for _, file := range files {
			var err error
			if file != "" {
				input, err = os.Open(file)
				if err != nil {
					log.Fatal(err)
				}
			}

			// parse entries
			got, err := eda.Parse(input)
			if err != nil {
				log.Fatal(err)
			}
			entries = append(entries, got...)
		}
	}

	// group by tags
	tagged := map[string]*Tag{}
	for _, e := range entries {
		if e == nil {
			continue
		}
		if groupByTag {
			for _, t := range e.Tags() {
				if _, found := tagged[t]; !found {
					tagged[t] = &Tag{}
				}
				tagged[t].Amount += e.Monthly()
				tagged[t].Count++
			}
		} else {
			t := strings.Join(e.Tags(), " ")
			tagged[t] = &Tag{}
			tagged[t].Amount = e.Monthly()
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

	// group tags
	keys := sortedKeys(tagged)
	group := make([]string, 0)
	for _, k := range keys {
		if tagged[k].Count == 1 {
			continue
		}
		group = append(group, k)
	}
	if len(group) > 0 {
		fmt.Println("---------- --------------------")
		for _, k := range group {
			t := tagged[k]
			write(t.Amount, k)
		}
	}
	fmt.Println("---------- --------------------")
	write(monthly, "sum")
	if shared > 1 {
		fmt.Printf("%10v people\n", shared)
		fmt.Println("---------- --------------------")
		write(monthly/int(shared), "each")
	}
}

type Tag struct {
	Count  int // number of tags
	Amount int
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
