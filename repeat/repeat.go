package main

import (
	"fmt"
	"strings"
)

func Repeat(u string, count int) string {
	return strings.Repeat(u, count)
}

func main() {
	fmt.Println(Repeat("x", 5))
}
