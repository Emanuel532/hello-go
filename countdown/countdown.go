package main

import (
	"fmt"
	"io"
	"iter"
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
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func countDownFrom(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(w io.Writer, s Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, &sleeper)
}
