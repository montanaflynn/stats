package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestWeightedMean(t *testing.T) {
	// (1*3 + 2*1 + 3*1) / (3 + 1 + 1) = 8 / 5 = 1.6
	data := stats.Float64Data{1, 2, 3}
	weights := stats.Float64Data{3, 1, 1}

	r, err := stats.WeightedMean(data, weights)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 1.6 {
		t.Errorf("got %v, want 1.6", r)
	}
}

func TestWeightedMeanOverflow(t *testing.T) {
	m := math.MaxFloat64
	for _, c := range []struct {
		data, weights stats.Float64Data
		out           float64
	}{
		// The products overflow, but the mean is the value itself
		{stats.Float64Data{m, m, m, m}, stats.Float64Data{1, 2, 3, 4}, m},
		{stats.Float64Data{-m, -m}, stats.Float64Data{1, 3}, -m},
		// The total weight overflows
		{stats.Float64Data{1, 3}, stats.Float64Data{m, m}, 2},
		{stats.Float64Data{1e308, -1e308, 1e308, 5}, stats.Float64Data{1, 2, 3, 4}, 2e307},
		{stats.Float64Data{m, math.Inf(1)}, stats.Float64Data{1, 1}, math.Inf(1)},
	} {
		if got, err := stats.WeightedMean(c.data, c.weights); err != nil || !close(got, c.out) {
			t.Errorf("WeightedMean(%v, %v) => %v, %v; want %v", c.data, c.weights, got, err, c.out)
		}
	}
	if got, _ := stats.WeightedMean(stats.Float64Data{m, math.NaN()}, stats.Float64Data{1, 1}); !math.IsNaN(got) {
		t.Errorf("WeightedMean with NaN => %v; want NaN", got)
	}
}

func TestWeightedMeanUniformWeights(t *testing.T) {
	// With uniform weights the weighted mean should equal the mean
	data := stats.Float64Data{1.5, 3.6, 7.9, 4.2}
	weights := stats.Float64Data{2, 2, 2, 2}

	r, err := stats.WeightedMean(data, weights)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	m, err := stats.Mean(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if math.Abs(r-m) > 0.0000001 {
		t.Errorf("got %v, want %v", r, m)
	}
}

func TestWeightedMeanEmptyInput(t *testing.T) {
	_, err := stats.WeightedMean(stats.Float64Data{}, stats.Float64Data{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestWeightedMeanSizeMismatch(t *testing.T) {
	_, err := stats.WeightedMean(stats.Float64Data{1, 2}, stats.Float64Data{1})
	if err != stats.ErrSize {
		t.Errorf("expected ErrSize, got %v", err)
	}
}

func TestWeightedMeanNegativeWeight(t *testing.T) {
	_, err := stats.WeightedMean(stats.Float64Data{1, 2, 3}, stats.Float64Data{1, -1, 1})
	if err != stats.ErrNegative {
		t.Errorf("expected ErrNegative, got %v", err)
	}
}

func TestWeightedMeanZeroWeights(t *testing.T) {
	_, err := stats.WeightedMean(stats.Float64Data{1, 2, 3}, stats.Float64Data{0, 0, 0})
	if err != stats.ErrZero {
		t.Errorf("expected ErrZero for all-zero weights, got %v", err)
	}
}

func TestFloat64DataWeightedMean(t *testing.T) {
	data := stats.Float64Data{1, 2, 3}
	weights := stats.Float64Data{3, 1, 1}

	r, err := data.WeightedMean(weights)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 1.6 {
		t.Errorf("got %v, want 1.6", r)
	}
}
