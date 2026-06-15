package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Model struct {
	Name string `validate:"min=3"`
	Age  int    `validate:"min=0,max=120"`
}

func main() {
	v := validator.New()
	err := v.Struct(Model{Name: "al", Age: 200})
	if err == nil {
		fmt.Println("ok")
		return
	}
	seen := map[string]bool{}
	var fields []string
	for _, e := range err.(validator.ValidationErrors) {
		f := strings.ToLower(e.Field())
		if !seen[f] {
			seen[f] = true
			fields = append(fields, f)
		}
	}
	sort.Strings(fields)
	fmt.Println(strings.Join(fields, "\n"))
}
