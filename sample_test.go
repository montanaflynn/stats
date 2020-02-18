package stats_test

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestSample(t *testing.T) {
	_, err := stats.Sample([]float64{}, 10, false)
	if err == nil {
		t.Errorf("should return an error")
	}

	_, err = stats.Sample([]float64{0.1, 0.2}, 10, false)
	if err == nil {
		t.Errorf("should return an error")
	}
}

func TestSampleWithoutReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result, _ := stats.Sample(arr, 5, false)
	checks := map[float64]bool{}
	for _, res := range result {
		_, ok := checks[res]
		if ok {
			t.Errorf("%v already seen", res)
		}
		checks[res] = true
	}
}

func TestSampleWithReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	numsamples := 100
	result, _ := stats.Sample(arr, numsamples, true)
	if len(result) != numsamples {
		t.Errorf("%v != %v", len(result), numsamples)
	}
}

func TestStableSample(t *testing.T) {
	_, err := stats.StableSample(stats.Float64Data{}, 10)
	if err != stats.EmptyInputErr {
		t.Errorf("should return EmptyInputError when sampling an empty data")
	}
	_, err = stats.StableSample(stats.Float64Data{1.0, 2.0}, 10)
	if err != stats.BoundsErr {
		t.Errorf("should return BoundsErr when sampling size exceeds the maximum element size of data")
	}
	arr := []float64{1.0, 3.0, 2.0, -1.0, 5.0}
	locations := map[float64]int{
		1.0:  0,
		3.0:  1,
		2.0:  2,
		-1.0: 3,
		5.0:  4,
	}
	ret, _ := stats.StableSample(arr, 3)
	if len(ret) != 3 {
		t.Errorf("returned wrong sample size")
	}
	for i := 1; i < 3; i++ {
		if locations[ret[i]] < locations[ret[i-1]] {
			t.Errorf("doesn't keep order")
		}
	}
}
