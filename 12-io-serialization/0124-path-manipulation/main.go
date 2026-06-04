package main

import (
	"fmt"
	"path"
)

func main() {
	full := path.Join("/tmp", "file.txt")
	base := path.Base(full)
	ext := path.Ext(full)
	fmt.Println(full, base, ext)
}
