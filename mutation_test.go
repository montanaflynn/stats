package stats_test

import (
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

// TestInputsNotMutated verifies that every exported function taking data
// slices leaves its inputs unchanged, so callers never see their data
// modified as a side effect. Each call gets a fresh unsorted slice a and a
// strictly increasing slice b (second sample, weights, or coordinates).
func TestInputsNotMutated(t *testing.T) {
	pcts := []float64{25, 50, 75}
	for _, c := range []struct {
		name string
		call func(a, b stats.Float64Data) error
	}{
		{"ArgMax", func(a, b stats.Float64Data) error { _, err := stats.ArgMax(a); return err }},
		{"ArgMin", func(a, b stats.Float64Data) error { _, err := stats.ArgMin(a); return err }},
		{"AutoCorrelation", func(a, b stats.Float64Data) error { _, err := stats.AutoCorrelation(a, 1); return err }},
		{"ChebyshevDistance", func(a, b stats.Float64Data) error { _, err := stats.ChebyshevDistance(a, b); return err }},
		{"Clip", func(a, b stats.Float64Data) error { _, err := stats.Clip(a, 2, 8); return err }},
		{"CoefficientOfVariation", func(a, b stats.Float64Data) error { _, err := stats.CoefficientOfVariation(a); return err }},
		{"Correlation", func(a, b stats.Float64Data) error { _, err := stats.Correlation(a, b); return err }},
		{"Covariance", func(a, b stats.Float64Data) error { _, err := stats.Covariance(a, b); return err }},
		{"CovariancePopulation", func(a, b stats.Float64Data) error { _, err := stats.CovariancePopulation(a, b); return err }},
		{"CumulativeMax", func(a, b stats.Float64Data) error { _, err := stats.CumulativeMax(a); return err }},
		{"CumulativeMin", func(a, b stats.Float64Data) error { _, err := stats.CumulativeMin(a); return err }},
		{"CumulativeProduct", func(a, b stats.Float64Data) error { _, err := stats.CumulativeProduct(a); return err }},
		{"CumulativeSum", func(a, b stats.Float64Data) error { _, err := stats.CumulativeSum(a); return err }},
		{"Describe", func(a, b stats.Float64Data) error { _, err := stats.Describe(a, false, &pcts); return err }},
		{"DescribePercentileFunc", func(a, b stats.Float64Data) error {
			_, err := stats.DescribePercentileFunc(a, false, &pcts, stats.Percentile)
			return err
		}},
		{"Diff", func(a, b stats.Float64Data) error { _, err := stats.Diff(a); return err }},
		{"Entropy", func(a, b stats.Float64Data) error { _, err := stats.Entropy(a); return err }},
		{"EuclideanDistance", func(a, b stats.Float64Data) error { _, err := stats.EuclideanDistance(a, b); return err }},
		{"EWMA", func(a, b stats.Float64Data) error { _, err := stats.EWMA(a, 0.5); return err }},
		{"GeometricMean", func(a, b stats.Float64Data) error { _, err := stats.GeometricMean(a); return err }},
		{"HarmonicMean", func(a, b stats.Float64Data) error { _, err := stats.HarmonicMean(a); return err }},
		{"Histogram", func(a, b stats.Float64Data) error { _, _, err := stats.Histogram(a, 3); return err }},
		{"Interp", func(a, b stats.Float64Data) error { _, err := stats.Interp(a, b, b); return err }},
		{"InterQuartileRange", func(a, b stats.Float64Data) error { _, err := stats.InterQuartileRange(a); return err }},
		{"KendallTau", func(a, b stats.Float64Data) error { _, err := stats.KendallTau(a, b); return err }},
		{"Kurtosis", func(a, b stats.Float64Data) error { _, err := stats.Kurtosis(a); return err }},
		{"ManhattanDistance", func(a, b stats.Float64Data) error { _, err := stats.ManhattanDistance(a, b); return err }},
		{"Max", func(a, b stats.Float64Data) error { _, err := stats.Max(a); return err }},
		{"Mean", func(a, b stats.Float64Data) error { _, err := stats.Mean(a); return err }},
		{"Median", func(a, b stats.Float64Data) error { _, err := stats.Median(a); return err }},
		{"MedianAbsoluteDeviation", func(a, b stats.Float64Data) error { _, err := stats.MedianAbsoluteDeviation(a); return err }},
		{"MedianAbsoluteDeviationPopulation", func(a, b stats.Float64Data) error {
			_, err := stats.MedianAbsoluteDeviationPopulation(a)
			return err
		}},
		{"Midhinge", func(a, b stats.Float64Data) error { _, err := stats.Midhinge(a); return err }},
		{"Min", func(a, b stats.Float64Data) error { _, err := stats.Min(a); return err }},
		{"MinkowskiDistance", func(a, b stats.Float64Data) error { _, err := stats.MinkowskiDistance(a, b, 3); return err }},
		{"Mode", func(a, b stats.Float64Data) error { _, err := stats.Mode(a); return err }},
		{"MovingAverage", func(a, b stats.Float64Data) error { _, err := stats.MovingAverage(a, 3); return err }},
		{"MovingMax", func(a, b stats.Float64Data) error { _, err := stats.MovingMax(a, 3); return err }},
		{"MovingMedian", func(a, b stats.Float64Data) error { _, err := stats.MovingMedian(a, 3); return err }},
		{"MovingMin", func(a, b stats.Float64Data) error { _, err := stats.MovingMin(a, 3); return err }},
		{"MovingStdDev", func(a, b stats.Float64Data) error { _, err := stats.MovingStdDev(a, 3); return err }},
		{"MovingSum", func(a, b stats.Float64Data) error { _, err := stats.MovingSum(a, 3); return err }},
		{"NormFit", func(a, b stats.Float64Data) error { stats.NormFit(a); return nil }},
		{"Pearson", func(a, b stats.Float64Data) error { _, err := stats.Pearson(a, b); return err }},
		{"PercentChange", func(a, b stats.Float64Data) error { _, err := stats.PercentChange(a); return err }},
		{"Percentile", func(a, b stats.Float64Data) error { _, err := stats.Percentile(a, 50); return err }},
		{"PercentileNearestRank", func(a, b stats.Float64Data) error { _, err := stats.PercentileNearestRank(a, 50); return err }},
		{"PercentileOfScore", func(a, b stats.Float64Data) error { _, err := stats.PercentileOfScore(a, 3); return err }},
		{"PercentileWeighted", func(a, b stats.Float64Data) error { _, err := stats.PercentileWeighted(a, b, 50); return err }},
		{"PopulationKurtosis", func(a, b stats.Float64Data) error { _, err := stats.PopulationKurtosis(a); return err }},
		{"PopulationSkewness", func(a, b stats.Float64Data) error { _, err := stats.PopulationSkewness(a); return err }},
		{"PopulationVariance", func(a, b stats.Float64Data) error { _, err := stats.PopulationVariance(a); return err }},
		{"Product", func(a, b stats.Float64Data) error { _, err := stats.Product(a); return err }},
		{"Quartile", func(a, b stats.Float64Data) error { _, err := stats.Quartile(a); return err }},
		{"QuartileOutliers", func(a, b stats.Float64Data) error { _, err := stats.QuartileOutliers(a); return err }},
		{"Range", func(a, b stats.Float64Data) error { _, err := stats.Range(a); return err }},
		{"Rank", func(a, b stats.Float64Data) error { _, err := stats.Rank(a); return err }},
		{"Rescale", func(a, b stats.Float64Data) error { _, err := stats.Rescale(a); return err }},
		{"RMS", func(a, b stats.Float64Data) error { _, err := stats.RMS(a); return err }},
		{"Sample", func(a, b stats.Float64Data) error { _, err := stats.Sample(a, 4, true); return err }},
		{"SampleKurtosis", func(a, b stats.Float64Data) error { _, err := stats.SampleKurtosis(a); return err }},
		{"SampleSkewness", func(a, b stats.Float64Data) error { _, err := stats.SampleSkewness(a); return err }},
		{"SampleVariance", func(a, b stats.Float64Data) error { _, err := stats.SampleVariance(a); return err }},
		{"SEM", func(a, b stats.Float64Data) error { _, err := stats.SEM(a); return err }},
		{"Sigmoid", func(a, b stats.Float64Data) error { _, err := stats.Sigmoid(a); return err }},
		{"Skewness", func(a, b stats.Float64Data) error { _, err := stats.Skewness(a); return err }},
		{"SoftMax", func(a, b stats.Float64Data) error { _, err := stats.SoftMax(a); return err }},
		{"Spearman", func(a, b stats.Float64Data) error { _, err := stats.Spearman(a, b); return err }},
		{"StableSample", func(a, b stats.Float64Data) error { _, err := stats.StableSample(a, 4); return err }},
		{"StandardDeviation", func(a, b stats.Float64Data) error { _, err := stats.StandardDeviation(a); return err }},
		{"StandardDeviationPopulation", func(a, b stats.Float64Data) error {
			_, err := stats.StandardDeviationPopulation(a)
			return err
		}},
		{"StandardDeviationSample", func(a, b stats.Float64Data) error { _, err := stats.StandardDeviationSample(a); return err }},
		{"StdDevP", func(a, b stats.Float64Data) error { _, err := stats.StdDevP(a); return err }},
		{"StdDevS", func(a, b stats.Float64Data) error { _, err := stats.StdDevS(a); return err }},
		{"Sum", func(a, b stats.Float64Data) error { _, err := stats.Sum(a); return err }},
		{"Trimean", func(a, b stats.Float64Data) error { _, err := stats.Trimean(a); return err }},
		{"TrimmedMean", func(a, b stats.Float64Data) error { _, err := stats.TrimmedMean(a, 0.1); return err }},
		{"TTest", func(a, b stats.Float64Data) error { _, _, err := stats.TTest(a, b, 0); return err }},
		{"Variance", func(a, b stats.Float64Data) error { _, err := stats.Variance(a); return err }},
		{"VarP", func(a, b stats.Float64Data) error { _, err := stats.VarP(a); return err }},
		{"VarS", func(a, b stats.Float64Data) error { _, err := stats.VarS(a); return err }},
		{"WeightedMean", func(a, b stats.Float64Data) error { _, err := stats.WeightedMean(a, b); return err }},
		{"Winsorize", func(a, b stats.Float64Data) error { _, err := stats.Winsorize(a, 0.1); return err }},
		{"ZScore", func(a, b stats.Float64Data) error { _, err := stats.ZScore(a); return err }},
		{"ZTest", func(a, b stats.Float64Data) error { _, _, err := stats.ZTest(a, b, 0, 1); return err }},
	} {
		a := stats.Float64Data{3.2, 1.1, 4.8, 2.5, 9.0, 7.3, 0.4, 5.5}
		b := stats.Float64Data{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}
		if err := c.call(a, b); err != nil {
			t.Errorf("%s returned an error: %v", c.name, err)
			continue
		}
		wantA := stats.Float64Data{3.2, 1.1, 4.8, 2.5, 9.0, 7.3, 0.4, 5.5}
		wantB := stats.Float64Data{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}
		if !reflect.DeepEqual(a, wantA) {
			t.Errorf("%s mutated its input: %v != %v", c.name, a, wantA)
		}
		if !reflect.DeepEqual(b, wantB) {
			t.Errorf("%s mutated its second input: %v != %v", c.name, b, wantB)
		}
	}
}
