package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// UUIDv5 is name-based (SHA-1): deterministic from (namespace, name).
	id := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
	fmt.Println(id)
}
