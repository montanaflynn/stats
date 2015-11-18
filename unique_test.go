package stats

import (
	"reflect"
	"sort"
	"testing"
)

func TestUnique(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		// {[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}},
		// {[]float64{1, 2, 2, 2, 3, 5}, []float64{1, 3, 5}},
		{[]float64{5, 1, 2, 3, 4, 5}, []float64{1, 2, 3, 4}},
	} {
		got, _ := Unique(c.in)
		sort.Float64s(got)
		if !reflect.DeepEqual(got, c.out) {
			t.Errorf("Unique(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Unique([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkUniqueSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique(makeFloatSlice(5))
	}
}

func BenchmarkUniqueLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(lf)
	}
}
