## Maintaining Integrity in Data Science Applications

Integity in data science applications is crucial, because data science applications exist to provide data-driven insights.  As soon as the integrity of a data science application breaks down, people lose trust in the output and, as a result, will refuse to make decisions based on the output.  Go helps us maintain integrity in terms of reproducibility and deployment, which are common struggles for data scientists.

## Notes

* Data science applications should consider integrity before performance or sophistication.
* A lack of reproducibility destroys the credibility of a data science application.
* Integrity cannot be maintained with a complicated deploy.
* If errors and edge cases are handled gracefully in Go, you can have confidence in how your application will behave.
* There are ways of deploying Go that maintain integrity, even if you utilize various dependencies for your statistics, ML, etc.

## Links

[Example python data science Dockerfile](https://github.com/wiseio/datascience-docker/blob/master/datascience-base/Dockerfile)  
[Example Go Dockerfile](https://www.iron.io/an-easier-way-to-create-tiny-golang-docker-images/)  

## Code Review

[Parse a clean CSV with python](example1/example1.py)  
[Parse a clean CSV with Go](example2/example2.go)  
[Force Integrity breakdown with python CSV parsing](example3/example3.py)  
[Maintain integrity in Go CSV parsing](example4/example4.go)  

## Exercises

### Exercise 1

Implement another way of handling the CSV parsing error we encountered above.  That is, handle the missing value in a way other than throwing an error.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)  
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
