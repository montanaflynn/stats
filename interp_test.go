package stats_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleInterp() {
	out, _ := stats.Interp([]float64{1.5}, []float64{1, 2}, []float64{10, 20})
	fmt.Println(out)
	// Output: [15]
}

func TestInterp(t *testing.T) {
	for _, c := range []struct {
		x   []float64
		xp  []float64
		fp  []float64
		out []float64
	}{
		{[]float64{1.5}, []float64{1, 2}, []float64{10, 20}, []float64{15}},
		{[]float64{0, 1, 1.5, 2.5, 4}, []float64{1, 2, 3}, []float64{10, 20, 40}, []float64{10, 10, 15, 30, 40}},
		{[]float64{2.5, 4}, []float64{1, 2, 3, 5}, []float64{3, 2, 0, 4}, []float64{1, 2}},
		{[]float64{-1, 0.5, 2}, []float64{1}, []float64{5}, []float64{5, 5, 5}},
		{[]float64{-10}, []float64{0, 1}, []float64{2, 4}, []float64{2}},
		{[]float64{10}, []float64{0, 1}, []float64{2, 4}, []float64{4}},
		// Exact knot hits return fp[j] exactly even when the naive
		// arithmetic would suffer catastrophic cancellation
		{[]float64{2}, []float64{1, 2, 3}, []float64{1e16, 1, 1}, []float64{1}},
		// Knot spacing that overflows float64
		{[]float64{0}, []float64{-math.MaxFloat64, math.MaxFloat64}, []float64{0, 100}, []float64{50}},
		{[]float64{math.MaxFloat64 / 2}, []float64{-math.MaxFloat64, math.MaxFloat64}, []float64{0, 100}, []float64{75}},
		// fp spread that overflows float64
		{[]float64{0.5}, []float64{0, 1}, []float64{-math.MaxFloat64, math.MaxFloat64}, []float64{0}},
	} {
		got, err := stats.Interp(c.x, c.xp, c.fp)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Interp(%.1f, %.1f, %.1f) => %.1f != %.1f", c.x, c.xp, c.fp, got, c.out)
		}
	}
}

func TestInterpInvalidInput(t *testing.T) {
	for _, c := range []struct {
		x   []float64
		xp  []float64
		fp  []float64
		err error
	}{
		{[]float64{}, []float64{1, 2}, []float64{10, 20}, stats.ErrEmptyInput},
		{[]float64{1.5}, []float64{}, []float64{}, stats.ErrEmptyInput},
		{[]float64{1.5}, []float64{1, 2}, []float64{10}, stats.ErrSize},
		{[]float64{1.5}, []float64{2, 1}, []float64{10, 20}, stats.ErrBounds},
		{[]float64{1.5}, []float64{1, 1}, []float64{10, 20}, stats.ErrBounds},
	} {
		_, err := stats.Interp(c.x, c.xp, c.fp)
		if err != c.err {
			t.Errorf("Interp(%.1f, %.1f, %.1f) => %v != %v", c.x, c.xp, c.fp, err, c.err)
		}
	}
}
