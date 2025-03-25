# Descriptive Statistics Guide

This guide provides a comprehensive overview of calculating descriptive statistics using the `stats` package. We'll cover various measures of central tendency, dispersion, and other summary statistics, along with examples and explanations of when to use each statistic.

## Table of Contents
1. [Measures of Central Tendency](#measures-of-central-tendency)
   - [Mean](#mean)
   - [Median](#median)
   - [Mode](#mode)
2. [Measures of Dispersion](#measures-of-dispersion)
   - [Variance](#variance)
   - [Standard Deviation](#standard-deviation)
3. [Percentiles and Quartiles](#percentiles-and-quartiles)
4. [Other Summary Statistics](#other-summary-statistics)
   - [Min and Max](#min-and-max)
   - [Sum](#sum)
   - [Geometric Mean](#geometric-mean)
   - [Harmonic Mean](#harmonic-mean)

## Measures of Central Tendency

### Mean

The mean is the average of a set of numbers. It's calculated by summing all values and dividing by the count of values.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}
mean, err := stats.Mean(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Mean: %.2f\n", mean)
// Output: Mean: 3.00
```

Use the mean when you want to represent the typical value in a dataset, especially when the data is symmetrically distributed.

### Median

The median is the middle value in a sorted list of numbers. For an even number of values, it's the average of the two middle numbers.

```go
data := stats.Float64Data{1, 2, 3, 4, 5, 6, 7}
median, err := stats.Median(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Median: %.2f\n", median)
// Output: Median: 4.00
```

Use the median when you want a measure of central tendency that's less affected by extreme values or when dealing with skewed distributions.

### Mode

The mode is the most frequent value(s) in a dataset.

```go
data := stats.Float64Data{1, 2, 2, 3, 3, 3, 4, 4, 5}
mode, err := stats.Mode(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Mode: %v\n", mode)
// Output: Mode: [3]
```

Use the mode when you want to find the most common value(s) in a dataset, especially for categorical or discrete data.

## Measures of Dispersion

### Variance

Variance measures the average squared deviation from the mean. The `stats` package provides both population and sample variance.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}

// Population variance
popVar, err := stats.PopulationVariance(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Population Variance: %.2f\n", popVar)

// Sample variance
sampleVar, err := stats.SampleVariance(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Sample Variance: %.2f\n", sampleVar)

// Output:
// Population Variance: 2.00
// Sample Variance: 2.50
```

Use variance to quantify the spread of data points. Population variance is used when you have data for an entire population, while sample variance is used when working with a sample of a larger population.

### Standard Deviation

Standard deviation is the square root of variance, providing a measure of dispersion in the same units as the original data.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}

// Population standard deviation
popStdDev, err := stats.StandardDeviationPopulation(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Population Standard Deviation: %.4f\n", popStdDev)

// Sample standard deviation
sampleStdDev, err := stats.StandardDeviationSample(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Sample Standard Deviation: %.4f\n", sampleStdDev)

// Output:
// Population Standard Deviation: 1.4142
// Sample Standard Deviation: 1.5811
```

Use standard deviation to describe the spread of data points around the mean. It's often preferred over variance because it's in the same units as the original data.

## Percentiles and Quartiles

Percentiles divide a dataset into 100 equal parts, while quartiles divide it into four parts.

```go
data := stats.Float64Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 75th percentile
p75, err := stats.Percentile(data, 75)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("75th Percentile: %.2f\n", p75)

// Quartiles
q, err := stats.Quartile(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Quartiles: Q1=%.2f, Q2=%.2f, Q3=%.2f\n", q.Q1, q.Q2, q.Q3)

// Output:
// 75th Percentile: 7.75
// Quartiles: Q1=3.00, Q2=5.50, Q3=8.00
```

Use percentiles and quartiles to understand the distribution of data and identify potential outliers.

## Other Summary Statistics

### Min and Max

Find the minimum and maximum values in a dataset.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}

min, err := stats.Min(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Minimum: %.2f\n", min)

max, err := stats.Max(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Maximum: %.2f\n", max)

// Output:
// Minimum: 1.00
// Maximum: 5.00
```

### Sum

Calculate the sum of all values in a dataset.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}
sum, err := stats.Sum(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Sum: %.2f\n", sum)
// Output: Sum: 15.00
```

### Geometric Mean

The geometric mean is calculated by multiplying all numbers and taking the nth root, where n is the count of numbers.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}
gm, err := stats.GeometricMean(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Geometric Mean: %.4f\n", gm)
// Output: Geometric Mean: 2.6052
```

Use the geometric mean when dealing with data that grows exponentially or when calculating average rates of growth.

### Harmonic Mean

The harmonic mean is the reciprocal of the arithmetic mean of reciprocals.

```go
data := stats.Float64Data{1, 2, 3, 4, 5}
hm, err := stats.HarmonicMean(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Harmonic Mean: %.4f\n", hm)
// Output: Harmonic Mean: 2.1898
```

Use the harmonic mean when working with rates or speeds, especially when you need to calculate an average rate for a given distance or time.

By utilizing these descriptive statistics, you can gain valuable insights into your datasets, understand their distributions, and make informed decisions based on the characteristics of your data.