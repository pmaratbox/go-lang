package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	mac := hmac.New(sha256.New, []byte("key"))
	mac.Write([]byte("hello"))
	sum := mac.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))
}
