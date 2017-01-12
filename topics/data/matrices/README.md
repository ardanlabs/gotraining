## Matrix Creation, Modification, and Access

Many modeling, machine learning, and optimization algorithms rely on linear algebra concepts.  These concepts include eigenvalues/vectors, matrix multiplication, matrix inversion, and more.  Thus, data utilized by data scientists often has to be represented in matrix form, and data scientists will likely need to employ matrix operations in their applications.  

## Notes

- A matrix is a rectangular array representation of numbers, expressions, etc.
- Elements in a matrix are referenced by a row and column index.
- `github.com/gonum/matrix/mat64` provides functionality to create, modify, and manipulate matrices made up of float64 values.

## Links

[The Matrix Cookbook](http://www.math.uwaterloo.ca/~hwolkowi/matrixcookbook.pdf)  
[Khan Academy - Matrices](https://www.khanacademy.org/math/algebra-home/precalculus/precalc-matrices)  
[Khan Academy - Linear Algebra](https://www.khanacademy.org/math/linear-algebra)

## Code Review

[github/gonum/matrix/mat64 docs](https://godoc.org/github.com/gonum/matrix/mat64)    
[Form a float64 matrix](example1/example1.go)  
[Modify a matrix](example2/example2.go)  
[Access values in a matrix](example3/example3.go)  
[Format matrix output](example4/example4.go)  

## Exercises

### Exercise 1

Create a matrix from [diabetes.csv](data/diabetes.csv) using `github.com/gonum/matrix/mat64`. Format and output the first 10 rows to standard out.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
