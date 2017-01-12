## Matrix Operations, Linear Algebra

As mentioned, matrix operations are ubiquitous in the data science world.  Most machine learning algorithms and many statistical/modeling techniques rely on iterative matrix operations including solving for eigenvalues/vectors, finding determinants, taking transposes, and more.  We will highlight a few here, such that you can have confidence in implementing your own custom models and understand various pre-existing implementations.

## Notes

- Matrix multiplication is in general not commutative.
- Many of the familiar rules for arithmetic with numbers hold in the case of matrix arithmetic.
- "Eigenvectors are a special set of vectors associated with a linear system of equations (i.e., a matrix equation) that are sometimes also known as characteristic vectors, proper vectors, or latent vectors... Each eigenvector is paired with a corresponding so-called eigenvalue" from [Wolfram Mathworld](http://mathworld.wolfram.com/Eigenvector.html).
- A norm is a function that assigns a strictly positive length or size to a vector or matrix.

## Links

[github/gonum/matrix/mat64 docs](https://godoc.org/github.com/gonum/matrix/mat64)  
[The Matrix Cookbook](http://www.math.uwaterloo.ca/~hwolkowi/matrixcookbook.pdf)  
[Khan Academy - Matrices](https://www.khanacademy.org/math/algebra-home/precalculus/precalc-matrices)  
[Khan Academy - Linear Algebra](https://www.khanacademy.org/math/linear-algebra)   
[Eigenvalues and eigenvectors explained visually](http://setosa.io/ev/eigenvectors-and-eigenvalues/)

## Code Review

[Matrix arithmetic](example1/example1.go)  
[Transpose, Det, Dot Product, Inverse](example2/example2.go)  
[Solve for eigenvalues/vectors](example3/example3.go)  
[Vector/Matrix Norms](example4/example4.go) 

## Exercises

### Exercise 1

Divide the following matrix by its norm:


    a = ⎡1  2  3⎤
        ⎢0  4  5⎥
        ⎣0  0  6⎦


[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
