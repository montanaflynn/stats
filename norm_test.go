package stats_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestNormPpf(t *testing.T) {
	if stats.NormPpf(0.5, 0, 1) != 0 {
		t.Error("Input 0.5, Expected 0")
	}
	if !veryclose(stats.NormPpf(0.1, 0, 1), -1.2815515655446004) {
		t.Error("Input 0.1, Expected -1.2815515655446004")
	}
	if stats.NormPpf(0.002423, 0, 1) != -2.817096255323953 {
		t.Error("Input 0.002423, Expected -2.817096255323953")
	}
	if !close(stats.NormPpf(1-0.002423, 0, 1), 2.817096255323956) {
		t.Error("Input 1 - 0.002423, Expected 2.817096255323956")
	}

	if !math.IsNaN(stats.NormPpf(1.1, 0, 1)) {
		t.Error("Input 1.1, Expected NaN")
	}
	if !math.IsNaN(stats.NormPpf(-1.1, 0, 1)) {
		t.Error("Input -0.1, Expected Nan")
	}
	if stats.NormPpf(0, 0, 1) != -math.Inf(1) {
		t.Error("Input 0, Expected -Inf")
	}
	if stats.NormPpf(1, 0, 1) != math.Inf(1) {
		t.Error("Input 1, Expected Inf")
	}
}

func TestNormCdf(t *testing.T) {
	if stats.NormCdf(0, 0, 1) != 0.5 {
		t.Error("Input 0, Expected 0.5")
	}
	if stats.NormCdf(0.5, 0, 1) != 0.6914624612740131 {
		t.Error("Input 0.5, Expected 0.6914624612740131")
	}
	if stats.NormCdf(-0.5, 0, 1) != 0.3085375387259869 {
		t.Error("Input -0.5, Expected 0.3085375387259869")
	}
}

func TestNormPdf(t *testing.T) {
	if stats.NormPdf(0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input 0.5, Expected 0.35206532676429947")
	}
	if stats.NormPdf(0, 0, 1) != 0.3989422804014327 {
		t.Error("Input 0, Expected 0.3989422804014327")
	}
	if stats.NormPdf(-0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input -0.5, Expected 0.35206532676429947")
	}
}

func TestNormLogPdf(t *testing.T) {
	if stats.NormLogPdf(0, 0, 1) != -0.9189385332046727 {
		t.Error("Input 0, Expected -0.9189385332046727")
	}
	if stats.NormPdf(0, 0, 1) != 0.3989422804014327 {
		t.Error("Input 0, Expected 0.3989422804014327")
	}
	if stats.NormPdf(-0.5, 0, 1) != 0.35206532676429947 {
		t.Error("Input -0.5, Expected 0.35206532676429947")
	}
}

func TestNormLogCdf(t *testing.T) {
	if stats.NormLogCdf(0.5, 0, 1) != -0.36894641528865635 {
		t.Error("Input 0.5, Expected -0.36894641528865635")
	}
}

func TestNormIsf(t *testing.T) {
	if stats.NormIsf(0.5, 0, 1) != 0 {
		t.Error("Input 0.5, Expected 0")
	}
	if !veryclose(stats.NormIsf(0.1, 0, 1), 1.2815515655446004) {
		t.Error("Input 0.1, Expected 1.2815515655446004")
	}
}

func TestNormSf(t *testing.T) {
	if stats.NormSf(0.5, 0, 1) != 0.3085375387259869 {
		t.Error("Input 0.5, Expected 0.3085375387259869")
	}
}

func TestNormLogSf(t *testing.T) {
	if stats.NormLogSf(0.5, 0, 1) != -1.1759117615936185 {
		t.Error("Input 0.5, Expected -1.1759117615936185")
	}
}

// normTailRef holds reference values for the standard normal at |z| where the
// 0.5*(1+Erf(z)) and 1-cdf forms lose the result to cancellation. Computed with
// an 80 digit evaluation of erfc(z/sqrt(2))/2; sf is 0 where it underflows.
var normTailRef = []struct {
	z, sf, logSf, logPdf float64
}{
	{8, 6.220960574271784e-16, -35.01343715991455, -32.918938533204674},
	{9, 1.1285884059538405e-19, -43.628149113332114, -41.418938533204674},
	{10, 7.619853024160525e-24, -53.23128515051247, -50.918938533204674},
	{20, 2.7536241186062337e-89, -203.91715537109727, -200.91893853320468},
	{30, 4.906713927148187e-198, -454.3212439563432, -450.91893853320465},
	{37, 5.725571222524577e-300, -689.0305855768906, -685.4189385332047},
	{100, 0, -5005.524208694205, -5000.918938533205},
	{1000, 0, -500007.82669481216, -500000.9189385332},
}

