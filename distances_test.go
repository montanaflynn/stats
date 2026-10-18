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

		_, err = stats.MinkowskiDistance([]float64{999, 999, 999}, []float64{1, 1, 1}, 1000)
		if err == nil {
			t.Errorf("Infinite distance should have resulted in an error")
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

func TestEuclideanDistanceScale(t *testing.T) {
	for _, scale := range []float64{1, 1e200, 1e-200} {
		t.Run(fmt.Sprintf("scale=%g", scale), func(t *testing.T) {
			for _, sign := range []float64{-1, 1} {
				got, err := stats.EuclideanDistance(
					stats.Float64Data{sign * 3 * scale, sign * 4 * scale},
					stats.Float64Data{0, 0},
				)
				if err != nil {
					t.Fatal(err)
				}
				want := 5 * scale
				if math.IsNaN(got) || math.IsInf(got, 0) || math.Abs(got/want-1) > 1e-15 {
					t.Errorf("distance = %g, want %g", got, want)
				}
			}
		})
	}
}

func TestEuclideanDistanceNonFinite(t *testing.T) {
	for _, point := range []stats.Float64Data{
		{math.NaN(), math.Inf(1)},
		{math.Inf(1), math.NaN()},
	} {
		got, err := stats.EuclideanDistance(point, stats.Float64Data{0, 0})
		if err != nil || !math.IsNaN(got) {
			t.Errorf("distance = (%g, %v), want (NaN, nil)", got, err)
		}
	}
	got, err := stats.EuclideanDistance(stats.Float64Data{math.Inf(1)}, stats.Float64Data{0})
	if err != nil || !math.IsInf(got, 1) {
		t.Errorf("distance = (%g, %v), want (+Inf, nil)", got, err)
	}
}

func ExampleChebyshevDistance() {
	d1 := []float64{2, 3, 4, 5, 6, 7, 8}
	d2 := []float64{8, 7, 6, 5, 4, 3, 2}
	cd, _ := stats.ChebyshevDistance(d1, d2)
	fmt.Println(cd)
	// Output: 6
}
