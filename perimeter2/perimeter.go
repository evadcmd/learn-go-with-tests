package main

import (
	"fmt"
	"math"
)

type Rect struct {
	W float64
	H float64
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.W + r.H)
}

func (r *Rect) Area() float64 {
	return r.W * r.H
}

type Circle struct {
	R float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

func main() {
	r := Rect{5, 2}
	fmt.Println(r.Perimeter())
	c := Circle{4}
	fmt.Println(c.Area())
}
