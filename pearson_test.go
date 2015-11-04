package stats

import "testing"

func TestPearson(t *testing.T) {
	s1 := []float64{1, 2, 3.5, 3.7, 8, 12}
	s2 := []float64{1, 2, 3, 5, 6}
	s3 := []float64{0.5, 1, 2.1, 3.4, 3.4, 4}
	s4 := []float64{}

	_, err := Pearson(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := Pearson(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 0.8437608593527455 {
		t.Errorf("Pearson %v != %v", a, 0.8437608593527455)
	}

	_, err = Pearson(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
