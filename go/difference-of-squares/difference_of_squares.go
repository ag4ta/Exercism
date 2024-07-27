// Package diffsquares provides functions to calculate Sum square difference
package diffsquares

// SquareOfSum calculates the square of sum of 1 to n and returns integer
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates sum of first n squares and returns integer
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between square of sum and sum of squares and returns integer
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
