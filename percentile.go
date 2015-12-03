package stats

import (
	"errors"
	"math"
)

// Percentile finds the relative standing in a slice of floats
func Percentile(input Float64Data, percent float64) (percentile float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	if percent == 0 {
		return 0, errors.New("Percent must not be empty")
	}

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// Multiply percent by length of input
	index := (percent / 100) * float64(len(c))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := float64ToInt(index)

		// Find the average of the index and following values
		percentile, _ = Mean(Float64Data{c[i-1], c[i]})

	} else {

		// Convert float to int
		i := float64ToInt(index)

		// Find the value at the index
		percentile = c[i-1]

	}

	return percentile, nil

}

// PercentileNearestRank finds the relative standing in a slice of floats using the Nearest Rank method
func PercentileNearestRank(input Float64Data, percent float64) (percentile float64, err error) {

	// Find the length of items in the slice
	il := input.Len()

	// Return an error for empty slices
	if il == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Return error for less than 0 percentages
	if percent <= 0 {
		return 0, errors.New("Percentage must be above 0")
	}

	// Return error for greater than 100 percentages
	if percent > 100 {
		return 0, errors.New("Percentage must not be above 100")
	}

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// Return the last item
	if percent == 100.0 {
		return c[il-1], nil
	}

	// Find ordinal ranking
	or := int(math.Ceil(float64(il) * percent / 100))

	// Return the item that is in the place of the ordinal rank
	return c[or-1], nil

}
