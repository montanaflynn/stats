package stats

import (
	"testing"
)

func TestCorrelation(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := Correlation(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := Correlation(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 0.9912407071619301 {
		t.Errorf("Correlation %v != %v", a, 0.9912407071619301)
	}

	_, err = Correlation(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}

}
