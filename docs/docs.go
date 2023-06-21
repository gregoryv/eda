package main

import (
	"log"
	"os/exec"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
	"github.com/gregoryv/web/theme"
	"github.com/gregoryv/web/toc"
)

func main() {
	nav := Nav()
	body := Body(
		H1("Expense Data - file format"),

		P(`Creating a simple budget for your family using plain text
        for easy version control. This project defines the file
        format, expense data (EDA) and a tool to process it with.`),

		nav,

		Code(`Source code: `, A(
			Href("https://github.com/gregoryv/eda"),
			"github.com/gregoryv/eda",
		)),

		H2("Quick start"),

		Pre(
			`$ go install github.com/gregoryv/eda/cmd/budget@latest
$ budget example.eda
`,
			shell("budget", "../example.eda"),
		),

		H2("File format"),
		Pre(
			files.MustLoadLines("../eda.go", 4, 22),
		),

		H3("Example"),
		Pre(
			files.MustLoad("../example.eda"),
		),
	)

	toc.MakeTOC(nav, body, "h2", "h3")
	page := NewPage(
		Html(
			Head(
				Style(theme.GoldenSpace()),
				Style(theme.GoishColors()),
				Style(myTheme()),
			),
			body,
		),
	)
	page.SaveAs("index.html")
}

func myTheme() *CSS {
	css := NewCSS()
	css.Style("li.h3", "margin-left: 1em")
	return css
}

func shell(cmd string, args ...string) *Element {
	c, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return Wrap(string(c))
}
