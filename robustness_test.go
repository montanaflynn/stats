package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

// robustCase adapts an exported function to a common shape so the property
// tests below can sweep the whole API. a is the primary data, b a second
// sample, weights, or coordinates of the same length.
type robustCase struct {
	name string
	call func(a, b stats.Float64Data) ([]float64, error)
	// scale gives the factor k(s) with f(s*a, s*b) = k(s) * f(a, b). nil
	// means the function is not homogeneous and the scaling test skips it.
	scale func(s float64) float64
}

// pow is s^d, for results like the mean or variance.
func pow(d int) func(float64) float64 {
	return func(s float64) float64 { return math.Pow(s, float64(d)) }
}

// absPow is |s|^d, for non-negative results like distances.
func absPow(d int) func(float64) float64 {
	return func(s float64) float64 { return math.Pow(math.Abs(s), float64(d)) }
}

// sign is the sign of s, for odd scale-free results like skewness.
func sign(s float64) float64 { return math.Copysign(1, s) }

func one(v float64, err error) ([]float64, error) { return []float64{v}, err }
func two(v, w float64, err error) ([]float64, error) {
	return []float64{v, w}, err
}

func robustCases() []robustCase {
	pcts := []float64{10, 25, 50, 75, 90}
	return []robustCase{
		{"AutoCorrelation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.AutoCorrelation(a, 1)) }, pow(0)},
		{"ChebyshevDistance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.ChebyshevDistance(a, b)) }, absPow(1)},
		{"CoefficientOfVariation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.CoefficientOfVariation(a)) }, sign},
		{"Correlation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Correlation(a, b)) }, pow(0)},
		{"Covariance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Covariance(a, b)) }, pow(2)},
		{"CovariancePopulation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.CovariancePopulation(a, b)) }, pow(2)},
		{"CumulativeMax", func(a, b stats.Float64Data) ([]float64, error) { return stats.CumulativeMax(a) }, nil},
		{"CumulativeMin", func(a, b stats.Float64Data) ([]float64, error) { return stats.CumulativeMin(a) }, nil},
		{"CumulativeProduct", func(a, b stats.Float64Data) ([]float64, error) { return stats.CumulativeProduct(a) }, nil},
		{"CumulativeSum", func(a, b stats.Float64Data) ([]float64, error) { return stats.CumulativeSum(a) }, pow(1)},
		{"Describe", func(a, b stats.Float64Data) ([]float64, error) {
			d, err := stats.Describe(a, false, &pcts)
			if err != nil {
				return nil, err
			}
			out := []float64{d.Mean, d.Std, d.Max, d.Min}
			for _, p := range d.DescriptionPercentiles {
				out = append(out, p.Value)
			}
			return out, nil
		}, nil},
		{"Diff", func(a, b stats.Float64Data) ([]float64, error) { return stats.Diff(a) }, pow(1)},
		{"EuclideanDistance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.EuclideanDistance(a, b)) }, absPow(1)},
		{"EWMA", func(a, b stats.Float64Data) ([]float64, error) { return stats.EWMA(a, 0.5) }, pow(1)},
		{"Histogram", func(a, b stats.Float64Data) ([]float64, error) {
			_, edges, err := stats.Histogram(a, 3)
			return edges, err
		}, nil},
		{"Interp", func(a, b stats.Float64Data) ([]float64, error) { return stats.Interp(a, b, b) }, nil},
		{"InterQuartileRange", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.InterQuartileRange(a)) }, absPow(1)},
		{"KendallTau", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.KendallTau(a, b)) }, pow(0)},
		{"Kurtosis", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Kurtosis(a)) }, pow(0)},
		{"ManhattanDistance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.ManhattanDistance(a, b)) }, absPow(1)},
		{"Max", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Max(a)) }, nil},
		{"Mean", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Mean(a)) }, pow(1)},
		{"Median", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Median(a)) }, pow(1)},
		{"MedianAbsoluteDeviation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.MedianAbsoluteDeviation(a)) }, absPow(1)},
		{"Midhinge", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Midhinge(a)) }, pow(1)},
		{"Min", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Min(a)) }, nil},
		{"MinkowskiDistance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.MinkowskiDistance(a, b, 3)) }, absPow(1)},
		{"Mode", func(a, b stats.Float64Data) ([]float64, error) { return stats.Mode(a) }, nil},
		{"MovingAverage", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingAverage(a, 2) }, pow(1)},
		{"MovingMax", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingMax(a, 2) }, nil},
		{"MovingMedian", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingMedian(a, 2) }, pow(1)},
		{"MovingMin", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingMin(a, 2) }, nil},
		{"MovingStdDev", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingStdDev(a, 2) }, absPow(1)},
		{"MovingSum", func(a, b stats.Float64Data) ([]float64, error) { return stats.MovingSum(a, 2) }, pow(1)},
		{"NormFit", func(a, b stats.Float64Data) ([]float64, error) {
			if len(a) == 0 {
				return nil, stats.ErrEmptyInput
			}
			f := stats.NormFit(a)
			return f[:], nil
		}, nil},
		{"Pearson", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Pearson(a, b)) }, pow(0)},
		{"PercentChange", func(a, b stats.Float64Data) ([]float64, error) { return stats.PercentChange(a) }, pow(0)},
		{"Percentile", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Percentile(a, 50)) }, pow(1)},
		{"PercentileNearestRank", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.PercentileNearestRank(a, 30)) }, nil},
		{"PercentileOfScore", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.PercentileOfScore(a, 3)) }, nil},
		{"PercentileWeighted", func(a, b stats.Float64Data) ([]float64, error) {
			w := make(stats.Float64Data, len(a))
			for i := range w {
				w[i] = 1
			}
			return one(stats.PercentileWeighted(a, w, 30))
		}, nil},
		{"PopulationKurtosis", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.PopulationKurtosis(a)) }, pow(0)},
		{"PopulationSkewness", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.PopulationSkewness(a)) }, sign},
		{"PopulationVariance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.PopulationVariance(a)) }, pow(2)},
		{"Product", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Product(a)) }, nil},
		{"Quartile", func(a, b stats.Float64Data) ([]float64, error) {
			q, err := stats.Quartile(a)
			return []float64{q.Q1, q.Q2, q.Q3}, err
		}, nil},
		{"Range", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Range(a)) }, absPow(1)},
		{"Rank", func(a, b stats.Float64Data) ([]float64, error) { return stats.Rank(a) }, nil},
		{"Rescale", func(a, b stats.Float64Data) ([]float64, error) { return stats.Rescale(a) }, nil},
		{"RMS", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.RMS(a)) }, absPow(1)},
		{"SampleKurtosis", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.SampleKurtosis(a)) }, pow(0)},
		{"SampleSkewness", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.SampleSkewness(a)) }, sign},
		{"SampleVariance", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.SampleVariance(a)) }, pow(2)},
		{"SEM", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.SEM(a)) }, absPow(1)},
		{"Sigmoid", func(a, b stats.Float64Data) ([]float64, error) { return stats.Sigmoid(a) }, nil},
		{"Skewness", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Skewness(a)) }, sign},
		{"SoftMax", func(a, b stats.Float64Data) ([]float64, error) { return stats.SoftMax(a) }, nil},
		{"Spearman", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Spearman(a, b)) }, pow(0)},
		{"StandardDeviation", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.StandardDeviation(a)) }, absPow(1)},
		{"StandardDeviationSample", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.StandardDeviationSample(a)) }, absPow(1)},
		{"Sum", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Sum(a)) }, pow(1)},
		{"Trimean", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.Trimean(a)) }, pow(1)},
		{"TrimmedMean", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.TrimmedMean(a, 0.1)) }, pow(1)},
		{"TTest", func(a, b stats.Float64Data) ([]float64, error) { return two(stats.TTest(a, b, 0)) }, nil},
		{"WeightedMean", func(a, b stats.Float64Data) ([]float64, error) { return one(stats.WeightedMean(a, b)) }, nil},
		{"Winsorize", func(a, b stats.Float64Data) ([]float64, error) { return stats.Winsorize(a, 0.1) }, pow(1)},
		{"ZScore", func(a, b stats.Float64Data) ([]float64, error) { return stats.ZScore(a) }, sign},
		{"ZTest", func(a, b stats.Float64Data) ([]float64, error) { return two(stats.ZTest(a, b, 0, 1)) }, nil},
	}
}

