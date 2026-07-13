package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleMovingMedian() {
	data := []float64{5, 1, 4, 2, 3}
	median, _ := stats.MovingMedian(data, 3)
	fmt.Println(median)
	// Output: [4 2 3]
}

func ExampleMovingMin() {
	data := []float64{5, 1, 4, 2, 3}
	min, _ := stats.MovingMin(data, 3)
	fmt.Println(min)
	// Output: [1 1 2]
}

func ExampleMovingMax() {
	data := []float64{5, 1, 4, 2, 3}
	max, _ := stats.MovingMax(data, 3)
	fmt.Println(max)
	// Output: [5 4 4]
}

func ExampleMovingSum() {
	data := []float64{1, 2, 3, 4, 5}
	sum, _ := stats.MovingSum(data, 3)
	fmt.Println(sum)
	// Output: [6 9 12]
}

func TestMovingMedian(t *testing.T) {
	for _, c := range []struct {
		in     []float64
		window int
		out    []float64
	}{
		{[]float64{5, 1, 4, 2, 3}, 3, []float64{4, 2, 3}},
		{[]float64{5, 1, 4, 2, 3}, 1, []float64{5, 1, 4, 2, 3}},
		{[]float64{1, 4, 9, 16}, 2, []float64{2.5, 6.5, 12.5}},
	} {
		got, err := stats.MovingMedian(c.in, c.window)
		if err != nil {
			t.Errorf("MovingMedian(%v, %d) returned an error", c.in, c.window)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("MovingMedian(%v, %d) => %v != %v", c.in, c.window, got, c.out)
		}
	}
}

func TestMovingMin(t *testing.T) {
	for _, c := range []struct {
		in     []float64
		window int
		out    []float64
	}{
		{[]float64{5, 1, 4, 2, 3}, 3, []float64{1, 1, 2}},
		{[]float64{5, 1, 4, 2, 3}, 1, []float64{5, 1, 4, 2, 3}},
		{[]float64{3, 2, 5, 1}, 2, []float64{2, 2, 1}},
	} {
		got, err := stats.MovingMin(c.in, c.window)
		if err != nil {
			t.Errorf("MovingMin(%v, %d) returned an error", c.in, c.window)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("MovingMin(%v, %d) => %v != %v", c.in, c.window, got, c.out)
		}
	}
}

func TestMovingMax(t *testing.T) {
	for _, c := range []struct {
		in     []float64
		window int
		out    []float64
	}{
		{[]float64{5, 1, 4, 2, 3}, 3, []float64{5, 4, 4}},
		{[]float64{5, 1, 4, 2, 3}, 1, []float64{5, 1, 4, 2, 3}},
		{[]float64{3, 2, 5, 1}, 2, []float64{3, 5, 5}},
	} {
		got, err := stats.MovingMax(c.in, c.window)
		if err != nil {
			t.Errorf("MovingMax(%v, %d) returned an error", c.in, c.window)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("MovingMax(%v, %d) => %v != %v", c.in, c.window, got, c.out)
		}
	}
}

func TestMovingSum(t *testing.T) {
	for _, c := range []struct {
		in     []float64
		window int
		out    []float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3, []float64{6, 9, 12}},
		{[]float64{1, 2, 3, 4, 5}, 1, []float64{1, 2, 3, 4, 5}},
		{[]float64{1, 4, 9}, 2, []float64{5, 13}},
	} {
		got, err := stats.MovingSum(c.in, c.window)
		if err != nil {
			t.Errorf("MovingSum(%v, %d) returned an error", c.in, c.window)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("MovingSum(%v, %d) => %v != %v", c.in, c.window, got, c.out)
		}
	}
}

