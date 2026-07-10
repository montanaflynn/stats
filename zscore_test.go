package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleZScore() {
	z, _ := stats.ZScore([]float64{2, 4, 6})
	fmt.Println(z)
	// Output: [-1 0 1]
}

func TestZScore(t *testing.T) {
	input := []float64{2, 4, 6}
	z, err := stats.ZScore(input)
	if err != nil {
		t.Error(err)
	}

	expected := []float64{-1, 0, 1}
	for i, v := range expected {
		if z[i] != v {
			t.Errorf("%v != %v", z[i], v)
		}
	}

	// Ensure the input slice was not mutated
	original := []float64{2, 4, 6}
	for i, v := range original {
		if input[i] != v {
			t.Errorf("input was mutated: %v != %v", input[i], v)
		}
	}
}

func TestZScoreZeroStandardDeviation(t *testing.T) {
	_, err := stats.ZScore([]float64{5, 5, 5})
	if err != stats.ErrZero {
		t.Errorf("expected ErrZero, got %v", err)
	}
}

func TestZScoreEmptyInput(t *testing.T) {
	_, err := stats.ZScore([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestFloat64DataZScore(t *testing.T) {
	d := stats.Float64Data{2, 4, 6}
	z, err := d.ZScore()
	if err != nil {
		t.Error(err)
	}

	expected := []float64{-1, 0, 1}
	for i, v := range expected {
		if z[i] != v {
			t.Errorf("%v != %v", z[i], v)
		}
	}
}