// safeCall runs c and converts a panic into an error-like report.
func safeCall(c robustCase, a, b stats.Float64Data) (out []float64, err error, panicked interface{}) {
	defer func() { panicked = recover() }()
	out, err = c.call(append(stats.Float64Data(nil), a...), append(stats.Float64Data(nil), b...))
	return
}

func inputs() map[string][2]stats.Float64Data {
	nan, inf := math.NaN(), math.Inf(1)
	return map[string][2]stats.Float64Data{
		"empty":       {{}, {}},
		"single":      {{3}, {2}},
		"pair":        {{3, 1}, {2, 4}},
		"constant":    {{2, 2, 2, 2}, {1, 2, 3, 4}},
		"zeros":       {{0, 0, 0, 0}, {0, 0, 0, 0}},
		"nan":         {{1, nan, 3, 4}, {1, 2, 3, 4}},
		"allnan":      {{nan, nan, nan}, {nan, nan, nan}},
		"inf":         {{1, inf, 3, 4}, {1, 2, 3, 4}},
		"neginf":      {{1, -inf, 3, 4}, {1, 2, 3, 4}},
		"bothinf":     {{-inf, inf, 3, 4}, {1, 2, 3, 4}},
		"huge":        {{1e308, -1e308, 1e308, 5}, {1, 2, 3, 4}},
		"tiny":        {{5e-324, 1e-320, 0, 2e-323}, {1, 2, 3, 4}},
		"mismatched":  {{1, 2, 3}, {1, 2}},
		"mixedsigns":  {{-3, 1, -4, 1, 5, -9}, {1, 2, 3, 4, 5, 6}},
		"maxfloat":    {{math.MaxFloat64, math.MaxFloat64}, {1, 2}},
		"negzero":     {{math.Copysign(0, -1), 0, 1}, {1, 2, 3}},
		"largeoffset": {{1e9 + 1, 1e9 + 2, 1e9 + 3}, {1, 2, 3}},
		// Different lengths with a single first element, for two-sample tests.
		"unevensingle": {{3}, {2, 4}},
	}
}

