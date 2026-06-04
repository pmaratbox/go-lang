package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "roundtrip-*.txt")
	if err != nil {
		panic(err)
	}
	name := f.Name()

	want := "ok-data"
	if _, err := f.WriteString(want); err != nil {
		panic(err)
	}
	f.Close()

	got, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	os.Remove(name)

	if string(got) != want {
		panic("mismatch")
	}
	fmt.Println("roundtrip: ok")
}
