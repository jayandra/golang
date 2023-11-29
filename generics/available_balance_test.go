package main

import (
	"reflect"
	"testing"
)

func TestAvailableBalance(t *testing.T) {
	invoices := []Invoice{
		Invoice{100.0, "Bob", 1},
		Invoice{50.0, "Bob", 1},
		Invoice{500.0, "Sally", 1},
		Invoice{200.0, "Bob", 2},
		Invoice{700.0, "Bob", 2},
		Invoice{400.0, "Sally", 2},
	}

	advances := []Advance{
		Advance{20, "Bob", 1},
		Advance{30, "Sally", 1},
		Advance{500, "Sally", 1},
		Advance{50, "Sally", 2},
	}

	t.Run("Testing ComputeTotal", func(t *testing.T) {
		invoices_acc := func(sum float64, x Invoice) float64 {
			return sum + x.NetEarning
		}
		got_inv := ComputeTotal(invoices, invoices_acc, 0)
		want_inv := 1950.0
		if got_inv != want_inv {
			t.Errorf("For Invoice: got %f, wanted %f", got_inv, want_inv)
		}

		advances_acc := func(sum float64, x Advance) float64 {
			return sum + x.WithdrawAmount
		}
		got_adv := ComputeTotal(advances, advances_acc, 0)
		want_adv := 600.0
		if got_adv != want_adv {
			t.Errorf("For Advances: got %f, wanted %f", got_adv, want_adv)
		}
	})

	t.Run("Testing FindAll", func(t *testing.T) {
		invoices_acc := func(x Invoice) bool {
			return x.User == "Bob"
		}
		got_inv := FindAll(invoices, invoices_acc)
		want_inv := []Invoice{}
		want_indices := []int{0, 1, 3, 4}
		for i := 0; i < len(want_indices); i++ {
			want_inv = append(want_inv, invoices[want_indices[i]])
		}
		if !reflect.DeepEqual(got_inv, want_inv) {
			t.Errorf("For Invoices: got %v, wanted %v", got_inv, want_inv)
		}

		advances_acc := func(x Advance) bool {
			return x.User == "Bob"
		}
		got_adv := FindAll(advances, advances_acc)
		want_adv := []Advance{}
		want_indices2 := []int{0}
		for i := 0; i < len(want_indices2); i++ {
			want_adv = append(want_adv, advances[want_indices2[i]])
		}
		if !reflect.DeepEqual(got_adv, want_adv) {
			t.Errorf("For Advances: got %v, wanted %v", got_adv, want_adv)
		}
	})

	t.Run("Testing AvailableBalance", func(t *testing.T) {
		got := AvailableBalance(invoices, advances, "Bob", 1)
		want := 130.0
		if got != want {
			t.Errorf("Invalid available balance got: %f, wanted %f", got, want)
		}
	})

}
