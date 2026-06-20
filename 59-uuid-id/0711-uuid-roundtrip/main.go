package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Parse an UPPERCASE UUID; uuid.Parse accepts mixed case.
	u := uuid.MustParse("550E8400-E29B-41D4-A716-446655440000")

	// String() always renders the canonical lowercase form.
	fmt.Println(u.String())
}
