package solutions

import "math"

func SumEvenFibNumbers() (sum int) {
	n := 1
	fibVal := 0
	for {
		//Every 3rd value is an even value
		fibVal = fibonacci(n * 3)
		if fibVal > 4000000 {
			break
		}
		sum += fibVal
		n++
	}

	return
}

//1, 1, 2, 3, 5, 8, 13, 21, 34, ...
//Notice pattern is every 3rd value is even.
//Use Binet's formula for O(1) fibonacci
func fibonacci(n int) (fib int) {
	goldenRatio := 1.618
	inverseGR := -0.618
	floatN := float64(n)
	floatVal := (math.Pow(goldenRatio, floatN) - math.Pow(inverseGR, floatN)) / math.Sqrt(5)
	return int(math.Round(floatVal))
}
