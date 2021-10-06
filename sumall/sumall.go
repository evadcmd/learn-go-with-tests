package main

import "fmt"

func sumall(slices ...[]int) []int {
	res := []int{}
	for _, slice := range slices {
		var sum int
		for _, s := range slice {
			sum += s
		}
		res = append(res, sum)
	}
	return res
}

func main() {
	fmt.Println(sumall([]int{1, 2, 3}, []int{2, 4, 6}))
}
