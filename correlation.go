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
