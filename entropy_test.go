package stats

import "testing"

func TestEntropy(t *testing.T) {
	for _, c := range []struct {
		in  Float64Data
		out float64
	}{
		{Float64Data{4, 8, 5, 1}, 1.2110440167801229},
		{Float64Data{0.8, 0.01, 0.4}, 0.6791185708986585},
		{Float64Data{0.8, 1.1, 0, 5}, 0.7759393943707658},
	} {
		got, err := Entropy(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Max(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Entropy([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}
