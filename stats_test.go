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
	if m != 3 {
		t.Errorf("%.1f != %.1f", m, 3)
	}

	m = Min([]float64{-5, 1, 5})
	if m != -5 {
		t.Errorf("%.1f != %.1f", m, -5)
	}
}

func TestMax(t *testing.T) {
	m := Max([]float64{1, 2, 3, 4, 5})
	if m != 5 {
		t.Errorf("%.1f != %.1f", m, 5)
	}

	m = Max([]float64{10.5, 3, 5, 7, 9})
	if m != 10.5 {
		t.Errorf("%.1f != %.1f", m, 10.5)
	}

	m = Max([]float64{-20, -1, -5.5})
	if m != -1 {
		t.Errorf("%.1f != %.1f", m, -1)
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
	if m != 6 {
		t.Errorf("%.1f != %.1f", m, 6)
	}

	m = Sum([]float64{1.0, 1.1, 1.2, 2.2})
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
}

func TestStdDev(t *testing.T) {
	m := Round(StdDev([]float64{1, 2, 3}), 2)
	if m != 0.82 {
		t.Errorf("%.10f != %.10f", m, 0.82)
	}

	m = Round(StdDev([]float64{-1, -2, -3.3}), 2)
	if m != 0.94 {
		t.Errorf("%.10f != %.10f", m, 0.94)
	}

	m = StdDev([]float64{})
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
	if m != 5 {
		t.Errorf("%.1f != %.1f", m, 5)
	}

	m = Round(math.NaN(), 2)
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}
