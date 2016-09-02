## CSV Cleaning and Organization

When dealing with CSV data or other forms of tabular data, you will likely want to do things like filter the data on certain fields, get subsets of the data, etc.  For example, you might just be interested in all rows where the Iris Species column has a certain value, or you maybe interested in splitting the dataset into training and test sets for a machine learning algorithm.  The Go data science community has produced a few great packages that can help you with these tasks.

## Notes

## Links

[github.com/kniren/gota](https://github.com/kniren/gota) - Dataframes package  
[github.com/go-hep/csvutil](https://github.com/go-hep/csvutil) - CSV library and utility for databases/sql

## Code Review

[Create and print a dataframe from a CSV file](example1/example1.go)  
[Filter/select/subset a dataframe](example2/example2.go)  
[Iterate over CSV records, reading data into a struct](example3/example3.go)  
[Register a CSV as a table, execute SQL statements on the CSV](example4/example4.go) 

## Exercises

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
