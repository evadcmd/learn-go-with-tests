package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dict map[string]string

func (d Dict) Get(k string) (string, error) {
	var err error
	v, ok := d[k]
	if !ok {
		err = ErrNotFound
	}
	return v, err
}

func (d Dict) Put(k string, v string) {
	d[k] = v
}

func main() {
	// or d := make(Dict)
	d := Dict{}
	v, err := d.Get("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	d.Put("hello", "hello world")
	v, err = d.Get("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
