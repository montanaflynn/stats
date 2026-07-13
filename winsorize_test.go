package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleWinsorize() {
	w, _ := stats.Winsorize([]float64{1, 2, 3, 4, 100}, 0.2)
	fmt.Println(w)
	// Output: [2 2 3 4 4]
}

func TestWinsorize(t *testing.T) {
	for _, c := range []struct {
		in      []float64
		percent float64
		out     []float64
	}{
		{[]float64{1, 2, 3, 4, 100}, 0.2, []float64{2, 2, 3, 4, 4}},
		// Verified against scipy.stats.mstats.winsorize and preserves
		// the original element order
		{
			[]float64{4.2, 1.1, 9.9, 3.3, 2.7, 8.8, 5.5, 7.4, 6.6, 100.5},
			0.2,
			[]float64{4.2, 3.3, 8.8, 3.3, 3.3, 8.8, 5.5, 7.4, 6.6, 8.8},
		},
		{[]float64{1, 2, 3, 4, 5}, 0.1, []float64{1, 2, 3, 4, 5}},
		{[]float64{42}, 0.4, []float64{42}},
	} {
		got, err := stats.Winsorize(c.in, c.percent)
		if err != nil {
			t.Errorf("Winsorize(%v, %v) returned an error: %v", c.in, c.percent, err)
			continue
		}
		for i := range c.out {
			if got[i] != c.out[i] {
				t.Errorf("Winsorize(%v, %v) => %v != %v", c.in, c.percent, got, c.out)
				break
			}
		}
	}
}

func TestWinsorizeZeroPercentReturnsCopy(t *testing.T) {
	data := []float64{3, 1, 2}
	w, err := stats.Winsorize(data, 0)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	for i := range data {
		if w[i] != data[i] {
			t.Errorf("Winsorize(%v, 0) => %v != %v", data, w, data)
			break
		}
	}
	w[0] = 99
	if data[0] != 3 {
		t.Errorf("Output should be a copy but input was mutated: %v", data)
	}
}

func TestWinsorizeInvalidInput(t *testing.T) {
	w, err := stats.Winsorize([]float64{}, 0.1)
	if err != stats.ErrEmptyInput {
		t.Errorf("Expected ErrEmptyInput, got %v", err)
	}
	if w != nil {
		t.Errorf("Expected nil output, got %v", w)
	}

	for _, percent := range []float64{-0.1, 0.5, 1.0, math.NaN()} {
		w, err = stats.Winsorize([]float64{1, 2, 3}, percent)
		if err != stats.ErrBounds {
			t.Errorf("Winsorize percent %v expected ErrBounds, got %v", percent, err)
		}
		if w != nil {
			t.Errorf("Winsorize percent %v expected nil output, got %v", percent, w)
		}
	}
}

func TestWinsorizeDoesNotMutateInput(t *testing.T) {
	data := []float64{100, 1, 4, 2, 3}
	_, err := stats.Winsorize(data, 0.2)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	want := []float64{100, 1, 4, 2, 3}
	for i := range want {
		if data[i] != want[i] {
			t.Errorf("Input was mutated: %v != %v", data, want)
			break
		}
	}
}

func TestFloat64DataWinsorize(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 100}
	w, err := data.Winsorize(0.2)
	if err != nil {
		t.Errorf("Should not have returned an error: %v", err)
	}
	want := []float64{2, 2, 3, 4, 4}
	for i := range want {
		if w[i] != want[i] {
			t.Errorf("Winsorize(0.2) => %v != %v", w, want)
			break
		}
	}
}
