package stats_test

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"testing"
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
	chance, _ := stats.ProbGeom(a, b, p)
	fmt.Println(chance)
	// Output: 1
}

func TestErrBoundsProbGeom(t *testing.T) {
	p := 0.5
	a := -1
	b := 4
	chance, err := stats.ProbGeom(a, b, p)
	if err == nil {
		t.Errorf("Did not return an error when expected")
	}
	fmt.Println(chance)
	// Output: NaN
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
	fmt.Println(exp)
	// Output: 2
}

func TestErrExpGeom(t *testing.T) {
	p := -1.0
	exp, err := stats.ExpGeom(p)
	if err == nil {
		t.Errorf("Expected Error")
	}
	fmt.Println(exp)
	// Output: NaN
}

func ExampleVarGeom() {
	p := 0.5
	exp, _ := stats.ExpGeom(p)
	fmt.Println(exp)
	// Output: 2
}

func TestVarGeom(t *testing.T) {
	p := 0.25
	exp, err := stats.VarGeom(p)
	if err != nil {
		t.Errorf("Returned an error when not expected")
	}
	fmt.Println(exp)
	// Output: 12.0
}

func TestErrVarGeom(t *testing.T) {
	p := -1.0
	exp, err := stats.VarGeom(p)
	if err == nil {
		t.Errorf("Expected Erorr")
	}
	fmt.Println(exp)
	// Output: NaN
}
