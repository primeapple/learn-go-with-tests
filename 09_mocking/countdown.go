package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
    Sleep()
}

type ConfigurableSleeper struct {
    duration time.Duration
    sleep func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
    c.sleep(c.duration)
}

func Countdown(sleeper Sleeper, out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
        sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(&ConfigurableSleeper{1 * time.Second, time.Sleep}, os.Stdout)
}
