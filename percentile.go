package stats

import (
	"math"
)

// Percentile finds the relative standing in a slice of floats
func Percentile(input Float64Data, percent float64) (percentile float64, err error) {
	length := input.Len()
	if length == 0 {
		return math.NaN(), EmptyInputErr
	}

	if percent < 0 || percent > 100 {
		return math.NaN(), BoundsErr
	}

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// Calculate rank
	rank := (percent / 100) * float64(len(c)-1)

	// Convert float to int
	ri := int(rank)

	// Check if the index is a whole number
	if rank == float64(ri) {

		// Find the value at the index
		percentile = c[ri]

	} else {

		// Calculate the fractional part of the rank
		rf := rank - float64(ri)

		// Interpolate
		percentile = c[ri] + rf*(c[ri+1]-c[ri])

	}

	return percentile, nil

}

// PercentileNearestRank finds the relative standing in a slice of floats using the Nearest Rank method
func PercentileNearestRank(input Float64Data, percent float64) (percentile float64, err error) {

	// Find the length of items in the slice
	il := input.Len()

	// Return an error for empty slices
	if il == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Return error for less than 0 or greater than 100 percentages
	if percent < 0 || percent > 100 {
		return math.NaN(), BoundsErr
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
	if or == 0 {
		return c[0], nil
	}
	return c[or-1], nil

}
