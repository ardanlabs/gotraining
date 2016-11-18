## Dimensionality Reduction

Data scientists use dimensionality reduction to transfrom high-dimensional data sets into more compact, low-dimensional data sets.  This process can be very useful when there is redundancy and correlation between features, when a data set includes irrelevant features, and when computational or modeling constraints necessitate lower dimensions. Principal Component Analysis is a widely used dimensionality reduction technique that we will explore here.

## Notes

- Principal components are (as stated [here](http://statweb.stanford.edu/~tibs/ElemStatLearn/)):
    - a sequence of projections of the data, 
    - mutually uncorrelated, and 
    - ordered in variance.
- The axis corresponding to the principal eigenvector/component is the one along which the data is most “spread out” (i.e., the axis along which the variance of the data is maximized).
- A PCA transformation, replaces high-dimensionality data by its projection onto it's most important axes.
- Although PCA is widely used, there are a variety of dimensionality reduction techniques 

## Links

[A tutorial on principal component analysis](http://faculty.iiit.ac.in/~mkrishna/PrincipalComponents.pdf)
[Another tutorial on principal component analysis](https://www.cs.princeton.edu/picasso/mats/PCA-Tutorial-Intuition_jp.pdf)
[A survey of dimensionality reduction techniques](http://computation.llnl.gov/casc/sapphire/pubs/148494.pdf)

## Code Review

[github.com/gonum/stat docs](https://godoc.org/github.com/gonum/stat)  
[Calculate Principal Components](example1/example1.go)
[Determine a Number of Target Dimensions](example2/example2.go)  
[Project the Data](example3/example3.go)  

## Exercises

### Exercise 1

Project the iris data set features on to three dimensions rather than four.  Output the results.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
