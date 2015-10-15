package stats

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	m, err := Round(0.1111, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.1 {
		t.Errorf("%.1f != %.1f", m, 0.1)
	}

	m, err = Round(-0.1111, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != -0.11 {
		t.Errorf("%.1f != %.1f", m, -0.11)
	}

	m, err = Round(5.3253, 3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 5.325 {
		t.Errorf("%.1f != %.1f", m, 5.325)
	}

	m, err = Round(5.3253, 0)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 5.0 {
		t.Errorf("%.1f != %.1f", m, 5.0)
	}

	m, err = Round(math.NaN(), 2)
	if err == nil {
		t.Errorf("Round should error on NaN")
	}
}

func BenchmarkRoundSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(0.1111, 1)
	}
}
