package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {

	var d stats.Float64Data = []float64{1, 2, 3, 4, 4, 5}

	// you could also use arbitrary types like this
	// var d = stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})

	min, _ := d.Min()
	fmt.Println(min) // 1

	max, _ := d.Max()
	fmt.Println(max) // 5

	sum, _ := d.Sum()
	fmt.Println(sum) // 19

	argmax, _ := d.ArgMax()
	fmt.Println(argmax) // 5

	argmin, _ := d.ArgMin()
	fmt.Println(argmin) // 0

	rng, _ := d.Range()
	fmt.Println(rng) // 4

	diff, _ := d.Diff()
	fmt.Println(diff) // [1 1 1 0 1]

	pc, _ := d.PercentChange()
	fmt.Println(pc) // [1 0.5 0.3333333333333333 0 0.25]

	cp, _ := d.CumulativeProduct()
	fmt.Println(cp) // [1 2 6 24 96 480]

	cmax, _ := d.CumulativeMax()
	fmt.Println(cmax) // [1 2 3 4 4 5]

	cmin, _ := d.CumulativeMin()
	fmt.Println(cmin) // [1 1 1 1 1 1]

	wm, _ := d.WeightedMean([]float64{1, 1, 1, 1, 1, 1})
	fmt.Println(wm) // 3.1666666666666665

	cv, _ := d.CoefficientOfVariation()
	fmt.Println(cv) // 0.46482951928041305

	zs, _ := d.ZScore()
	fmt.Println(zs) // [-1.4719601443879742 -0.7925939239012169 -0.11322770341445947 0.5661385170722979 0.5661385170722979 1.2455047375590553]

	rank, _ := d.Rank()
	fmt.Println(rank) // [1 2 3 4.5 4.5 6]

	ma, _ := d.MovingAverage(3)
	fmt.Println(ma) // [2 3 3.6666666666666665 4.333333333333333]

	msd, _ := d.MovingStdDev(3)
	fmt.Println(msd) // [1 1 0.5773502691896257 0.5773502691896257]

	kurt, _ := d.Kurtosis()
	fmt.Println(kurt) // -1.151715976331361

	pkurt, _ := d.PopulationKurtosis()
	fmt.Println(pkurt) // -1.151715976331361

	skurt, _ := d.SampleKurtosis()
	fmt.Println(skurt) // -0.8591715976331361

	cl, _ := d.Clip(2, 4)
	fmt.Println(cl) // [2 2 3 4 4 4]

	rsc, _ := d.Rescale()
	fmt.Println(rsc) // [0 0.25 0.5 0.75 0.75 1]

	tmean, _ := d.TrimmedMean(0.2)
	fmt.Println(tmean) // 3.25

	win, _ := d.Winsorize(0.2)
	fmt.Println(win) // [2 2 3 4 4 4]

	mmed, _ := d.MovingMedian(3)
	fmt.Println(mmed) // [2 3 4 4]

	mmin, _ := d.MovingMin(3)
	fmt.Println(mmin) // [1 2 3 4]

	mmax, _ := d.MovingMax(3)
	fmt.Println(mmax) // [3 4 4 5]

	msum, _ := d.MovingSum(3)
	fmt.Println(msum) // [6 9 11 13]

	ewma, _ := d.EWMA(0.5)
	fmt.Println(ewma) // [1 1.5 2.25 3.125 3.5625 4.28125]

	sem, _ := d.SEM()
	fmt.Println(sem) // 0.6009252125773317

	rms, _ := d.RMS()
	fmt.Println(rms) // 3.4399612400917157

	prod, _ := d.Product()
	fmt.Println(prod) // 480

	pos, _ := d.PercentileOfScore(4)
	fmt.Println(pos) // 66.66666666666667

	kt, _ := d.KendallTau([]float64{2, 1, 4, 3, 5, 6})
	fmt.Println(kt) // 0.6900655593423543

	counts, edges, _ := d.Histogram(4)
	fmt.Println(counts, edges) // [1 1 1 3] [1 2 3 4 5]

	// See https://godoc.org/github.com/montanaflynn/stats#Float64Data
	// or run godoc ./ Float64Data to view all available methods

}
