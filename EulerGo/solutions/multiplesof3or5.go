package solutions

func SumMultiplesOf3or5() (sum int) {
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}

	return
}
