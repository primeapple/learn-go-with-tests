package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
    sleeper := &SpySleeper{}

	Countdown(sleeper, buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

    if sleeper.Calls != 3 {
        t.Errorf("expected three calls got %d", sleeper.Calls)
    }
}
