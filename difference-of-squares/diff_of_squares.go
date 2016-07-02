package diffsquares

// SumOfSquares returns sum of squres
func SumOfSquares(n int) int {

	sumOfSquares := 0
	for num := 1; num <= n; num++ {

		square := num * num
		sumOfSquares += square
	}
	return sumOfSquares
}

// SquareOfSums returns sqare of sums
func SquareOfSums(n int) int {

	sums := 0
	for num := 1; num <= n; num++ {

		sums += num
	}
	return sums * sums
}

// Difference returns the difference between SqareOFSums and SumOfSquares
func Difference(n int) int {

	return SquareOfSums(n) - SumOfSquares(n)
}
