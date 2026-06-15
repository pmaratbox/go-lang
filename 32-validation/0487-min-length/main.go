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
	err := v.Struct(Model{Name: "al", Age: 30})
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
