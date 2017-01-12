## CSV Cleaning and Organization

When dealing with CSV data or other forms of tabular data, you will likely want to do things like filter the data on certain fields, get subsets of the data, etc.  For example, you might just be interested in all rows where the Iris Species column has a certain value, or you maybe interested in splitting the dataset into training and test sets for a machine learning algorithm.  The Go data science community has produced a few great packages that can help you with these tasks.

## Notes

* Use `encoding/csv` unless there is a need to do more complicated filtering, merging, etc.
* Dataframes are useful for quick filtering, subsetting, merging, etc. with your dataset in memory.
* The CSV driver for `databases/sql` is useful for iterating over your dataset, while cleaning/organizing it, without pulling it into memory. 

## Links

[github.com/kniren/gota](https://github.com/kniren/gota) - Dataframes package  
[github.com/go-hep/csvutil](https://github.com/go-hep/csvutil) - CSV library and utility for databases/sql

## Code Review

[Create and print a dataframe from a CSV file](example1/example1.go)  
[Filter/select/subset a dataframe](example2/example2.go)  
[Iterate over CSV records, reading data into a struct](example3/example3.go)  
[Register a CSV as a table, execute SQL statements on the CSV](example4/example4.go) 

## Exercises

### Exercise 1

Use Gota dataframes to read [iris.csv](data/iris.csv) and output three files corresponding to each Iris species (`setosa.csv`, `versicolor`, and `virginica.csv`), each of the three files containing only the rows corresponding to the respective species.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

Use csvutil/csvdriver to read [iris.csv](data/iris.csv), sum the float values in the first four columns, and output a processed CSV file with two columns delimited by semicolons, the first having the sum value for the row and the second having the respective species.

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
