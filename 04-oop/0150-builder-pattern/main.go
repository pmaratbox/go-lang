package main

import (
	"fmt"
	"strings"
)

type Pizza struct {
	size     string
	toppings []string
}

func (p Pizza) String() string {
	return fmt.Sprintf("Pizza(%s, %s)", p.size, strings.Join(p.toppings, ", "))
}

type PizzaBuilder struct {
	pizza Pizza
}

func (b *PizzaBuilder) setSize(size string) *PizzaBuilder {
	b.pizza.size = size
	return b
}

func (b *PizzaBuilder) addTopping(topping string) *PizzaBuilder {
	b.pizza.toppings = append(b.pizza.toppings, topping)
	return b
}

func (b *PizzaBuilder) build() Pizza {
	return b.pizza
}

func main() {
	pizza := (&PizzaBuilder{}).setSize("M").addTopping("cheese").build()
	fmt.Println(pizza)
}
