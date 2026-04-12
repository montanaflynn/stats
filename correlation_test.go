package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleCorrelation() {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{1, 2, 3, 5, 6}
	a, _ := stats.Correlation(s1, s2)
	rounded, _ := stats.Round(a, 5)
	fmt.Println(rounded)
	// Output: 0.99124
}

func TestCorrelation(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}
	s5 := []float64{0, 0, 0}
	testCases := []struct {
		name   string
		input  [][]float64
		output float64
		err    error
	}{
		{"Empty Slice Error", [][]float64{s4, s4}, math.NaN(), stats.EmptyInputErr},
		{"Different Length Error", [][]float64{s1, s2}, math.NaN(), stats.SizeErr},
		{"Correlation Value", [][]float64{s1, s3}, 0.9912407071619302, nil},
		{"Same Input Value", [][]float64{s5, s5}, 0.00, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := stats.Correlation(tc.input[0], tc.input[1])
			if err != nil {
				if err != tc.err {
					t.Errorf("Should have returned error %s", tc.err)
				}
			} else if !veryclose(a, tc.output) {
				t.Errorf("Result %.08f should be %.08f", a, tc.output)
			}
			a2, err2 := stats.Pearson(tc.input[0], tc.input[1])
			if err2 != nil {
				if err2 != tc.err {
					t.Errorf("Should have returned error %s", tc.err)
				}
			} else if !veryclose(a2, tc.output) {
				t.Errorf("Result %.08f should be %.08f", a2, tc.output)
			}
		})
	}
}

func ExampleSpearman() {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{5, 6, 7, 8, 7}
	a, _ := stats.Spearman(s1, s2)
	rounded, _ := stats.Round(a, 5)
	fmt.Println(rounded)
	// Output: 0.82078
}

func TestSpearman(t *testing.T) {
	testCases := []struct {
		name   string
		data1  []float64
		data2  []float64
		output float64
		err    error
	}{
		// Error cases
		{"Empty Slice Error", []float64{}, []float64{}, math.NaN(), stats.EmptyInputErr},
		{"Different Length Error", []float64{1, 2, 3, 4, 5}, []float64{10, -51.2, 8}, math.NaN(), stats.SizeErr},

		// Perfect correlation
		{"Perfect Positive", []float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}, 1.0, nil},
		{"Perfect Negative", []float64{1, 2, 3, 4, 5}, []float64{5, 4, 3, 2, 1}, -1.0, nil},

		// Verified against SciPy spearmanr and R cor.test(method="spearman")
		{
			"Basic with tie in y (SciPy/R)",
			[]float64{1, 2, 3, 4, 5},
			[]float64{5, 6, 7, 8, 7},
			0.8207826816681233,
			nil,
		},
		{
			"Negative with tie (SciPy mstats)",
			[]float64{5.05, 6.75, 3.21, 2.66},
			[]float64{1.65, 2.64, 2.64, 6.95},
			-0.632455532033676,
			nil,
		},
		{
			"15 elements with ties in both (SciPy/R)",
			[]float64{2.0, 47.4, 42.0, 10.8, 60.1, 1.7, 64.0, 63.1, 1.0, 1.4, 7.9, 0.3, 3.9, 0.3, 6.7},
			[]float64{22.6, 8.3, 44.4, 11.9, 24.6, 0.6, 5.7, 41.6, 0.0, 0.6, 6.7, 3.8, 1.0, 1.2, 1.4},
			0.6887298747763864,
			nil,
		},
		{
			"Adjacent swap (SciPy)",
			[]float64{0, 1, 2, 3, 4, 5},
			[]float64{0, 1, 2, 3, 5, 4},
			0.9428571428571429,
			nil,
		},
		{
			"Heavy ties in both (SciPy)",
			[]float64{1, 1, 1, 2, 2, 2},
			[]float64{1, 1, 2, 2, 3, 3},
			0.8164965809277261,
			nil,
		},

		// Zero standard deviation case (all same values)
		{"Same Input Value", []float64{0, 0, 0}, []float64{0, 0, 0}, 0.00, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := stats.Spearman(tc.data1, tc.data2)
			if err != nil {
				if err != tc.err {
					t.Errorf("Should have returned error %s, got %s", tc.err, err)
				}
			} else if !veryclose(a, tc.output) {
				t.Errorf("Result %.16f should be %.16f", a, tc.output)
			}

			// Also test the Float64Data method
			a2, err2 := stats.Float64Data(tc.data1).Spearman(tc.data2)
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

func BenchmarkSpearman(b *testing.B) {
	s1 := []float64{2.0, 47.4, 42.0, 10.8, 60.1, 1.7, 64.0, 63.1, 1.0, 1.4, 7.9, 0.3, 3.9, 0.3, 6.7}
	s2 := []float64{22.6, 8.3, 44.4, 11.9, 24.6, 0.6, 5.7, 41.6, 0.0, 0.6, 6.7, 3.8, 1.0, 1.2, 1.4}
	for i := 0; i < b.N; i++ {
		_, _ = stats.Spearman(s1, s2)
	}
}

func ExampleAutoCorrelation() {
	s1 := []float64{1, 2, 3, 4, 5}
	a, _ := stats.AutoCorrelation(s1, 1)
	fmt.Println(a)
	// Output: 0.4
}

func TestAutoCorrelation(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{}

	a, err := stats.AutoCorrelation(s1, 1)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}
	if a != 0.4 {
		t.Errorf("Should have returned 0.4")
	}

	_, err = stats.AutoCorrelation(s2, 1)
	if err != stats.EmptyInputErr {
		t.Errorf("Should have returned empty input error")
	}

	// Reference values cross-checked against statsmodels.tsa.stattools.acf,
	// R acf(), and MATLAB autocorr (issue #83).
	s3 := []float64{22, 24, 25, 25, 28, 29, 34, 37, 40, 44, 51, 48, 47, 50, 51}
	expected := []float64{
		1.0,
		0.83174224,
		0.65632458,
		0.49105012,
		0.27863962,
		0.03102625,
		-0.16527446,
		-0.30369928,
		-0.40095465,
		-0.45823389,
		-0.45047733,
	}
	for lag, want := range expected {
		got, err := stats.AutoCorrelation(s3, lag)
		if err != nil {
			t.Errorf("AutoCorrelation(s3, %d) returned error: %v", lag, err)
		}
		if math.Abs(got-want) > 1e-8 {
			t.Errorf("AutoCorrelation(s3, %d) = %v, want %v", lag, got, want)
		}
	}

	if _, err := stats.AutoCorrelation(s1, -1); err != stats.BoundsErr {
		t.Errorf("Should have returned bounds error for negative lag")
	}
	if _, err := stats.AutoCorrelation(s1, len(s1)); err != stats.BoundsErr {
		t.Errorf("Should have returned bounds error for lag >= len(data)")
	}

	constant := []float64{3, 3, 3, 3, 3}
	if a, err := stats.AutoCorrelation(constant, 1); err != nil || a != 0 {
		t.Errorf("AutoCorrelation of constant series should be 0, got %v err=%v", a, err)
	}
}
