package stats_test

import (
	"fmt"
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
