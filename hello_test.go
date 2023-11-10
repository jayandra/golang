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

func TestIterateHello(t *testing.T) {
	got := iterateHello(2, "jayandra", "Spanish")
	want := "Hola jayandra\nHola jayandra\n"
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
