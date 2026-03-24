package stats_test

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// Reference values verified against scipy.stats.skew
var skewnessTests = []struct {
	name string
	data stats.Float64Data
}{
	{"RightSkewed", stats.Float64Data{1, 2, 2, 3, 9}},
	{"LeftSkewed", stats.Float64Data{9, 8, 8, 7, 1}},
	{"Symmetric", stats.Float64Data{2, 4, 6, 8, 10}},
	{"UniformRange", stats.Float64Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{"HeavilyRightSkewed", stats.Float64Data{1, 1, 1, 1, 1, 1, 1, 1, 1, 100}},
}

func TestSkewness(t *testing.T) {
	// scipy.stats.skew(data, bias=True)
	expected := []float64{
		1.3210869678752710,
		-1.3210869678752706,
		0.0,
		0.0,
		2.6666666666666665,
	}

	for i, tt := range skewnessTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.Skewness(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestPopulationSkewness(t *testing.T) {
	// scipy.stats.skew(data, bias=True)
	expected := []float64{
		1.3210869678752710,
		-1.3210869678752706,
		0.0,
		0.0,
		2.6666666666666665,
	}

	for i, tt := range skewnessTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.PopulationSkewness(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestSampleSkewness(t *testing.T) {
	// scipy.stats.skew(data, bias=False)
	expected := []float64{
		1.9693601762387913,
		-1.9693601762387909,
		0.0,
		0.0,
		3.1622776601683795,
	}

	for i, tt := range skewnessTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.SampleSkewness(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestSkewness_ZeroVariance(t *testing.T) {
	for _, fn := range []func(stats.Float64Data) (float64, error){
		stats.Skewness, stats.PopulationSkewness, stats.SampleSkewness,
	} {
		_, err := fn(stats.Float64Data{5, 5, 5, 5})
		if err != stats.ErrEmptyInput {
			t.Fatalf("expected ErrEmptyInput for zero variance, got %v", err)
		}
	}
}

func TestSkewness_TooSmallDataset(t *testing.T) {
	_, err := stats.Skewness(stats.Float64Data{42})
	if err != stats.ErrEmptyInput {
		t.Fatalf("expected ErrEmptyInput for dataset length < 2, got %v", err)
	}
}

func TestSampleSkewness_TooSmallDataset(t *testing.T) {
	_, err := stats.SampleSkewness(stats.Float64Data{1, 2})
	if err != stats.ErrEmptyInput {
		t.Fatalf("expected ErrEmptyInput for dataset length < 3, got %v", err)
	}
}
