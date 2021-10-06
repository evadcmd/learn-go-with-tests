package main

import "testing"

func TestSayHelloFunc(t *testing.T) {
	want := "Hello, world"
	got := sayhello()
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func TestChangeHelloVar(t *testing.T) {
	saved := hello
	defer func() { hello = saved }()
	hello = "Hello, goodbye"
	got := sayhello()
	if got == "Hello, world" {
		t.Errorf("got '%q' want '%q'", got, hello)
	}
}
