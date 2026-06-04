package main

import (
	"fmt"
	"sync"
)

type config struct{}

var (
	instance *config
	once     sync.Once
)

func getInstance() *config {
	once.Do(func() {
		instance = &config{}
	})
	return instance
}

func main() {
	a := getInstance()
	b := getInstance()
	if a == b {
		fmt.Println("same: yes")
	} else {
		fmt.Println("same: no")
	}
}
