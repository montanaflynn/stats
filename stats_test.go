package stats

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	m := Min([]float64{1.1, 2, 3, 4, 5})
	if m != 1.1 {
		t.Errorf("%.1f != %.1f", m, 1.1)
	}

	m = Min([]float64{10, 3, 5, 7, 9})
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

	m = Max([]float64{10, 3, 5, 7, 9})
	if m != 10 {
		t.Errorf("%.1f != %.1f", m, 10)
	}

	m = Max([]float64{-20, -1, -5})
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
	a = []float64{5, 3}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}

	m = Mode([]float64{5, 5, 3, 3, 4, 2, 1, 1, 1})
	a = []float64{1}
	if !reflect.DeepEqual(m, a) {
		t.Errorf("%.1f != %.1f", m, a)
	}
}
