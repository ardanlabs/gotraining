regression
=======
[![Build Status](https://travis-ci.org/sajari/regression.svg?branch=master)](https://travis-ci.org/sajari/regression)
[![License][license-image]][license-url]

[license-image]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: LICENSE.txt


Multivariable Linear Regression in Go (golang)

installation
------------

    $ go get github.com/sajari/regression

example usage
-------------

Import the package, create a regression and add data to it. You can use as many variables as you like, in the below example there are 3 variables for each observation.

```go
package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	r := new(regression.Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")
	r.Train(
		regression.DataPoint(11.2, []float64{587000, 16.5, 6.2}),
		regression.DataPoint(13.4, []float64{643000, 20.5, 6.4}),
		regression.DataPoint(40.7, []float64{635000, 26.3, 9.3}),
		regression.DataPoint(5.3, []float64{692000, 16.5, 5.3}),
		regression.DataPoint(24.8, []float64{1248000, 19.2, 7.3}),
		regression.DataPoint(12.7, []float64{643000, 16.5, 5.9}),
		regression.DataPoint(20.9, []float64{1964000, 20.2, 6.4}),
		regression.DataPoint(35.7, []float64{1531000, 21.3, 7.6}),
		regression.DataPoint(8.7, []float64{713000, 17.2, 4.9}),
		regression.DataPoint(9.6, []float64{749000, 14.3, 6.4}),
		regression.DataPoint(14.5, []float64{7895000, 18.1, 6}),
		regression.DataPoint(26.9, []float64{762000, 23.1, 7.4}),
		regression.DataPoint(15.7, []float64{2793000, 19.1, 5.8}),
		regression.DataPoint(36.2, []float64{741000, 24.7, 8.6}),
		regression.DataPoint(18.1, []float64{625000, 18.6, 6.5}),
		regression.DataPoint(28.9, []float64{854000, 24.9, 8.3}),
		regression.DataPoint(14.9, []float64{716000, 17.9, 6.7}),
		regression.DataPoint(25.8, []float64{921000, 22.4, 8.6}),
		regression.DataPoint(21.7, []float64{595000, 20.2, 8.4}),
		regression.DataPoint(25.7, []float64{3353000, 16.9, 6.7}),
	)
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
}
```

Note: You can also add data points one by one.

Once calculated you can print the data, look at the R^2, Variance, residuals, etc. You can also access the coefficients directly to use elsewhere, e.g.

```go
// Get the coefficient for the "Inhabitants" variable 0:
c := r.Coeff(0)
```

You can also use the model to predict new data points

```go
prediction, err := r.Predict([]float64{587000, 16.5, 6.2})
```

Feature crosses are supported so your model can capture fixed non-linear relationships

```go

r.Train(
  regression.DataPoint(11.2, []float64{587000, 16.5, 6.2}),
)
//Add a new feature which is the first variable (index 0) to the power of 2
r.AddCross(PowCross(0, 2))
r.Run()

```
