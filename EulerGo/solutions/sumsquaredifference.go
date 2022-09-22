package solutions

func SumSquareDifference(n int) int {
	sum := (n * (n + 1)) / 2 // simply use the gauss formula and square
	squareSum := sum * sum
	sumOfSquares := (n * (n + 1) * (2*n + 1)) / 6 //Using another well know summation formula
	return squareSum - sumOfSquares
}
