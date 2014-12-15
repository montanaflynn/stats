package stats

import (
	"sort"
)

func Mean(input []int) (mean float64) {

	// Easy as pie, this one 
	if len(input) == 0 {
		return 0.0
	}

	// Add the total up
	for _, n := range input {
		mean += float64(n)
	}

	// Return the mean average
	return mean / float64(len(input))
}

func Median(input []int) (median float64) {

	// Sort the numbers
	sort.Ints(input)

	// Get the length
	l := len(input)

	// No math needed
	if l == 0 {
		return 0.0

		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
	} else if l%2 == 0 {
		median = Mean(input[l/2-1 : l/2+1])

		// For odd numbers we just use the middle number
	} else {
		median = float64(input[l/2])
	}

	return median
}

func Mode(input []int) []int {

	// Create a map to hold the counts
	m := make(map[int]int)
	for _, v := range input {
		m[v]++
	}

	// Find the highest counts to return as a slice
	// of ints to accomodate duplicate counts
	var mode []int
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
	l := len(input)
	lm := len(mode)

	if l == lm {
		return []int{}
	} else {
		return mode
	}
}