// TestNoPanics feeds every function adversarial inputs and fails on any
// panic. Functions must report bad input through their error return.
func TestNoPanics(t *testing.T) {
	for _, c := range robustCases() {
		for name, in := range inputs() {
			if _, _, p := safeCall(c, in[0], in[1]); p != nil {
				t.Errorf("%s(%s) panicked: %v", c.name, name, p)
			}
		}
	}
}

// knownIssues tracks property failures that are found but not yet fixed. A
// listed failure is logged instead of failing the test, and a listed key that
// no longer fails is an error, so each fix must also remove its entry.
type knownIssues struct {
	t     *testing.T
	known map[string]bool
	seen  map[string]bool
}

func newKnownIssues(t *testing.T, keys ...string) *knownIssues {
	k := &knownIssues{t: t, known: map[string]bool{}, seen: map[string]bool{}}
	for _, key := range keys {
		k.known[key] = true
	}
	return k
}

// failf fails the test for key unless key is a known issue.
func (k *knownIssues) failf(key, format string, args ...interface{}) {
	k.t.Helper()
	if k.known[key] {
		k.seen[key] = true
		k.t.Logf("known issue: "+format, args...)
		return
	}
	k.t.Errorf(format, args...)
}

// done fails the test for known issues that did not reproduce.
func (k *knownIssues) done() {
	k.t.Helper()
	for key := range k.known {
		if !k.seen[key] {
			k.t.Errorf("%s no longer fails; remove it from the known issues", key)
		}
	}
}

func hasNaN(xs []float64) bool {
	for _, x := range xs {
		if math.IsNaN(x) {
			return true
		}
	}
	return false
}

