package stats_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/montanaflynn/stats"
// )

// func Example() {
// 	// t.Parallel()
// 	t.Run("LoadRawData", func(t *testing.T) {
// 		// t.Parallel()
// 		data := stats.LoadRawData([]interface{}{1.1, "2", 3})
// 		fmt.Println(data)
// 		// Output: 1.1, 2.0, 3.0, 4
// 	})
// }

// // func Example() {

// // 	// start with some source data to use
// // 	data := []float64{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}
// // 	// you could also use different types like this
// // 	// data := stats.LoadRawData([]int{1, 2, 3, 4, 5})
// // 	// data := stats.LoadRawData([]interface{}{1.1, "2", 3})
// // 	// etc...

// // 	median, _ := Median(data)
// // 	fmt.Println(median)
// // 	// Output: 3.65

// // 	roundedMedian, _ := Round(median, 0)
// // 	fmt.Println(roundedMedian)
// // 	// Output: 4

// // 	a, _ := Mean([]float64{1, 2, 3, 4, 5})
// // 	fmt.Println(a)
// // 	// Output: 3

// // 	a, _ = Median([]float64{1, 2, 3, 4, 5, 6, 7})
// // 	fmt.Println(a)
// // 	// Output: 4

// // 	m, _ := Mode([]float64{5, 5, 3, 3, 4, 2, 1})
// // 	fmt.Println(m)
// // 	// Output: [5 3]

// // 	a, _ = PopulationVariance([]float64{1, 2, 3, 4, 5})
// // 	fmt.Println(a)
// // 	// Output: 2

// // 	a, _ = SampleVariance([]float64{1, 2, 3, 4, 5})
// // 	fmt.Println(a)
// // 	// Output: 2.5

// // 	a, _ = MedianAbsoluteDeviationPopulation([]float64{1, 2, 3})
// // 	fmt.Println(a)
// // 	// Output: 1

// // 	a, _ = StandardDeviationPopulation([]float64{1, 2, 3})
// // 	fmt.Println(a)
// // 	// Output: 0.816496580927726

// // 	a, _ = StandardDeviationSample([]float64{1, 2, 3})
// // 	fmt.Println(a)
// // 	// Output: 1

// // 	a, _ = Percentile([]float64{1, 2, 3, 4, 5}, 75)
// // 	fmt.Println(a)
// // 	// Output: 4

// // 	a, _ = PercentileNearestRank([]float64{35, 20, 15, 40, 50}, 75)
// // 	fmt.Println(a)
// // 	// Output: 40

// // 	c := []Coordinate{
// // 		{1, 2.3},
// // 		{2, 3.3},
// // 		{3, 3.7},
// // 		{4, 4.3},
// // 		{5, 5.3},
// // 	}

// // 	r, _ := LinearRegression(c)
// // 	fmt.Println(r)
// // 	// Output: [{1 2.3800000000000026} {2 3.0800000000000014} {3 3.7800000000000002} {4 4.479999999999999} {5 5.179999999999998}]

// // 	r, _ = ExponentialRegression(c)
// // 	fmt.Println(r)
// // 	// Output: [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]

// // 	r, _ = LogarithmicRegression(c)
// // 	fmt.Println(r)
// // 	// Output: [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]

// // 	s, _ := Sample([]float64{0.1, 0.2, 0.3, 0.4}, 3, false)
// // 	fmt.Println(s)
// // 	// Output: [0.2,0.4,0.3]

// // 	s, _ = Sample([]float64{0.1, 0.2, 0.3, 0.4}, 10, true)
// // 	fmt.Println(s)
// // 	// Output: [0.2,0.2,0.4,0.1,0.2,0.4,0.3,0.2,0.2,0.1]

// // 	q, _ := Quartile([]float64{7, 15, 36, 39, 40, 41})
// // 	fmt.Println(q)
// // 	// Output: {15 37.5 40}

// // 	iqr, _ := InterQuartileRange([]float64{102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118})
// // 	fmt.Println(iqr)
// // 	// Output: 10

// // 	mh, _ := Midhinge([]float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13})
// // 	fmt.Println(mh)
// // 	// Output: 7.5

// // 	tr, _ := Trimean([]float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13})
// // 	fmt.Println(tr)
// // 	// Output: 7.25

// // 	o, _ := QuartileOutliers([]float64{-1000, 1, 3, 4, 4, 6, 6, 6, 6, 7, 8, 15, 18, 100})
// // 	fmt.Printf("%+v\n", o)
// // 	// Output:  {Mild:[15 18] Extreme:[-1000 100]}

// // 	gm, _ := GeometricMean([]float64{10, 51.2, 8})
// // 	fmt.Println(gm)
// // 	// Output: 15.999999999999991

// // 	hm, _ := HarmonicMean([]float64{1, 2, 3, 4, 5})
// // 	fmt.Println(hm)
// // 	// Output: 2.18978102189781

// // 	a, _ = Round(2.18978102189781, 3)
// // 	fmt.Println(a)
// // 	// Output: 2.189

// // 	e, _ := ChebyshevDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
// // 	fmt.Println(e)
// // 	// Output: 6

// // 	e, _ = ManhattanDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
// // 	fmt.Println(e)
// // 	// Output: 24

// // 	e, _ = EuclideanDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2})
// // 	fmt.Println(e)
// // 	// Output: 10.583005244258363

// // 	e, _ = MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(1))
// // 	fmt.Println(e)
// // 	// Output: 24

// // 	e, _ = MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(2))
// // 	fmt.Println(e)
// // 	// Output: 10.583005244258363

// // 	e, _ = MinkowskiDistance([]float64{2, 3, 4, 5, 6, 7, 8}, []float64{8, 7, 6, 5, 4, 3, 2}, float64(99))
// // 	fmt.Println(e)
// // 	// Output: 6
// // }
