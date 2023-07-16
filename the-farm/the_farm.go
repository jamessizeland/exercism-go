package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, noCows int) (float64, error) {
	total, err := fc.FodderAmount(noCows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (total / float64(noCows)) * factor, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, noCows int) (float64, error) {
	if ValidateNumberOfCows(noCows) != nil {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, noCows)
}

type InvalidCowsError struct {
	number int
}

func (e *InvalidCowsError) Error() string {
	switch {
	case e.number == 0:
		return fmt.Sprintf("%d cows are invalid: no cows don't need food", e.number)
	case e.number < 0:
		return fmt.Sprintf("%d cows are invalid: there are no negative cows", e.number)
	default:
		return fmt.Sprintf("%d cows: unknown error", e.number)
	}
}

func ValidateNumberOfCows(noCows int) error {
	if noCows <= 0 {
		return &InvalidCowsError{
			number: noCows,
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
