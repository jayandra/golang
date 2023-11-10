package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
}
