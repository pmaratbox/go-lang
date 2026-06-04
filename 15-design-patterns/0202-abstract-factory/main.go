package main

import "fmt"

type button struct{ label string }
type checkbox struct{ label string }

type themeFactory interface {
	makeButton() button
	makeCheckbox() checkbox
}

type darkFactory struct{}

func (darkFactory) makeButton() button     { return button{"dark-button"} }
func (darkFactory) makeCheckbox() checkbox { return checkbox{"dark-checkbox"} }

func main() {
	var f themeFactory = darkFactory{}
	b := f.makeButton()
	c := f.makeCheckbox()
	fmt.Println(b.label, c.label)
}
