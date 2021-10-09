package main

import "testing"

func TestWalk(t *testing.T) {
	s := struct {
		X string
		Y string
		T int
		Z struct {
			A string
			B string
		}
	}{
		"X",
		"Y",
		0,
		struct {
			A string
			B string
		}{
			"A",
			"B",
		},
	}
	called := []string{}
	c := func(in string) {
		called = append(called, in)
	}
	Walk(s, c)
	for i, want := range []string{"X", "Y", "A", "B"} {
		got := called[i]
		if got != want {
			t.Errorf("got %s want %s\n", got, want)
		}
	}
}