// TestFiniteInputNoSilentNaN checks that finite inputs never produce a NaN
// result alongside a nil error. A NaN result must come with an error so
// callers that only check err are not handed garbage.
func TestFiniteInputNoSilentNaN(t *testing.T) {
	known := newKnownIssues(t,
		// Sample statistics divide by n-1, which is 0 for one element. numpy,
		// scipy and pandas also return NaN here.
		"CoefficientOfVariation(single)",
		"Covariance(single)",
		"SampleVariance(single)",
		"SEM(single)",
		"StandardDeviationSample(single)",
		"ZScore(single)",
		"CoefficientOfVariation(unevensingle)",
		"SampleVariance(unevensingle)",
		"SEM(unevensingle)",
		"StandardDeviationSample(unevensingle)",
		"ZScore(unevensingle)",
		// The mean or the squared deviations overflow for huge values, as they
		// do in numpy and pandas.
		"AutoCorrelation(huge)",
		"AutoCorrelation(maxfloat)",
		"CoefficientOfVariation(maxfloat)",
		"Correlation(maxfloat)",
		"Covariance(maxfloat)",
		"CovariancePopulation(maxfloat)",
		"Kurtosis(huge)",
		"Kurtosis(maxfloat)",
		"Pearson(maxfloat)",
		"PopulationKurtosis(huge)",
		"PopulationKurtosis(maxfloat)",
		"PopulationSkewness(huge)",
		"PopulationSkewness(maxfloat)",
		"SampleKurtosis(huge)",
		"SampleSkewness(huge)",
		"Skewness(huge)",
		"Skewness(maxfloat)",
		"TTest(maxfloat)",
		"ZScore(maxfloat)",
		// max - min overflows for values of opposite sign. numpy.histogram
		// raises an error for this range instead.
		"Histogram(huge)",
		"Rescale(huge)",
		// The two-sample t-test divides by a zero pooled variance, returning
		// NaN like scipy.stats.ttest_ind.
		"TTest(zeros)",
		// The two-sample t-test takes the sample variance of a single
		// element, which is NaN. scipy.stats.ttest_ind gives a finite result
		// because the pooled variance has n1+n2-2 = 1 degree of freedom.
		"TTest(unevensingle)",
	)
	defer known.done()
	finite := []string{"single", "pair", "constant", "zeros", "huge", "tiny", "mixedsigns", "maxfloat", "negzero", "largeoffset", "unevensingle"}
	all := inputs()
	for _, c := range robustCases() {
		// PercentChange documents IEEE 754 NaN for 0/0, like pandas.
		if c.name == "PercentChange" {
			continue
		}
		for _, name := range finite {
			in := all[name]
			out, err, p := safeCall(c, in[0], in[1])
			if p == nil && err == nil && hasNaN(out) {
				key := c.name + "(" + name + ")"
				known.failf(key, "%s = %v with nil error", key, out)
			}
		}
	}
}

// TestNaNInputPropagates checks that a NaN in the data yields a NaN result
// or an error, as Mean and Sum do, rather than a finite result computed as if
// the NaN were absent or sorted below every number.
func TestNaNInputPropagates(t *testing.T) {
	known := newKnownIssues(t,
		// Comparisons with NaN are false, so NaN is skipped by max and min
		// scans and sorted below every number. numpy and scipy return NaN,
		// and numpy.histogram raises.
		"ChebyshevDistance",
		"CumulativeMax",
		"CumulativeMin",
		"Histogram",
		"KendallTau",
		"Max",
		"Median",
		"MedianAbsoluteDeviation",
		"Min",
		"Percentile",
		"PercentileNearestRank",
		"PercentileOfScore",
		"Range",
		"Rank",
		"Spearman",
	)
	defer known.done()
	in := inputs()["nan"]
	for _, c := range robustCases() {
		// Mode reports no repeated value, which is true with a NaN.
		if c.name == "Mode" {
			continue
		}
		out, err, p := safeCall(c, in[0], in[1])
		if p == nil && err == nil && !hasNaN(out) {
			known.failf(c.name, "%s(%v) = %v with nil error", c.name, in[0], out)
		}
	}
}