func TestNormTail(t *testing.T) {
	for _, c := range normTailRef {
		if got := stats.NormSf(c.z, 0, 1); !tolerance(got, c.sf, 1e-12) {
			t.Errorf("NormSf(%v), got %v, want %v", c.z, got, c.sf)
		}
		if got := stats.NormCdf(-c.z, 0, 1); !tolerance(got, c.sf, 1e-12) {
			t.Errorf("NormCdf(%v), got %v, want %v", -c.z, got, c.sf)
		}
		if got := stats.NormLogSf(c.z, 0, 1); !tolerance(got, c.logSf, 1e-12) {
			t.Errorf("NormLogSf(%v), got %v, want %v", c.z, got, c.logSf)
		}
		if got := stats.NormLogCdf(-c.z, 0, 1); !tolerance(got, c.logSf, 1e-12) {
			t.Errorf("NormLogCdf(%v), got %v, want %v", -c.z, got, c.logSf)
		}
		if got := stats.NormLogPdf(-c.z, 0, 1); !tolerance(got, c.logPdf, 1e-12) {
			t.Errorf("NormLogPdf(%v), got %v, want %v", -c.z, got, c.logPdf)
		}
	}
}

// The near side of each pair is the one that used to be computed as 1-cdf.
func TestNormTailNearOne(t *testing.T) {
	if got, want := stats.NormCdf(8, 0, 1), 0.9999999999999993; got != want {
		t.Errorf("NormCdf(8), got %v, want %v", got, want)
	}
	if got, want := stats.NormLogCdf(8, 0, 1), -6.220960574271786e-16; !tolerance(got, want, 1e-12) {
		t.Errorf("NormLogCdf(8), got %v, want %v", got, want)
	}
	if got, want := stats.NormLogSf(-10, 0, 1), -7.619853024160525e-24; !tolerance(got, want, 1e-12) {
		t.Errorf("NormLogSf(-10), got %v, want %v", got, want)
	}
}

func TestNormCdfSfSymmetry(t *testing.T) {
	for x := -40.0; x <= 40.0; x += 0.25 {
		if a, b := stats.NormCdf(-x, 0, 1), stats.NormSf(x, 0, 1); a != b {
			t.Errorf("NormCdf(%v) = %v, NormSf(%v) = %v", -x, a, x, b)
		}
		if a, b := stats.NormLogCdf(-x, 0, 1), stats.NormLogSf(x, 0, 1); a != b {
			t.Errorf("NormLogCdf(%v) = %v, NormLogSf(%v) = %v", -x, a, x, b)
		}
		if sum := stats.NormCdf(x, 0, 1) + stats.NormSf(x, 0, 1); !veryclose(sum, 1) {
			t.Errorf("NormCdf(%v) + NormSf(%v) = %v, want 1", x, x, sum)
		}
	}
}

// loc and scale only enter through z, so the tail has to survive them too.
func TestNormTailScaled(t *testing.T) {
	if got, want := stats.NormCdf(-50, 5, 5), 1.9106595744986757e-28; !tolerance(got, want, 1e-12) {
		t.Errorf("NormCdf(-50, 5, 5), got %v, want %v", got, want)
	}
	if got, want := stats.NormSf(60, 5, 5), 1.9106595744986757e-28; !tolerance(got, want, 1e-12) {
		t.Errorf("NormSf(60, 5, 5), got %v, want %v", got, want)
	}
	if got, want := stats.NormLogCdf(-50, 5, 5), -63.82493409442372; !tolerance(got, want, 1e-12) {
		t.Errorf("NormLogCdf(-50, 5, 5), got %v, want %v", got, want)
	}
	if got, want := stats.NormLogPdf(-50, 5, 5), -63.02837644563877; !tolerance(got, want, 1e-12) {
		t.Errorf("NormLogPdf(-50, 5, 5), got %v, want %v", got, want)
	}
}

