package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleHistogram() {
	counts, edges, _ := stats.Histogram([]float64{1, 2, 1, 3, 4}, 3)
	fmt.Println(counts, edges)
	// Output: [2 1 2] [1 2 3 4]
}

func TestHistogram(t *testing.T) {
	for _, c := range []struct {
		in     stats.Float64Data
		bins   int
		counts []int
		edges  []float64
	}{
		{stats.Float64Data{1, 2, 1, 3, 4}, 3, []int{2, 1, 2}, []float64{1, 2, 3, 4}},
		{stats.Float64Data{1, 1.5, 2, 2.5, 3}, 4, []int{1, 1, 1, 2}, []float64{1, 1.5, 2, 2.5, 3}},
		{stats.Float64Data{5, 5, 5}, 4, []int{0, 0, 3, 0}, []float64{4.5, 4.75, 5, 5.25, 5.5}},
		{stats.Float64Data{7}, 1, []int{1}, []float64{6.5, 7.5}},
		{stats.Float64Data{0, 10}, 2, []int{1, 1}, []float64{0, 5, 10}},
	} {
		counts, edges, err := stats.Histogram(c.in, c.bins)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(counts, c.counts) {
			t.Errorf("Histogram(%v, %d) counts => %v != %v", c.in, c.bins, counts, c.counts)
		}
		if !reflect.DeepEqual(edges, c.edges) {
			t.Errorf("Histogram(%v, %d) edges => %v != %v", c.in, c.bins, edges, c.edges)
		}
		total := 0
		for _, count := range counts {
			total += count
		}
		if total != c.in.Len() {
			t.Errorf("Histogram(%v, %d) counts sum => %d != %d", c.in, c.bins, total, c.in.Len())
		}
	}
}

func TestHistogramMaxInLastBin(t *testing.T) {
	counts, _, err := stats.Histogram([]float64{1, 2, 1, 3, 4}, 3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if counts[len(counts)-1] != 2 {
		t.Errorf("Maximum value not counted in the last bin => %v", counts)
	}
}

func TestHistogramEmptyInput(t *testing.T) {
	_, _, err := stats.Histogram([]float64{}, 3)
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice didn't return expected error; got %v", err)
	}
}

func TestHistogramInvalidBins(t *testing.T) {
	for _, bins := range []int{0, -1} {
		_, _, err := stats.Histogram([]float64{1, 2, 3}, bins)
		if err != stats.ErrBounds {
			t.Errorf("Histogram with %d bins didn't return expected error; got %v", bins, err)
		}
	}
}

func TestHistogramDoesNotMutateInput(t *testing.T) {
	in := stats.Float64Data{4, 1, 3, 2}
	_, _, err := stats.Histogram(in, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual(in, stats.Float64Data{4, 1, 3, 2}) {
		t.Errorf("Input slice was mutated => %v", in)
	}
}

func TestFloat64DataHistogram(t *testing.T) {
	in := stats.Float64Data{1, 2, 1, 3, 4}
	counts, edges, err := in.Histogram(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual(counts, []int{2, 1, 2}) {
		t.Errorf("Histogram(3) counts => %v != %v", counts, []int{2, 1, 2})
	}
	if !reflect.DeepEqual(edges, []float64{1, 2, 3, 4}) {
		t.Errorf("Histogram(3) edges => %v != %v", edges, []float64{1, 2, 3, 4})
	}
}

func BenchmarkHistogramSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = stats.Histogram(makeFloatSlice(5), 3)
	}
}

func BenchmarkHistogramLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = stats.Histogram(lf, 10)
	}
}
