package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
		got := Sum(numbers)
		want := 15.0
		if got != want {
			t.Errorf("got %f wanted %f given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]float64{1, 2}, []float64{3, 4})
	want := []float64{3, 7}
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("got %v wanted %v", got, want)
	}
}
