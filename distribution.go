package stats

import (
	"errors"
	"math"
)

const (
	erfinvA3 = -0.140543331
	erfinvA2 = 0.914624893
	erfinvA1 = -1.645349621
	erfinvA0 = 0.886226899

	erfinvB4 = 0.012229801
	erfinvB3 = -0.329097515
	erfinvB2 = 1.442710462
	erfinvB1 = -2.118377725
	erfinvB0 = 1

	erfinvC3 = 1.641345311
	erfinvC2 = 3.429567803
	erfinvC1 = -1.62490649
	erfinvC0 = -1.970840454

	erfinvD2 = 1.637067800
	erfinvD1 = 3.543889200
	erfinvD0 = 1
)

func StandardErrorPooled(x1, n1, x2, n2 float64) (float64, error) {
	pHat, err := PHatPooled(x1, n1, x2, n2)
	radicand := ((pHat * (1 - pHat)) / n1) + ((pHat * (1 - pHat)) / n2)
	return math.Sqrt(radicand), err
}

func PHatPooled(x1, n1, x2, n2 float64) (float64, error) {
	return (x1 + x2) / (n1 + n2), nil
}

func NormalPDF(x, mean, sd float64) (float64, error) {
	exponent := -math.Pow(x-mean, 2) / (2 * math.Pow(sd, 2))
	numerator := math.Pow(math.E, exponent)
	denominator := sd * math.Sqrt(2*math.Pi)
	return numerator / denominator, nil
}

func NormalCDF(lower, upper, mean, sd float64) (float64, error) {
	upperCDF := (1 + math.Erf((upper-mean)/(sd*math.Sqrt2))) / 2
	lowerCDF := (1 + math.Erf((lower-mean)/(sd*math.Sqrt2))) / 2
	return upperCDF - lowerCDF, nil
}

func InvNorm(area, mean, sd float64) (float64, error) {
	erfinv, err := ErfInv((2 * area) - 1)
	invNorm := math.Sqrt2 * erfinv
	invNorm = mean + (sd * invNorm)
	return invNorm, err
}

// Inverse error function based off of libit - http://libit.sourceforge.net/math_8c-source.html#l00339
func ErfInv(x float64) (float64, error) {
	var x2, r, y float64
	var signX int
	if x < -1 || x > 1 {
		return 0, errors.New("outside of domain")
	}
	if x == 0 {
		return 0, nil
	}
	if x > 0 {
		signX = 1
	} else {
		signX = -1
		x = -x
	}
	if x <= 0.7 {
		x2 = x * x
		r = x * (((erfinvA3*x2+erfinvA2)*x2+erfinvA1)*x2 + erfinvA0)
		r /= (((erfinvB4*x2+erfinvB3)*x2+erfinvB2)*x2+erfinvB1)*x2 + erfinvB0
	} else {
		y = math.Sqrt(-math.Log((1 - x) / 2))
		r = (((erfinvC3*y+erfinvC2)*y+erfinvC1)*y + erfinvC0)
		r /= ((erfinvD2*y+erfinvD1)*y + erfinvD0)
	}

	r = r * float64(signX)
	x = x * float64(signX)

	r -= (math.Erf(r) - x) / (2 / math.SqrtPi * math.Exp(-r*r))
	r -= (math.Erf(r) - x) / (2 / math.SqrtPi * math.Exp(-r*r))

	return r, nil
}

func InvT(area float64) (float64, error) {
	if area < -1 || area > 1 {
		return 0, errors.New("outside of domain")
	}
	invt := math.Tan(math.Pi * (area - 0.5))
	return invt, nil
}

func TPDF(x, df float64) (float64, error) {
	exponent := -0.5 * (df + 1)
	numerator := math.Gamma((df+1)/2) * math.Pow(1+((x*x)/df), exponent)
	denominator := math.Sqrt(df*math.Pi) * math.Gamma(df/2)
	return numerator / denominator, nil
}

func TCDF(lower, upper float64) (float64, error) {
	return (math.Atan(upper) - math.Atan(lower)) / math.Pi, nil
}

func ChiSquaredPDF(x, df float64) (float64, error) {
	numerator := math.Pow(x, (df/2)-1) * math.Pow(math.E, -x/2)
	denominator := math.Pow(2, df/2) * math.Gamma(df/2)
	return numerator / denominator, nil
}

func ChiSquaredCDF(lower, upper, df float64) (float64, error) {
	u, err := ChiSquaredPDF(upper, df)
	if err != nil {
		return 0, err
	}
	l, err := ChiSquaredPDF(lower, df)
	return u - l, err
}

func FPDF(x, d1, d2 float64) (float64, error) {
	n1 := math.Pow((d1*x)/((d1*x)+d2), d1/2)
	n2 := math.Pow(1-math.Pow(n1, 2/d1), d2/2)
	numerator := n1 * n2
	b, _ := Beta(d1/2, d2/2)
	denominator := x * b
	return numerator / denominator, nil
}

func FCDF(lower, upper, d1, d2 float64) (float64, error) {
	u, err := FPDF(upper, d1, d2)
	if err != nil {
		return 0, err
	}
	l, err := FPDF(lower, d1, d2)
	return u - l, err
}

func Beta(x, y float64) (float64, error) {
	numerator := math.Gamma(x) * math.Gamma(y)
	denominator := math.Gamma(x + y)
	return numerator / denominator, nil
}

func BinomPDF(trials int64, p float64, x int64) (float64, error) {
	nck, err := NCK(int64(trials), int64(x))
	pdf := float64(nck) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(trials-x))
	return pdf, err
}

func BinomCDF(trials int64, p float64, x int64) (float64, error) {
	var sum float64
	for i := int64(0); i <= x; i++ {
		pdf, err := BinomPDF(trials, p, i)
		if err != nil {
			return sum, err
		}
		sum += pdf
	}
	return sum, nil
}

func NCK(n, k int64) (int64, error) {
	nFactorial, err := Factorial(n)
	if err != nil {
		return 0, err
	}
	kFactorial, err := Factorial(k)
	if err != nil {
		return 0, err
	}
	nkDiffFactorial, err := Factorial(n - k)
	return nFactorial / (kFactorial * nkDiffFactorial), err
}

func Factorial(n int64) (int64, error) {
	var tmp int64
	if n <= 1 {
		return 1, nil
	}
	f, _ := Factorial(n - 1)
	tmp = n * f
	return tmp, nil
}

func PoissonPDF(λ float64, x int64) (float64, error) {
	numerator := math.Pow(math.E, -λ) * math.Pow(λ, float64(x))
	denominator, err := Factorial(x)
	return numerator / float64(denominator), err
}

func PoissonCDF(λ float64, x int64) (float64, error) {
	var sum float64
	for i := int64(0); i <= x; i++ {
		pdf, err := PoissonPDF(λ, i)
		if err != nil {
			return sum, err
		}
		sum += pdf
	}
	return sum, nil
}

func GeometPDF(p, x float64) (float64, error) {
	return math.Pow(1-p, x-1) * p, nil
}

func GeometCDF(p, x float64) (float64, error) {
	return 1 - math.Pow(1-p, x), nil
}
