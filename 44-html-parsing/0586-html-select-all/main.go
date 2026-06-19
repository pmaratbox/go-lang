package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	DOC := `<html><body>
<h1>Hello</h1>
<span id="status">active</span>
<ul class="items">
<li class="item">apple</li>
<li class="item">banana</li>
<li class="item">cherry</li>
</ul>
<a href="https://example.com">site</a>
<div class="content"><p>first</p><p>second</p></div>
</body></html>`

	d, err := goquery.NewDocumentFromReader(strings.NewReader(DOC))
	if err != nil {
		panic(err)
	}
	// CSS class selector `.item` matches ALL list items; collect each one's
	// text so whitespace between elements never matters, then join with commas.
	var items []string
	d.Find(".item").Each(func(i int, s *goquery.Selection) {
		items = append(items, s.Text())
	})
	fmt.Println(strings.Join(items, ","))
}
