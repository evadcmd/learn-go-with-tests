package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("BTC: %d coins", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Withdraw(b Bitcoin) bool {
	if w.balance < b {
		return false
	}
	w.balance -= b
	return false
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{}
	w.Deposit(5)
	fmt.Println(w.Balance())
	fmt.Println(Bitcoin(5))
}
