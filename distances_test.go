package stats_test

import (
	"fmt"
	"math"
	"math/big"
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

// euclideanTestData returns n-element data sets: "sincos" has ordinary values
// and "onebig" has one large difference next to many tiny ones, which a plain
// running sum of squares loses.
func euclideanTestData(kind string, n int, scale float64) (x, y stats.Float64Data) {
	x, y = make(stats.Float64Data, n), make(stats.Float64Data, n)
	for i := range x {
		switch kind {
		case "sincos":
			x[i], y[i] = math.Sin(float64(i+1)), math.Cos(float64(i+1))
		case "onebig":
			x[i] = 1e-8 * math.Sin(float64(i+1))
		}
		x[i], y[i] = x[i]*scale, y[i]*scale
	}
	if kind == "onebig" {
		x[0] = scale
	}
	return x, y
}

func TestEuclideanDistanceAccuracy(t *testing.T) {
	for _, kind := range []string{"sincos", "onebig"} {
		for _, scale := range []float64{1, 1e200, 1e-200} {
			t.Run(fmt.Sprintf("%s/scale=%g", kind, scale), func(t *testing.T) {
				x, y := euclideanTestData(kind, 1003, scale)
				sum := new(big.Float).SetPrec(300)
				for i := range x {
					delta := new(big.Float).SetPrec(300).SetFloat64(x[i])
					delta.Sub(delta, new(big.Float).SetFloat64(y[i]))
					sum.Add(sum, delta.Mul(delta, delta))
				}
				want, _ := sum.Sqrt(sum).Float64()
				got, err := stats.EuclideanDistance(x, y)
				if err != nil {
					t.Fatal(err)
				}
				if math.Abs(got/want-1) > 1e-15 {
					t.Errorf("distance = %g, want %g", got, want)
				}
			})
		}
	}
}

func TestEuclideanDistanceExtremes(t *testing.T) {
	tiny := math.SmallestNonzeroFloat64
	for _, test := range []struct {
		x, y stats.Float64Data
		want float64
	}{
		{stats.Float64Data{1, 2, 3}, stats.Float64Data{1, 2, 3}, 0},
		{stats.Float64Data{-7.25}, stats.Float64Data{0}, 7.25},
		{stats.Float64Data{math.MaxFloat64}, stats.Float64Data{0}, math.MaxFloat64},
		{stats.Float64Data{math.MaxFloat64, math.MaxFloat64}, stats.Float64Data{0, 0}, math.Inf(1)},
		{stats.Float64Data{1e308}, stats.Float64Data{-1e308}, math.Inf(1)},
		{stats.Float64Data{tiny}, stats.Float64Data{0}, tiny},
		{stats.Float64Data{3 * tiny, 4 * tiny}, stats.Float64Data{0, 0}, 5 * tiny},
		{stats.Float64Data{1e300, 1}, stats.Float64Data{0, 0}, 1e300},
		{stats.Float64Data{1e-300, 1}, stats.Float64Data{0, 0}, 1},
		{stats.Float64Data{1, 2}, stats.Float64Data{math.Inf(-1), 0}, math.Inf(1)},
	} {
		got, err := stats.EuclideanDistance(test.x, test.y)
		if err != nil || got != test.want || math.Signbit(got) {
			t.Errorf("EuclideanDistance(%v, %v) = (%g, %v), want (%g, nil)", test.x, test.y, got, err, test.want)
		}
	}
}

func BenchmarkEuclideanDistance(b *testing.B) {
	for _, n := range []int{3, 10, 1000, 1000000} {
		x, y := euclideanTestData("sincos", n, 1)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = stats.EuclideanDistance(x, y)
			}
		})
	}
}

func TestEuclideanDistanceNonFinite(t *testing.T) {
	for _, point := range []stats.Float64Data{
		{math.NaN(), math.Inf(1)},
		{math.Inf(1), math.NaN()},
		{1e300, math.NaN()},
		{1e-300, math.NaN()},
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
