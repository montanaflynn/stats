package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleClip() {
	c, _ := stats.Clip([]float64{1, 5, 10}, 2, 8)
	fmt.Println(c)
	// Output: [2 5 8]
}

func TestClip(t *testing.T) {
	for _, c := range []struct {
		in       []float64
		min      float64
		max      float64
		expected []float64
	}{
		{[]float64{1, 5, 10}, 2, 8, []float64{2, 5, 8}},
		{[]float64{-1, 0, 1}, 0, 0, []float64{0, 0, 0}},
		{[]float64{3, 4}, 1, 10, []float64{3, 4}},
	} {
		got, err := stats.Clip(c.in, c.min, c.max)
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

func TestClipInputNotMutated(t *testing.T) {
	input := []float64{1, 5, 10}
	_, err := stats.Clip(input, 2, 8)
	if err != nil {
		t.Error(err)
	}

	original := []float64{1, 5, 10}
	for i, v := range original {
		if input[i] != v {
			t.Errorf("input was mutated: %v != %v", input[i], v)
		}
	}
}

func TestClipInvalidBounds(t *testing.T) {
	_, err := stats.Clip([]float64{1, 2, 3}, 8, 2)
	if err != stats.ErrBounds {
		t.Errorf("expected ErrBounds, got %v", err)
	}
}

func TestClipEmptyInput(t *testing.T) {
	_, err := stats.Clip([]float64{}, 0, 1)
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestFloat64DataClip(t *testing.T) {
	d := stats.Float64Data{1, 5, 10}
	c, err := d.Clip(2, 8)
	if err != nil {
		t.Error(err)
	}

	expected := []float64{2, 5, 8}
	for i, v := range expected {
		if c[i] != v {
			t.Errorf("%v != %v", c[i], v)
		}
	}
}
