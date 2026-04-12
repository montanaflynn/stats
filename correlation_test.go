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
