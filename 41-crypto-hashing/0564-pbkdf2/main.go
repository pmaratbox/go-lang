package main

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	key, err := pbkdf2.Key(sha256.New, "password", []byte("salt"), 1000, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(key))
}
