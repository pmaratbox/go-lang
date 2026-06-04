package main

import "fmt"

type account struct {
	balance int
}

func (a *account) deposit(amount int) {
	a.balance += amount
}

func (a *account) getBalance() int {
	return a.balance
}

func main() {
	a := &account{balance: 100}
	a.deposit(50)
	fmt.Println(a.getBalance())
}
