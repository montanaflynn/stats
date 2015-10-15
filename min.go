package stats

import (
	"errors"
)

// Min finds the lowest number in a set of data
func Min(input Float64Data) (min float64, err error) {

	// Get the count of numbers in the slice
	l := input.Len()

	// Return an error if there are no numbers
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the first value as the starting point
	min = input.Get(0)

	// Iterate until done checking for a lower value
	for i := 1; i < l; i++ {
		if input.Get(i) < min {
			min = input.Get(i)
		}
	}
	return min, nil
}
