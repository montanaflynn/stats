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
func ComputeChebyshevDistance( data_point_x, data_point_y []float64 ) ( distance float64, 
                                                                        err error ) {
	var (
		temp_distance float64
		i             int
	)
	if len(data_point_x) == 0 || len(data_point_y) == 0 {
		return math.NaN(), EmptyInput
	}

	if len(data_point_x) != len(data_point_y) {
		return math.NaN(), SizeErr
	}
	distance = 0
	for i = 0; i < len(data_point_y); i++ {
		temp_distance = math.Abs(data_point_x[i] - data_point_y[i])
		if distance < temp_distance {
			distance = temp_distance
		}
	}
	return distance, nil
}
