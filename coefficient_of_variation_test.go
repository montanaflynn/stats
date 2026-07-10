package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestCoefficientOfVariation(t *testing.T) {
	// mean = 4, sample standard deviation = 2, so cv = 0.5
	r, err := stats.CoefficientOfVariation(stats.Float64Data{2, 4, 6})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-0.5) > 0.0000001 {
		t.Errorf("got %v, want ~0.5", r)
	}
}

func TestCoefficientOfVariationEmptyInput(t *testing.T) {
	_, err := stats.CoefficientOfVariation(stats.Float64Data{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestCoefficientOfVariationZeroMean(t *testing.T) {
	_, err := stats.CoefficientOfVariation(stats.Float64Data{-1, 0, 1})
	if err != stats.ErrZero {
		t.Errorf("expected ErrZero for zero mean, got %v", err)
	}
}

func TestFloat64DataCoefficientOfVariation(t *testing.T) {
	data := stats.Float64Data{2, 4, 6}

	r, err := data.CoefficientOfVariation()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-0.5) > 0.0000001 {
		t.Errorf("got %v, want ~0.5", r)
	}
}
