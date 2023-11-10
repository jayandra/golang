package main

import (
	"reflect"
	"testing"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, e error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	// buffer := &bytes.Buffer{}
	spySleeper := &SpyCountdownOperations{}

	Countdown(spySleeper, spySleeper)

	want := []string{
		"write",
		"sleep",
		"write",
		"sleep",
		"write",
		"sleep",
	}

	if !reflect.DeepEqual(want, spySleeper.Calls) {
		t.Errorf("got %v want %v", want, spySleeper.Calls)
	}
}
