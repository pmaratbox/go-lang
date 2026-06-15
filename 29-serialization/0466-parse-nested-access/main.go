package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{"user":{"name":"alice","roles":["admin","user"]}}`)

	// Dynamic tree decode into generic Go values.
	var tree map[string]any
	if err := json.Unmarshal(data, &tree); err != nil {
		panic(err)
	}

	user := tree["user"].(map[string]any)
	name := user["name"].(string)
	roles := user["roles"].([]any)
	firstRole := roles[0].(string)

	fmt.Println(name, firstRole)
}
