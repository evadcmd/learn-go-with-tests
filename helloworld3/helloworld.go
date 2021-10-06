package main

import "fmt"

func sayhello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(sayhello(""))
}
