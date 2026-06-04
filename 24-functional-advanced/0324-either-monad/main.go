package main

import "fmt"

// Either holds a Right value or a Left error tag.
type Either struct {
	right int
	left  string
	isErr bool
}

func Right(v int) Either    { return Either{right: v} }
func Left(tag string) Either { return Either{left: tag, isErr: true} }

// bind continues only on Right; Left short-circuits.
func (e Either) bind(f func(int) Either) Either {
	if e.isErr {
		return e
	}
	return f(e.right)
}

func divBy(d int) func(int) Either {
	return func(x int) Either {
		if d == 0 {
			return Left("err")
		}
		return Right(x / d)
	}
}

func main() {
	ok := Right(8).bind(divBy(2)).bind(divBy(2))   // 8/2/2 = 2
	bad := Right(8).bind(divBy(0)).bind(divBy(2))   // divide by zero

	left := fmt.Sprintf("%d", ok.right)
	right := bad.left
	fmt.Println(left, right)
}
