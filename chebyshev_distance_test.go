package stats

import (
	"testing"
)

func TestChebyshevDistance(t *testing.T) {
	var err error
	dataPointX := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	dataPointY := []float64{2, 4, 4, 4, 5, 5, 7, 1}
	_, err = ComputeChebyshevDistance(dataPointX, dataPointY)
	if err != nil {
		t.Errorf("Failed to compute chebyshev disatance.")
	}
}
