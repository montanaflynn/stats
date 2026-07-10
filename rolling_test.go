package stats_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleMovingAverage() {
	data := []float64{1, 2, 3, 4, 5}
	avg, _ := stats.MovingAverage(data, 3)
	fmt.Println(avg)
	// Output: [2 3 4]
}

func TestMovingAverage(t *testing.T) {
	for _, c := range []struct {
		in     []float64
		window int
		out    []float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3, []float64{2, 3, 4}},
		{[]float64{1, 2, 3, 4, 5}, 1, []float64{1, 2, 3, 4, 5}},
		{[]float64{1, 4, 9}, 2, []float64{2.5, 6.5}},
	} {
		got, err := stats.MovingAverage(c.in, c.window)
		if err != nil {
			t.Errorf("MovingAverage(%v, %d) returned an error", c.in, c.window)
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("MovingAverage(%v, %d) => %v != %v", c.in, c.window, got, c.out)
		}
	}
}

func TestMovingAverageFullWindowEqualsMean(t *testing.T) {
	in := []float64{1.1, 2.2, 3.3, 4.4}
	got, err := stats.MovingAverage(in, len(in))
	if err != nil {
		t.Errorf("Returned an error")
	}
	if len(got) != 1 {
		t.Fatalf("Expected a single element, got %d", len(got))
	}
	mean, _ := stats.Mean(in)
	if got[0] != mean {
		t.Errorf("MovingAverage full window %v != Mean %v", got[0], mean)
	}
}

func TestMovingAverageWindowOneCopiesInput(t *testing.T) {
	in := []float64{5, 6, 7}
	got, err := stats.MovingAverage(in, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual(in, got) {
		t.Errorf("MovingAverage(%v, 1) => %v != input", in, got)
	}
	got[0] = 99
	if in[0] != 5 {
		t.Errorf("MovingAverage(_, 1) must return a copy, not the input slice")
	}
}

func TestMovingAverageInvalidInput(t *testing.T) {
	_, err := stats.MovingAverage([]float64{}, 1)
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty input should have returned ErrEmptyInput, got %v", err)
	}
	for _, window := range []int{0, -1, 4} {
		_, err := stats.MovingAverage([]float64{1, 2, 3}, window)
		if err != stats.ErrBounds {
			t.Errorf("Window %d should have returned ErrBounds, got %v", window, err)
		}
	}
}

func ExampleMovingStdDev() {
	data := []float64{1, 2, 3, 4}
	sdev, _ := stats.MovingStdDev(data, 4)
	fmt.Println(sdev)
	// Output: [1.2909944487358056]
}

func TestMovingStdDev(t *testing.T) {
	in := []float64{1, 2, 3, 4}
	got, err := stats.MovingStdDev(in, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	want := []float64{0.7071, 0.7071, 0.7071}
	if len(got) != len(want) {
		t.Fatalf("Expected %d elements, got %d", len(want), len(got))
	}
	for i := range want {
		if math.Abs(got[i]-want[i]) > 0.0001 {
			t.Errorf("MovingStdDev(%v, 2)[%d] => %v != %v", in, i, got[i], want[i])
		}
	}
}

func TestMovingStdDevFullWindowEqualsSample(t *testing.T) {
	in := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	got, err := stats.MovingStdDev(in, len(in))
	if err != nil {
		t.Errorf("Returned an error")
	}
	if len(got) != 1 {
		t.Fatalf("Expected a single element, got %d", len(got))
	}
	sdev, _ := stats.StandardDeviationSample(in)
	if got[0] != sdev {
		t.Errorf("MovingStdDev full window %v != StandardDeviationSample %v", got[0], sdev)
	}
}

func TestMovingStdDevInvalidInput(t *testing.T) {
	_, err := stats.MovingStdDev([]float64{}, 2)
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty input should have returned ErrEmptyInput, got %v", err)
	}
	for _, window := range []int{0, -1, 1, 4} {
		_, err := stats.MovingStdDev([]float64{1, 2, 3}, window)
		if err != stats.ErrBounds {
			t.Errorf("Window %d should have returned ErrBounds, got %v", window, err)
		}
	}
}

func TestFloat64DataMovingAverage(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4, 5}
	got, err := data.MovingAverage(3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{2, 3, 4}, got) {
		t.Errorf("Float64Data.MovingAverage(3) => %v != [2 3 4]", got)
	}
}

func TestFloat64DataMovingStdDev(t *testing.T) {
	data := stats.Float64Data{1, 2, 3, 4}
	got, err := data.MovingStdDev(4)
	if err != nil {
		t.Errorf("Returned an error")
	}
	sdev, _ := stats.StandardDeviationSample(data)
	if len(got) != 1 || got[0] != sdev {
		t.Errorf("Float64Data.MovingStdDev(4) => %v != [%v]", got, sdev)
	}
}

func BenchmarkMovingAverageSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.MovingAverage(makeFloatSlice(5), 3)
	}
}

func BenchmarkMovingAverageLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.MovingAverage(lf, 100)
	}
}
