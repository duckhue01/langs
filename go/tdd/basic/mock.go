package main

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}



func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
