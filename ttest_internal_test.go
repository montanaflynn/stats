package stats

import (
	"math"
	"testing"
)

func TestRegIncBetaBoundary(t *testing.T) {
	// x=0 should return 0
	if result := regIncBeta(1, 1, 0); result != 0 {
		t.Errorf("regIncBeta(1,1,0) = %v, want 0", result)
	}

	// x=1 should return 1
	if result := regIncBeta(1, 1, 1); result != 1 {
		t.Errorf("regIncBeta(1,1,1) = %v, want 1", result)
	}

	// Parameters that force near-zero intermediate values in Lentz's algorithm.
	// Very large a with x near 1 causes d = 1 - (a+b)*x/(a+1) to approach zero.
	result := regIncBeta(1e15, 0.5, 1-1e-16)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		t.Errorf("regIncBeta with extreme params should not be NaN/Inf, got %v", result)
	}

	// Exercise a wide range of parameters to cover numerical stability guards.
	// These extreme combinations force near-zero c and d values inside the
	// continued fraction loop, triggering the 1e-30 floor clamps.
	extremeCases := [][3]float64{
		{0.5, 0.5, 0.999999},
		{500, 0.5, 0.999},
		{0.001, 0.001, 0.5},
		{0.001, 1000, 0.999999},
		{1000, 0.001, 0.000001},
		{1e-10, 1e-10, 0.5},
		{1, 1e-15, 0.5},
		{1e-15, 1, 0.5},
		// Force initial d to zero: d = 1 - (a+b)*x/(a+1) = 0 when x = (a+1)/(a+b)
		{1e16, 1e16, (1e16 + 1) / 2e16},
		// Large symmetric parameters near the boundary
		{1e8, 1e8, 0.5},
	}
	for _, tc := range extremeCases {
		result := regIncBeta(tc[0], tc[1], tc[2])
		if math.IsNaN(result) || math.IsInf(result, 0) {
			t.Errorf("regIncBeta(%v, %v, %v) = %v, should not be NaN/Inf", tc[0], tc[1], tc[2], result)
		}
	}
}

func TestRegIncBetaLoopGuards(t *testing.T) {
	// Test that the continued fraction loop handles near-zero c and d values.
	// We test this by calling regIncBeta with parameters where the loop
	// numerators create values that cancel with c or d, driving them to zero.
	//
	// For even step: num = m*(b-m)*x / ((a+2m-1)*(a+2m))
	// When m > b (i.e., b-m < 0), num is negative. If d_prev is positive
	// and num*d_prev ≈ -1, then d = 1 + num*d ≈ 0.
	//
	// We test a variety of parameter combinations that stress these paths.
	cases := [][3]float64{
		{0.5, 1.5, 0.9999999999999},
		{1.5, 0.5, 0.0000000000001},
		{0.1, 0.1, 0.99999999999},
		{100, 100, 0.50000000000001},
		{1e-20, 0.5, 0.5},
		{0.5, 1e-20, 0.5},
	}
	for _, tc := range cases {
		result := regIncBeta(tc[0], tc[1], tc[2])
		if math.IsNaN(result) || math.IsInf(result, 0) {
			t.Errorf("regIncBeta(%v, %v, %v) should not be NaN/Inf, got %v", tc[0], tc[1], tc[2], result)
		}
		if result < -1e-10 || result > 1+1e-10 {
			t.Errorf("regIncBeta(%v, %v, %v) = %v, should be approximately in [0,1]", tc[0], tc[1], tc[2], result)
		}
	}
}

func TestClampTiny(t *testing.T) {
	// Values smaller than 1e-30 should be clamped
	if result := clampTiny(0); result != 1e-30 {
		t.Errorf("clampTiny(0) = %v, want 1e-30", result)
	}
	if result := clampTiny(1e-31); result != 1e-30 {
		t.Errorf("clampTiny(1e-31) = %v, want 1e-30", result)
	}
	if result := clampTiny(-1e-31); result != 1e-30 {
		t.Errorf("clampTiny(-1e-31) = %v, want 1e-30", result)
	}

	// Values at or above threshold should pass through
	if result := clampTiny(1e-30); result != 1e-30 {
		t.Errorf("clampTiny(1e-30) = %v, want 1e-30", result)
	}
	if result := clampTiny(0.5); result != 0.5 {
		t.Errorf("clampTiny(0.5) = %v, want 0.5", result)
	}
	if result := clampTiny(-0.5); result != -0.5 {
		t.Errorf("clampTiny(-0.5) = %v, want -0.5", result)
	}
}

func TestTSfZero(t *testing.T) {
	// t=0 should give survival = 0.5 (symmetric distribution)
	result := tSf(0, 10)
	if math.Abs(result-0.5) > 1e-10 {
		t.Errorf("tSf(0, 10) = %v, want 0.5", result)
	}
}
