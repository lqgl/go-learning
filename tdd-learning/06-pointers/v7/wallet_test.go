package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		fmt.Println("address of balance in test is", &wallet.balance)
		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}

		err := wallet.Withdraw(20)

		assertBalance(wallet, Bitcoin(80), t)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(wallet, startingBalance, t)

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

// go install github.com/kisielk/errcheck
