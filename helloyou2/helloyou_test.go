package main

import "testing"

func TestSayHelloFunc(t *testing.T) {
	names := [...]string{
		"TATA",
		"RURU",
		"chibi",
		"",
	}
	for _, name := range names {
		if name == "" {
			name = "world"
		}
		want := "Hello, " + name
		got := sayhello(name)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
}
