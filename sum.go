package main

func Sum(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]float64) []float64 {
	retval := make([]float64, len(numbers))
	for i, number := range numbers {
		retval[i] = Sum(number)
	}
	return retval
}
