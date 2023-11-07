package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("jayandra", "")
		want := "Hello jayandra"

		assertCorrectMessage(t, got, want)
	})
	t.Run("default behavior", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("jayandra", "Spanish")
		want := "Hola jayandra"
		assertCorrectMessage(t, got, want)
	})
}

// ASK: What does the string here indicate? If it was return type shouldn't having it outside bracket work?
// What does the t *testing.TB and the one without * indicate?
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
