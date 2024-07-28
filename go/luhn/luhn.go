package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	var result int
	secondDigit := len(id)%2 == 0

	for _, num := range id {
		integer, err := strconv.Atoi(string(num))
		if err != nil {
			return false
		}

		if secondDigit {
			integer *= 2

			if integer > 9 {
				integer -= 9
			}
		}
		result += integer
		secondDigit = !secondDigit
	}

	return result%10 == 0
}
