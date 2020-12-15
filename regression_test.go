package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleLinearRegression() {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
	}

	r, _ := stats.LinearRegression(data)
	fmt.Println(r)
	// Output: [{1 2.400000000000001} {2 3.1} {3 3.7999999999999994}]
}

func TestLinearRegression(t *testing.T) {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := stats.LinearRegression(data)
	a := 2.3800000000000026
	if !close(r[0].Y, a) {
		t.Errorf("%v != %v", r[0].Y, a)
	}
	a = 3.0800000000000014
	if !veryclose(r[1].Y, a) {
		t.Errorf("%v != %v", r[1].Y, a)
	}
	a = 3.7800000000000002
	if r[2].Y != a {
		t.Errorf("%v != %v", r[2].Y, a)
	}
	a = 4.479999999999999
	if !veryclose(r[3].Y, a) {
		t.Errorf("%v != %v", r[3].Y, a)
	}
	a = 5.179999999999998
	if !veryclose(r[4].Y, a) {
		t.Errorf("%v != %v", r[4].Y, a)
	}

	_, err := stats.LinearRegression([]stats.Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestExponentialRegression(t *testing.T) {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := stats.ExponentialRegression(data)
	a, _ := stats.Round(r[0].Y, 3)
	if a != 2.515 {
		t.Errorf("%v != %v", r[0].Y, 2.515)
	}
	a, _ = stats.Round(r[1].Y, 3)
	if a != 3.032 {
		t.Errorf("%v != %v", r[1].Y, 3.032)
	}
	a, _ = stats.Round(r[2].Y, 3)
	if a != 3.655 {
		t.Errorf("%v != %v", r[2].Y, 3.655)
	}
	a, _ = stats.Round(r[3].Y, 3)
	if a != 4.407 {
		t.Errorf("%v != %v", r[3].Y, 4.407)
	}
	a, _ = stats.Round(r[4].Y, 3)
	if a != 5.313 {
		t.Errorf("%v != %v", r[4].Y, 5.313)
	}

	_, err := stats.ExponentialRegression([]stats.Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestExponentialRegressionYCoordErr(t *testing.T) {
	c := []stats.Coordinate{{1, -5}, {4, 25}, {6, 5}}
	_, err := stats.ExponentialRegression(c)
	if err != stats.YCoordErr {
		t.Errorf(err.Error())
	}
}

func TestLogarithmicRegression(t *testing.T) {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := stats.LogarithmicRegression(data)
	a := 2.1520822363811702
	if !close(r[0].Y, a) {
		t.Errorf("%v != %v", r[0].Y, a)
	}
	a = 3.3305559222492214
	if !veryclose(r[1].Y, a) {
		t.Errorf("%v != %v", r[1].Y, a)
	}
	a = 4.019918836568674
	if !veryclose(r[2].Y, a) {
		t.Errorf("%v != %v", r[2].Y, a)
	}
	a = 4.509029608117273
	if !veryclose(r[3].Y, a) {
		t.Errorf("%v != %v", r[3].Y, a)
	}
	a = 4.888413396683663
	if !veryclose(r[4].Y, a) {
		t.Errorf("%v != %v", r[4].Y, a)
	}

	_, err := stats.LogarithmicRegression([]stats.Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
