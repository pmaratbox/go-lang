package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Model requires both Name and Age. Age uses the `required` constraint,
// which fails when the value is its zero value (0 / missing).
type Model struct {
	Name string `validate:"required,min=3"`
	Age  int    `validate:"required,min=0,max=120"`
}

func main() {
	v := validator.New()

	// Name is present; Age is missing (zero value).
	err := v.Struct(Model{Name: "alice"})

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
