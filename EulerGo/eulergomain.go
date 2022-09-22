package main

import (
	"fmt"

	"github.com/joshdavidg/EulerProject/EulerGo/solutions"
)

func main() {
	q1 := solutions.SumMultiplesOf3or5()
	fmt.Println("Q1 - Sum of multiples of 3 or 5 from 1 to 1000: ", q1)
	q2 := solutions.SumEvenFibNumbers()
	fmt.Println("Q2 - Sum of even fibonacci terms less than 4 mill: ", q2)
	q3 := solutions.LargestPrimeFactor(600851475143)
	fmt.Println("Q3 - The largest prime factor of 600851475143 is: ", q3)
	q4 := solutions.LargestPalidromeProduct()
	fmt.Println("Q4 - The largest palidrome product from two 3 digit numbers is: ", q4)
	q5 := solutions.SmallestMultiple()
	fmt.Println("Q5 - The smallest number divisible by all numbers 1 -> 20 is: ", q5)
	q6 := solutions.SumSquareDifference(100)
	fmt.Println("Q6 - The difference between the sum of squares and the square of the sum for the first 100 natural numbers is: ", q6)
}
