package main

import (
	"encoding/hex"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	b, err := msgpack.Marshal(map[string]int{"a": 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(b))
}
