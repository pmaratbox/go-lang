package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// UUIDv5 is name-based (SHA-1): it hashes the (namespace, name) pair,
	// so a different name yields a different — but still deterministic — UUID.
	u := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("test.com"))
	fmt.Println(u)
}
