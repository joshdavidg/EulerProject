package solutions

func LargestPalidromeProduct() (largestPal int) {
	var product int = 0
	for i := 999; i > 100; i-- {
		//Inner loop only needs to check from the current n value instead of max 999 value all the time
		//This is because we already checked all n * every number between 999 and 100, no need to use the same n again
		// n * i == i * j
		for j := i; j > 100; j-- {
			product = i * j
			if isPalidrome(product) && product > largestPal {
				largestPal = product
			}
		}
	}
	return
}

//Mathematically reverses the number and compares
func isPalidrome(n int) bool {
	numForward := n
	numReverse := 0

	for numForward > 0 {
		lastDigit := numForward % 10
		numReverse = (numReverse * 10) + lastDigit
		numForward = numForward / 10
	}

	return n == numReverse
}
