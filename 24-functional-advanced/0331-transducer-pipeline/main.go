package main

import (
	"fmt"
	"strings"
)

// reducer accumulates an int into a slice.
type reducer func([]int, int) []int

// transducer transforms a reducer into another reducer.
type transducer func(reducer) reducer

func mapping(f func(int) int) transducer {
	return func(next reducer) reducer {
		return func(acc []int, x int) []int { return next(acc, f(x)) }
	}
}

func filtering(pred func(int) bool) transducer {
	return func(next reducer) reducer {
		return func(acc []int, x int) []int {
			if pred(x) {
				return next(acc, x)
			}
			return acc
		}
	}
}

// compose chains transducers left-to-right.
func compose(ts ...transducer) transducer {
	return func(r reducer) reducer {
		for i := len(ts) - 1; i >= 0; i-- {
			r = ts[i](r)
		}
		return r
	}
}

func main() {
	xform := compose(
		mapping(func(x int) int { return x + 1 }),
		filtering(func(x int) bool { return x%2 == 0 }),
	)
	appendInt := func(acc []int, x int) []int { return append(acc, x) }
	step := xform(appendInt)

	acc := []int{}
	for _, x := range []int{1, 2, 3, 4} {
		acc = step(acc, x)
	}

	parts := make([]string, len(acc))
	for i, v := range acc {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
