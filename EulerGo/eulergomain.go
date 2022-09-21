package main

import (
	"fmt"

	"github.com/joshdavidg/EulerProject/EulerGo/solutions"
)

func main() {
	q1 := solutions.SumMultiplesOf3or5()
	fmt.Println("Sum of multiples of 3 or 5 from 1 to 1000: ", q1)
	q2 := solutions.SumEvenFibNumbers()
	fmt.Println("Sum of even fibonacci terms less than 4 mill: ", q2)
	q3 := solutions.LargestPrimeFactor(600851475143)
	fmt.Println("The largest prime factor of 600851475143 is: ", q3)
	q4 := solutions.LargestPalidromeProduct()
	fmt.Println("The largest palidrome product from two 3 digit numbers is: ", q4)
	q5 := solutions.SmallestMultiple()
	fmt.Println("The smallest number divisible by all numbers 1 -> 20 is: ", q5)
}
