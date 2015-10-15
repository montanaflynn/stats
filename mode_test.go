package stats

import (
	"reflect"
	"sort"
	"testing"
)

func TestMode(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{5, 3, 4, 2, 1}, []float64{}},
		{[]float64{5, 5, 3, 4, 2, 1}, []float64{5}},
		{[]float64{5, 5, 3, 3, 4, 2, 1}, []float64{3, 5}},
		{[]float64{1}, []float64{1}},
	} {
		got, err := Mode(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		sort.Float64s(got)
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Mode(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Mode([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkModeSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mode(makeFloatSlice(5))
	}
}

func BenchmarkModeLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mode(lf)
	}
}
