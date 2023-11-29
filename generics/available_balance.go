package main

type Invoice struct {
	NetEarning float64
	User       string
	PayWeek    int
}

type Advance struct {
	WithdrawAmount float64
	User           string
	WithdrawWeek   int
}

func ComputeTotal[A any](collection []A, accumulator func(acc float64, x A) float64, initialValue float64) float64 {
	var total = initialValue
	for _, x := range collection {
		total = accumulator(total, x)
	}
	return total
}

func FindAll[A any](collection []A, predicate func(A) bool) []A {
	return_slice := make([]A, 0)

	for _, x := range collection {
		if predicate(x) {
			return_slice = append(return_slice, x)
		}
	}

	return return_slice
}

func AvailableBalance(invoices []Invoice, advances []Advance, name string, payperiod int) float64 {
	invoices_matcher := func(x Invoice) bool {
		return x.User == name && x.PayWeek == payperiod
	}
	advances_matcher := func(x Advance) bool {
		return x.User == name && x.WithdrawWeek == payperiod
	}
	invoices_accmulator := func(sum float64, x Invoice) float64 {
		return sum + x.NetEarning
	}

	advances_accumulator := func(sum float64, x Advance) float64 {
		return sum + x.WithdrawAmount
	}

	matched_invoices := FindAll(invoices, invoices_matcher)
	matched_invoices_sum := ComputeTotal(matched_invoices, invoices_accmulator, 0)

	matched_advances := FindAll(advances, advances_matcher)
	matched_advances_sum := ComputeTotal(matched_advances, advances_accumulator, 0)

	return (matched_invoices_sum - matched_advances_sum)
}
