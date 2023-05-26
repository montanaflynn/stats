package stats

import (
	"math"
)

// MedianAbsoluteDeviation finds the median of the absolute deviations from the dataset median
func MedianAbsoluteDeviation(input Float64Data) (mad float64, err error) {
	return MedianAbsoluteDeviationPopulation(input)
}

// MedianAbsoluteDeviationPopulation finds the median of the absolute deviations from the population median
func MedianAbsoluteDeviationPopulation(input Float64Data) (mad float64, err error) {
	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}

	i := copyslice(input)
	m, _ := Median(i)

	for key, value := range i {
		i[key] = math.Abs(value - m)
	}

	return Median(i)
}

// StandardDeviation the amount of variation in the dataset
func StandardDeviation(input Float64Data) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

// StandardDeviationPopulation finds the amount of variation from the population
func StandardDeviationPopulation(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the population variance
	vp, _ := PopulationVariance(input)

	// Return the population standard deviation
	return math.Sqrt(vp), nil
}

// StandardDeviationSample finds the amount of variation from a sample
func StandardDeviationSample(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the sample variance
	vs, _ := SampleVariance(input)

	// Return the sample standard deviation
	return math.Sqrt(vs), nil
}

// StandardDeviationUnion finds the new Standard deviation of an array
// when you are appending new elemnts (without processing former elements)
func StandardDeviationUnion(input Float64Data, formerMean float64, formerLength int, formerSD float64) (sdev float64, err error) {

	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}
	inputMean, _ := Mean(input)
	UnionMean := ((float64(formerLength) * formerMean) + (float64(len(input)) * inputMean)) / float64((len(input) + formerLength))
	C := UnionMean / formerMean
	elem1 := formerSD * formerSD
	formerMeanPow2 := formerMean * formerMean
	elem2 := (C*C - 1) * formerMeanPow2
	elem3 := (1 - C) * 2 * formerMeanPow2
	elem4 := 0.0

	for _, n := range input {
		elem4 += (n - UnionMean) * (n - UnionMean)
	}
	elem4 /= float64(formerLength)

	UnionVariance := float64(formerLength) / float64(formerLength+len(input)) * (elem1 + elem2 + elem3 + elem4)
	// Return the sample standard deviation
	return math.Sqrt(UnionVariance), nil
}
