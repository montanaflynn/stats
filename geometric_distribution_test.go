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
	// Output: 0.25
}

func TestProbGeomLarge(t *testing.T) {
	p := 0.5
	a := 1
	b := 10000
	chance, err := stats.ProbGeom(a, b, p)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if chance != 0.5 {
		t.Errorf("ProbGeom(%d, %d, %.01f) => %.1f != %.1f", a, b, p, chance, 0.5)
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
