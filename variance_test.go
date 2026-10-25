package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestVariance(t *testing.T) {
	_, err := stats.Variance([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestPopulationVariance(t *testing.T) {
	e, err := stats.PopulationVariance([]float64{})
	if !math.IsNaN(e) {
		t.Errorf("%.1f != %.1f", e, math.NaN())
	}
	if err != stats.EmptyInputErr {
		t.Errorf("%v != %v", err, stats.EmptyInputErr)
	}

	pv, _ := stats.PopulationVariance([]float64{1, 2, 3})
	a, err := stats.Round(pv, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if a != 0.7 {
		t.Errorf("%.1f != %.1f", a, 0.7)
	}
}

func TestSampleVariance(t *testing.T) {
	m, err := stats.SampleVariance([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
	if err != stats.EmptyInputErr {
		t.Errorf("%v != %v", err, stats.EmptyInputErr)
	}
	m, _ = stats.SampleVariance([]float64{1, 2, 3})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestSampleStatisticsSingleElement(t *testing.T) {
	one := stats.Float64Data{3}
	for _, c := range []struct {
		name string
		call func() (float64, error)
	}{
		{"SampleVariance", func() (float64, error) { return stats.SampleVariance(one) }},
		{"VarS", func() (float64, error) { return stats.VarS(one) }},
		{"StandardDeviationSample", func() (float64, error) { return stats.StandardDeviationSample(one) }},
		{"StdDevS", func() (float64, error) { return stats.StdDevS(one) }},
		{"SEM", func() (float64, error) { return stats.SEM(one) }},
		{"CoefficientOfVariation", func() (float64, error) { return stats.CoefficientOfVariation(one) }},
		{"Covariance", func() (float64, error) { return stats.Covariance(one, stats.Float64Data{2}) }},
		{"ZScore", func() (float64, error) {
			z, err := stats.ZScore(one)
			if z != nil {
				t.Errorf("ZScore should return nil, got %v", z)
			}
			return math.NaN(), err
		}},
	} {
		got, err := c.call()
		if err != stats.ErrTooFew || !math.IsNaN(got) {
			t.Errorf("%s(%v) = %v, %v; want NaN, ErrTooFew", c.name, one, got, err)
		}
	}

	// Two elements are enough for the sample statistics
	if v, err := stats.SampleVariance(stats.Float64Data{1, 3}); err != nil || v != 2 {
		t.Errorf("SampleVariance([1 3]) = %v, %v; want 2, nil", v, err)
	}
}

func TestCovariance(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := stats.Covariance(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := stats.Covariance(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 3.2499999999999996 {
		t.Errorf("Covariance %v != %v", a, 3.2499999999999996)
	}

	_, err = stats.Covariance(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestCovariancePopulation(t *testing.T) {
	s1 := []float64{1, 2, 3.5, 3.7, 8, 12}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{0.5, 1, 2.1, 3.4, 3.4, 4}
	s4 := []float64{}

	_, err := stats.CovariancePopulation(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := stats.CovariancePopulation(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 4.191666666666666 {
		t.Errorf("CovariancePopulation %v != %v", a, 4.191666666666666)
	}

	_, err = stats.CovariancePopulation(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestVarianceExtremeScales(t *testing.T) {
	// Squaring deviations of 1e±150 overflows or underflows, but the
	// variances 1.25e±300 and 5/3e±300 do not
	for _, s := range []float64{1e150, 1e-150} {
		d := stats.Float64Data{1 * s, 2 * s, 3 * s, 4 * s}
		if v, err := stats.PopulationVariance(d); err != nil || !close(v, 1.25*s*s) {
			t.Errorf("PopulationVariance(%v) = %v, %v; want %v", d, v, err, 1.25*s*s)
		}
		if v, err := stats.SampleVariance(d); err != nil || !close(v, 5.0/3*s*s) {
			t.Errorf("SampleVariance(%v) = %v, %v; want %v", d, v, err, 5.0/3*s*s)
		}
	}
}
