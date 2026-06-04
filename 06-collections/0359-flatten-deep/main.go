package main

import (
	"fmt"
	"strconv"
	"strings"
)

func flatten(node any, out *[]string) {
	switch v := node.(type) {
	case int:
		*out = append(*out, strconv.Itoa(v))
	case []any:
		for _, child := range v {
			flatten(child, out)
		}
	}
}

func main() {
	nested := []any{1, []any{2, []any{3, 4}}, 5}
	var out []string
	flatten(nested, &out)
	fmt.Println(strings.Join(out, " "))
}
