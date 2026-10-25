package stats

import (
	"math"
)

// Validate data for distance calculation
func validateData(dataPointX, dataPointY Float64Data) error {
	if len(dataPointX) == 0 || len(dataPointY) == 0 {
		return EmptyInputErr
	}

	if len(dataPointX) != len(dataPointY) {
		return SizeErr
	}
	return nil
}

// ChebyshevDistance computes the Chebyshev distance between two data sets.
// If any coordinate difference is NaN the distance is NaN.
func ChebyshevDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {
	err = validateData(dataPointX, dataPointY)
	if err != nil {
		return math.NaN(), err
	}
	var tempDistance float64
	for i := 0; i < len(dataPointY); i++ {
		tempDistance = math.Abs(dataPointX[i] - dataPointY[i])
		// Keep a NaN difference, since NaN never compares greater
		if distance < tempDistance || math.IsNaN(tempDistance) {
			distance = tempDistance
		}
	}
	return distance, nil
}

// EuclideanDistance computes the Euclidean distance between two data sets
func EuclideanDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {

	err = validateData(dataPointX, dataPointY)
	if err != nil {
		return math.NaN(), err
	}
	distance = 0
	for i := 0; i < len(dataPointX); i++ {
		distance = distance + ((dataPointX[i] - dataPointY[i]) * (dataPointX[i] - dataPointY[i]))
	}
	return math.Sqrt(distance), nil
}

// ManhattanDistance computes the Manhattan distance between two data sets
func ManhattanDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {
	err = validateData(dataPointX, dataPointY)
	if err != nil {
		return math.NaN(), err
	}
	distance = 0
	for i := 0; i < len(dataPointX); i++ {
		distance = distance + math.Abs(dataPointX[i]-dataPointY[i])
	}
	return distance, nil
}

// MinkowskiDistance computes the Minkowski distance between two data sets
//
// Arguments:
//
//	dataPointX: First set of data points
//	dataPointY: Second set of data points. Length of both data
//	            sets must be equal.
//	lambda:     aka p or city blocks; With lambda = 1
//	            returned distance is manhattan distance and
//	            lambda = 2; it is euclidean distance. Lambda
//	            reaching to infinite - distance would be chebysev
//	            distance. A negative or NaN lambda returns
//	            ErrBounds.
//
// Return:
//
//	Distance or error
func MinkowskiDistance(dataPointX, dataPointY Float64Data, lambda float64) (distance float64, err error) {
	err = validateData(dataPointX, dataPointY)
	if err != nil {
		return math.NaN(), err
	}
	if math.IsNaN(lambda) || lambda < 0 {
		return math.NaN(), ErrBounds
	}
	for i := 0; i < len(dataPointY); i++ {
		distance = distance + math.Pow(math.Abs(dataPointX[i]-dataPointY[i]), lambda)
	}
	distance = math.Pow(distance, 1/lambda)

	// The powers overflow or underflow for differences far from unit
	// magnitude, so compute them again from scaled differences
	if (math.IsInf(distance, 1) || distance == 0) && lambda > 0 {
		distance = scaledMinkowskiDistance(dataPointX, dataPointY, lambda)
	}

	if math.IsInf(distance, 1) {
		return math.NaN(), InfValue
	}
	return distance, nil
}

// scaledMinkowskiDistance computes the Minkowski distance from differences
// scaled by powers of two, first so the differences cannot overflow and then
// so the largest is in [0.5, 1), and scales the result back. Scaling by a
// power of two is exact.
func scaledMinkowskiDistance(dataPointX, dataPointY Float64Data, lambda float64) float64 {
	e := unitExponent(append(append(Float64Data{}, dataPointX...), dataPointY...))
	f := math.Ldexp(1, -e)
	diffs := make(Float64Data, len(dataPointX))
	for i := range diffs {
		diffs[i] = math.Abs(dataPointX[i]*f - dataPointY[i]*f)
	}
	de := unitExponent(diffs)
	var sum float64
	for _, d := range diffs {
		sum += math.Pow(math.Ldexp(d, -de), lambda)
	}
	return math.Ldexp(math.Pow(sum, 1/lambda), e+de)
}
