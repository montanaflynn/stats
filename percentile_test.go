package stats_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func round(value float64) float64 {
	return math.Round(value*100) / 100
}

func TestPercentileOneToTen(t *testing.T) {
	m, _ := stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)
	if m != 1 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	if m != 1.45 {
		t.Errorf("%.2f != %.2f", m, 1.45)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	if m != 1.9 {
		t.Errorf("%.1f != %.1f", m, 1.9)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15)
	if round(m) != 2.35 {
		t.Errorf("%.2f != %.2f", m, 2.35)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20)
	if m != 2.8 {
		t.Errorf("%.1f != %.1f", m, 2.8)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25)
	if m != 3.25 {
		t.Errorf("%.1f != %.1f", m, 3.25)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 30)
	if round(m) != 3.7 {
		t.Errorf("%.1f != %.1f", m, 3.7)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 35)
	if m != 4.15 {
		t.Errorf("%.1f != %.1f", m, 4.15)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 40)
	if m != 4.6 {
		t.Errorf("%.1f != %.1f", m, 4.6)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 45)
	if m != 5.05 {
		t.Errorf("%.1f != %.1f", m, 5.05)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50)
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55)
	if m != 5.95 {
		t.Errorf("%.1f != %.1f", m, 5.95)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 60)
	if round(m) != 6.4 {
		t.Errorf("%.1f != %.1f", m, 6.4)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 65)
	if round(m) != 6.85 {
		t.Errorf("%.1f != %.1f", m, 6.85)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 70)
	if m != 7.3 {
		t.Errorf("%.1f != %.1f", m, 7.3)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 75)
	if m != 7.75 {
		t.Errorf("%.1f != %.1f", m, 7.75)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 80)
	if m != 8.2 {
		t.Errorf("%.1f != %.1f", m, 8.2)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 85)
	if round(m) != 8.65 {
		t.Errorf("%.1f != %.1f", m, 8.65)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 90)
	if m != 9.1 {
		t.Errorf("%.1f != %.1f", m, 9.1)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 95)
	if round(m) != 9.55 {
		t.Errorf("%.1f != %.1f", m, 9.55)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100)
	if m != 10.0 {
		t.Errorf("%.1f != %.1f", m, 10.0)
	}
}

func TestPercentile(t *testing.T) {
	m, _ := stats.Percentile([]float64{43, 54, 56, 61, 62, 66}, 90)
	if m != 64.0 {
		t.Errorf("%.1f != %.1f", m, 64.0)
	}
	m, _ = stats.Percentile([]float64{43}, 90)
	if m != 43.0 {
		t.Errorf("%.1f != %.1f", m, 43.0)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50)
	median, _ := stats.Median([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if m != median || m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 99.9)
	if round(m) != 9.99 {
		t.Errorf("%.2f != %.2f", m, 9.91)
	}
	m, _ = stats.Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100)
	if m != 10.0 {
		t.Errorf("%.1f != %.1f", m, 10.0)
	}
	_, err := stats.Percentile([]float64{}, 99.9)
	if err != stats.EmptyInputErr {
		t.Errorf("Empty slice didn't return expected error; got %v", err)
	}
	m, err = stats.Percentile([]float64{1, 2, 3, 4, 5}, 0)
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
	m, err = stats.Percentile([]float64{1, 2, 3, 4, 5}, 0.13)
	if round(m) != 1.01 {
		t.Errorf("%.2f != %.2f", m, 1.0)
	}
	_, err = stats.Percentile([]float64{1, 2, 3, 4, 5}, 101)
	if err != stats.BoundsErr {
		t.Errorf("Too high percent didn't return expected error; got %v", err)
	}
}

func TestPercentileSortSideEffects(t *testing.T) {
	s := []float64{43, 54, 56, 44, 62, 66}
	a := []float64{43, 54, 56, 44, 62, 66}
	_, _ = stats.Percentile(s, 90)
	if !reflect.DeepEqual(s, a) {
		t.Errorf("%.1f != %.1f", s, a)
	}
}

func BenchmarkPercentileSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.Percentile(makeFloatSlice(5), 50)
	}
}

func BenchmarkPercentileLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.Percentile(lf, 50)
	}
}

func TestPercentileNearestRank(t *testing.T) {
	f1 := []float64{35, 20, 15, 40, 50}
	f2 := []float64{20, 6, 7, 8, 8, 10, 13, 15, 16, 3}
	f3 := makeFloatSlice(101)

	for _, c := range []struct {
		sample  []float64
		percent float64
		result  float64
	}{
		{f1, 30, 20},
		{f1, 40, 20},
		{f1, 50, 35},
		{f1, 75, 40},
		{f1, 95, 50},
		{f1, 99, 50},
		{f1, 99.9, 50},
		{f1, 100, 50},
		{f2, 25, 7},
		{f2, 50, 8},
		{f2, 75, 15},
		{f2, 100, 20},
		{f3, 1, 100},
		{f3, 99, 9900},
		{f3, 100, 10000},
		{f3, 0, 0},
	} {
		got, err := stats.PercentileNearestRank(c.sample, c.percent)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}
		if got != c.result {
			t.Errorf("%v != %v", got, c.result)
		}
	}

	_, err := stats.PercentileNearestRank([]float64{}, 50)
	if err == nil {
		t.Errorf("Should have returned an empty slice error")
	}

	_, err = stats.PercentileNearestRank([]float64{1, 2, 3, 4, 5}, -0.01)
	if err == nil {
		t.Errorf("Should have returned an percentage must be above 0 error")
	}

	_, err = stats.PercentileNearestRank([]float64{1, 2, 3, 4, 5}, 110)
	if err == nil {
		t.Errorf("Should have returned an percentage must not be above 100 error")
	}

}

func BenchmarkPercentileNearestRankSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.PercentileNearestRank(makeFloatSlice(5), 50)
	}
}

func BenchmarkPercentileNearestRankLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.PercentileNearestRank(lf, 50)
	}
}
