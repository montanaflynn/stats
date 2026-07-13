package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleKendallTau() {
	s1 := []float64{1, 2, 2, 3}
	s2 := []float64{1, 2, 3, 3}
	a, _ := stats.KendallTau(s1, s2)
	rounded, _ := stats.Round(a, 5)
	fmt.Println(rounded)
	// Output: 0.8
}

func TestKendallTau(t *testing.T) {
	testCases := []struct {
		name   string
		data1  []float64
		data2  []float64
		output float64
		err    error
	}{
		// Error cases
		{"Empty Slice Error", []float64{}, []float64{}, math.NaN(), stats.ErrEmptyInput},
		{"Different Length Error", []float64{1, 2, 3, 4, 5}, []float64{10, -51.2, 8}, math.NaN(), stats.ErrSize},

		// Perfect correlation
		{"Perfect Concordant", []float64{1, 2, 3, 4, 5}, []float64{10, 20, 30, 40, 50}, 1.0, nil},
		{"Perfect Discordant", []float64{1, 2, 3, 4, 5}, []float64{5, 4, 3, 2, 1}, -1.0, nil},

		// Verified against SciPy kendalltau
		{"Ties in both variables (SciPy)", []float64{1, 2, 2, 3}, []float64{1, 2, 3, 3}, 0.8, nil},
		{"Adjacent swap (SciPy)", []float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 5, 4}, 0.8, nil},
		{"Pair tied in both (SciPy)", []float64{1, 1, 2, 3}, []float64{2, 2, 1, 4}, 0.2, nil},

		// Zero denominator cases (no variation in one or both variables)
		{"Constant First Input", []float64{3, 3, 3}, []float64{1, 2, 3}, 0.0, nil},
		{"Constant Second Input", []float64{1, 2, 3}, []float64{3, 3, 3}, 0.0, nil},
		{"Constant Both Inputs", []float64{0, 0, 0}, []float64{0, 0, 0}, 0.0, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := stats.KendallTau(tc.data1, tc.data2)
			if err != nil {
				if err != tc.err {
					t.Errorf("Should have returned error %s, got %s", tc.err, err)
				}
			} else if !veryclose(a, tc.output) {
				t.Errorf("Result %.16f should be %.16f", a, tc.output)
			}

			// Also test the Float64Data method
			a2, err2 := stats.Float64Data(tc.data1).KendallTau(tc.data2)
			if err2 != nil {
				if err2 != tc.err {
					t.Errorf("Should have returned error %s, got %s", tc.err, err2)
				}
			} else if !veryclose(a2, tc.output) {
				t.Errorf("Result %.16f should be %.16f", a2, tc.output)
			}
		})
	}
}

func TestKendallTauDoesNotMutateInput(t *testing.T) {
	s1 := []float64{3, 1, 2, 2}
	s2 := []float64{2, 3, 1, 1}
	_, _ = stats.KendallTau(s1, s2)
	if s1[0] != 3 || s1[1] != 1 || s1[2] != 2 || s1[3] != 2 {
		t.Errorf("First input slice was mutated: %v", s1)
	}
	if s2[0] != 2 || s2[1] != 3 || s2[2] != 1 || s2[3] != 1 {
		t.Errorf("Second input slice was mutated: %v", s2)
	}
}

func BenchmarkKendallTau(b *testing.B) {
	s1 := []float64{2.0, 47.4, 42.0, 10.8, 60.1, 1.7, 64.0, 63.1, 1.0, 1.4, 7.9, 0.3, 3.9, 0.3, 6.7}
	s2 := []float64{22.6, 8.3, 44.4, 11.9, 24.6, 0.6, 5.7, 41.6, 0.0, 0.6, 6.7, 3.8, 1.0, 1.2, 1.4}
	for i := 0; i < b.N; i++ {
		_, _ = stats.KendallTau(s1, s2)
	}
}