func TestMovingFullWindowEqualsScalar(t *testing.T) {
	in := []float64{2.2, 4.4, 1.1, 3.3}
	for _, c := range []struct {
		name   string
		moving func(stats.Float64Data, int) ([]float64, error)
		scalar func(stats.Float64Data) (float64, error)
	}{
		{"MovingMedian", stats.MovingMedian, stats.Median},
		{"MovingMin", stats.MovingMin, stats.Min},
		{"MovingMax", stats.MovingMax, stats.Max},
		{"MovingSum", stats.MovingSum, stats.Sum},
	} {
		got, err := c.moving(in, len(in))
		if err != nil {
			t.Errorf("%s returned an error", c.name)
		}
		if len(got) != 1 {
			t.Fatalf("%s expected a single element, got %d", c.name, len(got))
		}
		want, _ := c.scalar(in)
		if got[0] != want {
			t.Errorf("%s full window %v != scalar %v", c.name, got[0], want)
		}
	}
}

func TestMovingWindowOneCopiesInput(t *testing.T) {
	for _, c := range []struct {
		name   string
		moving func(stats.Float64Data, int) ([]float64, error)
	}{
		{"MovingMedian", stats.MovingMedian},
		{"MovingMin", stats.MovingMin},
		{"MovingMax", stats.MovingMax},
		{"MovingSum", stats.MovingSum},
	} {
		in := []float64{5, 6, 7}
		got, err := c.moving(in, 1)
		if err != nil {
			t.Errorf("%s returned an error", c.name)
		}
		if !reflect.DeepEqual(in, got) {
			t.Errorf("%s(%v, 1) => %v != input", c.name, in, got)
		}
		got[0] = 99
		if in[0] != 5 {
			t.Errorf("%s(_, 1) must return a copy, not the input slice", c.name)
		}
	}
}

func TestMovingInvalidInput(t *testing.T) {
	for _, c := range []struct {
		name   string
		moving func(stats.Float64Data, int) ([]float64, error)
	}{
		{"MovingMedian", stats.MovingMedian},
		{"MovingMin", stats.MovingMin},
		{"MovingMax", stats.MovingMax},
		{"MovingSum", stats.MovingSum},
	} {
		_, err := c.moving([]float64{}, 1)
		if err != stats.ErrEmptyInput {
			t.Errorf("%s empty input should have returned ErrEmptyInput, got %v", c.name, err)
		}
		for _, window := range []int{0, -1, 4} {
			_, err := c.moving([]float64{1, 2, 3}, window)
			if err != stats.ErrBounds {
				t.Errorf("%s window %d should have returned ErrBounds, got %v", c.name, window, err)
			}
		}
	}
}

func TestMovingDoesNotMutateInput(t *testing.T) {
	in := []float64{5, 1, 4, 2, 3}
	want := []float64{5, 1, 4, 2, 3}
	_, _ = stats.MovingMedian(in, 3)
	_, _ = stats.MovingMin(in, 3)
	_, _ = stats.MovingMax(in, 3)
	_, _ = stats.MovingSum(in, 3)
	if !reflect.DeepEqual(want, in) {
		t.Errorf("Input slice was mutated: %v != %v", in, want)
	}
}

func TestFloat64DataMovingMedian(t *testing.T) {
	data := stats.Float64Data{5, 1, 4, 2, 3}
	got, err := data.MovingMedian(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{4, 2, 3}, got) {
		t.Errorf("Float64Data.MovingMedian(3) => %v != [4 2 3]", got)
	}
}

func TestFloat64DataMovingMin(t *testing.T) {
	data := stats.Float64Data{5, 1, 4, 2, 3}
	got, err := data.MovingMin(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{1, 1, 2}, got) {
		t.Errorf("Float64Data.MovingMin(3) => %v != [1 1 2]", got)
	}
}

func TestFloat64DataMovingMax(t *testing.T) {
	data := stats.Float64Data{5, 1, 4, 2, 3}
	got, err := data.MovingMax(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{5, 4, 4}, got) {
		t.Errorf("Float64Data.MovingMax(3) => %v != [5 4 4]", got)
	}
}

func TestFloat64DataMovingSum(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 5}
	got, err := data.MovingSum(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{6, 9, 12}, got) {
		t.Errorf("Float64Data.MovingSum(3) => %v != [6 9 12]", got)
	}
}
