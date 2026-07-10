package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleArgMax() {
	d := []float64{1.1, 2.3, 3.2, 4.0, 4.01, 5.09}
	a, _ := stats.ArgMax(d)
	fmt.Println(a)
	// Output: 5
}

func ExampleArgMin() {
	d := []float64{1.1, 2.3, 3.2, 4.0, 4.01, 5.09}
	a, _ := stats.ArgMin(d)
	fmt.Println(a)
	// Output: 0
}

func ExampleRange() {
	d := []float64{4, 1, 9}
	a, _ := stats.Range(d)
	fmt.Println(a)
	// Output: 8
}

func TestArgMax(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out int
	}{
		{[]float64{1, 2, 3, 4, 5}, 4},
		{[]float64{10.5, 3, 5, 7, 9}, 0},
		{[]float64{-20, -1, -5.5}, 1},
		{[]float64{-1.0}, 0},
		{[]float64{1, 5, 2, 5}, 1},
		{[]float64{7, 7, 7}, 0},
	} {
		got, err := stats.ArgMax(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("ArgMax(%.1f) => %d != %d", c.in, got, c.out)
		}
	}
	_, err := stats.ArgMax([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice didn't return empty input error")
	}
}

func TestArgMin(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out int
	}{
		{[]float64{1, 2, 3, 4, 5}, 0},
		{[]float64{10.5, 3, 5, 7, 9}, 1},
		{[]float64{-20, -1, -5.5}, 0},
		{[]float64{-1.0}, 0},
		{[]float64{3, 1, 2, 1}, 1},
		{[]float64{7, 7, 7}, 0},
	} {
		got, err := stats.ArgMin(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("ArgMin(%.1f) => %d != %d", c.in, got, c.out)
		}
	}
	_, err := stats.ArgMin([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice didn't return empty input error")
	}
}

func TestRange(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{4, 1, 9}, 8.0},
		{[]float64{1, 2, 3, 4, 5}, 4.0},
		{[]float64{-20, -1, -5.5}, 19.0},
		{[]float64{-1.0}, 0.0},
		{[]float64{7, 7, 7}, 0.0},
	} {
		got, err := stats.Range(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Range(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Range([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice didn't return empty input error")
	}
}

func TestFloat64DataArgMaxArgMinRangeMethods(t *testing.T) {
	d := stats.Float64Data{1, 5, 2, 5}

	argMax, err := d.ArgMax()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if argMax != 1 {
		t.Errorf("ArgMax method => %d != 1", argMax)
	}

	argMin, err := d.ArgMin()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if argMin != 0 {
		t.Errorf("ArgMin method => %d != 0", argMin)
	}

	rng, err := d.Range()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if rng != 4.0 {
		t.Errorf("Range method => %.1f != 4.0", rng)
	}
}

func BenchmarkArgMaxSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.ArgMax(makeFloatSlice(5))
	}
}

func BenchmarkArgMinSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.ArgMin(makeFloatSlice(5))
	}
}

func BenchmarkRangeSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Range(makeFloatSlice(5))
	}
}
