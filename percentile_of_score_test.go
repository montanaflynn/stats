package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExamplePercentileOfScore() {
	rank, _ := stats.PercentileOfScore([]float64{1, 2, 3, 4}, 3)
	fmt.Println(rank)
	// Output: 62.5
}

func TestPercentileOfScore(t *testing.T) {
	for _, c := range []struct {
		in    []float64
		score float64
		out   float64
	}{
		{[]float64{1, 2, 3, 4}, 3, 62.5},
		{[]float64{1, 2, 3, 4}, 0, 0},
		{[]float64{1, 2, 3, 4}, 5, 100},
		{[]float64{1, 2, 3, 4}, 2.5, 50},
		{[]float64{1, 2, 2, 2, 3}, 2, 50},
		{[]float64{2, 2, 2}, 2, 50},
		{[]float64{1, 2, 3, 4}, 1, 12.5},
		{[]float64{1, 2, 3, 4}, 4, 87.5},
	} {
		got, err := stats.PercentileOfScore(c.in, c.score)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if math.Abs(got-c.out) > 0.0000001 {
			t.Errorf("PercentileOfScore(%v, %v) => %v != %v", c.in, c.score, got, c.out)
		}
	}
}

func TestPercentileOfScoreEmptyInput(t *testing.T) {
	r, err := stats.PercentileOfScore([]float64{}, 1)
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(r) {
		t.Errorf("expected NaN, got %v", r)
	}
}

func TestFloat64DataPercentileOfScore(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4}

	r, err := data.PercentileOfScore(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 62.5 {
		t.Errorf("got %v, want 62.5", r)
	}
}
