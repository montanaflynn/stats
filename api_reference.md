# API Reference

This page provides a complete API reference for all exported functions and types in the stats package. The reference is organized by categories.

## Table of Contents

- [Descriptive Statistics](#descriptive-statistics)
- [Regression Analysis](#regression-analysis)
- [Data Manipulation](#data-manipulation)
- [Correlation](#correlation)
- [Distribution Functions](#distribution-functions)
- [Types](#types)

## Descriptive Statistics

### Mean

```go
func Mean(input Float64Data) (float64, error)
```

Calculates the average of a slice of numbers.

Parameters:
- `input`: A slice of float64 values

Returns:
- The mean value
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5}
mean, err := stats.Mean(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Mean: %.2f\n", mean)
```

### Median

```go
func Median(input Float64Data) (float64, error)
```

Calculates the median number in a slice of numbers.

Parameters:
- `input`: A slice of float64 values

Returns:
- The median value
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 3, 5, 7, 9}
median, err := stats.Median(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Median: %.2f\n", median)
```

### Mode

```go
func Mode(input Float64Data) ([]float64, error)
```

Calculates the mode (most frequent value(s)) of a slice of numbers.

Parameters:
- `input`: A slice of float64 values

Returns:
- A slice containing the mode value(s)
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 2, 3, 3, 3, 4, 5}
mode, err := stats.Mode(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Mode: %v\n", mode)
```

### Variance

```go
func Variance(input Float64Data) (float64, error)
```

Calculates the population variance of a dataset.

Parameters:
- `input`: A slice of float64 values

Returns:
- The variance value
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5}
variance, err := stats.Variance(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Variance: %.2f\n", variance)
```

### StandardDeviation

```go
func StandardDeviation(input Float64Data) (float64, error)
```

Calculates the standard deviation of a dataset.

Parameters:
- `input`: A slice of float64 values

Returns:
- The standard deviation value
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5}
stdDev, err := stats.StandardDeviation(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Standard Deviation: %.2f\n", stdDev)
```

## Regression Analysis

### LinearRegression

```go
func LinearRegression(s Series) (regressions Series, err error)
```

Performs least squares linear regression on a data series.

Parameters:
- `s`: A Series of Coordinate points

Returns:
- A Series of Coordinate points representing the regression line
- An error if the input is empty

Example:
```go
data := stats.Series{
    {X: 1, Y: 2},
    {X: 2, Y: 3},
    {X: 3, Y: 5},
    {X: 4, Y: 6},
}
regression, err := stats.LinearRegression(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Linear Regression: %v\n", regression)
```

### ExponentialRegression

```go
func ExponentialRegression(s Series) (regressions Series, err error)
```

Performs exponential regression on a data series.

Parameters:
- `s`: A Series of Coordinate points

Returns:
- A Series of Coordinate points representing the regression curve
- An error if the input is empty or contains negative Y values

Example:
```go
data := stats.Series{
    {X: 1, Y: 2},
    {X: 2, Y: 4},
    {X: 3, Y: 8},
    {X: 4, Y: 16},
}
regression, err := stats.ExponentialRegression(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Exponential Regression: %v\n", regression)
```

### LogarithmicRegression

```go
func LogarithmicRegression(s Series) (regressions Series, err error)
```

Performs logarithmic regression on a data series.

Parameters:
- `s`: A Series of Coordinate points

Returns:
- A Series of Coordinate points representing the regression curve
- An error if the input is empty

Example:
```go
data := stats.Series{
    {X: 1, Y: 0},
    {X: 2, Y: 0.69},
    {X: 3, Y: 1.09},
    {X: 4, Y: 1.38},
}
regression, err := stats.LogarithmicRegression(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Logarithmic Regression: %v\n", regression)
```

## Data Manipulation

### Sample

```go
func Sample(input Float64Data, takenum int, replacement bool) ([]float64, error)
```

Returns a sample from the input data with or without replacement.

Parameters:
- `input`: A slice of float64 values
- `takenum`: The number of samples to take
- `replacement`: Whether to sample with replacement

Returns:
- A slice of sampled values
- An error if the input is empty or if `takenum` is greater than the input length (when sampling without replacement)

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sample, err := stats.Sample(data, 5, false)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Sample: %v\n", sample)
```

### CumulativeSum

```go
func CumulativeSum(input Float64Data) ([]float64, error)
```

Calculates the cumulative sum of the input slice.

Parameters:
- `input`: A slice of float64 values

Returns:
- A slice containing the cumulative sum
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5}
cumSum, err := stats.CumulativeSum(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Cumulative Sum: %v\n", cumSum)
```

## Correlation

### Correlation

```go
func Correlation(data1, data2 Float64Data) (float64, error)
```

Calculates the correlation coefficient between two sets of data.

Parameters:
- `data1`: First set of float64 values
- `data2`: Second set of float64 values

Returns:
- The correlation coefficient
- An error if the inputs are empty or have different lengths

Example:
```go
data1 := stats.Float64Data{1, 2, 3, 4, 5}
data2 := stats.Float64Data{2, 4, 5, 4, 5}
correlation, err := stats.Correlation(data1, data2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Correlation: %.2f\n", correlation)
```

### AutoCorrelation

```go
func AutoCorrelation(data Float64Data, lags int) (float64, error)
```

Calculates the auto-correlation of a signal with a delayed copy of itself.

Parameters:
- `data`: A slice of float64 values representing the signal
- `lags`: The number of lags to consider

Returns:
- The auto-correlation value
- An error if the input is empty

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5, 4, 3, 2, 1}
autoCorr, err := stats.AutoCorrelation(data, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Auto-correlation: %.2f\n", autoCorr)
```

## Distribution Functions

### NormPdf

```go
func NormPdf(x float64, loc float64, scale float64) float64
```

Calculates the probability density function for a normal distribution.

Parameters:
- `x`: The input value
- `loc`: The mean of the distribution
- `scale`: The standard deviation of the distribution

Returns:
- The probability density at x

Example:
```go
pdf := stats.NormPdf(0, 0, 1)
fmt.Printf("Normal PDF at x=0: %.4f\n", pdf)
```

### NormCdf

```go
func NormCdf(x float64, loc float64, scale float64) float64
```

Calculates the cumulative distribution function for a normal distribution.

Parameters:
- `x`: The input value
- `loc`: The mean of the distribution
- `scale`: The standard deviation of the distribution

Returns:
- The cumulative probability at x

Example:
```go
cdf := stats.NormCdf(1.96, 0, 1)
fmt.Printf("Normal CDF at x=1.96: %.4f\n", cdf)
```

## Types

### Float64Data

```go
type Float64Data []float64
```

A named type for `[]float64` with helper methods for common statistical operations.

Example:
```go
data := stats.Float64Data{1, 2, 3, 4, 5}
mean, _ := data.Mean()
median, _ := data.Median()
fmt.Printf("Mean: %.2f, Median: %.2f\n", mean, median)
```

### Series

```go
type Series []Coordinate
```

A container for a series of data points used in regression analysis.

### Coordinate

```go
type Coordinate struct {
    X, Y float64
}
```

Holds the X and Y values for a single data point in a Series.

Example:
```go
point := stats.Coordinate{X: 1, Y: 2}
series := stats.Series{point}
```

This API reference covers the main functions and types exported by the stats package. For more detailed information on specific functions or advanced usage, please refer to the package documentation and source code.