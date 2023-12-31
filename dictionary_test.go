package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("known not avaiable", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is failing test"}
		got := dictionary.Search("exam")
		want := "n/a"
		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
