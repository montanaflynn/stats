package stats

import (
	"math"
	"sort"
)

// Find the lowest number in a slice
func Min(input []float64) (min float64) {
	if len(input) > 0 {
		min = input[0]
	}
	for i := 1; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

// Find the highest number in a slice
func Max(input []float64) (max float64) {
	if len(input) > 0 {
		max = input[0]
	}
	for i := 1; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

// Add all the numbers of a slice together
func Sum(input []float64) (sum float64) {
	for _, n := range input {
		sum += float64(n)
	}
	return sum
}

// Get the average of a slice of numbers
func Mean(input []float64) (mean float64) {

	if len(input) == 0 {
		return 0.0
	}

	sum := Sum(input)

	return sum / float64(len(input))
}

// Get the median number in a slice of numbers
func Median(input []float64) (median float64) {

	sort.Float64s(input)

	l := len(input)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	if l == 0 {
		return 0.0
	} else if l%2 == 0 {
		median = Mean(input[l/2-1 : l/2+1])
	} else {
		median = float64(input[l/2])
	}

	return median
}

// Get the mode of a slice of numbers
func Mode(input []float64) (mode []float64) {

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
	l := len(input)
	lm := len(mode)
	if l == lm {
		return []float64{}
	} else {
		return mode
	}
}

func Variance(input []float64, popfactor int) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}

	// Get the mean and then subtract that from each number
	// and then square the result
	m := Mean(input)
	for _, n := range input {
		varp += (float64(n) - m) * (float64(n) - m)
	}

	// Get the mean of the squared differences
	// "popfactor" will allow us to
	return varp / float64((len(input)-(1*popfactor)))
}

// Find the amount of variance within a population
func VarP(input []float64) (sdev float64) {
	// Get the mean of the squared differences
	return Variance(input,0)
}

// Find the amount of variance within the sample (when entire population is not available)
func Var(input []float64) (sdev float64) {
	// Get the mean of the squared differences
	return Variance(input,1)
}

// Find the amount of variation from the average also
// known as standard deviation or the greek symbol σ
func StdDevP(input []float64) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}
	//Get the Population Variance (VarP) and take the square root
	m := VarP(input,0)
	return math.Pow(m, 0.5)
}

func StdDevP(input []float64) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}
	//Get the Sample Variance (Var) and take the square root
	m := VarP(input,1)
	return math.Pow(m, 0.5)
}

// Round a float to a specific decimal place or precision
func Round(input float64, places int) (rounded float64) {

	// If the float is not a number we just return it
	// @todo add an error
	if math.IsNaN(input) {
		return input
	}

	// Find out the actual sign and correct the input for later
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	// Use the places arg to get the amount of precision wanted
	// by getting it's power of 10 to multiple by
	precision := math.Pow(10, float64(places))

	// Find the decimal place we are looking to round
	digit := input * precision

	// Get the actual decimal number as a fraction to be compared
	_, decimal := math.Modf(digit)

	// If the decimal is less than .5 we round down otherwise up
	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	// Finally we do the math to actually create a rounded number
	// by dividing the rounded number by the precision and then
	// multiply by the sign to keep the correct sign
	return rounded / precision * sign
}