func TestNormIntervalSymmetry(t *testing.T) {
	for _, alpha := range []float64{0.5, 0.9, 0.99, 1 - 1e-8, 1 - 1e-15} {
		iv := stats.NormInterval(alpha, 0, 1)
		if math.IsInf(iv[0], 0) || math.IsInf(iv[1], 0) {
			t.Errorf("NormInterval(%v, 0, 1) = %v, want finite endpoints", alpha, iv)
		}
		if iv[0] != -iv[1] {
			t.Errorf("NormInterval(%v, 0, 1) = %v, want symmetric endpoints", alpha, iv)
		}
		want := [2]float64{3 + 2*iv[0], 3 + 2*iv[1]}
		if got := stats.NormInterval(alpha, 3, 2); !close(got[0], want[0]) || !close(got[1], want[1]) {
			t.Errorf("NormInterval(%v, 3, 2) = %v, want %v", alpha, got, want)
		}
	}
}

func TestNormMoment(t *testing.T) {
	if stats.NormMoment(4, 0, 1) != 3 {
		t.Error("Input 3, Expected 3")
	}
	if stats.NormMoment(4, 0, 1) != 3 {
		t.Error("Input 3, Expected 3")
	}
}

func TestNormStats(t *testing.T) {
	if !reflect.DeepEqual(stats.NormStats(0, 1, "m"), []float64{0}) {
		t.Error("Input 'm' , Expected 0")
	}
	if !reflect.DeepEqual(stats.NormStats(0, 1, "v"), []float64{1}) {
		t.Error("Input 'v' , Expected 1")
	}
	if !reflect.DeepEqual(stats.NormStats(0, 1, "s"), []float64{0}) {
		t.Error("Input 's' , Expected 0")
	}
	if !reflect.DeepEqual(stats.NormStats(0, 1, "k"), []float64{0}) {
		t.Error("Input 'k' , Expected 0")
	}
}

func TestNormEntropy(t *testing.T) {
	if stats.NormEntropy(0, 1) != 1.4189385332046727 {
		t.Error("Input ( 0 , 1 ), Expected 1.4189385332046727")
	}
}

func TestNormFit(t *testing.T) {
	if !reflect.DeepEqual(stats.NormFit([]float64{0, 2, 3, 4}), [2]float64{2.25, 1.479019945774904}) {
		t.Error("Input (0,2,3,4), Expected {2.25, 1.479019945774904}")
	}
}

func TestNormInterval(t *testing.T) {
	if !reflect.DeepEqual(stats.NormInterval(0.5, 0, 1), [2]float64{-0.6744897501960818, 0.6744897501960818}) {
		t.Error("Input (50 % ), Expected {-0.6744897501960818, 0.6744897501960818}")
	}
}

func TestNormMean(t *testing.T) {
	if stats.NormMean(0, 1) != 0 {
		t.Error("Input (0, 1), Expected 0")
	}
}

func TestNormMedian(t *testing.T) {
	if stats.NormMedian(0, 1) != 0 {
		t.Error("Input (0, 1), Expected 0")
	}
}

func TestNormVar(t *testing.T) {
	if stats.NormVar(0, 1) != 1 {
		t.Error("Input (0, 1), Expected 1")
	}
}

func TestNormStd(t *testing.T) {
	if stats.NormStd(0, 1) != 1 {
		t.Error("Input (0, 1), Expected 1")
	}
}

func TestNormSample(t *testing.T) {
	samples := stats.NormSample(0, 1, 100)
	if len(samples) != 100 {
		t.Errorf("Input size=100, got %d", len(samples))
	}

	samples = stats.NormSample(5, 2, 50)
	if len(samples) != 50 {
		t.Errorf("Input size=50, got %d", len(samples))
	}
}

func TestNormPpfRvs(t *testing.T) {
	if len(stats.NormPpfRvs(0, 1, 101)) != 101 {
		t.Error("Input size=101, Expected 101")
	}
}

func TestNormBoxMullerRvs(t *testing.T) {
	if len(stats.NormBoxMullerRvs(0, 1, 101)) != 101 {
		t.Error("Input size=101, Expected 101")
	}
}

func TestNcr(t *testing.T) {
	if stats.Ncr(4, 1) != 4 {
		t.Error("Input 4 choose 1, Expected 4")
	}
	if stats.Ncr(4, 3) != 4 {
		t.Error("Input 4 choose 3, Expected 4")
	}
}
