package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleCumulativeSum() {
	data := []float64{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}
	csum, _ := stats.CumulativeSum(data)
	fmt.Println(csum)
	// Output: [1 3.1 6.300000000000001 11.123000000000001 15.223 21.023]
}

func TestCumulativeSum(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{1, 2, 3}, []float64{1, 3, 6}},
		{[]float64{1.0, 1.1, 1.2, 2.2}, []float64{1.0, 2.1, 3.3, 5.5}},
		{[]float64{-1, -1, 2, -3}, []float64{-1, -2, 0, -3}},
	} {
		got, err := stats.CumulativeSum(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("CumulativeSum(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.CumulativeSum([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkCumulativeSumSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeSum(makeFloatSlice(5))
	}
}

func BenchmarkCumulativeSumLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeSum(lf)
	}
}
