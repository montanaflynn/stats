package stats_test

import (
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/montanaflynn/stats"
)

var data1 = stats.Float64Data{-10, -10.001, 5, 1.1, 2, 3, 4.20, 5}
var data2 = stats.Float64Data{-9, -9.001, 4, .1, 1, 2, 3.20, 5}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func checkResult(result float64, err error, name string, f float64, t *testing.T) {
	if err != nil {
		t.Errorf("%s returned an error", name)
	}
	if !veryclose(result, f) {
		t.Errorf("%s() => %v != %v", name, result, f)
	}
}

// makeFloatSlice makes a slice of float64s
func makeFloatSlice(c int) []float64 {
	lf := make([]float64, 0, c)
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func makeRandFloatSlice(c int) []float64 {
	lf := make([]float64, 0, c)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func TestInterfaceMethods(t *testing.T) {
	// Test Get
	a := data1.Get(1)
	if a != -10.001 {
		t.Errorf("Get(2) => %.1f != %.1f", a, -10.001)
	}

	// Test Len
	l := data1.Len()
	if l != 8 {
		t.Errorf("Len() => %v != %v", l, 8)
	}

	// Test Less
	b := data1.Less(0, 5)
	if !b {
		t.Errorf("Less() => %v != %v", b, true)
	}

	// Test Swap
	data1.Swap(0, 2)
	if data1.Get(0) != 5 {
		t.Errorf("Len() => %v != %v", l, 8)
	}
}

func TestHelperMethods(t *testing.T) {

	// Test Min
	m, _ := data1.Min()
	if m != -10.001 {
		t.Errorf("Min() => %v != %v", m, -10.001)
	}

	// Test Max
	m, _ = data1.Max()
	if m != 5 {
		t.Errorf("Max() => %v != %v", m, 5)
	}

	// Test Sum
	m, _ = data1.Sum()
	if m != 0.2990000000000004 {
		t.Errorf("Sum() => %v != %v", m, 0.2990000000000004)
	}

	// Test CumulativeSum
	cs, _ := data1.CumulativeSum()
	want := []float64{5, -5.0009999999999994, -15.001, -13.901, -11.901, -8.901, -4.701, 0.2990000000000004}
	if !reflect.DeepEqual(cs, want) {
		t.Errorf("CumulativeSum() => %v != %v", cs, want)
	}

	// Test Mean
	m, _ = data1.Mean()
	if m != 0.03737500000000005 {
		t.Errorf("Mean() => %v != %v", m, 0.03737500000000005)
	}

	// Test GeometricMean
	m, _ = data1.GeometricMean()
	if m != 4.028070682618703 {
		t.Errorf("GeometricMean() => %v != %v", m, 4.028070682618703)
	}

	// Test HarmonicMean
	m, _ = data1.HarmonicMean()
	if !math.IsNaN(m) {
		t.Errorf("HarmonicMean() => %v != %v", m, math.NaN())
	}

	// Test Median
	m, _ = data1.Median()
	if m != 2.5 {
		t.Errorf("Median() => %v != %v", m, 2.5)
	}

	// Test Mode
	mo, _ := data1.Mode()
	if !reflect.DeepEqual(mo, []float64{5.0}) {
		t.Errorf("Mode() => %.1f != %.1f", mo, []float64{5.0})
	}

	// Test InterQuartileRange
	iqr, _ := data1.InterQuartileRange()
	if iqr != 9.05 {
		t.Errorf("InterQuartileRange() => %v != %v", iqr, 9.05)
	}
}

func assertFloat64(fn func() (float64, error), f float64, t *testing.T) {
	res, err := fn()
	checkResult(res, err, getFunctionName(fn), f, t)
}

func TestMedianAbsoluteDeviationMethods(t *testing.T) {
	assertFloat64(data1.MedianAbsoluteDeviation, 2.1, t)
	assertFloat64(data1.MedianAbsoluteDeviationPopulation, 2.1, t)
}

func TestStandardDeviationMethods(t *testing.T) {
	assertFloat64(data1.StandardDeviation, 5.935684731720091, t)
	assertFloat64(data1.StandardDeviationPopulation, 5.935684731720091, t)
	assertFloat64(data1.StandardDeviationSample, 6.345513892000508, t)
}

func TestVarianceMethods(t *testing.T) {
	assertFloat64(data1.Variance, 35.232353234375005, t)
	assertFloat64(data1.PopulationVariance, 35.232353234375005, t)
	assertFloat64(data1.SampleVariance, 40.26554655357143, t)

}

func assertPercentiles(fn func(i float64) (float64, error), i float64, f float64, t *testing.T) {
	res, err := fn(i)
	checkResult(res, err, getFunctionName(fn), f, t)
}

func TestPercentileMethods(t *testing.T) {
	assertPercentiles(data1.Percentile, 75, 4.2, t)
	assertPercentiles(data1.PercentileNearestRank, 75, 4.2, t)

}

func assertOtherDataMethods(fn func(d stats.Float64Data) (float64, error), d stats.Float64Data, f float64, t *testing.T) {
	res, err := fn(d)
	checkResult(res, err, getFunctionName(fn), f, t)
}

func TestOtherDataMethods(t *testing.T) {
	assertOtherDataMethods(data1.Correlation, data2, 0.20875473597605448, t)
	assertOtherDataMethods(data1.Pearson, data2, 0.20875473597605448, t)
	assertOtherDataMethods(data1.Midhinge, data2, -0.42500000000000004, t)
	assertOtherDataMethods(data1.Trimean, data2, 0.5375, t)
	assertOtherDataMethods(data1.Covariance, data2, 7.3814215535714265, t)
	assertOtherDataMethods(data1.CovariancePopulation, data2, 6.458743859374998, t)
}

func TestAutoCorrelationMethod(t *testing.T) {
	_, err := data1.AutoCorrelation(1)
	if err != nil {
		t.Error("stats.Float64Data.AutoCorrelation returned an error")
	}
}

func TestSampleMethod(t *testing.T) {
	// Test Sample method
	_, err := data1.Sample(5, true)
	if err != nil {
		t.Errorf("%s returned an error", getFunctionName(data1.Sample))
	}
}

func TestQuartileMethods(t *testing.T) {
	// Test QuartileOutliers method
	_, err := data1.QuartileOutliers()
	if err != nil {
		t.Errorf("%s returned an error", getFunctionName(data1.QuartileOutliers))
	}

	// Test Quartile method
	_, err = data1.Quartile(data2)
	if err != nil {
		t.Errorf("%s returned an error", getFunctionName(data1.Quartile))
	}
}

func TestSigmoidMethod(t *testing.T) {
	d := stats.LoadRawData([]float64{3.0, 1.0, 2.1})
	a := []float64{0.9525741268224334, 0.7310585786300049, 0.8909031788043871}
	s, _ := d.Sigmoid()
	if !reflect.DeepEqual(s, a) {
		t.Errorf("Sigmoid() => %g != %g", s, a)
	}
}

func TestSoftMaxMethod(t *testing.T) {
	d := stats.LoadRawData([]float64{3.0, 1.0, 0.2})
	a := []float64{0.8360188027814407, 0.11314284146556013, 0.05083835575299916}
	s, _ := d.SoftMax()
	if !reflect.DeepEqual(s, a) {
		t.Errorf("SoftMax() => %g != %g", s, a)
	}
}

func TestEntropyMethod(t *testing.T) {
	d := stats.LoadRawData([]float64{3.0, 1.0, 0.2})
	a := 0.7270013625470586
	e, _ := d.Entropy()
	if e != a {
		t.Errorf("Entropy() => %v != %v", e, a)
	}
}

// Here we show the regular way of doing it
// with a plain old slice of float64s
func BenchmarkRegularAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []float64{-10, -7, -3.11, 5, 1.1, 2, 3, 4.20, 5, 18}
		_, _ = stats.Min(data)
		_, _ = stats.Max(data)
		_, _ = stats.Sum(data)
		_, _ = stats.Mean(data)
		_, _ = stats.Median(data)
		_, _ = stats.Mode(data)
	}
}

// Here's where things get interesting
// and we start to use the included
// stats.Float64Data type and methods
func BenchmarkMethodsAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := stats.Float64Data{-10, -7, -3.11, 5, 1.1, 2, 3, 4.20, 5, 18}
		_, _ = data.Min()
		_, _ = data.Max()
		_, _ = data.Sum()
		_, _ = data.Mean()
		_, _ = data.Median()
		_, _ = data.Mode()
	}
}

func TestQuartilesMethods(t *testing.T) {
	_, err := data1.Quartiles()
	if err != nil {
		t.Errorf("%s returned an error", getFunctionName(data1.Quartiles))
	}
}
