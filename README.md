# Stats [![Build Status](https://img.shields.io/travis/montanaflynn/stats.svg)](https://travis-ci.org/montanaflynn/stats) [![Coverage Status](https://img.shields.io/coveralls/montanaflynn/stats.svg)](https://coveralls.io/r/montanaflynn/stats?branch=master) [![GoDoc](https://godoc.org/github.com/montanaflynn/stats?status.svg)](https://godoc.org/github.com/montanaflynn/stats)

A stats package including common functions that are missing from the Golang standard library. 

__Pull requests are always welcome and much appreciated.__

### Install

```go
go get github.com/montanaflynn/stats
```

### Usage

The [entire API documentation](http://godoc.org/github.com/montanaflynn/stats) is available on GoDoc.

### Examples

```go
package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	a, _ := Min([]float64{1.1, 2, 3, 4, 5})
	fmt.Println(a) // 1.1

	a, _ = Max([]float64{1.1, 2, 3, 4, 5})
	fmt.Println(a) // 5

	a, _ = Sum([]float64{1.1, 2.2, 3.3})
	fmt.Println(a) // 6.6

	a, _ = Mean([]float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 3

	a, _ = Median([]float64{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a) // 4

	m, _ := Mode([]float64{5, 5, 3, 3, 4, 2, 1})
	fmt.Println(m) // [5 3]

	a, _ = VarP([]float64{1,2,3,4,5})
	fmt.Println(a) // 2

	a, _ = VarS([]float64{1,2,3,4,5})
	fmt.Println(a) // 2.5

	a, _ = StdDevP([]float64{1, 2, 3})
	fmt.Println(a) // 0.816496580927726

	a, _ = StdDevS([]float64{1, 2, 3})
	fmt.Println(a) // 1

	a, _ = Percentile([]float64{1, 2, 3, 4, 5}, 75)
	fmt.Println(a) // 4

	a, _ = Round(5.3253543, 3)
	fmt.Println(a) // 5.325

	c := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := LinReg(c)
	fmt.Println(r) // [{1 2.3800000000000026} {2 3.0800000000000014} {3 3.7800000000000002} {4 4.479999999999999} {5 5.179999999999998}]

	r, _ = ExpReg(c)
	fmt.Println(r) // [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]

	r, _ = LogReg(c)
	fmt.Println(r) // [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]

	s, _ := Sample([]float64{0.1,0.2,0.3,0.4}, 3, false)
	fmt.Println(s) //[0.2,0.4,0.3]

	s, _ = Sample([]float64{0.1,0.2,0.3,0.4}, 10, true)
	fmt.Println(s) //[0.2,0.2,0.4,0.1,0.2,0.4,0.3,0.2,0.2,0.1]
}
```

### Todos

- Trimean, Chi-squared test, and [more](https://en.wikipedia.org/wiki/Statistics)
- Seperate example directory with real world projects
- Performance improvements plus a complete benchmark suite

### MIT license

Copyright (c) 2014-2015, Montana Flynn (http://anonfunction.com/)
