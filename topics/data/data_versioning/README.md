## Data Versioning

Even if you have code versioned (e.g., via git), you simply can’t reproduce an analysis if you don’t run the code on the same data. This means that you need to have a plan and tooling in place to retrieve the state of both your analysis and your data at certain points in history. Data science prior to data versioning is a little bit like software engineering before Git.

## Notes

## Links

[Pachyderm - the system we will use for data versioning](http://pachyderm.io/)    
[github.com/pachyderm/pachyderm/src/client docs](https://godoc.org/github.com/pachyderm/pachyderm/src/client)    

## Code Review

[Connecting to a running instance of Pachyderm](example1/example1.go)   
[Creating a data repository](example2/example2.go)    
[Committing data into a repository](example3/example3.go)   
[Retrieving data from a repository](example4/example4.go)      

## Exercises

### Exercise 1

Create another data repository called "diabetes."  We will use this repository to version other data that we will use throughout the course.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

Make a commit of the data in [diabetes.csv](data/diabetes.csv) to the newly created "diabetes" data repository.

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
