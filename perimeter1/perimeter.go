package main

import (
	"fmt"
)

func calperi(w float64, h float64) float64 {
	return 2 * (w + h)
}

func main() {
	fmt.Println(calperi(5, 2))
}
