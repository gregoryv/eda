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
		cli = cmdline.NewBasicParser()
	)
	cli.Parse()

	log.SetFlags(0)
	entries, err := eda.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var monthly int
	for _, e := range entries {
		monthly += e.Monthly()
	}
	fmt.Printf("%v/m\n", monthly)
}
