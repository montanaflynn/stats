package stats_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleDiff() {
	data := []float64{1, 4, 9, 16}
	diff, _ := stats.Diff(data)
	fmt.Println(diff)
	// Output: [3 5 7]
}

func ExamplePercentChange() {
	data := []float64{100, 110, 99}
	change, _ := stats.PercentChange(data)
	fmt.Println(change)
	// Output: [0.1 -0.1]
}

func TestDiff(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{1, 4, 9, 16}, []float64{3, 5, 7}},
		{[]float64{1.5, 1.0, 2.5}, []float64{-0.5, 1.5}},
		{[]float64{-1, -3, 2}, []float64{-2, 5}},
		{[]float64{42}, []float64{}},
	} {
		got, err := stats.Diff(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got == nil {
			t.Errorf("Diff(%.1f) returned a nil slice", c.in)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Diff(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Diff([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice should have returned an empty input error")
	}
}

func TestDiffMethod(t *testing.T) {
	got, err := stats.Float64Data{1, 4, 9, 16}.Diff()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{3, 5, 7}, got) {
		t.Errorf("Float64Data.Diff() => %.1f != [3 5 7]", got)
	}
}

func TestPercentChange(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{100, 110, 99}, []float64{0.1, -0.1}},
		{[]float64{2, 1, 4}, []float64{-0.5, 3}},
		{[]float64{-10, -5}, []float64{-0.5}},
		{[]float64{42}, []float64{}},
	} {
		got, err := stats.PercentChange(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got == nil {
			t.Errorf("PercentChange(%.1f) returned a nil slice", c.in)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("PercentChange(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.PercentChange([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice should have returned an empty input error")
	}
}

func TestPercentChangeZeroDenominator(t *testing.T) {
	got, err := stats.PercentChange([]float64{0, 5, 0, -5, 0, 0})
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !math.IsInf(got[0], 1) {
		t.Errorf("Change from 0 to 5 should be +Inf, got %v", got[0])
	}
	if got[1] != -1 {
		t.Errorf("Change from 5 to 0 should be -1, got %v", got[1])
	}
	if !math.IsInf(got[2], -1) {
		t.Errorf("Change from 0 to -5 should be -Inf, got %v", got[2])
	}
	if !math.IsNaN(got[4]) {
		t.Errorf("Change from 0 to 0 should be NaN, got %v", got[4])
	}
}

func TestPercentChangeMethod(t *testing.T) {
	got, err := stats.Float64Data{100, 110, 99}.PercentChange()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{0.1, -0.1}, got) {
		t.Errorf("Float64Data.PercentChange() => %.1f != [0.1 -0.1]", got)
	}
}

func BenchmarkDiffLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Diff(lf)
	}
}

func BenchmarkPercentChangeLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.PercentChange(lf)
	}
}
