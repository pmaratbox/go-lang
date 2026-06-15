package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type Model struct {
	Password string `validate:"hasdigit"`
}

// hasDigit is a custom rule: the value must contain at least one digit.
func hasDigit(fl validator.FieldLevel) bool {
	for _, r := range fl.Field().String() {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func main() {
	v := validator.New()
	v.RegisterValidation("hasdigit", hasDigit)

	err := v.Struct(Model{Password: "abcdef"})
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
