package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	DOC := `<table><tr><td>r1c1</td><td>r1c2</td></tr><tr><td>r2c1</td><td>r2c2</td></tr></table>`

	d, err := goquery.NewDocumentFromReader(strings.NewReader(DOC))
	if err != nil {
		panic(err)
	}
	// CSS tag selector `td` selects every table cell in row-major order.
	var cells []string
	d.Find("td").Each(func(i int, s *goquery.Selection) {
		cells = append(cells, s.Text())
	})
	fmt.Println(strings.Join(cells, ","))
}
