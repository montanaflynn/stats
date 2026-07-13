package stats_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleProduct() {
	product, _ := stats.Product([]float64{2, 3, 4})
	fmt.Println(product)
	// Output: 24
}

func TestProduct(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{2, 3, 4}, 24},
		{[]float64{2, 0.5, 4}, 4},
		{[]float64{-1, 2, -3}, 6},
		{[]float64{5, 0, 3}, 0},
		{[]float64{7}, 7},
	} {
		got, err := stats.Product(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if math.Abs(got-c.out) > 0.0000001 {
			t.Errorf("Product(%v) => %v != %v", c.in, got, c.out)
		}
	}
}

func TestProductEmptyInput(t *testing.T) {
	r, err := stats.Product([]float64{})
	if err != stats.ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
	if !math.IsNaN(r) {
		t.Errorf("expected NaN, got %v", r)
	}
}

func TestFloat64DataProduct(t *testing.T) {
	data := stats.Float64Data{2, 3, 4}

	r, err := data.Product()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r != 24 {
		t.Errorf("got %v, want 24", r)
	}
}
