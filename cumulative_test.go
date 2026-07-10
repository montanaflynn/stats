package stats_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleCumulativeProduct() {
	data := []float64{1.0, 2.0, 3.0, 4.0}
	cprod, _ := stats.CumulativeProduct(data)
	fmt.Println(cprod)
	// Output: [1 2 6 24]
}

func TestCumulativeProduct(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{1, 2, 3, 4}, []float64{1, 2, 6, 24}},
		{[]float64{2, 0.5, 4}, []float64{2, 1, 4}},
		{[]float64{-1, 2, -3}, []float64{-1, -2, 6}},
		{[]float64{5, 0, 3}, []float64{5, 0, 0}},
	} {
		got, err := stats.CumulativeProduct(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("CumulativeProduct(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.CumulativeProduct([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice should have returned an empty input error")
	}
}

func ExampleCumulativeMax() {
	data := []float64{3.0, 1.0, 4.0, 1.0, 5.0}
	cmax, _ := stats.CumulativeMax(data)
	fmt.Println(cmax)
	// Output: [3 3 4 4 5]
}

func TestCumulativeMax(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{3, 1, 4, 1, 5}, []float64{3, 3, 4, 4, 5}},
		{[]float64{1, 2, 3}, []float64{1, 2, 3}},
		{[]float64{-1, -3, -2}, []float64{-1, -1, -1}},
		{[]float64{2.5}, []float64{2.5}},
	} {
		got, err := stats.CumulativeMax(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("CumulativeMax(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.CumulativeMax([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice should have returned an empty input error")
	}
}

func ExampleCumulativeMin() {
	data := []float64{3.0, 1.0, 4.0, 1.0, 5.0}
	cmin, _ := stats.CumulativeMin(data)
	fmt.Println(cmin)
	// Output: [3 1 1 1 1]
}

func TestCumulativeMin(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{3, 1, 4, 1, 5}, []float64{3, 1, 1, 1, 1}},
		{[]float64{3, 2, 1}, []float64{3, 2, 1}},
		{[]float64{-1, -3, -2}, []float64{-1, -3, -3}},
		{[]float64{2.5}, []float64{2.5}},
	} {
		got, err := stats.CumulativeMin(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("CumulativeMin(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := stats.CumulativeMin([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("Empty slice should have returned an empty input error")
	}
}

func TestCumulativeMethods(t *testing.T) {
	data := stats.Float64Data{3, 1, 4, 1, 5}

	cprod, err := data.CumulativeProduct()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{3, 3, 12, 12, 60}, cprod) {
		t.Errorf("Float64Data.CumulativeProduct() => %.1f != [3 3 12 12 60]", cprod)
	}

	cmax, err := data.CumulativeMax()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{3, 3, 4, 4, 5}, cmax) {
		t.Errorf("Float64Data.CumulativeMax() => %.1f != [3 3 4 4 5]", cmax)
	}

	cmin, err := data.CumulativeMin()
	if err != nil {
		t.Errorf("Returned an error")
	}
	if !reflect.DeepEqual([]float64{3, 1, 1, 1, 1}, cmin) {
		t.Errorf("Float64Data.CumulativeMin() => %.1f != [3 1 1 1 1]", cmin)
	}
}

func TestCumulativeDoesNotMutateInput(t *testing.T) {
	in := []float64{3, 1, 4, 1, 5}
	want := []float64{3, 1, 4, 1, 5}
	_, _ = stats.CumulativeProduct(in)
	_, _ = stats.CumulativeMax(in)
	_, _ = stats.CumulativeMin(in)
	if !reflect.DeepEqual(want, in) {
		t.Errorf("Input slice was mutated: %.1f != %.1f", in, want)
	}
}

func BenchmarkCumulativeProductSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeProduct(makeFloatSlice(5))
	}
}

func BenchmarkCumulativeProductLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeProduct(lf)
	}
}

func BenchmarkCumulativeMaxSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeMax(makeFloatSlice(5))
	}
}

func BenchmarkCumulativeMaxLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeMax(lf)
	}
}

func BenchmarkCumulativeMinSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeMin(makeFloatSlice(5))
	}
}

func BenchmarkCumulativeMinLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stats.CumulativeMin(lf)
	}
}
