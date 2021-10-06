package main

import "fmt"

func sum(slice []int) (res int) {
	if len(slice) < 1 {
		return 0
	}
	for _, i := range slice[1:] {
		res += i
	}
	return res
}

func sumall(slices ...[]int) []int {
	res := []int{}
	for _, slice := range slices {
		res = append(res, sum(slice))
	}
	return res
}

func main() {
	fmt.Println(sumall([]int{1, 2, 3}, []int{2, 4, 6}))
}
