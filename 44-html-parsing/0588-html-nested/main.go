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

	// Descendant selector ".content p": every <p> inside .content.
	var ps []string
	d.Find(".content p").Each(func(i int, s *goquery.Selection) {
		ps = append(ps, s.Text())
	})
	fmt.Println(strings.Join(ps, ","))
}
