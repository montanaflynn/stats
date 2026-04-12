package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestTTestOneSample(t *testing.T) {
	// Test against known values from scipy.stats.ttest_1samp
	data := stats.Float64Data{2.5, 3.1, 2.8, 3.0, 2.9, 3.2, 2.7, 3.3, 2.6, 3.1}

	tstat, p, err := stats.TTest(data, nil, 2.5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// mean=2.92, sample sd≈0.2658, se=0.2658/sqrt(10)=0.08406
	// t=(2.92-2.5)/0.08406 ≈ 4.997
	if math.Abs(tstat-4.997) > 0.01 {
		t.Errorf("t statistic: got %v, want ~4.997", tstat)
	}

	if p > 0.001 {
		t.Errorf("p-value should be very small, got %v", p)
	}
}

func TestTTestTwoSample(t *testing.T) {
	data1 := stats.Float64Data{5.1, 5.5, 4.8, 5.2, 5.0}
	data2 := stats.Float64Data{4.2, 4.8, 4.0, 4.5, 4.3}

	tstat, p, err := stats.TTest(data1, data2, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Should detect significant difference between groups
	if math.Abs(tstat) < 1.0 {
		t.Errorf("t statistic should indicate significant difference, got %v", tstat)
	}

	if p > 0.05 {
		t.Errorf("p-value should be < 0.05 for these clearly different groups, got %v", p)
	}
}

func TestTTestTwoSampleEqual(t *testing.T) {
	// Two samples from the same distribution should not be significant
	data1 := stats.Float64Data{5.0, 5.1, 4.9, 5.0, 5.1}
	data2 := stats.Float64Data{5.0, 4.9, 5.1, 5.0, 4.9}

	_, p, err := stats.TTest(data1, data2, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if p < 0.05 {
		t.Errorf("p-value should be > 0.05 for similar groups, got %v", p)
	}
}

func TestTTestEmptyInput(t *testing.T) {
	_, _, err := stats.TTest(stats.Float64Data{}, nil, 0)
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestTTestOneSampleTooFewValues(t *testing.T) {
	_, _, err := stats.TTest(stats.Float64Data{1.0}, nil, 0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for single value one-sample test, got %v", err)
	}
}

func TestTTestTwoSampleTooFewValues(t *testing.T) {
	_, _, err := stats.TTest(stats.Float64Data{1.0}, stats.Float64Data{2.0}, 0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for too few values in two-sample test, got %v", err)
	}
}

func TestTTestOneSampleNoEffect(t *testing.T) {
	// Sample mean matches population mean
	data := stats.Float64Data{10.0, 10.0, 10.0, 10.0}

	tstat, p, err := stats.TTest(data, nil, 10.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if tstat != 0 {
		t.Errorf("t statistic should be 0 when sample mean equals population mean, got %v", tstat)
	}

	if p != 1.0 {
		t.Errorf("p-value should be 1.0 when t=0, got %v", p)
	}
}

func TestTTestOneSampleZeroVarianceDifferentMean(t *testing.T) {
	// All identical values but different from population mean
	data := stats.Float64Data{10.0, 10.0, 10.0, 10.0}

	_, _, err := stats.TTest(data, nil, 5.0)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds for zero variance with different mean, got %v", err)
	}
}

func TestTTestLargeStatistic(t *testing.T) {
	// Very different groups to produce extreme t value
	data1 := stats.Float64Data{100, 101, 102, 103, 104}
	data2 := stats.Float64Data{1, 2, 3, 4, 5}

	tstat, p, err := stats.TTest(data1, data2, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if tstat < 50 {
		t.Errorf("t statistic should be very large, got %v", tstat)
	}

	if p > 0.0001 {
		t.Errorf("p-value should be extremely small, got %v", p)
	}
}
