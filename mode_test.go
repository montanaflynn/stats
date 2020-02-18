package stats_test

import (
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestMode(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{2, 2, 2, 2}, []float64{2}},
		{[]float64{5, 3, 4, 2, 1}, []float64{}},
		{[]float64{5, 5, 3, 3, 4, 4, 2, 2, 1, 1}, []float64{}},
		{[]float64{5, 5, 3, 4, 2, 1}, []float64{5}},
		{[]float64{5, 5, 3, 3, 4, 2, 1}, []float64{3, 5}},
		{[]float64{1}, []float64{1}},
		{[]float64{-50, -46.325, -46.325, -.87, 1, 2.1122, 3.20, 5, 15, 15, 15.0001}, []float64{-46.325, 15}},
		{[]float64{1, 2, 3, 4, 4, 4, 4, 4, 5, 3, 6, 7, 5, 0, 8, 8, 7, 6, 9, 9}, []float64{4}},
		{[]float64{76, 76, 110, 76, 76, 76, 76, 119, 76, 76, 76, 76, 31, 31, 31, 31, 83, 83, 83, 78, 78, 78, 78, 78, 78, 78, 78}, []float64{76}},
	} {
		got, err := stats.Mode(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Mode(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Mode([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkModeSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mode(makeFloatSlice(5))
	}
}

func BenchmarkModeSmallRandFloatSlice(b *testing.B) {
	lf := makeRandFloatSlice(5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mode(lf)
	}
}

func BenchmarkModeLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mode(lf)
	}
}

func BenchmarkModeLargeRandFloatSlice(b *testing.B) {
	lf := makeRandFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Mode(lf)
	}
}
