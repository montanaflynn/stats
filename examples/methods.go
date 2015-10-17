package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {

	var d stats.Float64Data = []float64{1, 2, 3, 4, 4, 5}

	// could also use type directly like this
	// var d = stats.Float64Data{1, 2, 3, 4, 4, 5}

	min, _ := d.Min()
	fmt.Println(min) // 1

	max, _ := d.Max()
	fmt.Println(max) // 5

	sum, _ := d.Sum()
	fmt.Println(sum) // 19

	// See https://godoc.org/github.com/montanaflynn/stats#Float64Data
	// or run godoc ./ Float64Data to view all available methods

}
