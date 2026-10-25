package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

// TestOrderStatisticsPropagateNaN checks the exact results of functions
// that sort or compare their input when that input contains a NaN, with
// the NaN first, in the middle, and last.
func TestOrderStatisticsPropagateNaN(t *testing.T) {
	nan := math.NaN()
	y := stats.Float64Data{1, 2, 3, 4}
	for _, d := range []stats.Float64Data{
		{nan, 1, 3, 4},
		{1, nan, 3, 4},
		{1, 3, 4, nan},
	} {
		for _, c := range []struct {
			name string
			call func() (float64, error)
		}{
			{"Max", func() (float64, error) { return stats.Max(d) }},
			{"Min", func() (float64, error) { return stats.Min(d) }},
			{"Range", func() (float64, error) { return stats.Range(d) }},
			{"Median", func() (float64, error) { return stats.Median(d) }},
			{"MedianAbsoluteDeviation", func() (float64, error) { return stats.MedianAbsoluteDeviation(d) }},
			{"Percentile", func() (float64, error) { return stats.Percentile(d, 50) }},
			{"PercentileNearestRank", func() (float64, error) { return stats.PercentileNearestRank(d, 50) }},
			{"PercentileOfScore", func() (float64, error) { return stats.PercentileOfScore(d, 2) }},
			{"ChebyshevDistance", func() (float64, error) { return stats.ChebyshevDistance(d, y) }},
			{"ChebyshevDistance swapped", func() (float64, error) { return stats.ChebyshevDistance(y, d) }},
			{"Spearman", func() (float64, error) { return stats.Spearman(d, y) }},
			{"Spearman swapped", func() (float64, error) { return stats.Spearman(y, d) }},
			{"KendallTau", func() (float64, error) { return stats.KendallTau(d, y) }},
			{"KendallTau swapped", func() (float64, error) { return stats.KendallTau(y, d) }},
		} {
			got, err := c.call()
			if err != nil || !math.IsNaN(got) {
				t.Errorf("%s(%v) = %v, %v; want NaN, nil", c.name, d, got, err)
			}
		}

		if got, err := stats.Rank(d); err != nil || len(got) != len(d) || !allNaN(got) {
			t.Errorf("Rank(%v) = %v, %v; want all NaN", d, got, err)
		}
		if _, _, err := stats.Histogram(d, 2); err != stats.ErrNaN {
			t.Errorf("Histogram(%v) error = %v; want ErrNaN", d, err)
		}
	}
}

// TestArgExtremesNaN checks that ArgMax and ArgMin return the index of the
// first NaN, matching numpy.argmax and numpy.argmin.
func TestArgExtremesNaN(t *testing.T) {
	nan := math.NaN()
	for _, c := range []struct {
		data stats.Float64Data
		want int
	}{
		{stats.Float64Data{nan, 1, 3}, 0},
		{stats.Float64Data{1, nan, 3}, 1},
		{stats.Float64Data{1, 3, nan}, 2},
		{stats.Float64Data{1, nan, 3, nan}, 1},
	} {
		if i, err := stats.ArgMax(c.data); err != nil || i != c.want {
			t.Errorf("ArgMax(%v) = %d, %v; want %d, nil", c.data, i, err, c.want)
		}
		if i, err := stats.ArgMin(c.data); err != nil || i != c.want {
			t.Errorf("ArgMin(%v) = %d, %v; want %d, nil", c.data, i, err, c.want)
		}
	}
}

// TestCumulativeExtremesPropagateNaN checks that every value from the first
// NaN onward is NaN, like NumPy's maximum.accumulate and minimum.accumulate.
func TestCumulativeExtremesPropagateNaN(t *testing.T) {
	nan := math.NaN()
	d := stats.Float64Data{2, 1, nan, 3, 0}
	for _, c := range []struct {
		name string
		call func(stats.Float64Data) ([]float64, error)
		want []float64
	}{
		{"CumulativeMax", stats.CumulativeMax, []float64{2, 2, nan, nan, nan}},
		{"CumulativeMin", stats.CumulativeMin, []float64{2, 1, nan, nan, nan}},
	} {
		got, err := c.call(d)
		if err != nil || len(got) != len(c.want) {
			t.Errorf("%s(%v) = %v, %v; want %v", c.name, d, got, err, c.want)
			continue
		}
		for i := range got {
			if got[i] != c.want[i] && !(math.IsNaN(got[i]) && math.IsNaN(c.want[i])) {
				t.Errorf("%s(%v) = %v; want %v", c.name, d, got, c.want)
				break
			}
		}
	}
}
