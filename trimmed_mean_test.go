package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleTrimmedMean() {
	tm, _ := stats.TrimmedMean([]float64{1, 2, 3, 4, 100}, 0.2)
	fmt.Println(tm)
	// Output: 3
}

func TestTrimmedMean(t *testing.T) {
	for _, c := range []struct {
		in      []float64
		percent float64
		out     float64
	}{
		{[]float64{1, 2, 3, 4, 100}, 0.2, 3.0},
		// Verified against scipy.stats.trim_mean
		{[]float64{4.2, 1.1, 9.9, 3.3, 2.7, 8.8, 5.5, 7.4, 6.6, 100.5}, 0.2, 5.966666666666666},
		{[]float64{1, 2, 3, 4, 5}, 0.1, 3.0},
		{[]float64{1, 2, 3, 4, 5, 6}, 0.25, 3.5},
		{[]float64{5}, 0.4, 5.0},
	} {
		got, err := stats.TrimmedMean(c.in, c.percent)
		if err != nil {
			t.Errorf("TrimmedMean(%v, %v) returned an error: %v", c.in, c.percent, err)
		}
		if math.Abs(got-c.out) > 1e-12 {
			t.Errorf("TrimmedMean(%v, %v) => %v != %v", c.in, c.percent, got, c.out)
		}
	}
}

func TestTrimmedMeanZeroPercentEqualsMean(t *testing.T) {
	data := []float64{-2.5, 7.5, 1.25, 4.75}
	tm, err := stats.TrimmedMean(data, 0)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	mean, _ := stats.Mean(data)
	if tm != mean {
		t.Errorf("TrimmedMean(%v, 0) => %v != Mean %v", data, tm, mean)
	}
}

func TestTrimmedMeanInvalidInput(t *testing.T) {
	tm, err := stats.TrimmedMean([]float64{}, 0.1)
	if err != stats.ErrEmptyInput {
		t.Errorf("Expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(tm) {
		t.Errorf("Expected NaN, got %v", tm)
	}

	for _, percent := range []float64{-0.1, 0.5, 1.0, math.NaN()} {
		tm, err = stats.TrimmedMean([]float64{1, 2, 3}, percent)
		if err != stats.ErrBounds {
			t.Errorf("TrimmedMean percent %v expected ErrBounds, got %v", percent, err)
		}
		if !math.IsNaN(tm) {
			t.Errorf("TrimmedMean percent %v expected NaN, got %v", percent, tm)
		}
	}
}

func TestTrimmedMeanDoesNotMutateInput(t *testing.T) {
	data := []float64{100, 1, 4, 2, 3}
	_, err := stats.TrimmedMean(data, 0.2)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	want := []float64{100, 1, 4, 2, 3}
	for i := range want {
		if data[i] != want[i] {
			t.Errorf("Input was mutated: %v != %v", data, want)
			break
		}
	}
}

func TestFloat64DataTrimmedMean(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 100}
	tm, err := data.TrimmedMean(0.2)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	if tm != 3.0 {
		t.Errorf("TrimmedMean(0.2) => %v != 3", tm)
	}
}
