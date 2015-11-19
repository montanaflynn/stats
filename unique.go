package stats

import (
	"errors"
)

// Unique returns any unique numbers in the data
func Unique(input Float64Data) (uniques []float64, err error) {

	// Return an error if there are no numbers
	if input.Len() == 0 {
		return nil, errors.New("Input must not be empty")
	}

	// Get the counts and only include the uniques
	c := counts(input)
	for key, value := range c {
		if value == 1 {
			uniques = append(uniques, key)
		}
	}

	return uniques, nil
}
