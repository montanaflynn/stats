package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func assertRegressionY(
	t *testing.T,
	regression func(stats.Series) (stats.Series, error),
	data stats.Series,
	expected []float64,
	tolerance float64,
) {
	t.Helper()
	result, err := regression(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != len(expected) {
		t.Fatalf("got %d results, expected %d", len(result), len(expected))
	}
	for i, coordinate := range result {
		if math.IsNaN(coordinate.Y) || math.IsInf(coordinate.Y, 0) || math.Abs(coordinate.Y-expected[i]) > tolerance {
			t.Errorf("result[%d].Y = %v, expected %v", i, coordinate.Y, expected[i])
		}
	}
}

func ExampleLinearRegression() {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
	}

	r, _ := stats.LinearRegression(data)
	fmt.Println(r)
	// Output: [{1 2.4} {2 3.1} {3 3.8000000000000003}]
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
	a = 3.08
	if r[1].Y != a {
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

func TestLinearRegressionWithLargeXOffset(t *testing.T) {
	data := stats.Series{{1e9, 1}, {1e9 + 1, 2}, {1e9 + 2, 3}}
	assertRegressionY(t, stats.LinearRegression, data, []float64{1, 2, 3}, 1e-12)
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

func TestExponentialRegressionWithLargeXOffset(t *testing.T) {
	data := stats.Series{{1e9, 1}, {1e9 + 1, 2}, {1e9 + 2, 4}}
	assertRegressionY(t, stats.ExponentialRegression, data, []float64{1, 2, 4}, 1e-12)
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
	a = 3.33055592224922
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
	a = 4.888413396683665
	if !veryclose(r[4].Y, a) {
		t.Errorf("%v != %v", r[4].Y, a)
	}

	_, err := stats.LogarithmicRegression([]stats.Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestLogarithmicRegressionWithLargeLogOffset(t *testing.T) {
	data := stats.Series{
		{math.Exp(700), 1},
		{math.Exp(700.000001), 2},
		{math.Exp(700.000002), 3},
	}
	assertRegressionY(t, stats.LogarithmicRegression, data, []float64{1, 2, 3}, 1e-7)
}

func TestRegressionsRejectInvalidDomains(t *testing.T) {
	type regression func(stats.Series) (stats.Series, error)
	tests := []struct {
		name       string
		regression regression
		data       stats.Series
		expected   error
	}{
		{"linear single point", stats.LinearRegression, stats.Series{{2, 7}}, stats.ErrBounds},
		{"linear constant x", stats.LinearRegression, stats.Series{{2, 1}, {2, 3}, {2, 5}}, stats.ErrBounds},
		{"exponential zero y issue 31", stats.ExponentialRegression, stats.Series{{1, 0}, {2, 3.3}}, stats.ErrYCoord},
		{"exponential single point", stats.ExponentialRegression, stats.Series{{2, 7}}, stats.ErrBounds},
		{"exponential constant x", stats.ExponentialRegression, stats.Series{{2, 1}, {2, 3}, {2, 5}}, stats.ErrBounds},
		{"logarithmic zero x", stats.LogarithmicRegression, stats.Series{{0, 1}, {2, 3.3}}, stats.ErrBounds},
		{"logarithmic negative x", stats.LogarithmicRegression, stats.Series{{-1, 1}, {2, 3.3}}, stats.ErrBounds},
		{"logarithmic later zero x", stats.LogarithmicRegression, stats.Series{{1, 1}, {0, 3.3}}, stats.ErrBounds},
		{"logarithmic single point", stats.LogarithmicRegression, stats.Series{{2, 7}}, stats.ErrBounds},
		{"logarithmic constant x", stats.LogarithmicRegression, stats.Series{{2, 1}, {2, 3}, {2, 5}}, stats.ErrBounds},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			regressions, err := test.regression(test.data)
			if err != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, err)
			}
			if regressions != nil {
				t.Fatalf("expected no regressions, got %v", regressions)
			}
		})
	}
}
