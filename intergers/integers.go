package main

import "fmt"

func sum(ints ...int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 3, 5, 6))
}
