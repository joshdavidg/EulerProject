package solutions

import "fmt"

func sumMultiplesOf3or5() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
