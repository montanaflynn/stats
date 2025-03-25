# Regression Analysis Guide

This guide explains how to perform regression analysis using the `stats` package. We'll cover linear, exponential, and logarithmic regression, including data preparation, analysis execution, and result interpretation.

## Table of Contents

1. [Introduction](#introduction)
2. [Data Preparation](#data-preparation)
3. [Linear Regression](#linear-regression)
4. [Exponential Regression](#exponential-regression)
5. [Logarithmic Regression](#logarithmic-regression)
6. [Interpreting Results](#interpreting-results)
7. [Examples](#examples)

## Introduction

Regression analysis is a statistical method used to model the relationship between variables. The `stats` package provides functions for three types of regression:

- Linear Regression
- Exponential Regression
- Logarithmic Regression

These methods can help you understand and predict relationships in your data.

## Data Preparation

Before performing regression analysis, you need to prepare your data. The `stats` package uses the `Series` type, which is a slice of `Coordinate` structs:

```go
type Series []Coordinate

type Coordinate struct {
    X, Y float64
}
```

To create a series, you can do the following:

```go
data := stats.Series{
    {X: 1, Y: 2.3},
    {X: 2, Y: 3.3},
    {X: 3, Y: 3.7},
    {X: 4, Y: 4.3},
    {X: 5, Y: 5.3},
}
```

## Linear Regression

Linear regression finds the best-fitting straight line through the points in your data.

### Function Signature

```go
func LinearRegression(s Series) (regressions Series, err error)
```

### Usage

```go
regressions, err := stats.LinearRegression(data)
if err != nil {
    // Handle error
}
```

## Exponential Regression

Exponential regression fits a curve of the form y = a * e^(bx) to your data.

### Function Signature

```go
func ExponentialRegression(s Series) (regressions Series, err error)
```

### Usage

```go
regressions, err := stats.ExponentialRegression(data)
if err != nil {
    // Handle error
}
```

Note: All Y values must be positive for exponential regression.

## Logarithmic Regression

Logarithmic regression fits a curve of the form y = a + b * ln(x) to your data.

### Function Signature

```go
func LogarithmicRegression(s Series) (regressions Series, err error)
```

### Usage

```go
regressions, err := stats.LogarithmicRegression(data)
if err != nil {
    // Handle error
}
```

## Interpreting Results

All regression functions return a `Series` of `Coordinate` structs representing the fitted values. You can use these to:

1. Plot the regression line alongside your original data
2. Make predictions for new X values
3. Assess the fit of the model visually

To evaluate the quality of the fit quantitatively, you may need to implement additional metrics such as R-squared or mean squared error.

## Examples

Here's a complete example demonstrating how to perform and compare different types of regression:

```go
package main

import (
    "fmt"
    "github.com/montanaflynn/stats"
)

func main() {
    // Prepare data
    data := stats.Series{
        {X: 1, Y: 2.3},
        {X: 2, Y: 3.3},
        {X: 3, Y: 3.7},
        {X: 4, Y: 4.3},
        {X: 5, Y: 5.3},
    }

    // Perform linear regression
    linear, err := stats.LinearRegression(data)
    if err != nil {
        fmt.Println("Linear regression error:", err)
        return
    }

    // Perform exponential regression
    exponential, err := stats.ExponentialRegression(data)
    if err != nil {
        fmt.Println("Exponential regression error:", err)
        return
    }

    // Perform logarithmic regression
    logarithmic, err := stats.LogarithmicRegression(data)
    if err != nil {
        fmt.Println("Logarithmic regression error:", err)
        return
    }

    // Print results
    fmt.Println("Original data:", data)
    fmt.Println("Linear regression:", linear)
    fmt.Println("Exponential regression:", exponential)
    fmt.Println("Logarithmic regression:", logarithmic)
}
```

This example demonstrates how to use all three regression types on the same dataset, allowing you to compare their outputs and choose the best fit for your data.

Remember to visualize your results for a better understanding of how well each regression type fits your data. You may want to use a plotting library to create scatter plots of your original data with the regression lines overlaid.