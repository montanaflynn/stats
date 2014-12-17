# Stats [![Build Status](https://img.shields.io/wercker/ci/548fca786b3ba8733d7f219d.svg?style=flat-square)](https://app.wercker.com/project/bykey/2eafc5c6f7c702b53d967aef3b2bb65e) [![Coverage Status](https://img.shields.io/coveralls/montanaflynn/stats.svg?style=flat-square)](https://coveralls.io/r/montanaflynn/stats?branch=master)

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
    fmt.Println(s) // [3.0 5.0]

    m = stats.StdDev([]float64{1, 2, 3})
    fmt.Println(m) // 0.816496580927726

    m = stats.Round(5.3253543, 3)
    fmt.Println(m) // 5.325

    m = stats.Percentile([]float64{1, 2, 3, 4, 5}, 75)
    fmt.Println(m) // 4.0

    i := stats.Float64ToInt(-234.0234)
    fmt.Println(m) // --234
}
```

### API Documentation

#### func  Min

```go
func Min(input []float64) (min float64)
```
Min finds the lowest number in a slice

#### func  Max

```go
func Max(input []float64) (max float64)
```
Max finds the highest number in a slice

#### func  Sum

```go
func Sum(input []float64) (sum float64)
```
Sum adds all the numbers of a slice together

#### func  Mean

```go
func Mean(input []float64) (mean float64)
```
Mean gets the average of a slice of numbers

#### func  Median

```go
func Median(input []float64) (median float64)
```
Median gets the median number in a slice of numbers

#### func  Mode

```go
func Mode(input []float64) (mode []float64)
```
Mode gets the mode of a slice of numbers

#### func  StdDev

```go
func StdDev(input []float64) (sdev float64)
```
StdDev gets the amount of variation from the average

#### func  Round

```go
func Round(input float64, places int) (rounded float64)
```
Round a float to a specific decimal place or precision

#### func  Percentile

```go
func Percentile(input []float64, percent float64) (percentile float64)
```
Percentile finds the relative standing in a slice of floats

#### func  Float64ToInt

```go
func Float64ToInt(input float64) (output int)
```
Float64ToInt rounds a float64 to an int

### Todos

- Error checking in idiomatic Go style
- Add linear and exponential regression 

### MIT license

Copyright (c) 2014, Montana Flynn (http://anonfunction.com/)
