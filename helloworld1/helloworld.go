package main

import "fmt"

var hello = "Hello, world"

func sayhello() string {
	return hello
}

func main() {
	fmt.Println(sayhello())
}
