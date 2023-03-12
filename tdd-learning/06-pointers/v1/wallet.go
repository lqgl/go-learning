package main

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
