package main

import "testing"

func TestDeposit(t *testing.T) {
	bitcoins := []Bitcoin{
		3, 4, 5, 10, 11,
	}
	w := Wallet{}
	var want Bitcoin
	for _, bitcoin := range bitcoins {
		w.Deposit(bitcoin)
		want += bitcoin
		got := w.Balance()
		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	}
}

func TestWithdraw(t *testing.T) {
	testcases := []struct {
		bitcoin Bitcoin
		status  bool
	}{
		{bitcoin: 3, status: true},
		{bitcoin: 4, status: true},
		{bitcoin: 5, status: true},
		{bitcoin: 10, status: true},
		{bitcoin: 11, status: false},
	}
	w := Wallet{32}
	want := Bitcoin(32)
	for _, test := range testcases {
		w.Withdraw(test.bitcoin)
		if want > test.bitcoin {
			want -= test.bitcoin
		}
		got := w.Balance()
		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	}
}
