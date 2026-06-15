package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha1.Sum([]byte("hello"))
	fmt.Println(hex.EncodeToString(h[:]))
}
