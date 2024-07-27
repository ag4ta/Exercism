package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fd FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := fd.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount * fatteningFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(fd FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fd, cows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}

	return nil
}
