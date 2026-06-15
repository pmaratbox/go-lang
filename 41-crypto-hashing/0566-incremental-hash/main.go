package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("foo"))
	h.Write([]byte("bar"))
	sum := h.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))
}
