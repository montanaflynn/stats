package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestPercentileWeighted(t *testing.T) {
	// Example from the GitHub issue, matching Python statsmodels:
	// data = [1, 2, 9, 3.2, 4], weights = [0.0, 0.5, 1.0, 0.3, 0.5]
	// quantile(0.1) = 2, quantile(0.9) = 9
	data := stats.Float64Data{1, 2, 9, 3.2, 4}
	weights := stats.Float64Data{0.0, 0.5, 1.0, 0.3, 0.5}

	r, err := stats.PercentileWeighted(data, weights, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-2.0) > 0.01 {
		t.Errorf("10th percentile: got %v, want ~2.0", r)
	}

	r, err = stats.PercentileWeighted(data, weights, 90)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-9.0) > 0.01 {
		t.Errorf("90th percentile: got %v, want ~9.0", r)
	}
}

func TestPercentileWeightedUniform(t *testing.T) {
	// With uniform weights, should match regular Percentile behavior closely
	data := stats.Float64Data{1, 2, 3, 4, 5}
	weights := stats.Float64Data{1, 1, 1, 1, 1}

	r, err := stats.PercentileWeighted(data, weights, 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-3.0) > 0.01 {
		t.Errorf("50th percentile: got %v, want 3.0", r)
	}
}

func TestPercentileWeightedSingleElement(t *testing.T) {
	data := stats.Float64Data{42.0}
	weights := stats.Float64Data{1.0}

	r, err := stats.PercentileWeighted(data, weights, 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 42.0 {
		t.Errorf("got %v, want 42.0", r)
	}
}

func TestPercentileWeightedHeavyTail(t *testing.T) {
	// Heavy weight on last element should pull high percentiles toward it
	data := stats.Float64Data{1, 2, 3, 4, 100}
	weights := stats.Float64Data{1, 1, 1, 1, 10}

	r, err := stats.PercentileWeighted(data, weights, 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// With heavy weight on 100, the 50th percentile should be > simple median (3)
	if r <= 3 {
		t.Errorf("50th percentile with heavy tail: got %v, want > 3", r)
	}
}

func TestPercentileWeightedEmptyInput(t *testing.T) {
	_, err := stats.PercentileWeighted(stats.Float64Data{}, stats.Float64Data{}, 50)
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestPercentileWeightedSizeMismatch(t *testing.T) {
	_, err := stats.PercentileWeighted(stats.Float64Data{1, 2}, stats.Float64Data{1}, 50)
	if err != stats.ErrSize {
		t.Errorf("expected ErrSize, got %v", err)
	}
}

func TestPercentileWeightedInvalidPercent(t *testing.T) {
	data := stats.Float64Data{1, 2, 3}
	weights := stats.Float64Data{1, 1, 1}

	_, err := stats.PercentileWeighted(data, weights, 0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for percent=0, got %v", err)
	}

	_, err = stats.PercentileWeighted(data, weights, 101)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for percent=101, got %v", err)
	}

	_, err = stats.PercentileWeighted(data, weights, -5)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for percent=-5, got %v", err)
	}
}

func TestPercentileWeightedNegativeWeight(t *testing.T) {
	_, err := stats.PercentileWeighted(stats.Float64Data{1, 2, 3}, stats.Float64Data{1, -1, 1}, 50)
	if err != stats.ErrNegative {
		t.Errorf("expected ErrNegative, got %v", err)
	}
}

func TestPercentileWeighted100(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 5}
	weights := stats.Float64Data{1, 1, 1, 1, 1}

	r, err := stats.PercentileWeighted(data, weights, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 5.0 {
		t.Errorf("100th percentile: got %v, want 5.0", r)
	}
}

func TestPercentileWeightedZeroWeights(t *testing.T) {
	_, err := stats.PercentileWeighted(stats.Float64Data{1, 2, 3}, stats.Float64Data{0, 0, 0}, 50)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for all-zero weights, got %v", err)
	}
}
