package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleSEM() {
	sem, _ := stats.SEM([]float64{1, 2, 3, 4, 5})
	fmt.Printf("%.4f", sem)
	// Output: 0.7071
}

func TestSEM(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 0.7071067811865476},
		{[]float64{2, 4, 6}, 1.1547005383792515},
		{[]float64{5, 5, 5, 5}, 0},
	} {
		got, err := stats.SEM(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if math.Abs(got-c.out) > 0.0000001 {
			t.Errorf("SEM(%v) => %v != %v", c.in, got, c.out)
		}
	}
}

func TestSEMEmptyInput(t *testing.T) {
	r, err := stats.SEM([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(r) {
		t.Errorf("expected NaN, got %v", r)
	}
}

func TestFloat64DataSEM(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 5}

	r, err := data.SEM()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-0.7071067811865476) > 0.0000001 {
		t.Errorf("got %v, want ~0.7071", r)
	}
}
