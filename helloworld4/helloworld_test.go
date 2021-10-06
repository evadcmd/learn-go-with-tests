package main

import "testing"

func TestSayHelloFunc(t *testing.T) {
	// helper
	assertEqual := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("say hello with Spanish to someone", func(t *testing.T) {
		names := [...]string{"TATA", "RURU", "chibi"}
		for _, name := range names {
			want := "Hola, " + name
			got := sayhello(name, "Spanish")
			assertEqual(t, got, want)
		}
	})

	t.Run("say hello with English to someone", func(t *testing.T) {
		names := [...]string{"TATA", "RURU", "chibi"}
		for _, name := range names {
			want := "Hello, " + name
			got := sayhello(name, "English")
			assertEqual(t, got, want)
		}
	})

	t.Run("say hello to the world with Spanish", func(t *testing.T) {
		got := sayhello("", "Spanish")
		want := "Hola, world"
		assertEqual(t, got, want)
	})

	t.Run("say hello to the world with English", func(t *testing.T) {
		got := sayhello("", "English")
		want := "Hello, world"
		assertEqual(t, got, want)
	})
}
