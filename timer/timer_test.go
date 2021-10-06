package main

import (
	"bytes"
	"fmt"
	"testing"
)

type mockTimer struct {
	t     *testing.T
	buf   *bytes.Buffer
	count int
}

// We intercept timer.Sleep to check previous output number
func (tm *mockTimer) Sleep() {
	tm.t.Helper()
	got := tm.buf.String()
	want := fmt.Sprintf("%d\n", tm.count)
	if got != want {
		tm.t.Errorf("got %s want %s", got, want)
	}
	tm.buf.Reset()
	tm.count--
}
func TestCountdown(t *testing.T) {
	from := 10
	buf := bytes.Buffer{}
	Countdown(from, &mockTimer{t, &buf, from}, &buf)
	want := "go!\n"
	got := buf.String()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
