# Stats [![Build Status][travis-svg]][travis-url] [![Coverage Status][coveralls-svg]][coveralls-url] [![API Documentation][godoc-svg]][godoc-url]

A statistics package with common functions that are missing from the Golang standard library. 

Currently only `float64` and `[]float64` data is supported due to the lack of generics. However, if the library becomes popular and there is demand I could also add support for the other number types. See the [CHANGELOG.md](https://github.com/montanaflynn/stats/blob/master/CHANGELOG.md) for API changes and tagged releases you can vendor into your projects.

> Statistics are used much like a drunk uses a lamppost: for support, not illumination. **- Vin Scully**

### Installation

```
go get github.com/montanaflynn/stats
```

### Usage

Examples of using all the functions can be seen in [examples/main.go](https://github.com/montanaflynn/stats/blob/master/examples/main.go).

### Documentation

The [entire API documentation](http://godoc.org/github.com/montanaflynn/stats) is available on GoDoc. 

Types: [`Float64Data`](http://godoc.org/github.com/montanaflynn/stats#Float64Data), [`Series`](http://godoc.org/github.com/montanaflynn/stats#Series), [`Coordinate`](http://godoc.org/github.com/montanaflynn/stats#Coordinate), [`Quartiles`](http://godoc.org/github.com/montanaflynn/stats#Quartiles), [`Outliers`](http://godoc.org/github.com/montanaflynn/stats#Outliers)

Functions: [`Min`](http://godoc.org/github.com/montanaflynn/stats#Min), [`Max`](http://godoc.org/github.com/montanaflynn/stats#Max), [`Sum`](http://godoc.org/github.com/montanaflynn/stats#Sum), [`Mean`](http://godoc.org/github.com/montanaflynn/stats#Mean), [`Median`](http://godoc.org/github.com/montanaflynn/stats#Median), [`Mode`](http://godoc.org/github.com/montanaflynn/stats#Mode), [`Sample`](http://godoc.org/github.com/montanaflynn/stats#Sample), [`Round`](http://godoc.org/github.com/montanaflynn/stats#Round), [`StandardDeviation`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviation), [`StandardDeviationPopulation`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviationPopulation), [`StandardDeviationSample`](http://godoc.org/github.com/montanaflynn/stats#StandardDeviationSample), [`Percentile`](http://godoc.org/github.com/montanaflynn/stats#Percentile), [`PercentileNearestRank`](http://godoc.org/github.com/montanaflynn/stats#PercentileNearestRank), [`LinearRegression`](http://godoc.org/github.com/montanaflynn/stats#LinearRegression), [`ExponentialRegression`](http://godoc.org/github.com/montanaflynn/stats#ExponentialRegression), [`LogarithmicRegression`](http://godoc.org/github.com/montanaflynn/stats#LogarithmicRegression), [`Variance`](http://godoc.org/github.com/montanaflynn/stats#Variance), [`PopulationVariance`](http://godoc.org/github.com/montanaflynn/stats#PopulationVariance), [`SampleVariance`](http://godoc.org/github.com/montanaflynn/stats#SampleVariance), [`Quartile`](http://godoc.org/github.com/montanaflynn/stats#Quartile), [`InterQuartileRange`](http://godoc.org/github.com/montanaflynn/stats#InterQuartileRange), [`Midhinge`](http://godoc.org/github.com/montanaflynn/stats#Midhinge), [`Trimean`](http://godoc.org/github.com/montanaflynn/stats#Trimean), [`QuartileOutliers`](http://godoc.org/github.com/montanaflynn/stats#QuartileOutliers), [`GeometricMean`](http://godoc.org/github.com/montanaflynn/stats#GeometricMean), [`HarmonicMean`](http://godoc.org/github.com/montanaflynn/stats#HarmonicMean), [`Covariance`](http://godoc.org/github.com/montanaflynn/stats#Covariance), [`Correlation`](http://godoc.org/github.com/montanaflynn/stats#Correlation)

### Contributing

If you have any suggestions, criticism or bug reports please [create an issue](https://github.com/montanaflynn/stats/issues) and I'll do my best to accommodate you. 

Pull requests are much appreciated, you may want to read the [CONTRIBUTING.md](https://github.com/montanaflynn/stats/blob/master/CONTRIBUTING.md) document to ensure a seamless merge.

Check out the [Makefile](https://github.com/montanaflynn/stats/blob/master/Makefile) for some helpful targets to common actions such as linting and testing code. 

**Protip: `watch -n 0.5 make check`**

### MIT License

Copyright (c) 2014-2015 Montana Flynn <http://anonfunction.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[travis-url]: https://travis-ci.org/montanaflynn/stats
[travis-svg]: https://img.shields.io/travis/montanaflynn/stats.svg

[coveralls-url]: https://coveralls.io/r/montanaflynn/stats?branch=master
[coveralls-svg]: https://img.shields.io/coveralls/montanaflynn/stats.svg

[godoc-url]: https://godoc.org/github.com/montanaflynn/stats
[godoc-svg]: https://godoc.org/github.com/montanaflynn/stats?status.svg