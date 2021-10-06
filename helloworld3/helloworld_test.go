package main

import "testing"

func TestSayHelloFunc(t *testing.T) {

	assertEqual := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		names := [...]string{
			"TATA",
			"RURU",
			"chibi",
		}
		for _, name := range names {
			want := "Hello, " + name
			got := sayhello(name)
			assertEqual(t, got, want)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		want := "Hello, world"
		got := sayhello("")
		assertEqual(t, got, want)
	})
}
