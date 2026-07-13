package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleRMS() {
	rms, _ := stats.RMS([]float64{3, 4})
	fmt.Printf("%.4f", rms)
	// Output: 3.5355
}

func TestRMS(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{3, 4}, 3.5355339059327378},
		{[]float64{1, 1, 1}, 1},
		{[]float64{-3, 4}, 3.5355339059327378},
		{[]float64{0, 0, 0}, 0},
		{[]float64{2}, 2},
	} {
		got, err := stats.RMS(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if math.Abs(got-c.out) > 0.0000001 {
			t.Errorf("RMS(%v) => %v != %v", c.in, got, c.out)
		}
	}
}

func TestRMSEmptyInput(t *testing.T) {
	r, err := stats.RMS([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(r) {
		t.Errorf("expected NaN, got %v", r)
	}
}

func TestFloat64DataRMS(t *testing.T) {
	data := stats.Float64Data{3, 4}

	r, err := data.RMS()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-3.5355339059327378) > 0.0000001 {
		t.Errorf("got %v, want ~3.5355", r)
	}
}
