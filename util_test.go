package stats

import (
	"testing"
)

// makeFloatSlice makes a slice of float64s
func makeFloatSlice(c int) []float64 {
	lf := []float64{}
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func TestFloat64ToInt(t *testing.T) {
	m := float64ToInt(234.0234)
	if m != 234 {
		t.Errorf("%x != %x", m, 234)
	}
	m = float64ToInt(-234.0234)
	if m != -234 {
		t.Errorf("%x != %x", m, -234)
	}
	m = float64ToInt(1)
	if m != 1 {
		t.Errorf("%x != %x", m, 1)
	}
}