// TestScaleEquivariance checks f(s*x) = s^d * f(x) for scales that keep
// every input and result representable. Intermediate overflow or underflow
// (squaring 1e200, say) shows up as Inf, 0, or NaN at the extreme scales.
func TestScaleEquivariance(t *testing.T) {
	known := newKnownIssues(t,
		// Correlation divides by the product of overflowing standard
		// deviations. scipy.stats.pearsonr normalizes first and stays exact.
		"Correlation",
		"Pearson",
		// Squared deviations overflow or underflow at 1e±200, as they do in
		// numpy and scipy.
		"AutoCorrelation",
		"CoefficientOfVariation",
		"EuclideanDistance",
		"MovingStdDev",
		"RMS",
		"SEM",
		"StandardDeviation",
		"StandardDeviationSample",
		"ZScore",
		// Cubed and fourth-power deviations overflow or underflow from 1e±100,
		// as they do in scipy, except skewness that underflows to 0 rather
		// than NaN.
		"Kurtosis",
		"PopulationKurtosis",
		"PopulationSkewness",
		"SampleKurtosis",
		"SampleSkewness",
		"Skewness",
		// |x-y|^lambda overflows or underflows, as in scipy.
		"MinkowskiDistance",
	)
	defer known.done()
	a := stats.Float64Data{3.2, 1.1, 4.8, 2.5, 9.0, 7.3, 0.4, 5.5}
	b := stats.Float64Data{1.5, 2.7, 3.1, 4.9, 5.2, 6.8, 7.4, 8.6}
	for _, c := range robustCases() {
		if c.scale == nil {
			continue
		}
		base, err := c.call(a, b)
		if err != nil {
			t.Errorf("%s base call: %v", c.name, err)
			continue
		}
		for _, s := range []float64{-1, 1e100, 1e-100, 1e150, 1e-150, 1e200, 1e-200} {
			// Skip scales whose expected result is not representable.
			factor := c.scale(s)
			if f := math.Abs(math.Log10(math.Abs(factor))); f > 300 {
				continue
			}
			sa, sb := make(stats.Float64Data, len(a)), make(stats.Float64Data, len(b))
			for i := range a {
				sa[i], sb[i] = a[i]*s, b[i]*s
			}
			got, err := c.call(sa, sb)
			if err != nil {
				known.failf(c.name, "%s at scale %g: %v", c.name, s, err)
				continue
			}
			for i := range base {
				want := base[i] * factor
				if i >= len(got) || !tolerance(got[i], want, 1e-9) {
					known.failf(c.name, "%s at scale %g: got %v, want %v", c.name, s, got, scaled(base, factor))
					break
				}
			}
		}
	}
}

func scaled(xs []float64, f float64) []float64 {
	out := make([]float64, len(xs))
	for i, x := range xs {
		out[i] = x * f
	}
	return out
}

// TestInvalidParameters checks that a NaN or out-of-domain scalar parameter
// is rejected with an error, or at least propagates as NaN, rather than
// producing a plausible finite result with a nil error.
func TestInvalidParameters(t *testing.T) {
	known := newKnownIssues(t,
		// Range checks like min > max are false for NaN, so NaN slips past.
		// numpy.clip and scipy.stats.percentileofscore return NaN.
		"Clip(max=NaN)",
		"Clip(min=NaN)",
		"EWMA(alpha=NaN)",
		"PercentileOfScore(NaN)",
		// Out-of-domain values are not validated. scipy raises for p <= 0 in
		// minkowski and returns NaN for p outside [0, 1] in geom.
		"MinkowskiDistance(lambda=-1)",
		"ProbGeom(p=-0.5)",
	)
	defer known.done()
	a := stats.Float64Data{3.2, 1.1, 4.8, 2.5, 9.0, 7.3, 0.4, 5.5}
	b := stats.Float64Data{1.5, 2.7, 3.1, 4.9, 5.2, 6.8, 7.4, 8.6}
	nan := math.NaN()
	for _, c := range []struct {
		name string
		call func() ([]float64, error)
	}{
		{"Clip(min=NaN)", func() ([]float64, error) { return stats.Clip(a, nan, 8) }},
		{"Clip(max=NaN)", func() ([]float64, error) { return stats.Clip(a, 2, nan) }},
		{"EWMA(alpha=NaN)", func() ([]float64, error) { return stats.EWMA(a, nan) }},
		{"ExpGeom(p=NaN)", func() ([]float64, error) { return one(stats.ExpGeom(nan)) }},
		{"MinkowskiDistance(lambda=NaN)", func() ([]float64, error) { return one(stats.MinkowskiDistance(a, b, nan)) }},
		{"MinkowskiDistance(lambda=0)", func() ([]float64, error) { return one(stats.MinkowskiDistance(a, b, 0)) }},
		{"MinkowskiDistance(lambda=-1)", func() ([]float64, error) { return one(stats.MinkowskiDistance(a, b, -1)) }},
		{"Percentile(NaN)", func() ([]float64, error) { return one(stats.Percentile(a, nan)) }},
		{"PercentileNearestRank(NaN)", func() ([]float64, error) { return one(stats.PercentileNearestRank(a, nan)) }},
		{"PercentileOfScore(NaN)", func() ([]float64, error) { return one(stats.PercentileOfScore(a, nan)) }},
		{"PercentileWeighted(NaN)", func() ([]float64, error) { return one(stats.PercentileWeighted(a, b, nan)) }},
		{"ProbGeom(p=NaN)", func() ([]float64, error) { return one(stats.ProbGeom(1, 3, nan)) }},
		{"ProbGeom(p=1.5)", func() ([]float64, error) { return one(stats.ProbGeom(1, 3, 1.5)) }},
		{"ProbGeom(p=-0.5)", func() ([]float64, error) { return one(stats.ProbGeom(1, 3, -0.5)) }},
		{"Round(NaN)", func() ([]float64, error) { return one(stats.Round(nan, 2)) }},
		{"TrimmedMean(NaN)", func() ([]float64, error) { return one(stats.TrimmedMean(a, nan)) }},
		{"TTest(mean=NaN)", func() ([]float64, error) { return two(stats.TTest(a, nil, nan)) }},
		{"VarGeom(p=NaN)", func() ([]float64, error) { return one(stats.VarGeom(nan)) }},
		{"Winsorize(NaN)", func() ([]float64, error) { return stats.Winsorize(a, nan) }},
		{"ZTest(mean=NaN)", func() ([]float64, error) { return two(stats.ZTest(a, nil, nan, 1)) }},
		{"ZTest(std=NaN)", func() ([]float64, error) { return two(stats.ZTest(a, nil, 0, nan)) }},
		{"ZTest(std=0)", func() ([]float64, error) { return two(stats.ZTest(a, nil, 0, 0)) }},
		{"ZTest(std=-1)", func() ([]float64, error) { return two(stats.ZTest(a, nil, 0, -1)) }},
		{"WeightedMean(weight=NaN)", func() ([]float64, error) {
			return one(stats.WeightedMean(stats.Float64Data{1, 2, 3}, stats.Float64Data{1, nan, 1}))
		}},
	} {
		func() {
			defer func() {
				if p := recover(); p != nil {
					t.Errorf("%s panicked: %v", c.name, p)
				}
			}()
			out, err := c.call()
			if err == nil && !allNaN(out) {
				known.failf(c.name, "%s = %v with nil error", c.name, out)
			}
		}()
	}
}

