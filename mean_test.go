package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestMean(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := stats.Mean(c.in)
		if got != c.out {
			t.Errorf("Mean(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Mean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestMeanOverflow(t *testing.T) {
	m := math.MaxFloat64
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{m, m, m}, m},
		{[]float64{-m, -m, -m}, -m},
		{[]float64{1e308, 1e308, -1e308}, 1e308 / 3},
		{[]float64{m, math.Inf(1)}, math.Inf(1)},
	} {
		if got, err := stats.Mean(c.in); err != nil || !veryclose(got, c.out) {
			t.Errorf("Mean(%v) => %v, %v; want %v", c.in, got, err, c.out)
		}
	}
	if got, _ := stats.Mean([]float64{math.Inf(1), math.Inf(-1)}); !math.IsNaN(got) {
		t.Errorf("Mean(+Inf, -Inf) => %v; want NaN", got)
	}
}

func BenchmarkMeanSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mean(makeFloatSlice(5))
	}
}

func BenchmarkMeanLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mean(lf)
	}
}

func TestGeometricMean(t *testing.T) {
	s1 := []float64{2, 18}
	s2 := []float64{10, 51.2, 8}
	s3 := []float64{1, 3, 9, 27, 81}
	s4 := []float64{5}
	s5 := []float64{10, -5, 2}
	s6 := []float64{1, 0, 9}

	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{s1, 6},
		{s2, 16},
		{s3, 9},
		{s4, 5},
	} {
		gm, err := stats.GeometricMean(c.in)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}

		gm, _ = stats.Round(gm, 0)
		if gm != c.out {
			t.Errorf("Geometric Mean %v != %v", gm, c.out)
		}
	}

	_, err := stats.GeometricMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}

	_, err = stats.GeometricMean(s5)
	if err == nil {
		t.Errorf("Should have returned a negative number error")
	}

	_, err = stats.GeometricMean(s6)
	if err == nil {
		t.Errorf("Should have returned a zero number error")
	}
}

func TestHarmonicMean(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 0, 9, 27, 81}

	hm, err := stats.HarmonicMean(s1)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	hm, _ = stats.Round(hm, 2)
	if hm != 2.19 {
		t.Errorf("Harmonic Mean %v != %v", hm, 2.19)
	}

	_, err = stats.HarmonicMean(s2)
	if err == nil {
		t.Errorf("Should have returned a negative number error")
	}

	_, err = stats.HarmonicMean(s3)
	if err == nil {
		t.Errorf("Should have returned a zero number error")
	}

	_, err = stats.HarmonicMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
