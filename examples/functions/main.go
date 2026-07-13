package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {

	// d := stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})
	d := stats.LoadRawData([]int{1, 2, 3, 4, 5})

	a, _ := stats.Min(d)
	fmt.Println(a)
	// Output: 1.1

	a, _ = stats.Max(d)
	fmt.Println(a)
	// Output: 5

	argmax, _ := stats.ArgMax([]float64{2, 8, 4, 6})
	fmt.Println(argmax)
	// Output: 1

	argmin, _ := stats.ArgMin([]float64{2, 8, 4, 6})
	fmt.Println(argmin)
	// Output: 0

	rng, _ := stats.Range([]float64{2, 8, 4, 6})
	fmt.Println(rng)
	// Output: 6

	a, _ = stats.Sum([]float64{1.1, 2.2, 3.3})
	fmt.Println(a)
	// Output: 6.6

	cs, _ := stats.CumulativeSum([]float64{1.1, 2.2, 3.3})
	fmt.Println(cs) // [1.1 3.3000000000000003 6.6]

	cp, _ := stats.CumulativeProduct([]float64{1, 2, 3, 4})
	fmt.Println(cp)
	// Output: [1 2 6 24]

	cmax, _ := stats.CumulativeMax([]float64{1, 3, 2, 5, 4})
	fmt.Println(cmax)
	// Output: [1 3 3 5 5]

	cmin, _ := stats.CumulativeMin([]float64{5, 3, 4, 1, 2})
	fmt.Println(cmin)
	// Output: [5 3 3 1 1]

	df, _ := stats.Diff([]float64{1, 2, 4, 7})
	fmt.Println(df)
	// Output: [1 2 3]

	pc, _ := stats.PercentChange([]float64{1, 2, 4, 8})
	fmt.Println(pc)
	// Output: [1 1 1]

	a, _ = stats.Mean([]float64{1, 2, 3, 4, 5})
	fmt.Println(a)
	// Output: 3

	wm, _ := stats.WeightedMean([]float64{1, 2, 3}, []float64{1, 1, 2})
	fmt.Println(wm)
	// Output: 2.25

	a, _ = stats.Median([]float64{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a)
	// Output: 4

	m, _ := stats.Mode([]float64{5, 5, 3, 3, 4, 2, 1})
	fmt.Println(m)
	// Output: [5 3]

	a, _ = stats.PopulationVariance([]float64{1, 2, 3, 4, 5})
	fmt.Println(a)
	// Output: 2

	a, _ = stats.SampleVariance([]float64{1, 2, 3, 4, 5})
	fmt.Println(a)
	// Output: 2.5

	a, _ = stats.MedianAbsoluteDeviationPopulation([]float64{1, 2, 3})
	fmt.Println(a)
	// Output: 1

	a, _ = stats.StandardDeviationPopulation([]float64{1, 2, 3})
	fmt.Println(a)
	// Output: 0.816496580927726

	a, _ = stats.StandardDeviationSample([]float64{1, 2, 3})
	fmt.Println(a)
	// Output: 1

	cv, _ := stats.CoefficientOfVariation([]float64{1, 2, 3})
	fmt.Println(cv)
	// Output: 0.5

	zs, _ := stats.ZScore([]float64{1, 2, 3})
	fmt.Println(zs)
	// Output: [-1 0 1]

	rk, _ := stats.Rank([]float64{10, 20, 20, 30})
	fmt.Println(rk)
	// Output: [1 2.5 2.5 4]

	ma, _ := stats.MovingAverage([]float64{1, 2, 3, 4, 5}, 3)
	fmt.Println(ma)
	// Output: [2 3 4]

	msd, _ := stats.MovingStdDev([]float64{1, 2, 3, 4, 5}, 3)
	fmt.Println(msd)
	// Output: [1 1 1]

	a, _ = stats.Percentile([]float64{1, 2, 3, 4, 5}, 75)
	fmt.Println(a)
	// Output: 4

	a, _ = stats.PercentileNearestRank([]float64{35, 20, 15, 40, 50}, 75)
	fmt.Println(a)
	// Output: 40

	wp, _ := stats.PercentileWeighted([]float64{1, 2, 9, 3.2, 4}, []float64{0, 0.5, 1, 0.3, 0.5}, 90)
	fmt.Println(wp)
	// Output: 9

	c := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := stats.LinearRegression(c)
	fmt.Println(r)
	// Output: [{1 2.3800000000000026} {2 3.0800000000000014} {3 3.7800000000000002} {4 4.479999999999999} {5 5.179999999999998}]

	r, _ = stats.ExponentialRegression(c)
	fmt.Println(r)
	// Output: [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]

	r, _ = stats.LogarithmicRegression(c)
	fmt.Println(r)
	// Output: [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]

	s, _ := stats.Sample([]float64{0.1, 0.2, 0.3, 0.4}, 3, false)
	fmt.Println(s)
	// Output: [0.2,0.4,0.3]

	s, _ = stats.Sample([]float64{0.1, 0.2, 0.3, 0.4}, 10, true)
	fmt.Println(s)
	// Output: [0.2,0.2,0.4,0.1,0.2,0.4,0.3,0.2,0.2,0.1]

	q, _ := stats.Quartile([]float64{7, 15, 36, 39, 40, 41})
	fmt.Println(q)
	// Output: {15 37.5 40}

	iqr, _ := stats.InterQuartileRange([]float64{102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118})
	fmt.Println(iqr)
	// Output: 10

	mh, _ := stats.Midhinge([]float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13})
	fmt.Println(mh)
	// Output: 7.5

	tr, _ := stats.Trimean([]float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13})
	fmt.Println(tr)
	// Output: 7.25

	o, _ := stats.QuartileOutliers([]float64{-1000, 1, 3, 4, 4, 6, 6, 6, 6, 7, 8, 15, 18, 100})
	fmt.Printf("%+v\n", o)
	// Output:  {Mild:[15 18] Extreme:[-1000 100]}

	gm, _ := stats.GeometricMean([]float64{10, 51.2, 8})
	fmt.Println(gm)
	// Output: 15.999999999999991

	hm, _ := stats.HarmonicMean([]float64{1, 2, 3, 4, 5})
	fmt.Println(hm)
	// Output: 2.18978102189781

	a, _ = stats.Round(2.18978102189781, 3)
	fmt.Println(a)
	// Output: 2.19

	e, _ := stats.ChebyshevDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
	fmt.Println(e)
	// Output: 6

	e, _ = stats.ManhattanDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
	fmt.Println(e)
	// Output: 24

	e, _ = stats.EuclideanDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
	fmt.Println(e)
	// Output: 10.583005244258363

	e, _ = stats.MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(1))
	fmt.Println(e)
	// Output: 24

	e, _ = stats.MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(2))
	fmt.Println(e)
	// Output: 10.583005244258363

	e, _ = stats.MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(99))
	fmt.Println(e)
	// Output: 6

	cor, _ := stats.Correlation([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 5, 6})
	fmt.Println(cor)
	// Output: 0.9912407071619302

	sp, _ := stats.Spearman([]float64{1, 2, 3, 4, 5}, []float64{5, 6, 7, 8, 7})
	fmt.Println(sp)
	// Output: 0.8207826816681233

	ac, _ := stats.AutoCorrelation([]float64{1, 2, 3, 4, 5}, 1)
	fmt.Println(ac)
	// Output: 0.4

	// Sample from a normal distribution with mean=0, std=1
	normSamples := stats.NormSample(0, 1, 5)
	fmt.Println(len(normSamples))
	// Output: 5

	sig, _ := stats.Sigmoid([]float64{3.0, 1.0, 2.1})
	fmt.Println(sig)
	// Output: [0.9525741268224334 0.7310585786300049 0.8909031788043871]

	sm, _ := stats.SoftMax([]float64{3.0, 1.0, 0.2})
	fmt.Println(sm)
	// Output: [0.8360188027814407 0.11314284146556013 0.05083835575299916]

	e, _ = stats.Entropy([]float64{1.1, 2.2, 3.3})
	fmt.Println(e)
	// Output: 1.0114042647073518

	p := 0.5
	begin := 1
	end := 2
	chance, _ := stats.ProbGeom(begin, end, p)
	fmt.Println(chance)
	// Output: 0.25

	prob1 := 0.5
	exp, _ := stats.ExpGeom(prob1)
	fmt.Println(exp)
	// Output:

	prob2 := 0.5
	vari, _ := stats.VarGeom(prob2)
	fmt.Println(vari)
	// Output: 2

	// Z-test: test if sample mean differs from known population mean
	z, zp, _ := stats.ZTest([]float64{2.5, 3.1, 2.8, 3.0, 2.9}, nil, 2.5, 0.3)
	fmt.Println(z, zp)
	// Output: 2.981... 0.002...

	// T-test: test if two sample means differ
	tt, tp, _ := stats.TTest([]float64{5.1, 5.5, 4.8, 5.2, 5.0}, []float64{4.2, 4.8, 4.0, 4.5, 4.3}, 0)
	fmt.Println(tt, tp)
	// Output: 3.279... 0.011...

	description, _ := stats.Describe([]float64{1.0, 2.0, 3.0}, true, &[]float64{25.0, 50.0, 75.0})
	fmt.Println(description.String(2))

	// -----------------------------
	// Skewness (Shape of distribution)
	// -----------------------------

	// Symmetric data
	skew, _ := stats.Skewness([]float64{2, 4, 6, 8, 10})
	fmt.Println("Skewness (symmetric):", skew)
	// Output: ~0 (distribution is symmetric)

	// Right-skewed data (long tail on the right)
	skew, _ = stats.Skewness([]float64{1, 2, 2, 3, 9})
	fmt.Println("Skewness (right-skewed):", skew)
	// Output: positive value (> 0)

	// Left-skewed data (long tail on the left)
	skew, _ = stats.Skewness([]float64{9, 8, 8, 7, 1})
	fmt.Println("Skewness (left-skewed):", skew)
	// Output: negative value (< 0)

	// -----------------------------
	// Kurtosis (Tailedness of distribution)
	// -----------------------------

	kurt, _ := stats.Kurtosis([]float64{2, 4, 6, 8, 10})
	fmt.Println(kurt)
	// Output: -1.3

	kurt, _ = stats.PopulationKurtosis([]float64{2, 4, 6, 8, 10})
	fmt.Println(kurt)
	// Output: -1.3

	kurt, _ = stats.SampleKurtosis([]float64{2, 4, 6, 8, 10})
	fmt.Println(kurt)
	// Output: -1.2000000000000002

	cl, _ := stats.Clip([]float64{-2, 3, 7, 12}, 0, 10)
	fmt.Println(cl)
	// Output: [0 3 7 10]

	rsc, _ := stats.Rescale([]float64{2, 4, 6, 8, 10})
	fmt.Println(rsc)
	// Output: [0 0.25 0.5 0.75 1]

	tmean, _ := stats.TrimmedMean([]float64{1, 2, 3, 4, 100}, 0.2)
	fmt.Println(tmean)
	// Output: 3

	win, _ := stats.Winsorize([]float64{1, 2, 3, 4, 100}, 0.2)
	fmt.Println(win)
	// Output: [2 2 3 4 4]

	mmed, _ := stats.MovingMedian([]float64{1, 9, 2, 8, 3}, 3)
	fmt.Println(mmed)
	// Output: [2 8 3]

	mmin, _ := stats.MovingMin([]float64{1, 9, 2, 8, 3}, 3)
	fmt.Println(mmin)
	// Output: [1 2 2]

	mmax, _ := stats.MovingMax([]float64{1, 9, 2, 8, 3}, 3)
	fmt.Println(mmax)
	// Output: [9 9 8]

	msum, _ := stats.MovingSum([]float64{1, 2, 3, 4, 5}, 3)
	fmt.Println(msum)
	// Output: [6 9 12]

	ewma, _ := stats.EWMA([]float64{1, 2, 3, 4, 5}, 0.5)
	fmt.Println(ewma)
	// Output: [1 1.5 2.25 3.125 4.0625]

	sem, _ := stats.SEM([]float64{1, 2, 3, 4, 5})
	fmt.Println(sem)
	// Output: 0.7071067811865476

	rms, _ := stats.RMS([]float64{1, 2, 3, 4, 5})
	fmt.Println(rms)
	// Output: 3.3166247903554

	prod, _ := stats.Product([]float64{1, 2, 3, 4})
	fmt.Println(prod)
	// Output: 24

	pos, _ := stats.PercentileOfScore([]float64{1, 2, 3, 4, 5}, 4)
	fmt.Println(pos)
	// Output: 70

	kt, _ := stats.KendallTau([]float64{1, 2, 3, 4, 5}, []float64{1, 3, 2, 4, 5})
	fmt.Println(kt)
	// Output: 0.8

	counts, edges, _ := stats.Histogram([]float64{1, 2, 2, 3, 5}, 4)
	fmt.Println(counts, edges)
	// Output: [1 2 1 1] [1 2 3 4 5]

	interp, _ := stats.Interp([]float64{0, 1.5, 3}, []float64{1, 2, 3}, []float64{10, 20, 30})
	fmt.Println(interp)
	// Output: [10 15 30]

}
