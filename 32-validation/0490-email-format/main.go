package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

type EmailFormat struct {
	Email string `validate:"email"`
}

func main() {
	v := validator.New()
	input := EmailFormat{Email: "not-an-email"}

	err := v.Struct(input)
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
