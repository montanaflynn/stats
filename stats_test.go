package stats

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestMin(t *testing.T) {
	m := Min([]float64{1.1, 2, 3, 4, 5})
	if m != 1.1 {
		t.Errorf("%.1f != %.1f", m, 1.1)
	}

	m = Min([]float64{10.534, 3, 5, 7, 9})
	if m != 3.0 {
		t.Errorf("%.1fx != %.1f", m, 3.0)
	}

	m = Min([]float64{-5, 1, 5})
	if m != -5.0 {
		t.Errorf("%.1f != %.1f", m, -5.0)
	}
}

func TestMax(t *testing.T) {
	m := Max([]float64{1, 2, 3, 4, 5})
	if m != 5.0 {
		t.Errorf("%.1f != %.1f", m, 5.0)
	}

	m = Max([]float64{10.5, 3, 5, 7, 9})
	if m != 10.5 {
		t.Errorf("%.1f != %.1f", m, 10.5)
	}

	m = Max([]float64{-20, -1, -5.5})
	if m != -1.0 {
		t.Errorf("%.1f != %.1f", m, -1.0)
	}
}

func TestMean(t *testing.T) {
	m := Mean([]float64{1, 2, 3, 4, 5})
	if m != 3.0 {
		t.Errorf("%.1f != %.1f", m, 3.0)
	}

	m = Mean([]float64{1, 2, 3, 4, 5, 6})
	if m != 3.5 {
		t.Errorf("%.1f != %.1f", m, 3.5)
	}

	m = Mean([]float64{1})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}

	m = Mean([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestMedian(t *testing.T) {
	m := Median([]float64{5, 3, 4, 2, 1})
	if m != 3.0 {
		t.Errorf("%.1f != %.1f", m, 3.0)
	}

	m = Median([]float64{6, 3, 2, 4, 5, 1})
	if m != 3.5 {
		t.Errorf("%.1f != %.1f", m, 3.5)
	}

	m = Median([]float64{1})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}

	m = Median([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestMode(t *testing.T) {
	m := Mode([]float64{5, 3, 4, 2, 1})
	a := []float64{}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}

	m = Mode([]float64{5, 5, 3, 4, 2, 1})
	a = []float64{5}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}

	m = Mode([]float64{5, 5, 3, 3, 4, 2, 1})
	sort.Float64s(m)
	a = []float64{3, 5}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}

	m = Mode([]float64{5, 5, 3, 3, 4, 2, 1, 1, 1})
	a = []float64{1}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}
}

func TestSum(t *testing.T) {
	m := Sum([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}

	m = Sum([]float64{1, 2, 3})
	if m != 6.0 {
		t.Errorf("%.1f != %.1f", m, 6.0)
	}

	m = Sum([]float64{1.0, 1.1, 1.2, 2.2})
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
}

func TestVariance(t *testing.T) {
	m := Variance([]float64{}, 0)
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m = Variance([]float64{}, 1)
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m = Round(Variance([]float64{1, 2, 3}, 0), 1)
	if m != 0.7 {
		t.Errorf("%.1f != %.1f", m, 0.7)
	}
	m = Variance([]float64{1, 2, 3}, 1)
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestVarP(t *testing.T) {
	m := VarP([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m = Round(VarP([]float64{1, 2, 3}), 1)
	if m != 0.7 {
		t.Errorf("%.1f != %.1f", m, 0.7)
	}
}

func TestVarS(t *testing.T) {
	m := VarS([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m = VarS([]float64{1, 2, 3})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestStdDevP(t *testing.T) {
	m := Round(StdDevP([]float64{1, 2, 3}), 2)
	if m != 0.82 {
		t.Errorf("%.10f != %.10f", m, 0.82)
	}

	m = Round(StdDevP([]float64{-1, -2, -3.3}), 2)
	if m != 0.94 {
		t.Errorf("%.10f != %.10f", m, 0.94)
	}

	m = StdDevP([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestStdDevS(t *testing.T) {
	m := Round(StdDevS([]float64{1, 2, 3}), 2)
	if m != 1.0 {
		t.Errorf("%.10f != %.10f", m, 1.0)
	}

	m = Round(StdDevS([]float64{-1, -2, -3.3}), 2)
	if m != 1.15 {
		t.Errorf("%.10f != %.10f", m, 1.15)
	}

	m = StdDevS([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestRound(t *testing.T) {
	m := Round(0.1111, 1)
	if m != 0.1 {
		t.Errorf("%.1f != %.1f", m, 0.1)
	}

	m = Round(-0.1111, 2)
	if m != -0.11 {
		t.Errorf("%.1f != %.1f", m, -0.11)
	}

	m = Round(5.3253, 3)
	if m != 5.325 {
		t.Errorf("%.1f != %.1f", m, 5.325)
	}

	m = Round(5.3253, 0)
	if m != 5.0 {
		t.Errorf("%.1f != %.1f", m, 5.0)
	}

	m = Round(math.NaN(), 2)
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestPercentile(t *testing.T) {
	m := Percentile([]float64{43, 54, 56, 61, 62, 66}, 90)
	if m != 62.0 {
		t.Errorf("%.1f != %.1f", m, 62.0)
	}
	m = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50)
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
	m = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 99.9)
	if m != 10.0 {
		t.Errorf("%.1f != %.1f", m, 10.0)
	}
}

func TestFloat64ToInt(t *testing.T) {
	m := Float64ToInt(234.0234)
	if m != 234 {
		t.Errorf("%x != %x", m, 234)
	}
	m = Float64ToInt(-234.0234)
	if m != -234 {
		t.Errorf("%x != %x", m, -234)
	}
	m = Float64ToInt(1)
	if m != 1 {
		t.Errorf("%x != %x", m, 1)
	}
}
