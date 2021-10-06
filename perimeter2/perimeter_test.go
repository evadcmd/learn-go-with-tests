package main

import (
	"math"
	"testing"
)

const epsilon = 1e-10

func floatEqual(x float64, y float64) bool {
	return math.Abs(x-y) < epsilon
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func TestPeri(t *testing.T) {
	tests := []struct {
		shape     Shape
		perimeter float64
	}{
		{&Rect{}, 0},
		{&Rect{5, 2}, 14},
		{&Rect{5.2, 2.5}, 2 * (5.2 + 2.5)},
		{&Circle{5}, 2 * math.Pi * 5},
		{&Circle{2.5}, 2 * math.Pi * 2.5},
	}

	for _, test := range tests {
		got := test.shape.Perimeter()
		want := test.perimeter
		t.Logf("got %v want %v\n", got, want)
		if !floatEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		shape Shape
		area  float64
	}{
		{&Rect{}, 0},
		{&Rect{5, 2}, 10},
		{&Rect{5.2, 2.5}, 5.2 * 2.5},
		{&Circle{5}, math.Pi * 5 * 5},
		{&Circle{2.5}, math.Pi * 2.5 * 2.5},
	}

	for _, test := range tests {
		got := test.shape.Area()
		want := test.area
		t.Logf("got %v want %v\n", got, want)
		if !floatEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	}
}
