package stats

import (
	"reflect"
	"testing"
)

func TestCumulativeSum(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{1, 2, 3}, []float64{1, 3, 6}},
		{[]float64{1.0, 1.1, 1.2, 2.2}, []float64{1.0, 2.1, 3.3, 5.5}},
		{[]float64{-1, -1, 2, -3}, []float64{-1, -2, 0, -3}},
	} {
		got, err := CumulativeSum(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("CumulativeSum(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := CumulativeSum([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkCumulativeSumSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CumulativeSum(makeFloatSlice(5))
	}
}

func BenchmarkCumulativeSumLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CumulativeSum(lf)
	}
}
