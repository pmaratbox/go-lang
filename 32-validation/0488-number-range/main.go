package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Model uses go-playground/validator tags: Age must satisfy min=0 and max=120,
// so the value 200 is out of range and fails the max constraint.
type Model struct {
	Name string `validate:"min=3"`
	Age  int    `validate:"min=0,max=120"`
}

func main() {
	v := validator.New()

	err := v.Struct(Model{Name: "alice", Age: 200})
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
