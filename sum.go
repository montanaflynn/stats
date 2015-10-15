package stats

import (
	"errors"
)

// Sum adds all the numbers of a slice together
func Sum(input Float64Data) (sum float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Add em up
	for _, n := range input {
		sum += n
	}

	return sum, nil
}
