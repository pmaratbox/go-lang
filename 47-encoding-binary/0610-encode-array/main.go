package main

import (
	"encoding/hex"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	b, err := msgpack.Marshal([]int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(b))
}
