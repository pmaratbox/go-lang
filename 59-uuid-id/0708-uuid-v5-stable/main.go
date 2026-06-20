package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// UUIDv5 is name-based (SHA-1) and deterministic: the same
	// (namespace, name) pair always produces the same UUID.
	a := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
	b := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
	fmt.Println(a == b)
}
