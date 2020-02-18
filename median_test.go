package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleMedian() {
	data := []float64{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}
	median, _ := stats.Median(data)
	fmt.Println(median)
	// Output: 3.65
}

func TestMedian(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{5, 3, 4, 2, 1}, 3.0},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := stats.Median(c.in)
		if got != c.out {
			t.Errorf("Median(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.Median([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkMedianSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Median(makeFloatSlice(5))
	}
}

func BenchmarkMedianLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Median(lf)
	}
}

func TestMedianSortSideEffects(t *testing.T) {
	s := []float64{0.1, 0.3, 0.2, 0.4, 0.5}
	a := []float64{0.1, 0.3, 0.2, 0.4, 0.5}
	_, _ = stats.Median(s)
	if !reflect.DeepEqual(s, a) {
		t.Errorf("%.1f != %.1f", s, a)
	}
}
