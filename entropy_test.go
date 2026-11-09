package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleEntropy() {
	d := []float64{1.1, 2.2, 3.3}
	e, _ := stats.Entropy(d)
	fmt.Println(e)
	// Output: 1.0114042647073518
}

func TestEntropy(t *testing.T) {
	for _, c := range []struct {
		in  stats.Float64Data
		out float64
	}{
		{stats.Float64Data{4, 8, 5, 1}, 1.2110440167801229},
		{stats.Float64Data{0.8, 0.01, 0.4}, 0.6791185708986585},
		{stats.Float64Data{0.8, 1.1, 0, 5}, 0.7759393943707658},
	} {
		got, err := stats.Entropy(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !veryclose(got, c.out) {
			t.Errorf("Max(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Entropy([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func TestEntropyErrors(t *testing.T) {
	// Empty input
	if got, err := stats.Entropy([]float64{}); err != stats.ErrEmptyInput || !math.IsNaN(got) {
		t.Errorf("Entropy([]) = (%v, %v), want (NaN, ErrEmptyInput)", got, err)
	}

	// Negative values
	for _, in := range []stats.Float64Data{
		{-1, 2},
		{4, -0.01, 5},
		{-10},
	} {
		if got, err := stats.Entropy(in); err != stats.ErrNegative || !math.IsNaN(got) {
			t.Errorf("Entropy(%v) = (%v, %v), want (NaN, ErrNegative)", in, got, err)
		}
	}

	// Zero sum / all zero values
	for _, in := range []stats.Float64Data{
		{0},
		{0, 0, 0},
	} {
		if got, err := stats.Entropy(in); err != stats.ErrZero || !math.IsNaN(got) {
			t.Errorf("Entropy(%v) = (%v, %v), want (NaN, ErrZero)", in, got, err)
		}
	}
}

func TestEntropyNaN(t *testing.T) {
	for _, in := range []stats.Float64Data{
		{1, math.NaN(), 3},
		{math.NaN()},
	} {
		if got, err := stats.Entropy(in); err != nil || !math.IsNaN(got) {
			t.Errorf("Entropy(%v) = (%v, %v), want (NaN, nil)", in, got, err)
		}
	}
}

func TestEntropyOverflow(t *testing.T) {
	want := math.Ln2
	for _, in := range []stats.Float64Data{
		{math.MaxFloat64, math.MaxFloat64},
		{1e308, 1e308},
		{1e-300, 1e-300},
		{math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64},
	} {
		got, err := stats.Entropy(in)
		if err != nil || !close(got, want) {
			t.Errorf("Entropy(%v) = (%v, %v), want (%v, nil)", in, got, err, want)
		}
	}
}

func TestEntropyScale(t *testing.T) {
	base := stats.Float64Data{4, 8, 5, 1}
	want, err := stats.Entropy(base)
	if err != nil {
		t.Fatalf("unexpected error on base data: %v", err)
	}

	for _, scale := range []float64{1e100, 1e200, 1e300, 1e-100, 1e-200, 1e-300} {
		t.Run(fmt.Sprintf("scale=%g", scale), func(t *testing.T) {
			scaled := make(stats.Float64Data, len(base))
			for i, v := range base {
				scaled[i] = v * scale
			}
			got, err := stats.Entropy(scaled)
			if err != nil {
				t.Fatalf("Entropy at scale %g failed: %v", scale, err)
			}
			if !close(got, want) {
				t.Errorf("Entropy at scale %g = %v, want %v", scale, got, want)
			}
		})
	}
}

func TestFloat64DataEntropy(t *testing.T) {
	d := stats.Float64Data{4, 8, 5, 1}
	got1, err1 := d.Entropy()
	got2, err2 := stats.Entropy(d)
	if err1 != err2 || got1 != got2 {
		t.Errorf("d.Entropy() = (%v, %v), want (%v, %v)", got1, err1, got2, err2)
	}
}

func TestEntropyInputNotMutated(t *testing.T) {
	input := []float64{4, 8, 5, 1}
	_, err := stats.Entropy(input)
	if err != nil {
		t.Error(err)
	}

	original := []float64{4, 8, 5, 1}
	for i, v := range original {
		if input[i] != v {
			t.Errorf("input was mutated: %v != %v", input[i], v)
		}
	}
}

func BenchmarkEntropySmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Entropy(makeFloatSlice(5))
	}
}

func BenchmarkEntropyLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Entropy(lf)
	}
}
