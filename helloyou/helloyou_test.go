package main

import (
	"fmt"
	"testing"
)

func TestSayHelloFunc(t *testing.T) {
	names := [...]string{
		"TATA", "RURU", "chibi",
	}
	for _, name := range names {
		got := sayhello(name)
		want := fmt.Sprintf("Hello, %s", name)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
}
