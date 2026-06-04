package main

import "fmt"

type Option struct {
	value   int
	present bool
}

func some(v int) Option { return Option{value: v, present: true} }
func none() Option      { return Option{} }

func (o Option) mapInt(f func(int) int) Option {
	if !o.present {
		return o
	}
	return some(f(o.value))
}

func (o Option) unwrapOr(s string) string {
	if o.present {
		return fmt.Sprintf("%d", o.value)
	}
	return s
}

func main() {
	add2 := func(x int) int { return x + 2 }
	a := some(10).mapInt(add2)
	b := none().mapInt(add2)
	fmt.Println(a.unwrapOr("none"), b.unwrapOr("none"))
}
