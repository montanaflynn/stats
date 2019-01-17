package stats

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

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
	rand.Seed(unixnano())
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
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

func ExampleCounts() {
	d := []float64{1, 1, 2, 3, 4, 4, 4, 4}
	c := Counts(d)
	fmt.Println(c[1], c[2], c[3], c[4])
	// Output: 2 1 1 4
}

func TestCounts(t *testing.T) {
	d := []float64{1, math.NaN(), math.NaN(), 1, 2, 3, 4, 4, 4, 4, math.Inf(-1), math.Inf(1)}
	c := Counts(d)
	if c[1] != 2 {
		t.Errorf("%x != %x", c[1], 2)
	}
	if c[2] != 1 {
		t.Errorf("%x != %x", c[2], 1)
	}
	if c[4] != 4 {
		t.Errorf("%x != %x", c[4], 4)
	}
	if c[math.Inf(-1)] != 1 {
		t.Errorf("%x != %x", c[1], 1)
	}
	if c[math.Inf(1)] != 1 {
		t.Errorf("%x != %x", c[1], 1)
	}
}
