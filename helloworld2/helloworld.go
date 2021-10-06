package main

import "fmt"

func sayhello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(sayhello("TATA"))
}
