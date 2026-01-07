package stats

import (
	"math"
	"testing"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestSkewness_RightSkewed(t *testing.T) {
	data := Float64Data{1, 2, 2, 3, 9}

	skew, err := Skewness(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if skew <= 0 {
		t.Errorf("expected right skewed (>0), got %f", skew)
	}
}

func TestSkewness_LeftSkewed(t *testing.T) {
	data := Float64Data{9, 8, 8, 7, 1}

	skew, err := Skewness(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if skew >= 0 {
		t.Errorf("expected left skewed (<0), got %f", skew)
	}
}

func TestSkewness_Symmetric(t *testing.T) {
	data := Float64Data{2, 4, 6, 8, 10}

	skew, err := Skewness(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !almostEqual(skew, 0, 1e-9) {
		t.Errorf("expected skewness ~0, got %f", skew)
	}
}

func TestSkewness_ZeroVariance(t *testing.T) {
	data := Float64Data{5, 5, 5, 5}

	_, err := Skewness(data)
	if err == nil {
		t.Fatal("expected error for zero variance, got nil")
	}
}

func TestSkewness_TooSmallDataset(t *testing.T) {
	data := Float64Data{42}

	_, err := Skewness(data)
	if err == nil {
		t.Fatal("expected error for dataset length < 2")
	}
}
