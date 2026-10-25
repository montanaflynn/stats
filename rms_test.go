package stats_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleRMS() {
	rms, _ := stats.RMS([]float64{3, 4})
	fmt.Printf("%.4f", rms)
	// Output: 3.5355
}

func TestRMS(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{3, 4}, 3.5355339059327378},
		{[]float64{1, 1, 1}, 1},
		{[]float64{-3, 4}, 3.5355339059327378},
		{[]float64{0, 0, 0}, 0},
		{[]float64{2}, 2},
	} {
		got, err := stats.RMS(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if math.Abs(got-c.out) > 0.0000001 {
			t.Errorf("RMS(%v) => %v != %v", c.in, got, c.out)
		}
	}
}

func TestRMSEmptyInput(t *testing.T) {
	r, err := stats.RMS([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(r) {
		t.Errorf("expected NaN, got %v", r)
	}
}

func TestFloat64DataRMS(t *testing.T) {
	data := stats.Float64Data{3, 4}

	r, err := data.RMS()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(r-3.5355339059327378) > 0.0000001 {
		t.Errorf("got %v, want ~3.5355", r)
	}
}

func TestRMSAccuracy(t *testing.T) {
	for _, kind := range []string{"sincos", "onebig"} {
		for _, scale := range []float64{1, 1e200, 1e-200} {
			t.Run(fmt.Sprintf("%s/scale=%g", kind, scale), func(t *testing.T) {
				x, _ := euclideanTestData(kind, 1003, scale)
				sum := new(big.Float).SetPrec(300)
				for _, v := range x {
					square := new(big.Float).SetPrec(300).SetFloat64(v)
					sum.Add(sum, square.Mul(square, square))
				}
				sum.Quo(sum, new(big.Float).SetInt64(int64(len(x))))
				want, _ := sum.Sqrt(sum).Float64()
				got, err := stats.RMS(x)
				if err != nil {
					t.Fatal(err)
				}
				if math.Abs(got/want-1) > 1e-15 {
					t.Errorf("RMS = %g, want %g", got, want)
				}
			})
		}
	}
}

func TestRMSExtremes(t *testing.T) {
	tiny := math.SmallestNonzeroFloat64
	for _, test := range []struct {
		in   stats.Float64Data
		want float64
	}{
		{stats.Float64Data{3e200, -4e200}, 5e200 / math.Sqrt2},
		{stats.Float64Data{3e-200, -4e-200}, 5e-200 / math.Sqrt2},
		{stats.Float64Data{math.MaxFloat64, math.MaxFloat64}, math.MaxFloat64},
		{stats.Float64Data{-math.MaxFloat64}, math.MaxFloat64},
		{stats.Float64Data{tiny, tiny}, tiny},
		{stats.Float64Data{1e300, 1}, 1e300 / math.Sqrt2},
		{stats.Float64Data{math.Inf(-1), 1}, math.Inf(1)},
		{stats.Float64Data{0, 0, 0}, 0},
	} {
		got, err := stats.RMS(test.in)
		ok := got == test.want || (!math.IsInf(test.want, 0) && math.Abs(got-test.want) <= 1e-15*test.want)
		if err != nil || !ok || math.Signbit(got) {
			t.Errorf("RMS(%v) = (%g, %v), want (%g, nil)", test.in, got, err, test.want)
		}
	}
	for _, in := range []stats.Float64Data{
		{math.NaN(), math.Inf(1)},
		{1e300, math.NaN()},
		{1e-300, math.NaN()},
	} {
		if got, err := stats.RMS(in); err != nil || !math.IsNaN(got) {
			t.Errorf("RMS(%v) = (%g, %v), want (NaN, nil)", in, got, err)
		}
	}
}

func BenchmarkRMS(b *testing.B) {
	for _, n := range []int{3, 10, 1000, 1000000} {
		x, _ := euclideanTestData("sincos", n, 1)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkSink, _ = stats.RMS(x)
			}
		})
	}
}
