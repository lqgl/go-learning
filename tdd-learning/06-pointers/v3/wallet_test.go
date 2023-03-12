package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()

		fmt.Println("address of balance in test is", &wallet.balance)

		want := Bitcoin(10.0) // 生成 Bitcoin（比特币）
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}

		wallet.Withdraw(20)

		got := wallet.Balance()

		want := Bitcoin(80)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
