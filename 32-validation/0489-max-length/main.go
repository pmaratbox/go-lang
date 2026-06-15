package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Model uses go-playground/validator tags: Code must satisfy max=5,
// so the 7-character value "ABCDEFG" violates the max constraint.
type Model struct {
	Code string `validate:"max=5"`
}

func main() {
	v := validator.New()

	err := v.Struct(Model{Code: "ABCDEFG"})
	if err == nil {
		fmt.Println("ok")
		return
	}

	var fields []string
	for _, e := range err.(validator.ValidationErrors) {
		fields = append(fields, strings.ToLower(e.Field()))
	}
	sort.Strings(fields)
	fmt.Println(strings.Join(fields, "\n"))
}
