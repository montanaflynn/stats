package stats

import (
	"errors"
	"math"
)

// StandardDeviation the amount of variation in the dataset
func StandardDeviation(input Float64Data) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

// StandardDeviationPopulation finds the amount of variation from the population
func StandardDeviationPopulation(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the population variance
	vp, _ := PopulationVariance(input)

	// Return the population standard deviation
	return math.Pow(vp, 0.5), nil
}

// StandardDeviationSample finds the amount of variation from a sample
func StandardDeviationSample(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the sample variance
	vs, _ := SampleVariance(input)

	// Return the sample standard deviation
	return math.Pow(vs, 0.5), nil
}
