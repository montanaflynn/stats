package stats

import "math"

// Mean gets the average of a slice of numbers
func Mean(input Float64Data) (float64, error) {

	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}

	sum, _ := input.Sum()

	return sum / float64(input.Len()), nil
}

// GeometricMean gets the geometric mean for a slice of numbers
func GeometricMean(input Float64Data) (float64, error) {

	l := input.Len()
	if l == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Traditional definition is:
	// geoMean := (a_1 * a_2 * ... * a_n)^(1/n)
	// however, that overflows for larger numbers.
	// We can rewrite it to this if all numbers are greater than 0; otherwise see (1)
	// geoMean := e^(log((a_1 * a_2 * ... * a_n)^1/n)) =
	//  	= e^(1/n * log(a_1 * a_2 * ... * a_n)) =
	//  	= e^(1/n * (log(a_1) + log(a_2) + ... + log(a_n)))
	// Using this, we may get less precision, but we'll avoid overflow.
	//
	// Note (1):
	// - if there is 0 in input, result is 0 no matter what.
	// - if count of negative numbers is even, we can take absolute value
	//	 of all the numbers as negatives cancel each other.
	// - if count of negative numbers is odd:
	//	  -> if there are even data points, we have to return NaN as even root
	//		 of negative number doesn't exist in real numbers.
	//	  -> if there are odd data points, we can take absolute value
	//		 of all the numbers and add sign minus to the result.
	var p float64
	var negative int64
	for _, n := range input {
		switch {
		case n < 0:
			negative++
			p += math.Log(-n)
		case n == 0:
			return 0, nil
		case n > 0:
			p += math.Log(n)
		}
	}

	if negative % 2 == 0 {
		return math.Exp(p / float64(input.Len())), nil
	} else {
		if input.Len() % 2 == 0 {
			return math.NaN(), nil
		} else {
			return -math.Exp(p / float64(input.Len())), nil
		}
	}
}

// HarmonicMean gets the harmonic mean for a slice of numbers
func HarmonicMean(input Float64Data) (float64, error) {

	l := input.Len()
	if l == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the sum of all the numbers reciprocals and return an
	// error for values that cannot be included in harmonic mean
	var p float64
	for _, n := range input {
		if n < 0 {
			return math.NaN(), NegativeErr
		} else if n == 0 {
			return math.NaN(), ZeroErr
		}
		p += (1 / n)
	}

	return float64(l) / p, nil
}
