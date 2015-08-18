# Stats [![Build Status][travis-svg]][travis-url] [![Coverage Status][coveralls-svg]][coveralls-url] [![API Documentation][godoc-svg]][godoc-url]

A statistics package with common functions that are missing from the Golang standard library. 

Currently only `float64` and `[]float64` data is supported due to the lack of generics. However, if the library becomes popular and there is demand I could also add support for the other number types. 

### Installation

```
go get github.com/montanaflynn/stats
```

### Documentation

The [entire API documentation](http://godoc.org/github.com/montanaflynn/stats) is available on GoDoc. 

### Example Usage

All the functions can be seen in [examples/main.go](https://github.com/montanaflynn/stats/blob/master/examples/main.go).

### Contributing

If you have suggestions, criticism or bug reports please [create an issue](https://github.com/montanaflynn/stats/issues).

__Pull requests are always welcome and much appreciated.__ 

### MIT license

Copyright (c) 2014-2015, Montana Flynn (http://anonfunction.com/)

[travis-url]: https://travis-ci.org/montanaflynn/stats
[travis-svg]: https://img.shields.io/travis/montanaflynn/stats.svg

[coveralls-url]: https://coveralls.io/r/montanaflynn/stats?branch=master
[coveralls-svg]: https://img.shields.io/coveralls/montanaflynn/stats.svg

[godoc-url]: https://godoc.org/github.com/montanaflynn/stats
[godoc-svg]: https://godoc.org/github.com/montanaflynn/stats?status.svg
