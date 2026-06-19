package main

import (
	"encoding/hex"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	bytes, err := hex.DecodeString("a568656c6c6f")
	if err != nil {
		panic(err)
	}
	var s string
	if err := msgpack.Unmarshal(bytes, &s); err != nil {
		panic(err)
	}
	fmt.Println(s)
}
