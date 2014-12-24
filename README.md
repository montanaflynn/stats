# Stats [![Build Status](https://img.shields.io/wercker/ci/548fca786b3ba8733d7f219d.svg)](https://app.wercker.com/project/bykey/2eafc5c6f7c702b53d967aef3b2bb65e) [![Coverage Status](https://img.shields.io/coveralls/montanaflynn/stats.svg)](https://coveralls.io/r/montanaflynn/stats?branch=master) [![GoDoc](https://godoc.org/github.com/montanaflynn/stats?status.svg)](https://godoc.org/github.com/montanaflynn/stats)

A simple stats package for Go that includes many common functions that are missing from the standard library. Pull requests are always welcome.

__The package is still in the early stages so expect the API to change.__ 

### Install

```go
go get github.com/montanaflynn/stats
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/montanaflynn/stats"
)

func main() {
    m := stats.Min([]float64{1.1, 2, 3, 4, 5})
    fmt.Println(m) // 1.1

    m = stats.Max([]float64{1.1, 2, 3, 4, 5})
    fmt.Println(m) // 5

    m = stats.Sum([]float64{1.1, 2.2, 3.3})
    fmt.Println(m) // 6.6

    m = stats.Mean([]float64{1, 2, 3, 4, 5})
    fmt.Println(m) // 3

    m = stats.Median([]float64{1, 2, 3, 4, 5, 6, 7})
    fmt.Println(m) // 4

    s := stats.Mode([]float64{5, 5, 3, 3, 4, 2, 1})
    fmt.Println(s) // [5 3]

    m = stats.VarP([]float64{1,2,3,4,5})
    fmt.Println(m) // 2

    m = stats.VarS([]float64{1,2,3,4,5})
    fmt.Println(m) // 2.5

    m = stats.StdDevP([]float64{1, 2, 3})
    fmt.Println(m) // 0.816496580927726

    m = stats.StdDevS([]float64{1, 2, 3})
    fmt.Println(m) // 1

    m = stats.Percentile([]float64{1, 2, 3, 4, 5}, 75)
    fmt.Println(m) // 4

    m = stats.Round(5.3253543, 3)
    fmt.Println(m) // 5.325

    d := []Coordinate{
        {1, 2.3},
        {2, 3.3},
        {3, 3.7},
        {4, 4.3},
        {5, 5.3},
    }

    r := stats.LinReg(d)
    fmt.Println(r) // [{1 2.3800000000000026} {2 3.0800000000000014} {3 3.7800000000000002} {4 4.479999999999999} {5 5.179999999999998}]

    r = stats.ExpReg(d)
    fmt.Println(r) // [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]

    r = stats.LogReg(d)
    fmt.Println(r) // [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]
}
```

### API

The [entire API documentation](http://godoc.org/github.com/montanaflynn/stats) is available on GoDoc.

### Todos

- Error checking in idiomatic Go style
- Regression projection modeling
- [Mathmatical Constants](http://en.wikipedia.org/wiki/Mathematical_constant)
- Seperate usage example directory

### MIT license

Copyright (c) 2014, Montana Flynn (http://anonfunction.com/)
