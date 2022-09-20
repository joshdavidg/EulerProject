package solutions

func LargestPrimeFactor(number int) int {
	largestFactor := 1
	largestPrimeFactor := 1
	//Start with trivial primes
	for number%2 == 0 {
		largestFactor = 2
		number = number / 2
	}
	for number%3 == 0 {
		largestFactor = 3
		number = number / 3
	}

	//Loop through prime numbers and check divisibilty
	//Start at 5 (prime) and skip all evens, and numbers divisible by 2,3 or 5 by following 6k +- 1 check
	for i := 5; i*i <= number; i += 6 {
		for number%i == 0 {
			largestFactor = i
			number = number / i
		}

		for number%(i+2) == 0 {
			largestFactor = i
			number = number / (i + 2)
		}
	}

	largestPrimeFactor = largestFactor

	//Check that the current factor is prime and if not take the product as largest prime factor
	if !isPrime(largestFactor) {
		largestPrimeFactor = number
	}
	return largestPrimeFactor
}

//Helper Check for last divisor to determine if product or devisor is the largest prime factor
func isPrime(n int) bool {
	if n <= 3 {
		return n > 0
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
