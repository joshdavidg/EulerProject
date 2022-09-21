package solutions

//Brute forced method - probably better method to improve performance.
func SmallestMultiple() (smallestMul int) {
	i := 2520 // already know this is the smallest multiple for 1 -> 10
	var j int
	for {
		//Only really need to check 11 to 20 because 1 to 10 will all be divisible by 11 to 20.
		for j = 11; j <= 20; j++ {
			if i%j != 0 {
				break
			}
		}
		if j == 21 {
			break
		}
		i++
	}
	return i
}
