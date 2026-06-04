package main

import "fmt"

type UserId int
type ProductId int

func (u UserId) String() string    { return fmt.Sprintf("user-%d", int(u)) }
func (p ProductId) String() string { return fmt.Sprintf("prod-%d", int(p)) }

func main() {
	u := UserId(1)
	p := ProductId(2)
	fmt.Println(u, p)
}
