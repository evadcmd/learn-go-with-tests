package main

import "testing"

func TestSayHelloFunc(t *testing.T) {
	/*
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
	*/
	t.Run("saying hello to people", func(t *testing.T) {
		names := [...]string{
			"TATA",
			"RURU",
			"chibi",
		}
		for _, name := range names {
			want := "Hello, " + name
			got := sayhello(name)
			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		want := "Hello, world"
		got := sayhello("")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
