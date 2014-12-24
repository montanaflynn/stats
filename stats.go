package stats

import (
	"math"
	"sort"
)

// Min finds the lowest number in a slice
func Min(input []float64) (min float64) {

	// Get the initial value
	// @todo add error if no length
	if len(input) > 0 {
		min = input[0]
	}

	// Iterate until done checking for a lower value
	for i := 1; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

// Max finds the highest number in a slice
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

// Sum adds all the numbers of a slice together
func Sum(input []float64) (sum float64) {
	for _, n := range input {
		sum += float64(n)
	}
	return sum
}

// Mean gets the average of a slice of numbers
func Mean(input []float64) (mean float64) {
	if len(input) == 0 {
		return 0.0
	}

	sum := Sum(input)

	return sum / float64(len(input))
}

// Median gets the median number in a slice of numbers
func Median(input []float64) (median float64) {

	// Start by sorting the slice
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

// Mode gets the mode of a slice of numbers
func Mode(input []float64) (mode []float64) {

	// Return the input if there's only one number
	// @todo add error for empty slice
	l := len(input)
	if l == 1 {
		return input
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
		return []float64{}
	}
	return mode
}

// Variance finds the variance for both population and sample data
func Variance(input []float64, sample int) (variance float64) {
	if len(input) == 0 {
		return 0.0
	}

	// Sum the square of the mean subtracted from each number
	m := Mean(input)
	for _, n := range input {
		variance += (float64(n) - m) * (float64(n) - m)
	}

	// When getting the mean of the squared differences
	// "sample" will allow us to know if it's a sample
	// or population and wether to subtract by one or not
	return variance / float64((len(input) - (1 * sample)))
}

// VarP finds the amount of variance within a population
func VarP(input []float64) (sdev float64) {
	return Variance(input, 0)
}

// VarS finds the amount of variance within a sample
func VarS(input []float64) (sdev float64) {
	return Variance(input, 1)
}

// StdDevP finds the amount of variation from the population
func StdDevP(input []float64) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}

	// Get the population variance
	m := VarP(input)

	// Return the population standard deviation
	return math.Pow(m, 0.5)
}

// StdDevS finds the amount of variation from a sample
func StdDevS(input []float64) (sdev float64) {
	if len(input) == 0 {
		return 0.0
	}

	// Get the sample variance
	m := VarS(input)

	// Return the sample standard deviation
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
	return rounded / precision * sign
}

// Percentile finds the relative standing in a slice of floats
func Percentile(input []float64, percent float64) (percentile float64) {

	// Sort the data
	sort.Float64s(input)

	// Multiple percent by length of input
	index := (percent / 100) * float64(len(input))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := Float64ToInt(index)

		// Find the average of the index and following values
		percentile = Mean([]float64{input[i-1], input[i]})

	} else {

		// Convert float to int
		i := Float64ToInt(index)

		// Find the value at the index
		percentile = input[i-1]

	}

	return percentile

}

// Float64ToInt rounds a float64 to an int
func Float64ToInt(input float64) (output int) {

	// Round input to nearest whole number and convert to int
	return int(Round(input, 0))

}

// Coordinate holds the data in a series
type Coordinate struct {
	X, Y float64
}

// LinReg finds the least squares linear regression on data series
func LinReg(s []Coordinate) (regressions []Coordinate) {

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)

	// Create the new regression series
	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: s[j].X*gradient + intercept,
		})
	}

	return

}

// ExpReg returns an exponential regression on data series
func ExpReg(s []Coordinate) (regressions []Coordinate) {

	var sum [6]float64

	for i := 0; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X * s[i].Y
		sum[3] += s[i].Y * math.Log(s[i].Y)
		sum[4] += s[i].X * s[i].Y * math.Log(s[i].Y)
		sum[5] += s[i].X * s[i].Y
	}

	denominator := (sum[1]*sum[2] - sum[5]*sum[5])
	a := math.Pow(math.E, (sum[2]*sum[3]-sum[5]*sum[4])/denominator)
	b := (sum[1]*sum[4] - sum[5]*sum[3]) / denominator

	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: a * math.Pow(2.718281828459045, b*s[j].X),
		})
	}

	return

}

// LogReg returns an logarithmic regression on data series
func LogReg(s []Coordinate) (regressions []Coordinate) {

	var sum [4]float64

	i := 0
	for ; i < len(s); i++ {
		sum[0] += math.Log(s[i].X)
		sum[1] += s[i].Y * math.Log(s[i].X)
		sum[2] += s[i].Y
		sum[3] += math.Pow(math.Log(s[i].X), 2)
	}

	f := float64(i)
	a := (f*sum[1] - sum[2]*sum[0]) / (f*sum[3] - sum[0]*sum[0])
	b := (sum[2] - a*sum[0]) / f

	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: b + a*math.Log(s[j].X),
		})
	}

	return

}
