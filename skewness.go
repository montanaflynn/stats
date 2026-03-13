package stats

import (
	"errors"
	"math"
)

func Skewness(dataset Float64Data) (float64, error) {
	if len(dataset) < 2 {
		return 0.0, errors.New("skewness requires at least two data points")
	}
	mean, err := Mean(dataset)

	if err != nil {
		return 0.0, err
	}
	differences := Differences(dataset, mean)
	sumofcubes, err := Sum(PowerofN(differences, 3))
	if err != nil {
		return 0, err
	}

	sumofsquares, err := Sum(PowerofN(differences, 2))

	if err != nil {
		return 0, err
	}

	if sumofsquares == 0 {
		return 0, errors.New("skewness undefined for zero variance")
	}

	if sumofcubes == 0 {
		return 0.0, nil
	}

	variance := sumofsquares / float64(len(dataset))
	stdDevCubed := math.Pow(variance, 3.0/2.0)
	skewness := sumofcubes / stdDevCubed

	return skewness, nil
}

func Differences(dataset Float64Data, mean float64) Float64Data {
	diff := make([]float64, len(dataset))
	for i, a := range dataset {
		diff[i] = a - mean
	}
	return diff
}

func PowerofN(dataset Float64Data, power float64) []float64 {
	result := make([]float64, len(dataset))
	for i, v := range dataset {
		result[i] = math.Pow(v, power)
	}
	return result
}
