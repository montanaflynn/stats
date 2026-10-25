package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestMedianAbsoluteDeviation(t *testing.T) {
	_, err := stats.MedianAbsoluteDeviation([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestMedianAbsoluteDeviationPopulation(t *testing.T) {
	s, _ := stats.MedianAbsoluteDeviation([]float64{1, 2, 3})
	m, err := stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.00 {
		t.Errorf("%.10f != %.10f", m, 1.00)
	}

	s, _ = stats.MedianAbsoluteDeviation([]float64{-2, 0, 4, 5, 7})
	m, err = stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 3.00 {
		t.Errorf("%.10f != %.10f", m, 3.00)
	}

	m, _ = stats.MedianAbsoluteDeviation([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestStandardDeviation(t *testing.T) {
	_, err := stats.StandardDeviation([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestStandardDeviationPopulation(t *testing.T) {
	s, _ := stats.StandardDeviationPopulation([]float64{1, 2, 3})
	m, err := stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.82 {
		t.Errorf("%.10f != %.10f", m, 0.82)
	}
	s, _ = stats.StandardDeviationPopulation([]float64{-1, -2, -3.3})
	m, err = stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.94 {
		t.Errorf("%.10f != %.10f", m, 0.94)
	}

	m, _ = stats.StandardDeviationPopulation([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestStandardDeviationSample(t *testing.T) {
	s, _ := stats.StandardDeviationSample([]float64{1, 2, 3})
	m, err := stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.0 {
		t.Errorf("%.10f != %.10f", m, 1.0)
	}
	s, _ = stats.StandardDeviationSample([]float64{-1, -2, -3.3})
	m, err = stats.Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.15 {
		t.Errorf("%.10f != %.10f", m, 1.15)
	}

	m, _ = stats.StandardDeviationSample([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestStandardDeviationExtremeScales(t *testing.T) {
	// The variance overflows or underflows, but the deviation does not
	for _, s := range []float64{1e200, 1e-200, 1e300} {
		d := stats.Float64Data{1 * s, 2 * s, 3 * s, 4 * s}
		if sd, err := stats.StandardDeviationPopulation(d); err != nil || !close(sd, math.Sqrt(1.25)*s) {
			t.Errorf("StandardDeviationPopulation(%v) = %v, %v; want %v", d, sd, err, math.Sqrt(1.25)*s)
		}
		if sd, err := stats.StandardDeviationSample(d); err != nil || !close(sd, math.Sqrt(5.0/3)*s) {
			t.Errorf("StandardDeviationSample(%v) = %v, %v; want %v", d, sd, err, math.Sqrt(5.0/3)*s)
		}
	}
	if sd, err := stats.StandardDeviationSample(stats.Float64Data{}); err != stats.ErrEmptyInput || !math.IsNaN(sd) {
		t.Errorf("StandardDeviationSample of empty input = %v, %v; want NaN, ErrEmptyInput", sd, err)
	}
	if sd, err := stats.StandardDeviationSample(stats.Float64Data{1, math.Inf(1)}); err != nil || !math.IsNaN(sd) {
		t.Errorf("StandardDeviationSample with Inf = %v, %v; want NaN", sd, err)
	}
}
