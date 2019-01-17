package stats

import (
	"math"
	"testing"
)

func equalSlice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] && !(math.IsNaN(v) && math.IsNaN(b[i])) {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	var testCases = []struct {
		in  []float64
		out []float64
	}{
		{
			[]float64{-4, 2, -4, -2, 2, 4, 4},
			[]float64{-4, 2, -2, 4},
		},
		{
			[]float64{59, math.Inf(-1), -959.3, math.NaN(), -784, 59, 74.3, math.NaN(), 38.2, math.Inf(-1), 905, math.Inf(1)},
			[]float64{59, math.Inf(-1), -959.3, -784, 74.3, 38.2, 905, math.Inf(1)},
		},
	}
	for _, tc := range testCases {
		output := RemoveDuplicates(tc.in)
		if !equalSlice(output, tc.out) {
			t.Errorf("got %v, want %v", output, tc.out)
		}
	}
}

func BenchmarkRemoveDuplicatesSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(makeFloatSlice(5))
	}
}

func BenchmarkRemoveDuplicatesLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(lf)
	}
}

func TestFindUniques(t *testing.T) {
	var testCases = []struct {
		in  []float64
		out []float64
	}{
		{
			[]float64{-4, -42, 2, math.Inf(-1), -4, -2, 2, math.Inf(1), 4, math.Inf(1), 4, 5},
			[]float64{-42, math.Inf(-1), -2, 5},
		},
		{
			[]float64{59, math.Inf(-1), math.NaN(), -959.7485, -784, 59, 74.3, 238.2, math.Inf(-1), 905, math.Inf(1)},
			[]float64{-959.7485, -784, 74.3, 238.2, 905, math.Inf(1)},
		},
		{
			[]float64{1, math.NaN(), 2, 3, 4, 1, 2, 3, math.NaN()},
			[]float64{4},
		},
		{
			[]float64{1, math.NaN(), 2, 3, 4, 1, 2, 3},
			[]float64{4},
		},
	}
	for _, tc := range testCases {
		output := FindUniques(tc.in)
		if !equalSlice(output, tc.out) {
			t.Errorf("got %v, want %v", output, tc.out)
		}
	}
}
func BenchmarkFindUniquesSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindUniques(makeFloatSlice(5))
	}
}

func BenchmarkFindUniquesLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindUniques(lf)
	}
}
