## SQL-like Databases

Although there is a good bit of hype around interesting NoSQL databases and key/values stores, SQL-like databases are still ubiquitous.  You can access, query, and otherwise interact with such databases via the `databases/sql` package in Go stdlib.  There are a variety of libraries for `databases/sql` that allow you to connect to MySQL, Postgres, etc.  In the following, we will be using a Postgres database hosted by [ElephantSQL](https://www.elephantsql.com/) to illustrate database/sql functionality. 

## Get a Postgres server

- Create an account on [ElephantSQL](https://www.elephantsql.com/).
- Create a free "tiny turtle" instance.
- Get the URL from the details page of your "tiny turtle" instance.
- Export this URL to a environmental variable `PGURL`.
- Install the `psql` CLI tool.

## Notes

* A `sql.DB` value is not a database connection.   
* Using the `sql.DB` value, `databases/sql` (i) opens and closes connections to the actual underlying database, via the specified driver, and (ii) manages a pool of connections as needed.
* If you fail to release connections, you can cause `databases/sql` to open a lot of connections.  
* After creating a `sql.DB` value, you can use it to query the database, execute statements, and commit transactions. 
* When you iterate over rows and scan them into variables, Go performs data type conversions behind the scenes.

## Links

[Go database/sql tutorial](http://go-database-sql.org/)   
[Common Go pitfalls when working with database/sql](https://www.vividcortex.com/blog/2015/09/22/common-pitfalls-go/)

## Code Review

[Open a database, ping the connection](example1/example1.go)  
[Load data into a database](example2/example2.go)  
[Retrieve data from a database](example3/example3.go)  
[Modify data in a database](example4/example4.go)

## Exercises

### Exercise 1

Query `data/iris.db` returning the sum of values in each of the sepal length, sepal width, petal length, and petal width columns grouped by species.  Output the results to standard out.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

Remove any rows in `data/iris.db` with `sepal_length` greater than 6.0.

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
