package main

import "fmt"

type memento struct{ state int }

type originator struct{ state int }

func (o *originator) save() memento     { return memento{state: o.state} }
func (o *originator) restore(m memento) { o.state = m.state }

func main() {
	o := &originator{state: 1}
	saved := o.save()
	o.state = 2
	current := o.state
	o.restore(saved)
	fmt.Println(current, o.state)
}
