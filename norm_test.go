package stats

import (
	"math"
	"reflect"
	"testing"
)

func testEq(a, b []float64) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestNormPpf(t *testing.T) {
	if NormPpf(0.5, 0, 1) != 0 {
		t.Error("Input 0.5, Expected 0")
	}
	if NormPpf(0.1, 0, 1) != -1.2815515655446004 {
		t.Error("Input 0.1, Expected -1.2815515655446004")
	}
	if NormPpf(0.002423, 0, 1) != -2.817096255323953 {
		t.Error("Input 0.002423, Expected -2.817096255323953")
	}
	if NormPpf(1 - 0.002423, 0, 1) != 2.817096255323956 {
		t.Error("Input 1 - 0.002423, Expected 2.817096255323956")
	}
	if NormPpf(1.1, 0, 1) == math.NaN() {
		t.Error("Input 1.1, Expected NaN")
	}
	if NormPpf(-0.1, 0, 1) == math.NaN() {
		t.Error("Input -0.1, Expected Nan")
	}
	if NormPpf(0, 0, 1) != -math.Inf(1) {
		t.Error("Input 0, Expected -Inf")
	}
	if NormPpf(1, 0, 1) != math.Inf(1) {
		t.Error("Input 1, Expected Inf")
	}
}

func TestNormCdf(t *testing.T) {
	if NormCdf(0, 0, 1) != 0.5 {
		t.Error("Input 0, Expected 0.5")
	}
	if NormCdf(0.5, 0, 1) != 0.6914624612740131 {
		t.Error("Input 0.5, Expected 0.6914624612740131")
	}
	if NormCdf(-0.5, 0, 1) != 0.3085375387259869 {
		t.Error("Input -0.5, Expected 0.3085375387259869")
	}
}

func TestNormPdf(t *testing.T) {
	if NormPdf(0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input 0.5, Expected 0.35206532676429947")
	}
	if NormPdf(0, 0, 1) != 0.3989422804014327 {
		t.Error("Input 0, Expected 0.3989422804014327")
	}
	if NormPdf(-0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input -0.5, Expected 0.35206532676429947")
	}
}

func TestNormLogPdf(t *testing.T) {
	if NormLogPdf(0, 0, 1) != -0.9189385332046727 {
		t.Error("Input 0, Expected -0.9189385332046727")
	}
	if NormPdf(0, 0, 1) != 0.3989422804014327 {
		t.Error("Input 0, Expected 0.3989422804014327")
	}
	if NormPdf(-0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input -0.5, Expected 0.35206532676429947")
	}
}

func TestNormLogCdf(t *testing.T) {
	if NormLogCdf(0.5, 0, 1) != -0.36894641528865635 {
		t.Error("Input 0.5, Expected -0.36894641528865635")
	}
}

func TestNormIsf(t *testing.T) {
	if NormIsf(0.5, 0, 1) != 0 {
		t.Error("Input 0.5, Expected 0")
	}
	if NormIsf(0.1, 0, 1) != 1.2815515655446004 {
		t.Error("Input 0.1, Expected 1.2815515655446004")
	}
}

func TestNormSf(t *testing.T) {
	if NormSf(0.5, 0, 1) != 0.3085375387259869 {
		t.Error("Input 0.5, Expected 0.3085375387259869")
	}
}

func TestNormLogSf(t *testing.T) {
	if NormLogSf(0.5, 0, 1) != -1.1759117615936185 {
		t.Error("Input 0.5, Expected -1.1759117615936185")
	}
}

func TestNormMoment(t *testing.T) {
	if NormMoment(4, 0, 1) != 3 {
		t.Error("Input 3, Expected 3")
	}
	if NormMoment(4, 0, 1) != 3 {
		t.Error("Input 3, Expected 3")
	}
}

func TestNormStats(t *testing.T) {
	if !reflect.DeepEqual(NormStats(0, 1, "m"), []float64{0}){
		t.Error("Input 'm' , Expected 0")
	}
	if !reflect.DeepEqual(NormStats(0, 1, "v"), []float64{1}){
		t.Error("Input 'v' , Expected 1")
	}
	if !reflect.DeepEqual(NormStats(0, 1, "s"), []float64{0}){
		t.Error("Input 's' , Expected 0")
	}
	if !reflect.DeepEqual(NormStats(0, 1, "k"), []float64{0}){
		t.Error("Input 'k' , Expected 0")
	}
}

func TestNormEntropy(t *testing.T) {
	if NormEntropy( 0, 1) != 1.4189385332046727 {
		t.Error("Input ( 0 , 1 ), Expected 1.4189385332046727")
	}
}

func TestNormFit(t *testing.T) {
	if !reflect.DeepEqual(NormFit( []float64{0,2,3,4}), [2]float64{2.25, 1.479019945774904}) {
		t.Error("Input (0,2,3,4), Expected {2.25, 1.479019945774904}")
	}
}

func TestNormInterval(t *testing.T) {
	if !reflect.DeepEqual(NormInterval(0.5, 0, 1), [2]float64{-0.6744897501960818, 0.674489750196082}){
		t.Error("Input (50 % ), Expected {-0.6744897501960818, 0.674489750196082}")
	}
}

func TestNormMean(t *testing.T) {
	if NormMean(0, 1) != 0 {
		t.Error("Input (0, 1), Expected 0")
	}
}

func TestNormMedian(t *testing.T) {
	if NormMedian(0, 1) != 0 {
		t.Error("Input (0, 1), Expected 0")
	}
}

func TestNormVar(t *testing.T) {
	if NormVar(0, 1) != 1 {
		t.Error("Input (0, 1), Expected 1")
	}
}

func TestNormStd(t *testing.T) {
	if NormStd(0, 1) != 1 {
		t.Error("Input (0, 1), Expected 1")
	}
}

func TestNormPpfRvs(t *testing.T) {
	if len(NormPpfRvs(0, 1, 101)) != 101 {
		t.Error("Input size=101, Expected 101")
	}
}

func TestNormBoxMullerRvs(t *testing.T) {
	if len(NormBoxMullerRvs(0, 1, 101)) != 101 {
		t.Error("Input size=101, Expected 101")
	}
}

func TestNcr(t *testing.T) {
	if Ncr(4, 1) != 4 {
		t.Error("Input 4 choose 1, Expected 4")
	}
	if Ncr(4, 3) != 4{
		t.Error("Input 4 choose 3, Expected 4")
	}
}
