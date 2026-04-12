package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestZTestOneSample(t *testing.T) {
	// Test data: sample with known population std dev
	data := stats.Float64Data{2.5, 3.1, 2.8, 3.0, 2.9, 3.2, 2.7, 3.3, 2.6, 3.1}

	z, p, err := stats.ZTest(data, nil, 2.5, 0.3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Expected: mean=2.92, se=0.3/sqrt(10)=0.09487, z=(2.92-2.5)/0.09487 ≈ 4.427
	if math.Abs(z-4.427) > 0.01 {
		t.Errorf("z statistic: got %v, want ~4.427", z)
	}

	if p > 0.001 {
		t.Errorf("p-value should be very small, got %v", p)
	}
}

func TestZTestTwoSample(t *testing.T) {
	data1 := stats.Float64Data{5.1, 5.5, 4.8, 5.2, 5.0}
	data2 := stats.Float64Data{4.2, 4.8, 4.0, 4.5, 4.3}

	z, p, err := stats.ZTest(data1, data2, 0, 0.5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// data1 mean=5.12, data2 mean=4.36, diff=0.76
	// se = 0.5 * sqrt(1/5 + 1/5) = 0.5 * 0.6325 = 0.3162
	// z = 0.76 / 0.3162 ≈ 2.404
	if math.Abs(z-2.404) > 0.01 {
		t.Errorf("z statistic: got %v, want ~2.404", z)
	}

	if p > 0.05 || p < 0.001 {
		t.Errorf("p-value should be between 0.001 and 0.05, got %v", p)
	}
}

func TestZTestEmptyInput(t *testing.T) {
	_, _, err := stats.ZTest(stats.Float64Data{}, nil, 0, 1)
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestZTestInvalidStdDev(t *testing.T) {
	data := stats.Float64Data{1.0, 2.0, 3.0}

	_, _, err := stats.ZTest(data, nil, 0, 0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for zero std dev, got %v", err)
	}

	_, _, err = stats.ZTest(data, nil, 0, -1)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for negative std dev, got %v", err)
	}
}

func TestZTestTwoSampleInvalidStdDev(t *testing.T) {
	data1 := stats.Float64Data{1.0, 2.0, 3.0}
	data2 := stats.Float64Data{4.0, 5.0, 6.0}

	_, _, err := stats.ZTest(data1, data2, 0, 0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for zero std dev, got %v", err)
	}
}
