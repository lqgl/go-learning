package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()

	fmt.Println("address of balance in test is", &wallet.balance)

	want := 10.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
