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

	counts := map[float64]int{}
	for _, value := range input {
		counts[value]++
	}

	for key, value := range counts {
		if value == 1 {
			uniques = append(uniques, key)
		}
	}

	return uniques, nil
}
