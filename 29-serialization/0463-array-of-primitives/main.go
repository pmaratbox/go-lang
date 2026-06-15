package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	b, err := json.Marshal(nums)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
