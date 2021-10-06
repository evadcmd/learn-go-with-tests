package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("BTC: %d coins", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance < b {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
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
