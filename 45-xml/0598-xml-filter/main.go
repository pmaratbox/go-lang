package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Book struct {
	ID    string `xml:"id,attr"`
	Lang  string `xml:"lang,attr"`
	Title string `xml:"title"`
	Price int    `xml:"price"`
}

type Catalog struct {
	Books []Book `xml:"book"`
}

const DOC = `<catalog>
  <book id="b1" lang="en"><title>Go</title><price>30</price></book>
  <book id="b2" lang="fr"><title>Rust</title><price>45</price></book>
</catalog>`

func main() {
	var c Catalog
	if err := xml.Unmarshal([]byte(DOC), &c); err != nil {
		panic(err)
	}
	var en []string
	for _, b := range c.Books {
		if b.Lang == "en" {
			en = append(en, b.Title)
		}
	}
	fmt.Println(strings.Join(en, ","))
}
