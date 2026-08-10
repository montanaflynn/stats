package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleProbGeom() {
	p := 0.5
	a := 1
	b := 2
	chance, _ := stats.ProbGeom(a, b, p)
	fmt.Println(chance)
	// Output: 0.75
}

func TestProbGeomLarge(t *testing.T) {
	p := 0.5
	a := 1
	b := 10000
	chance, err := stats.ProbGeom(a, b, p)
	if err != nil {
		t.Errorf("Returned an error")
	}
	// total probability mass over the whole support converges to 1
	if chance != 1.0 {
		t.Errorf("ProbGeom(%d, %d, %.01f) => %.1f != %.1f", a, b, p, chance, 1.0)
	}
}

// P(interval [a,a]) is the single-trial mass p*q^(a-1); the empty loop returned 0.
func TestProbGeomSinglePoint(t *testing.T) {
	cases := []struct {
		a    int
		p    float64
		want float64
	}{
		{1, 0.5, 0.5},
		{2, 0.5, 0.25},
		{3, 0.25, 0.140625},
	}
	for _, c := range cases {
		got, err := stats.ProbGeom(c.a, c.a, c.p)
		if err != nil {
			t.Errorf("ProbGeom(%d, %d, %v) returned an error", c.a, c.a, c.p)
		}
		if got != c.want {
			t.Errorf("ProbGeom(%d, %d, %v) => %v != %v", c.a, c.a, c.p, got, c.want)
		}
	}
}

// Total mass over the support converges to 1 for any valid p.
func TestProbGeomTotalMass(t *testing.T) {
	for _, p := range []float64{0.3, 0.1, 0.9} {
		got, err := stats.ProbGeom(1, 100000, p)
		if err != nil {
			t.Errorf("ProbGeom total mass returned an error for p=%v", p)
		}
		if math.Abs(got-1.0) > 1e-9 {
			t.Errorf("ProbGeom(1, 100000, %v) => %v, want ~1", p, got)
		}
	}
}

// Splitting an interval at m must sum back to the whole: [a,b] = [a,m] + [m+1,b].
func TestProbGeomAdditivity(t *testing.T) {
	whole, _ := stats.ProbGeom(1, 10, 0.3)
	lo, _ := stats.ProbGeom(1, 4, 0.3)
	hi, _ := stats.ProbGeom(5, 10, 0.3)
	if math.Abs(whole-(lo+hi)) > 1e-12 {
		t.Errorf("ProbGeom additivity: %v != %v + %v", whole, lo, hi)
	}
}

func TestErrBoundsProbGeom(t *testing.T) {
	p := 0.5
	a := -1
	b := 4
	chance, err := stats.ProbGeom(a, b, p)
	if err == nil {
		t.Errorf("Did not return an error when expected")
	}
	if !math.IsNaN(chance) {
		t.Errorf("ProbGeom(%d, %d, %.01f) => %.1f != %.1f", a, b, p, chance, math.NaN())
	}
}

func ExampleExpGeom() {
	p := 0.5
	exp, _ := stats.ExpGeom(p)
	fmt.Println(exp)
	// Output: 2
}

func TestExpGeom(t *testing.T) {
	p := 0.5
	exp, err := stats.ExpGeom(p)
	if err != nil {
		t.Errorf("Returned an error when not expected")
	}
	if exp != 2.0 {
		t.Errorf("ExpGeom(%.01f) => %.1f != %.1f", p, exp, 2.0)
	}
}

func TestErrExpGeom(t *testing.T) {
	p := -1.0
	exp, err := stats.ExpGeom(p)
	if err == nil {
		t.Errorf("Did not return an error")
	}
	if !math.IsNaN(exp) {
		t.Errorf("ExpGeom(%.01f) => %.1f != %.1f", p, exp, math.NaN())
	}
}

func ExampleVarGeom() {
	p := 0.5
	vari, _ := stats.VarGeom(p)
	fmt.Println(vari)
	// Output: 2
}

func TestVarGeom(t *testing.T) {
	p := 0.25
	vari, err := stats.VarGeom(p)
	if err != nil {
		t.Errorf("Returned an error when not expected")
	}
	if vari != 12.0 {
		t.Errorf("VarGeom(%.01f) => %.1f != %.1f", p, vari, 12.0)
	}
}

func TestErrVarGeom(t *testing.T) {
	p := -1.0
	vari, err := stats.VarGeom(p)
	if err == nil {
		t.Errorf("Did not return an error")
	}
	if !math.IsNaN(vari) {
		t.Errorf("VarGeom(%.01f) => %.1f != %.1f", p, vari, math.NaN())
	}
}
