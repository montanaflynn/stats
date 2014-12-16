# Stats

A simple stats package for Go that includes many common functions that are missing from the standard library. The package is still in the early stages so expect the API to change. Pull requests are always welcome.

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

    m = stats.Mode([]float64{5, 5, 3, 3, 4, 2, 1})
    fmt.Println(m) // [3.0 5.0]

    m = stats.StandardDev([]float64{1, 2, 3})
    fmt.Println(m) // 0.816496580927726

    m = stats.Round(5.3253543, 3)
    fmt.Println(m) // 5.325
}
```

### API Documentation


#### func  Min

```go
func Min(input []float64) (min float64)
```
Find the lowest number in a slice

#### func  Max

```go
func Max(input []float64) (max float64)
```
Find the highest number in a slice

#### func  Sum

```go
func Sum(input []float64) (sum float64)
```
Add all the numbers of a slice together

#### func  Mean

```go
func Mean(input []float64) (mean float64)
```
Get the average of a slice of numbers

#### func  Median

```go
func Median(input []float64) (median float64)
```
Get the median number in a slice of numbers

#### func  Mode

```go
func Mode(input []float64) (mode []float64)
```
Get the mode of a slice of numbers

#### func  StandardDev

```go
func StandardDev(input []float64) (sdev float64)
```
Find the amount of variation from the average also known as standard deviation

#### func  Round

```go
func Round(input float64, places int) (rounded float64)
```
Round a float to a specific decimal place or precision

### Todos

- Error checking in idiomatic Go style
- Add Percentiles function
- Add linear and exponential regression 

### MIT license

Copyright (c) 2014, Montana Flynn (http://anonfunction.com/)