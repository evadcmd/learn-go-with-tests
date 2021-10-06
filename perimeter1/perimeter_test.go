package main

import (
	"math"
	"testing"
)

const epsilon = 1e-10

func floatEqual(x float64, y float64) bool {
	return math.Abs(x-y) < epsilon
}

func TestCalPeri(t *testing.T) {
	got := calperi(6.7, 8.9)
	want := 2 * (6.7 + 8.9)
	if !floatEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
