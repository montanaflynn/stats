# Getting Started with Stats

Stats is a comprehensive Golang statistics package with no dependencies. This guide will help you get started with installing and using the package in your projects.

## Installation

To install the Stats package, use the following command:

```bash
go get github.com/montanaflynn/stats
```

## Basic Usage

Here are some basic examples to get you started with the Stats package:

### Importing the Package

First, import the package in your Go file:

```go
import "github.com/montanaflynn/stats"
```

### Loading Data

You can load data into the `Float64Data` type using various methods:

```go
// Using a slice of float64
data := stats.Float64Data{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}

// Using LoadRawData for mixed types
mixedData := stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})

// Using LoadRawData for integers
intData := stats.LoadRawData([]int{1, 2, 3, 4, 5})
```

### Basic Statistical Functions

Here are some examples of basic statistical functions:

```go
data := stats.Float64Data{1, 2, 3, 4, 5}

// Calculate mean
mean, _ := stats.Mean(data)
fmt.Println("Mean:", mean)

// Calculate median
median, _ := stats.Median(data)
fmt.Println("Median:", median)

// Calculate mode
mode, _ := stats.Mode(data)
fmt.Println("Mode:", mode)

// Calculate standard deviation
stdDev, _ := stats.StandardDeviation(data)
fmt.Println("Standard Deviation:", stdDev)
```

### Percentiles and Quartiles

```go
data := stats.Float64Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Calculate 75th percentile
percentile, _ := stats.Percentile(data, 75)
fmt.Println("75th Percentile:", percentile)

// Calculate quartiles
quartiles, _ := stats.Quartile(data)
fmt.Printf("Quartiles: Q1=%.2f, Q2=%.2f, Q3=%.2f\n", quartiles.Q1, quartiles.Q2, quartiles.Q3)
```

### Regression Analysis

The Stats package also provides functions for regression analysis:

```go
coordinates := []stats.Coordinate{
    {1, 2.3},
    {2, 3.3},
    {3, 3.7},
    {4, 4.3},
    {5, 5.3},
}

// Linear regression
linearRegression, _ := stats.LinearRegression(coordinates)
fmt.Println("Linear Regression:", linearRegression)

// Exponential regression
expRegression, _ := stats.ExponentialRegression(coordinates)
fmt.Println("Exponential Regression:", expRegression)
```

## Advanced Usage

For more advanced usage and a complete list of available functions, please refer to the [full documentation](https://godoc.org/github.com/montanaflynn/stats) or the [pkg.go.dev](https://pkg.go.dev/github.com/montanaflynn/stats) page.

## Examples

You can find more examples in the `examples/main.go` file in the Stats repository. These examples cover a wide range of statistical functions and use cases.

## Contributing

If you'd like to contribute to the Stats package, please check the [README.md](https://github.com/montanaflynn/stats) file for contribution guidelines and instructions.

## License

Stats is released under the MIT License. For more details, see the [LICENSE](https://github.com/montanaflynn/stats/blob/master/LICENSE) file in the repository.