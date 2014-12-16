package stats

import (
	"math"
	"sort"
)

func Min(input []float64) (min float64) {
	sort.Float64s(input)
	return input[0]
}

func Max(input []float64) (max float64) {
	sort.Float64s(input)
	return input[len(input)-1]
}

func Sum(input []float64) (sum float64) {
	for _, n := range input {
		sum += float64(n)
	}
	return sum
}

func Mean(input []float64) (mean float64) {

	// Easy as pie, this one
	if len(input) == 0 {
		return 0.0
	}

	// Get the total sum
	sum := Sum(input)

	// Return the mean average
	return sum / float64(len(input))
}

func Median(input []float64) (median float64) {

	// Sort the numbers
	sort.Float64s(input)

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

func Mode(input []float64) []float64 {

	// Create a map to hold the counts
	m := make(map[float64]int)
	for _, v := range input {
		m[v]++
	}

	// Find the highest counts to return as a slice
	// of ints to accomodate duplicate counts
	var mode []float64
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
		return []float64{}
	} else {
		return mode
	}
}

func StandardDev(input []float64) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}

	m := Mean(input)
	for _, n := range input {
		sdev += (float64(n) - m) * (float64(n) - m)
	}

	sdev = math.Pow(sdev/float64(len(input)), 0.5)
	return sdev
}

func Round(input float64, places int) (rounded float64) {
	if math.IsNaN(input) || math.IsInf(input, 0) {
		return input
	}

	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	precision := math.Pow(10, float64(places))
	digit := input * precision
	_, decimal := math.Modf(digit)

	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	return rounded / precision * sign
}
