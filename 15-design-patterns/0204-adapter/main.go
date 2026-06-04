package main

import "fmt"

type celsius struct{ degrees int }

type fahrenheitSource interface {
	fahrenheit() int
}

type celsiusAdapter struct{ src celsius }

func (a celsiusAdapter) fahrenheit() int { return a.src.degrees*9/5 + 32 }

func main() {
	var f fahrenheitSource = celsiusAdapter{src: celsius{degrees: 100}}
	fmt.Println(f.fahrenheit())
}
