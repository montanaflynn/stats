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

	a, _ = stats.Sum([]float64{1.1, 2.2, 3.3})
	fmt.Println(a)
	// Output: 6.6

	cs, _ := stats.CumulativeSum([]float64{1.1, 2.2, 3.3})
	fmt.Println(cs) // [1.1 3.3000000000000003 6.6]

	a, _ = stats.Mean([]float64{1, 2, 3, 4, 5})
	fmt.Println(a)
	// Output: 3

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

	a, _ = stats.Percentile([]float64{1, 2, 3, 4, 5}, 75)
	fmt.Println(a)
	// Output: 4

	a, _ = stats.PercentileNearestRank([]float64{35, 20, 15, 40, 50}, 75)
	fmt.Println(a)
	// Output: 40

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
	// Output: 2.189

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

	ac, _ := stats.AutoCorrelation([]float64{1, 2, 3, 4, 5}, 1)
	fmt.Println(ac)
	// Output: 0.4

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

	description, _ := stats.Describe([]float64{1.0, 2.0, 3.0}, true, &[]float64{25.0, 50.0, 75.0})
	fmt.Println(description.String(2))
}
