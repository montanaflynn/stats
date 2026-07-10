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

	// See https://godoc.org/github.com/montanaflynn/stats#Float64Data
	// or run godoc ./ Float64Data to view all available methods

}
