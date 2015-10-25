package stats

import (
	"errors"
	"math"
)

// Correlation describes the degree of relationship between two sets of data
func Correlation(data1, data2 Float64Data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return 0, errors.New("Input data must not be empty")
	}

	if l1 != l2 {
		return 0, errors.New("Input data must be same length")
	}

	var sumX, sumY, sumCross float64

	meanX := data1.Get(0)
	meanY := data2.Get(0)

	for i := 1; i < l1; i++ {
		ratio := float64(i) / float64(i+1)
		deltaX := data1.Get(i) - meanX
		deltaY := data2.Get(i) - meanY
		sumX += deltaX * deltaX * ratio
		sumY += deltaY * deltaY * ratio
		sumCross += deltaX * deltaY * ratio
		meanX += deltaX / float64(i+1)
		meanY += deltaY / float64(i+1)
	}

	return sumCross / (math.Sqrt(sumX) * math.Sqrt(sumY)), nil
}

func AutoCorrelation(lags int, data Float64Data) (float64, error) {
	if len(data) < 1 {
		return 0, errors.New("Input data must not be empty")
	}

	mean, _ := Mean(data)

	var result, q float64

	for i := 0; i < lags; i++ {
		v := (data[0] - mean) * (data[0] - mean)
		for i := 1; i < len(data); i++ {
			delta0 := data[i-1] - mean
			delta1 := data[i] - mean
			q += (delta0*delta1 - q) / float64(i+1)
			v += (delta1*delta1 - v) / float64(i+1)
		}

		result = q / v
	}

	return result, nil
}
