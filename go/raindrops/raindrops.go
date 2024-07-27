// Package raindrops provides one method to convert a number to corresponding raindrop sounds
package raindrops

import "strconv"

// Convert returns a string based on given number.
// If a given number:
// - is divisible by 3, add "Pling" to the result.
// - is divisible by 5, add "Plang" to the result.
// - is divisible by 7, add "Plong" to the result.
func Convert(number int) string {
	var result string

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
