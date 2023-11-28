// Package bank provided a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // recieve balance

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller gorutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
