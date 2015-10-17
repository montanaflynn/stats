# Stats [![Build Status][travis-svg]][travis-url] [![Coverage Status][coveralls-svg]][coveralls-url] [![API Documentation][godoc-svg]][godoc-url] [![MIT License][license-svg]][license-url]

A statistics package with many functions missing from the Golang standard library.

Currently only `float64` and `[]float64` data is supported due to the lack of generics. However, if the library becomes popular and there is demand I could also add support for the other number types. See the [CHANGELOG.md](https://github.com/montanaflynn/stats/blob/master/CHANGELOG.md) for API changes and tagged releases you can vendor into your projects.

> Statistics are used much like a drunk uses a lamppost: for support, not illumination. **- Vin Scully**

## Installation

```
go get github.com/montanaflynn/stats
```

**Protip:** `go get -u github.com/montanaflynn/stats` updates stats to the latest version.

## Usage

All the functions can be seen in [examples/main.go](https://github.com/montanaflynn/stats/blob/master/examples/main.go) but here's a little taste:

```go
// start with the some source data to use
var data = []float64{1, 2, 3, 4, 4, 5}

median, _ := stats.Median(data)
fmt.Println(median) // 3.5

roundedMedian, _ := stats.Round(median, 0)
fmt.Println(roundedMedian) // 4
```

**Protip:** You can call methods directly on the data ([example](https://github.com/montanaflynn/stats/blob/master/examples/methods.go)) if using the Float64Data type.

```
var d stats.Float64Data = data

max, _ := d.Max()
fmt.Println(max) // 5
```

## Documentation

The [entire API documentation](http://godoc.org/github.com/montanaflynn/stats) is available on GoDoc. 

Types: [`Float64Data`](http://godoc.org/github.com/montanaflynn/stats#Float64Data), [`Series`](http://godoc.org/github.com/montanaflynn/stats#Series), [`Coordinate`](http://godoc.org/github.com/montanaflynn/stats#Coordinate), [`Quartiles`](http://godoc.org/github.com/montanaflynn/stats#Quartiles), [`Outliers`](http://godoc.org/github.com/montanaflynn/stats#Outliers)

Functions: [`Min`](http://godoc.org/github.com/montanaflynn/stats#Min), [`Max`](http://godoc.org/github.com/montanaflynn/stats#Max), [`Sum`](http://godoc.org/github.com/montanaflynn/stats#Sum), [`Mean`](http://godoc.org/github.com/montanaflynn/stats#Mean), [`Median`](http://godoc.org/github.com/montanaflynn/stats#Median), [`Mode`](http://godoc.org/github.com/montanaflynn/stats#Mode), [`Sample`](http://godoc.org/github.com/montanaflynn/stats#Sample), [`Round`](http://godoc.org/github.com/montanaflynn/stats#Round), [`StandardDeviation`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviation), [`StandardDeviationPopulation`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviationPopulation), [`StandardDeviationSample`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviationSample), [`Percentile`](http://godoc.org/github.com/montanaflynn/stats#Percentile), [`PercentileNearestRank`](http://godoc.org/github.com/montanaflynn/stats#PercentileNearestRank), [`LinearRegression`](http://godoc.org/github.com/montanaflynn/stats#LinearRegression), [`ExponentialRegression`](http://godoc.org/github.com/montanaflynn/stats#ExponentialRegression), [`LogarithmicRegression`](http://godoc.org/github.com/montanaflynn/stats#LogarithmicRegression), [`Variance`](http://godoc.org/github.com/montanaflynn/stats#Variance), [`PopulationVariance`](http://godoc.org/github.com/montanaflynn/stats#PopulationVariance), [`SampleVariance`](http://godoc.org/github.com/montanaflynn/stats#SampleVariance), [`Quartile`](http://godoc.org/github.com/montanaflynn/stats#Quartile), [`InterQuartileRange`](http://godoc.org/github.com/montanaflynn/stats#InterQuartileRange), [`Midhinge`](http://godoc.org/github.com/montanaflynn/stats#Midhinge), [`Trimean`](http://godoc.org/github.com/montanaflynn/stats#Trimean), [`QuartileOutliers`](http://godoc.org/github.com/montanaflynn/stats#QuartileOutliers), [`GeometricMean`](http://godoc.org/github.com/montanaflynn/stats#GeometricMean), [`HarmonicMean`](http://godoc.org/github.com/montanaflynn/stats#HarmonicMean), [`Covariance`](http://godoc.org/github.com/montanaflynn/stats#Covariance), [`Correlation`](http://godoc.org/github.com/montanaflynn/stats#Correlation)

**Protip:** You can view go docs offline in the terminal or browser:

```
# Show current package docs
godoc ./ 

# Show current package's function or type doc
godoc ./ Min

# Create a local web server with the docs on port 4444
godoc -http=:4444
```

## Contributing

If you have any suggestions, criticism or bug reports please [create an issue](https://github.com/montanaflynn/stats/issues) and I'll do my best to accommodate you. In addition simply starring the repo would show your support for the project and be very much appreciated! 

### Pull Requests

Pull request are always welcome no matter how big or small. Here's an easy way to do it:

1. Fork it and clone your fork
2. Create new branch (`git checkout -b cool-new-thing`)
3. Make the desired changes
4. Ensure tests pass (`go test -cover` or `make test`)
5. Commit changes (`git commit -am 'Add cool new thing'`)
6. Push branch (`git push origin cool-new-thing`)
7. Submit pull request

#### Extra Credit 

To make things as seamless as possible please also consider the following steps:

- Update `README.md` to include new public types or functions in the documentation section.

- Update `examples/main.go` with a simple example of the new feature.

- Keep 100% code coverage, here's an easy way to see what code is not covered:

```go test -coverprofile=coverage.out; go tool cover -html="coverage.out"```

- Run [`gometalinter`](https://github.com/alecthomas/gometalinter) and make your code pass.

- Squash needless commits into single units of work with `git rebase -i new-feature`.


#### Makefile

I've included a [Makefile](https://github.com/montanaflynn/stats/blob/master/Makefile) that has a lot of helper targets for common actions such as linting, testing, coverage reporting and more.

**Protip:** `watch -n 1 make check` will continuously format and test your code.

## MIT License

Copyright (c) 2014-2015 Montana Flynn <http://anonfunction.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORpublicS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[travis-url]: https://travis-ci.org/montanaflynn/stats
[travis-svg]: https://img.shields.io/travis/montanaflynn/stats.svg

[coveralls-url]: https://coveralls.io/r/montanaflynn/stats?branch=master
[coveralls-svg]: https://img.shields.io/coveralls/montanaflynn/stats.svg

[godoc-url]: https://godoc.org/github.com/montanaflynn/stats
[godoc-svg]: https://godoc.org/github.com/montanaflynn/stats?status.svg

[license-url]: https://github.com/montanaflynn/stats/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg