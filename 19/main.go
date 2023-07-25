package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunction(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}

	withdrawFunction(a, 30)
	fmt.Println("withdrawFunction : ", a.balance)

	a.withdrawMethod(20)
	fmt.Println("withdrawMethod : ", a.balance)
}
