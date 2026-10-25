package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

type distanceFunctionType func(stats.Float64Data, stats.Float64Data) (float64, error)

var minkowskiDistanceTestMatrix = []struct {
	dataPointX []float64
	dataPointY []float64
	lambda     float64
	distance   float64
}{
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 1, 24},
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 2, 10.583005244258363},
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 99, 6},
}

var distanceTestMatrix = []struct {
	dataPointX       []float64
	dataPointY       []float64
	distance         float64
	distanceFunction distanceFunctionType
}{
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 6, stats.ChebyshevDistance},
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 24, stats.ManhattanDistance},
	{[]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, 10.583005244258363, stats.EuclideanDistance},
}

func TestMinkowskiDistanceInvalidLambda(t *testing.T) {
	for _, lambda := range []float64{-1, -0.5, math.NaN(), math.Inf(-1)} {
		d, err := stats.MinkowskiDistance([]float64{1, 2, 3}, []float64{4, 5, 6}, lambda)
		if err != stats.ErrBounds {
			t.Errorf("MinkowskiDistance(lambda=%v) returned %v, want ErrBounds", lambda, err)
		}
		if !math.IsNaN(d) {
			t.Errorf("MinkowskiDistance(lambda=%v) = %v, want NaN", lambda, d)
		}
	}
	// lambda = 0 keeps its existing ErrInfValue.
	if _, err := stats.MinkowskiDistance([]float64{1, 2, 3}, []float64{4, 5, 6}, 0); err != stats.ErrInfValue {
		t.Errorf("MinkowskiDistance(lambda=0) returned %v, want ErrInfValue", err)
	}
}

func TestDataSetDistances(t *testing.T) {

	// Test Minkowski Distance with different lambda values.
	for _, testData := range minkowskiDistanceTestMatrix {
		distance, err := stats.MinkowskiDistance(testData.dataPointX, testData.dataPointY, testData.lambda)
		if err != nil && distance != testData.distance {
			t.Errorf("Failed to compute Minkowski distance.")
		}

		_, err = stats.MinkowskiDistance([]float64{}, []float64{}, 3)
		if err == nil {
			t.Errorf("Empty slices should have resulted in an error")
		}

		_, err = stats.MinkowskiDistance([]float64{1, 2, 3}, []float64{1, 4}, 3)
		if err == nil {
			t.Errorf("Different length slices should have resulted in an error")
		}

		_, err = stats.MinkowskiDistance([]float64{math.Inf(1), 999, 999}, []float64{1, 1, 1}, 3)
		if err == nil {
			t.Errorf("Infinite distance should have resulted in an error")
		}

		// 998^1000 overflows, but the distance 998 * 3^(1/1000) does not
		d, err := stats.MinkowskiDistance([]float64{999, 999, 999}, []float64{1, 1, 1}, 1000)
		if err != nil || !close(d, 999.0970175522357) {
			t.Errorf("MinkowskiDistance with overflowing powers = %v, %v; want 999.0970175522357", d, err)
		}
	}

	// Compute distance with the help of all algorithms.
	for _, testSet := range distanceTestMatrix {
		distance, err := testSet.distanceFunction(testSet.dataPointX, testSet.dataPointY)
		if err != nil && testSet.distance != distance {
			t.Errorf("Failed to compute distance.")
		}

		_, err = testSet.distanceFunction([]float64{}, []float64{})
		if err == nil {
			t.Errorf("Empty slices should have resulted in an error")
		}
	}
}

func ExampleChebyshevDistance() {
	d1 := []float64{2, 3, 4, 5, 6, 7, 8}
	d2 := []float64{8, 7, 6, 5, 4, 3, 2}
	cd, _ := stats.ChebyshevDistance(d1, d2)
	fmt.Println(cd)
	// Output: 6
}

func TestMinkowskiDistanceExtremeScales(t *testing.T) {
	want := math.Cbrt(8 + 27)
	for _, s := range []float64{1e200, 1e-200} {
		x, y := stats.Float64Data{1 * s, 2 * s}, stats.Float64Data{3 * s, 5 * s}
		if d, err := stats.MinkowskiDistance(x, y, 3); err != nil || !close(d, want*s) {
			t.Errorf("MinkowskiDistance(%v, %v, 3) = %v, %v; want %v", x, y, d, err, want*s)
		}
	}
	// Differences of opposite-signed coordinates overflow before scaling
	m := math.MaxFloat64
	if d, err := stats.MinkowskiDistance(stats.Float64Data{m}, stats.Float64Data{-m}, 2); err != stats.ErrInfValue {
		t.Errorf("MinkowskiDistance(MaxFloat64, -MaxFloat64) = %v, %v; want ErrInfValue", d, err)
	}
	if d, err := stats.MinkowskiDistance(stats.Float64Data{1, 2}, stats.Float64Data{1, 2}, 3); err != nil || d != 0 {
		t.Errorf("MinkowskiDistance of equal points = %v, %v; want 0", d, err)
	}
}
