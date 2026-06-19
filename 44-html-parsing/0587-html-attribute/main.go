package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const DOC = `<html><body>
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

func main() {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(DOC))
	if err != nil {
		panic(err)
	}
	// Select the <a> element and read its href attribute.
	href, _ := d.Find("a").Attr("href")
	fmt.Println(href)
}