func allNaN(xs []float64) bool {
	for _, x := range xs {
		if !math.IsNaN(x) {
			return false
		}
	}
	return len(xs) > 0
}

// TestConstantInput checks that for data where every value is c, location
// statistics return c and spread statistics return 0 at any magnitude. Near
// math.MaxFloat64 a naive sum or midpoint overflows even though the answer
// is c itself.
func TestConstantInput(t *testing.T) {
	known := newKnownIssues(t,
		// Mean sums before dividing, so n values near MaxFloat64 overflow, and
		// Median, the quartiles and the variance family inherit the overflow.
		// numpy and pandas overflow the same way; Python's statistics module
		// and scipy.stats.iqr do not.
		"InterQuartileRange",
		"Mean",
		"Median",
		"MedianAbsoluteDeviation",
		"Midhinge",
		"MovingAverage",
		"MovingMedian",
		"MovingStdDev",
		"NormFit.loc",
		"NormFit.scale",
		"PopulationVariance",
		"Quartile",
		"SampleVariance",
		"SEM",
		"StandardDeviation",
		"StandardDeviationSample",
		"Trimean",
		"TrimmedMean",
		"WeightedMean",
	)
	defer known.done()
	location := []struct {
		name string
		call func(d stats.Float64Data) ([]float64, error)
	}{
		{"CumulativeMax", func(d stats.Float64Data) ([]float64, error) { return stats.CumulativeMax(d) }},
		{"CumulativeMin", func(d stats.Float64Data) ([]float64, error) { return stats.CumulativeMin(d) }},
		{"EWMA", func(d stats.Float64Data) ([]float64, error) { return stats.EWMA(d, 0.3) }},
		{"Max", func(d stats.Float64Data) ([]float64, error) { return one(stats.Max(d)) }},
		{"Mean", func(d stats.Float64Data) ([]float64, error) { return one(stats.Mean(d)) }},
		{"Median", func(d stats.Float64Data) ([]float64, error) { return one(stats.Median(d)) }},
		{"Midhinge", func(d stats.Float64Data) ([]float64, error) { return one(stats.Midhinge(d)) }},
		{"Min", func(d stats.Float64Data) ([]float64, error) { return one(stats.Min(d)) }},
		{"Mode", func(d stats.Float64Data) ([]float64, error) { return stats.Mode(d) }},
		{"MovingAverage", func(d stats.Float64Data) ([]float64, error) { return stats.MovingAverage(d, 2) }},
		{"MovingMax", func(d stats.Float64Data) ([]float64, error) { return stats.MovingMax(d, 2) }},
		{"MovingMedian", func(d stats.Float64Data) ([]float64, error) { return stats.MovingMedian(d, 2) }},
		{"MovingMin", func(d stats.Float64Data) ([]float64, error) { return stats.MovingMin(d, 2) }},
		{"NormFit.loc", func(d stats.Float64Data) ([]float64, error) { return []float64{stats.NormFit(d)[0]}, nil }},
		{"Percentile", func(d stats.Float64Data) ([]float64, error) { return one(stats.Percentile(d, 30)) }},
		{"PercentileNearestRank", func(d stats.Float64Data) ([]float64, error) { return one(stats.PercentileNearestRank(d, 30)) }},
		{"PercentileWeighted", func(d stats.Float64Data) ([]float64, error) {
			return one(stats.PercentileWeighted(d, stats.Float64Data{1, 2, 3, 4}, 30))
		}},
		{"Quartile", func(d stats.Float64Data) ([]float64, error) {
			q, err := stats.Quartile(d)
			return []float64{q.Q1, q.Q2, q.Q3}, err
		}},
		{"Trimean", func(d stats.Float64Data) ([]float64, error) { return one(stats.Trimean(d)) }},
		{"TrimmedMean", func(d stats.Float64Data) ([]float64, error) { return one(stats.TrimmedMean(d, 0.25)) }},
		{"WeightedMean", func(d stats.Float64Data) ([]float64, error) {
			return one(stats.WeightedMean(d, stats.Float64Data{1, 2, 3, 4}))
		}},
		{"Winsorize", func(d stats.Float64Data) ([]float64, error) { return stats.Winsorize(d, 0.25) }},
	}
	spread := []struct {
		name string
		call func(d stats.Float64Data) ([]float64, error)
	}{
		{"Diff", func(d stats.Float64Data) ([]float64, error) { return stats.Diff(d) }},
		{"InterQuartileRange", func(d stats.Float64Data) ([]float64, error) { return one(stats.InterQuartileRange(d)) }},
		{"MedianAbsoluteDeviation", func(d stats.Float64Data) ([]float64, error) { return one(stats.MedianAbsoluteDeviation(d)) }},
		{"MovingStdDev", func(d stats.Float64Data) ([]float64, error) { return stats.MovingStdDev(d, 2) }},
		{"NormFit.scale", func(d stats.Float64Data) ([]float64, error) { return []float64{stats.NormFit(d)[1]}, nil }},
		{"PopulationVariance", func(d stats.Float64Data) ([]float64, error) { return one(stats.PopulationVariance(d)) }},
		{"Range", func(d stats.Float64Data) ([]float64, error) { return one(stats.Range(d)) }},
		{"SampleVariance", func(d stats.Float64Data) ([]float64, error) { return one(stats.SampleVariance(d)) }},
		{"SEM", func(d stats.Float64Data) ([]float64, error) { return one(stats.SEM(d)) }},
		{"StandardDeviation", func(d stats.Float64Data) ([]float64, error) { return one(stats.StandardDeviation(d)) }},
		{"StandardDeviationSample", func(d stats.Float64Data) ([]float64, error) { return one(stats.StandardDeviationSample(d)) }},
	}
	for _, c := range []float64{2, -7.5, 1e308, math.MaxFloat64, -math.MaxFloat64, 5e-324} {
		d := stats.Float64Data{c, c, c, c}
		for _, f := range location {
			out, err := f.call(d)
			for _, v := range out {
				// EWMA may round by an ulp, so compare with a tolerance.
				if err != nil || !tolerance(v, c, 1e-15) {
					known.failf(f.name, "%s(%g x4) = %v, %v; want %g", f.name, c, out, err, c)
					break
				}
			}
		}
		for _, f := range spread {
			out, err := f.call(d)
			for _, v := range out {
				if err != nil || v != 0 {
					known.failf(f.name, "%s(%g x4) = %v, %v; want 0", f.name, c, out, err)
					break
				}
			}
		}
	}
}
