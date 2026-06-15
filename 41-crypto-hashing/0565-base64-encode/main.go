package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("hello"))
	fmt.Println(encoded)
}
