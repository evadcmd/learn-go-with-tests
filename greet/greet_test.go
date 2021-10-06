package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	greet(&buf, "RURU")
	want := "Hello, RURU\n"
	got := buf.String()
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}
