package stats

import (
	"testing"
)

func TestDataSetDistances(t *testing.T) {
	var err error
	dataPointX := []float64{2, 3, 4, 5, 6, 7, 8}
	dataPointY := []float64{8, 7, 6, 5, 4, 3, 2}
	_, err = ChebyshevDistance(dataPointX, dataPointY)
	if err != nil {
		t.Errorf("Failed to compute Chebyshev disatance.")
	}
	_, err = ManhattanDistance(dataPointX, dataPointY)
	if err != nil {
		t.Errorf("Failed to compute Manhattan disatance.")
	}
	_, err = EuclideanDistance(dataPointX, dataPointY)
	if err != nil {
		t.Errorf("Failed to compute Euclidean disatance.")
	}
	_, err = MinkowskiDistance(dataPointX, dataPointY, float64(1))
	if err != nil {
		t.Errorf("Failed to compute Minkowski disatance with lambda 1.")
	}
	_, err = MinkowskiDistance(dataPointX, dataPointY, float64(2))
	if err != nil {
		t.Errorf("Failed to compute Minkowski disatance with lambda 2.")
	}
	_, err = MinkowskiDistance(dataPointX, dataPointY, float64(300))
	if err != nil {
		t.Errorf("Failed to compute Minkowski disatance with large value of lambda.")
	}
}
