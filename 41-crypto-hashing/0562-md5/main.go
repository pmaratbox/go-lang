package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	sum := md5.Sum([]byte("hello"))
	fmt.Println(hex.EncodeToString(sum[:]))
}
