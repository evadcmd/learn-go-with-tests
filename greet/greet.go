package main

import (
	"fmt"
	"io"
	"os"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func main() {
	greet(os.Stdout, "TATA")
}
