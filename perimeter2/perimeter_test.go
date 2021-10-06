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
		{shape: &Rect{}, perimeter: 0},
		{shape: &Rect{5, 2}, perimeter: 14},
		{shape: &Rect{5.2, 2.5}, perimeter: 2 * (5.2 + 2.5)},
		{shape: &Circle{5}, perimeter: 2 * math.Pi * 5},
		{shape: &Circle{2.5}, perimeter: 2 * math.Pi * 2.5},
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
		{shape: &Rect{}, area: 0},
		{shape: &Rect{5, 2}, area: 10},
		{shape: &Rect{5.2, 2.5}, area: 5.2 * 2.5},
		{shape: &Circle{5}, area: math.Pi * 5 * 5},
		{shape: &Circle{2.5}, area: math.Pi * 2.5 * 2.5},
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
