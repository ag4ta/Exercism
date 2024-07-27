// Package collatzconjecture provides function to calculate Collatz Conjecture
package collatzconjecture

import "errors"

// CollatzConjecture returns a number of steps to reach number 1 based on given number as an argument applying therules:
// - If the number is even, divide it by two.
// - If the number is odd, triple it and add one.
func CollatzConjecture(n int) (int, error) {
	var result int

	if n < 1 {
		return result, errors.New("given number is non-positive")
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		result++
	}

	return result, nil
}
