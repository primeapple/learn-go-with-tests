package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Toni")

    got := buffer.String()
    want := "Hello, Toni"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

