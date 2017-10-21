//
// Compiled with Go version go1.9 linux/amd64
// Contributor : Shivendra Mishra
// Contact : shivendra_mishra@live.com
//
package stats

import (
	"math"
)

// Compute Chebyshev distance between two data points
func ComputeChebyshevDistance(dataPointX, dataPointY []float64) (distance float64, err error) {
	if len(dataPointX) == 0 || len(dataPointY) == 0 {
		return math.NaN(), EmptyInput
	}

	if len(dataPointX) != len(dataPointY) {
		return math.NaN(), SizeErr
	}

	var tempDistance float64
	for i := 0; i < len(dataPointY); i++ {
		tempDistance = math.Abs(dataPointX[i] - dataPointY[i])
		if distance < tempDistance {
			distance = tempDistance
		}
	}
	return distance, nil
}
