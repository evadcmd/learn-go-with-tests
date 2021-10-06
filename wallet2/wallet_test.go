package main

import "testing"

func assertBalanceEqual(t *testing.T, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}

func assertErrEqual(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{50}
		w.Deposit(30)
		assertBalanceEqual(t, w.Balance(), 80)
	})
	t.Run("withdraw with sufficient fund", func(t *testing.T) {
		w := Wallet{50}
		w.Withdraw(30)
		assertBalanceEqual(t, w.Balance(), 20)
	})
	t.Run("withdraw with insufficient fund", func(t *testing.T) {
		w := Wallet{50}
		err := w.Withdraw(100)
		assertErrEqual(t, err, ErrInsufficientFunds)
		assertBalanceEqual(t, w.Balance(), 50)
	})
}
