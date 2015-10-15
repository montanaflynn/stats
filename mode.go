package stats

import (
	"errors"
)

// Mode gets the mode of a slice of numbers
func Mode(input Float64Data) (mode []float64, err error) {

	// Return the input if there's only one number
	l := input.Len()
	if l == 1 {
		return input, nil
	} else if l == 0 {
		return nil, errors.New("Input must not be empty")
	}

	// Create a map with the counts for each number
	m := make(map[float64]int)
	for _, v := range input {
		m[v]++
	}

	// Find the highest counts to return as a slice
	// of ints to accomodate duplicate counts
	var current int
	for k, v := range m {

		// Depending if the count is lower, higher
		// or equal to the current numbers count
		// we return nothing, start a new mode or
		// append to the current mode
		switch {
		case v < current:
		case v > current:
			current = v
			mode = append(mode[:0], k)
		default:
			mode = append(mode, k)
		}
	}

	// Finally we check to see if there actually was
	// a mode by checking the length of the input and
	// mode against eachother
	lm := len(mode)
	if l == lm {
		return Float64Data{}, nil
	}

	return mode, nil
}
