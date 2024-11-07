package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Toni", "")
		want := "Hello, Toni"
        assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World', when called with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
        assertCorrectMessage(t, got, want)
	})

    t.Run("in spanish", func(t *testing.T) {
        got := Hello("Toni", "Spanish")
        want := "Hola, Toni"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in french", func(t *testing.T) {
        got := Hello("Toni", "French")
        want := "Bonjour, Toni"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
