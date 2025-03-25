# Data Loading Guide

This guide explains how to load data into the stats package, including supported data types and the use of the `LoadRawData` function. We'll provide examples of loading data from various sources and formats.

## Supported Data Types

The stats package supports loading data from a wide variety of data types. The `LoadRawData` function can handle the following types:

- Slices of numeric types (int, uint, float64, etc.)
- Slices of bool
- Slices of string
- Slices of time.Duration
- Maps with int keys and numeric, bool, or string values
- String (space-separated values)
- io.Reader (for reading from files or other input streams)

## Using LoadRawData

The `LoadRawData` function is the primary method for loading data into the stats package. It converts various input types into a `Float64Data` slice, which can be used with other functions in the package.

### Basic Usage

Here's a simple example of using `LoadRawData`:

```go
import "github.com/montanaflynn/stats"

// Load data from a slice of integers
data := stats.LoadRawData([]int{1, 2, 3, 4, 5})

// Now you can use this data with other stats functions
mean, _ := stats.Mean(data)
fmt.Println(mean) // Output: 3
```

### Loading Different Data Types

Let's look at examples of loading different data types:

1. Slice of mixed types:

```go
mixedData := stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})
```

2. Slice of booleans:

```go
boolData := stats.LoadRawData([]bool{true, false, true, true})
```

3. Map with int keys:

```go
mapData := stats.LoadRawData(map[int]float64{1: 1.1, 2: 2.2, 3: 3.3})
```

4. String (space-separated values):

```go
stringData := stats.LoadRawData("1 2 3 4 5")
```

5. Reading from a file:

```go
file, _ := os.Open("data.txt")
defer file.Close()
fileData := stats.LoadRawData(file)
```

## Working with Float64Data

After loading data using `LoadRawData`, you get a `Float64Data` type, which is a slice of float64 values. This type has several methods that can be called directly:

```go
var d stats.Float64Data = stats.LoadRawData([]int{1, 2, 3, 4, 4, 5})

min, _ := d.Min()
fmt.Println(min) // Output: 1

max, _ := d.Max()
fmt.Println(max) // Output: 5

sum, _ := d.Sum()
fmt.Println(sum) // Output: 19
```

## Best Practices

1. Always check for errors when calling `LoadRawData` or any of the `Float64Data` methods.
2. When working with large datasets, consider loading data in chunks to manage memory usage.
3. If you're frequently working with a specific data format, you might want to create a custom loader function that uses `LoadRawData` internally.

## Conclusion

The `LoadRawData` function provides a flexible way to load various data types into the stats package. By converting different input types to `Float64Data`, it allows you to use a consistent data format across all the statistical functions provided by the package.

For more information on available methods and functions, refer to the [stats package documentation](https://godoc.org/github.com/montanaflynn/stats).