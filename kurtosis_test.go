package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

// Reference values verified against scipy.stats.kurtosis and pandas .kurt()
var kurtosisTests = []struct {
	name string
	data stats.Float64Data
}{
	{"RightSkewed", stats.Float64Data{1, 2, 2, 3, 9}},
	{"LeftSkewed", stats.Float64Data{9, 8, 8, 7, 1}},
	{"Symmetric", stats.Float64Data{2, 4, 6, 8, 10}},
	{"UniformRange", stats.Float64Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{"HeavilyRightSkewed", stats.Float64Data{1, 1, 1, 1, 1, 1, 1, 1, 1, 100}},
}

func TestKurtosis(t *testing.T) {
	// scipy.stats.kurtosis(data, fisher=True, bias=True)
	expected := []float64{
		0.017296634932605137,
		0.017296634932604391,
		-1.3,
		-1.2242424242424241,
		5.1111111111111107,
	}

	for i, tt := range kurtosisTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.Kurtosis(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestPopulationKurtosis(t *testing.T) {
	// scipy.stats.kurtosis(data, fisher=True, bias=True)
	expected := []float64{
		0.017296634932605137,
		0.017296634932604391,
		-1.3,
		-1.2242424242424241,
		5.1111111111111107,
	}

	for i, tt := range kurtosisTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.PopulationKurtosis(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestSampleKurtosis(t *testing.T) {
	// pandas .kurt() / scipy.stats.kurtosis(data, fisher=True, bias=False)
	expected := []float64{
		4.0691865397304179,
		4.0691865397304179,
		-1.2,
		-1.2,
		10.0,
	}

	for i, tt := range kurtosisTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stats.SampleKurtosis(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !close(got, expected[i]) {
				t.Errorf("got %v, want %v", got, expected[i])
			}
		})
	}
}

func TestKurtosis_Methods(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 5}

	got, err := data.Kurtosis()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !close(got, -1.3) {
		t.Errorf("Kurtosis got %v, want %v", got, -1.3)
	}

	got, err = data.PopulationKurtosis()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !close(got, -1.3) {
		t.Errorf("PopulationKurtosis got %v, want %v", got, -1.3)
	}

	got, err = data.SampleKurtosis()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !close(got, -1.2) {
		t.Errorf("SampleKurtosis got %v, want %v", got, -1.2)
	}
}

func TestKurtosis_ZeroVariance(t *testing.T) {
	for _, fn := range []func(stats.Float64Data) (float64, error){
		stats.Kurtosis, stats.PopulationKurtosis, stats.SampleKurtosis,
	} {
		_, err := fn(stats.Float64Data{5, 5, 5, 5})
		if err != stats.ErrZero {
			t.Fatalf("expected ErrZero for zero variance, got %v", err)
		}
	}
}

func TestKurtosis_EmptyInput(t *testing.T) {
	for _, fn := range []func(stats.Float64Data) (float64, error){
		stats.Kurtosis, stats.PopulationKurtosis, stats.SampleKurtosis,
	} {
		_, err := fn(stats.Float64Data{})
		if err != stats.ErrEmptyInput {
			t.Fatalf("expected ErrEmptyInput for empty input, got %v", err)
		}
	}
}

func TestKurtosis_TooSmallDataset(t *testing.T) {
	_, err := stats.Kurtosis(stats.Float64Data{42})
	if err != stats.ErrEmptyInput {
		t.Fatalf("expected ErrEmptyInput for dataset length < 2, got %v", err)
	}
}

func TestSampleKurtosis_TooSmallDataset(t *testing.T) {
	_, err := stats.SampleKurtosis(stats.Float64Data{1, 2, 3})
	if err != stats.ErrEmptyInput {
		t.Fatalf("expected ErrEmptyInput for dataset length < 4, got %v", err)
	}
}

func ExampleKurtosis() {
	k, _ := stats.Kurtosis([]float64{1, 2, 3, 4, 5})
	fmt.Println(k)
	// Output: -1.3
}

func ExampleSampleKurtosis() {
	k, _ := stats.SampleKurtosis([]float64{1, 2, 3, 4, 5})
	fmt.Printf("%.1f\n", k)
	// Output: -1.2
}
