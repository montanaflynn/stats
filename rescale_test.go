package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleRescale() {
	r, _ := stats.Rescale([]float64{2, 4, 6})
	fmt.Println(r)
	// Output: [0 0.5 1]
}

func TestRescale(t *testing.T) {
	for _, c := range []struct {
		in       []float64
		expected []float64
	}{
		{[]float64{2, 4, 6}, []float64{0, 0.5, 1}},
		{[]float64{10, 0}, []float64{1, 0}},
		{[]float64{-2, 0, 2}, []float64{0, 0.5, 1}},
	} {
		got, err := stats.Rescale(c.in)
		if err != nil {
			t.Error(err)
		}
		for i, v := range c.expected {
			if got[i] != v {
				t.Errorf("%v != %v", got[i], v)
			}
		}
	}
}

func TestRescaleInputNotMutated(t *testing.T) {
	input := []float64{2, 4, 6}
	_, err := stats.Rescale(input)
	if err != nil {
		t.Error(err)
	}

	original := []float64{2, 4, 6}
	for i, v := range original {
		if input[i] != v {
			t.Errorf("input was mutated: %v != %v", input[i], v)
		}
	}
}

func TestRescaleConstantInput(t *testing.T) {
	_, err := stats.Rescale([]float64{5, 5, 5})
	if err != stats.ErrZero {
		t.Errorf("expected ErrZero, got %v", err)
	}
}

func TestRescaleEmptyInput(t *testing.T) {
	_, err := stats.Rescale([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestFloat64DataRescale(t *testing.T) {
	d := stats.Float64Data{2, 4, 6}
	r, err := d.Rescale()
	if err != nil {
		t.Error(err)
	}

	expected := []float64{0, 0.5, 1}
	for i, v := range expected {
		if r[i] != v {
			t.Errorf("%v != %v", r[i], v)
		}
	}
}
