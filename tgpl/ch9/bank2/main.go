// Package bank provided a concurrency-safe bank with one account.
package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // accquire token
	balance += amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // accquire token
	b := balance
	<-sema // release token
	return b
}
