package main

import "fmt"

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	if amount > b.Balance {
		fmt.Println("Insufficient funds")
		return
	}
	b.Balance -= amount
}

func main() {
	acc := BankAccount{Balance: 00}
	acc.Deposit(30)
	fmt.Println("Balance after deposit:", acc.Balance)
	acc.Withdraw(30)
	fmt.Println("Balance after withdrawal:", acc.Balance)
}
