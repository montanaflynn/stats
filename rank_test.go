package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleRank() {
	r, _ := stats.Rank([]float64{3, 1, 4, 1})
	fmt.Println(r)
	// Output: [3 1.5 4 1.5]
}

func TestRank(t *testing.T) {
	input := []float64{3, 1, 4, 1}
	r, err := stats.Rank(input)
	if err != nil {
		t.Error(err)
	}

	expected := []float64{3, 1.5, 4, 1.5}
	for i, v := range expected {
		if r[i] != v {
			t.Errorf("%v != %v", r[i], v)
		}
	}

	// Ensure the input slice was not mutated
	original := []float64{3, 1, 4, 1}
	for i, v := range original {
		if input[i] != v {
			t.Errorf("input was mutated: %v != %v", input[i], v)
		}
	}
}

func TestRankEmptyInput(t *testing.T) {
	_, err := stats.Rank([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestFloat64DataRank(t *testing.T) {
	d := stats.Float64Data{3, 1, 4, 1}
	r, err := d.Rank()
	if err != nil {
		t.Error(err)
	}

	expected := []float64{3, 1.5, 4, 1.5}
	for i, v := range expected {
		if r[i] != v {
			t.Errorf("%v != %v", r[i], v)
		}
	}
}
