package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestDescribeValidDataset(t *testing.T) {
	_, err := stats.Describe([]float64{1.0, 2.0, 3.0}, false, &[]float64{25.0, 50.0, 75.0})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestDescribeEmptyDataset(t *testing.T) {
	_, err := stats.Describe([]float64{}, false, nil)
	if err != stats.ErrEmptyInput {
		t.Errorf("Did not return empty input error")
	}
}

func TestDescribeEmptyDatasetNaN(t *testing.T) {
	describe, err := stats.Describe([]float64{}, true, nil)
	if err != nil {
		t.Errorf("Returned an error")
	}

	if !math.IsNaN(describe.Max) || !math.IsNaN(describe.Mean) || !math.IsNaN(describe.Min) || !math.IsNaN(describe.Std) {
		t.Errorf("Was not NaN")
	}
}

func TestDescribeValidDatasetNaN(t *testing.T) {
	describe, err := stats.Describe([]float64{1.0, 2.0, 3.0}, true, &[]float64{25.0, 50.0, 75.0})
	if err != nil {
		t.Errorf("Returned an error")
	}

	if math.IsNaN(describe.Max) {
		t.Errorf("Was NaN")
	}
}

func TestDescribeValues(t *testing.T) {
	dataSet := []float64{1.0, 2.0, 3.0}
	percentiles := []float64{25.0, 50.0, 75.0}
	describe, _ := stats.Describe(dataSet, true, &percentiles)
	if describe.Count != len(dataSet) {
		t.Errorf("Count was not == length of dataset")
	}
	if len(describe.DescriptionPercentiles) != len(percentiles) {
		t.Errorf("Percentiles length was not == length of input percentiles")
	}

	max, _ := stats.Max(dataSet)
	if max != describe.Max {
		t.Errorf("Max was not equal to Max(dataset)")
	}

	min, _ := stats.Min(dataSet)
	if min != describe.Min {
		t.Errorf("Min was not equal to Min(dataset)")
	}

	mean, _ := stats.Mean(dataSet)
	if mean != describe.Mean {
		t.Errorf("Mean was not equal to Mean(dataset)")
	}

	std, _ := stats.StandardDeviation(dataSet)
	if std != describe.Std {
		t.Errorf("Std was not equal to StandardDeviation(dataset)")
	}
}

func TestDescribeString(t *testing.T) {
	describe, _ := stats.Describe([]float64{1.0, 2.0, 3.0}, true, &[]float64{25.0, 50.0, 75.0})
	if describe.String(2) != "count\t3\nmean\t2.00\nstd\t0.82\nmax\t3.00\nmin\t1.00\n25.00%\tNaN\n50.00%\t1.50\n75.00%\t2.50\nNaN OK\ttrue" {
		t.Errorf("String output is not correct")
	}
}
