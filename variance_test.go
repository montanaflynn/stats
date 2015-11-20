package stats

import (
	"testing"
)

func TestVariance(t *testing.T) {
	_, err := Variance([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestPopulationVariance(t *testing.T) {
	e, _ := PopulationVariance([]float64{})
	if e != 0.0 {
		t.Errorf("%.1f != %.1f", e, 0.0)
	}
	pv, _ := PopulationVariance([]float64{1, 2, 3})
	a, err := Round(pv, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if a != 0.7 {
		t.Errorf("%.1f != %.1f", a, 0.7)
	}
}

func TestSampleVariance(t *testing.T) {
	m, _ := SampleVariance([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m, _ = SampleVariance([]float64{1, 2, 3})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestCovariance(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := Covariance(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := Covariance(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 3.2499999999999996 {
		t.Errorf("Covariance %v != %v", a, 3.2499999999999996)
	}

	_, err = Covariance(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestCovariancePopulation(t *testing.T) {
	s1 := []float64{1, 2, 3.5, 3.7, 8, 12}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{0.5, 1, 2.1, 3.4, 3.4, 4}
	s4 := []float64{}

	_, err := CovariancePopulation(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := CovariancePopulation(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 4.191666666666666 {
		t.Errorf("CovariancePopulation %v != %v", a, 4.191666666666666)
	}

	_, err = CovariancePopulation(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
