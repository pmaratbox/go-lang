package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Parse a UUID and report its version number.
	u := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")

	// u.Version().String() is "VERSION_4"; int() gives the bare number.
	fmt.Println(int(u.Version()))
}
