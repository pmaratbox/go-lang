package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	sum := sha512.Sum512([]byte("hello"))
	fmt.Println(hex.EncodeToString(sum[:]))
}
