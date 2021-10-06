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
		name      string
		shape     Shape
		perimeter float64
	}{
		{name: "Rect1", shape: &Rect{}, perimeter: 0},
		{name: "Rect2", shape: &Rect{5, 2}, perimeter: 14},
		{name: "Rect3", shape: &Rect{5.2, 2.5}, perimeter: 2 * (5.2 + 2.5)},
		{name: "Circ1", shape: &Circle{5}, perimeter: 2 * math.Pi * 5},
		{name: "Circ2", shape: &Circle{2.5}, perimeter: 2 * math.Pi * 2.5},
	}

	for _, test := range tests {
		// run a single test
		// go test -run ${FUNC_NAME}/${test.name}
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Perimeter()
			want := test.perimeter
			t.Logf("got %v want %v\n", got, want)
			if !floatEqual(got, want) {
				t.Errorf("got %v want %v\n", got, want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rect1", shape: &Rect{}, area: 0},
		{name: "Rect2", shape: &Rect{5, 2}, area: 10},
		{name: "Rect3", shape: &Rect{5.2, 2.5}, area: 5.2 * 2.5},
		{name: "Circ1", shape: &Circle{5}, area: math.Pi * 5 * 5},
		{name: "Circ2", shape: &Circle{2.5}, area: math.Pi * 2.5 * 2.5},
	}

	for _, test := range tests {
		// run a single test
		// go test -run ${FUNC_NAME}/${test.name}
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			want := test.area
			t.Logf("got %v want %v\n", got, want)
			if !floatEqual(got, want) {
				t.Errorf("got %v want %v\n", got, want)
			}
		})
	}
}
