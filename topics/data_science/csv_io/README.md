## CSV Data Input/Output/Parsing

Although the CSV data you will work with is not likely to be "big" (whatever that means), it is important.  As a data scientist, you will most definitely have to work with CSV data, because it is such a common format for data sets downloaded from the government, survey sites, marketers, etc.  Also, it is common to provide results, aggregations, etc. to colleagues in CSV format.

## Notes

* `encoding/csv` reads in CSV records as strings.
* Unexpected types should be handled while parsing records.
* Set an expected fields per record whenever possible.
* Check for empty values in a string column with the zero string value `""`.

## Links

[encoding/csv docs](https://golang.org/pkg/encoding/csv/)  

## Code Review

[Read in CSV records](example1/example1.go)  
[Handle unexpected fields](example2/example2.go)   
[Handle unexpected types](example3/example3.go)   
[Save a CSV file](example4/example4.go)  

## Exercises

### Exercise 1

Parse [iris_multiple_mixed_types.csv](data/iris_multiple_mixed_types.csv). Define expected types for all of the columns in the CSV file and log any errors indicating unexpected types.  

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)  

### Exercise 2

Save the parsed and cleaned [iris_multiple_mixed_types.csv](data/iris_multiple_mixed_types.csv) to a file called `processed.csv`.  

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)  

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
