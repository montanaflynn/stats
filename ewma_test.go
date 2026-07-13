package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleEWMA() {
	data := []float64{1, 2, 3}
	ewma, _ := stats.EWMA(data, 0.5)
	fmt.Println(ewma)
	// Output: [1 1.5 2.25]
}

func TestEWMA(t *testing.T) {
	for _, c := range []struct {
		in    []float64
		alpha float64
		out   []float64
	}{
		{[]float64{1, 2, 3}, 0.5, []float64{1, 1.5, 2.25}},
		{[]float64{4}, 0.25, []float64{4}},
		{[]float64{2, 2, 2}, 0.75, []float64{2, 2, 2}},
	} {
		got, err := stats.EWMA(c.in, c.alpha)
		if err != nil {
			t.Errorf("EWMA(%v, %v) returned an error", c.in, c.alpha)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("EWMA(%v, %v) => %v != %v", c.in, c.alpha, got, c.out)
		}
	}
}

func TestEWMAAlphaOneCopiesInput(t *testing.T) {
	in := []float64{5, 6, 7}
	got, err := stats.EWMA(in, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual(in, got) {
		t.Errorf("EWMA(%v, 1) => %v != input", in, got)
	}
	got[0] = 99
	if in[0] != 5 {
		t.Errorf("EWMA(_, 1) must return a copy, not the input slice")
	}
}

func TestEWMAInvalidInput(t *testing.T) {
	_, err := stats.EWMA([]float64{}, 0.5)
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty input should have returned ErrEmptyInput, got %v", err)
	}
	for _, alpha := range []float64{0, -0.5, 1.5} {
		_, err := stats.EWMA([]float64{1, 2, 3}, alpha)
		if err != stats.ErrBounds {
			t.Errorf("Alpha %v should have returned ErrBounds, got %v", alpha, err)
		}
	}
}

func TestEWMADoesNotMutateInput(t *testing.T) {
	in := []float64{1, 2, 3}
	want := []float64{1, 2, 3}
	_, _ = stats.EWMA(in, 0.5)
	if !reflect.DeepEqual(want, in) {
		t.Errorf("Input slice was mutated: %v != %v", in, want)
	}
}

func TestFloat64DataEWMA(t *testing.T) {
	data := stats.Float64Data{1, 2, 3}
	got, err := data.EWMA(0.5)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{1, 1.5, 2.25}, got) {
		t.Errorf("Float64Data.EWMA(0.5) => %v != [1 1.5 2.25]", got)
	}
}
