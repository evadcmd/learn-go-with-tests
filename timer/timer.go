package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Timer interface {
	Sleep()
}

type timer struct {
	period time.Duration
}

func (t *timer) Sleep() {
	time.Sleep(t.period)
}

func Countdown(i int, t Timer, out io.Writer) {
	for ; i > 0; i-- {
		fmt.Fprintln(out, i)
		t.Sleep()
	}
	fmt.Fprintln(out, "go!")
}

func main() {
	Countdown(5, &timer{time.Second}, os.Stdout)
}
