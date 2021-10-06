package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(2, 4, 6)
	want := 2 + 4 + 6
	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}
