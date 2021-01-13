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
}
