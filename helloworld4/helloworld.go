package main

import "fmt"

var lang2hello = map[string]string{
	"":        "Hello",
	"English": "Hello",
	"Spanish": "Hola",
	"French":  "Bonjour",
}

func sayhello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s, %s", lang2hello[lang], name)
}

func main() {
	fmt.Println(sayhello("", ""))
}
