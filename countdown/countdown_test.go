package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
    timeSlept time.Duration
}
func (s *SpyTime) Sleep(d time.Duration) {
    s.timeSlept = d
}

func TestCountdown(t *testing.T) {
	t.Run("prints correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		Countdown(sleeper, buffer)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleeps in between each call", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		got := spy.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected calls %v got %v", want, got)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.timeSlept != sleepTime {
        t.Errorf("Expected to sleep %v but slept %v", sleepTime, spyTime.timeSlept)
    }
}
