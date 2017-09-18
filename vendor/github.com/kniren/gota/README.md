Gota: DataFrames, Series and Data Wrangling for Go
==================================================

This is an implementation of DataFrames, Series and data wrangling
methods for the Go programming language. The API is still in flux so
*use at your own risk*.

DataFrame
---------

The term DataFrame typically refers to a tabular dataset that can be
viewed as a two dimensional table. Often the columns of this dataset
refers to a list of features, while the rows represent a number of
measurements. As the data on the real world is not perfect, DataFrame
supports non measurements or NaN elements.

Common examples of DataFrames can be found on Excel sheets, CSV files
or SQL database tables, but this data can come on a variety of other
formats, like a collection of JSON objects or XML files.

The utility of DataFrames resides on the ability to subset them, merge
them, summarize the data for individual features or apply functions to
entire rows or columns, all while keeping column type integrity.

### Usage
#### Loading data

DataFrames can be constructed passing Series to the dataframe.New constructor
function:

```
df := dataframe.New(
	series.New([]string{"b", "a"}, series.String, "COL.1"),
	series.New([]int{1, 2}, series.Int, "COL.2"),
	series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
)
```

But as a general rule it is easier to load the data directly from
other formats. The base loading function takes some records in the
form `[][]string` and returns a new DataFrame from there:

```
df := dataframe.LoadRecords(
    [][]string{
        []string{"A", "B", "C", "D"},
        []string{"a", "4", "5.1", "true"},
        []string{"k", "5", "7.0", "true"},
        []string{"k", "4", "6.0", "true"},
        []string{"a", "2", "7.1", "false"},
    },
)
```

By default, the column types will be auto detected but this can be
configured. For example, if we wish the default type to be `Float` but
columns `A` and `D` are `String` and `Bool` respectively:

```
df := dataframe.LoadRecords(
    [][]string{
        []string{"A", "B", "C", "D"},
        []string{"a", "4", "5.1", "true"},
        []string{"k", "5", "7.0", "true"},
        []string{"k", "4", "6.0", "true"},
        []string{"a", "2", "7.1", "false"},
    },
    dataframe.DetectTypes(false),
    dataframe.DefaultType(series.Float),
    dataframe.WithTypes(map[string]series.Type{
        "A": series.String,
        "D": series.Bool,
    }),
)
```

Similarly, you can load the data stored on a `[]map[string]interface{}`:

```
df := dataframe.LoadMaps(
    []map[string]interface{}{
        map[string]interface{}{
            "A": "a",
            "B": 1,
            "C": true,
            "D": 0,
        },
        map[string]interface{}{
            "A": "b",
            "B": 2,
            "C": true,
            "D": 0.5,
        },
    },
)
```

You can also pass an `io.Reader` to the functions `ReadCSV`/`ReadJSON`
and it will work as expected given that the data is correct:

```
csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
df := dataframe.ReadCSV(strings.NewReader(csvStr))
```

```
jsonStr := `[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`
df := dataframe.ReadJSON(strings.NewReader(jsonStr))
```

#### Subsetting

We can subset our DataFrames with the Subset method. For example if we
want the first and third rows we can do the following:

```
sub := df.Subset([]int{0, 2})
```

#### Column selection

If instead of subsetting the rows we want to select specific columns,
by an index or column name:

```
sel1 := df.Select([]int{0, 2})
sel2 := df.Select([]string{"A", "C"})
```

#### Updating values

In order to update the values of a DataFrame we can use the Set
method:

```
df2 := df.Set(
    []int{0, 2},
    dataframe.LoadRecords(
        [][]string{
            []string{"A", "B", "C", "D"},
            []string{"b", "4", "6.0", "true"},
            []string{"c", "3", "6.0", "false"},
        },
    ),
)
```

#### Filtering

For more complex row subsetting we can use the Filter method. For
example, if we want the rows where the column "A" is equal to "a" or
column "B" is greater than 4:

```
fil := df.Filter(
    dataframe.F{"A", series.Eq, "a"},
    dataframe.F{"B", series.Greater, 4},
) 
fil2 := fil.Filter(
    dataframe.F{"D", series.Eq, true},
)
```

Filters inside Filter are combined as OR operations whereas if we chain
Filter methods, they will behave as AND.

#### Arrange

With Arrange a DataFrame can be sorted by the given column names:

```
sorted := df.Arrange(
    dataframe.Sort("A"),    // Sort in ascending order
    dataframe.RevSort("B"), // Sort in descending order
)
```

#### Mutate

If we want to modify a column or add one based on a given Series at
the end we can use the Mutate method:

```
// Change column C with a new one
mut := df.Mutate(
    series.New([]string{"a", "b", "c", "d"}, series.String, "C"),
)
// Add a new column E
mut2 := df.Mutate(
    series.New([]string{"a", "b", "c", "d"}, series.String, "E"),
)
```

#### Joins

Different Join operations are supported (`InnerJoin`, `LeftJoin`,
`RightJoin`, `CrossJoin`). In order to use these methods you have to
specify which are the keys to be used for joining the DataFrames:

```
df := dataframe.LoadRecords(
    [][]string{
        []string{"A", "B", "C", "D"},
        []string{"a", "4", "5.1", "true"},
        []string{"k", "5", "7.0", "true"},
        []string{"k", "4", "6.0", "true"},
        []string{"a", "2", "7.1", "false"},
    },
)
df2 := dataframe.LoadRecords(
    [][]string{
        []string{"A", "F", "D"},
        []string{"1", "1", "true"},
        []string{"4", "2", "false"},
        []string{"2", "8", "false"},
        []string{"5", "9", "false"},
    },
)
join := df.InnerJoin(df2, "D")
```

#### Function application

Functions can be applied to the rows or columns of a DataFrame,
casting the types as necessary:

```
mean := func(s series.Series) series.Series {
    floats := s.Float()
    sum := 0.0
    for _, f := range floats {
        sum += f
    }
    return series.Floats(sum / float64(len(floats)))
}
df.Cbind(mean)
df.Rbind(mean)
```

#### Chaining operations

DataFrames support a number of methods for wrangling the data,
filtering, subsetting, selecting columns, adding new columns or
modifying existing ones. All these methods can be chained one after
another and at the end of the procedure check if there has been any
errors by the DataFrame Err field. If any of the methods in the chain
returns an error, the remaining operations on the chain will become
a no-op.

```
a = a.Rename("Origin", "Country").
    Filter(dataframe.F{"Age", "<", 50}).
    Filter(dataframe.F{"Origin", "==", "United States"}).
    Select("Id", "Origin", "Date").
    Subset([]int{1, 3})
if a.Err != nil {
    log.Fatal("Oh noes!")
}
```

#### Print to console

```
fmt.Println(flights)

> [336776x20] DataFrame
> 
>     X0    year  month day   dep_time sched_dep_time dep_delay arr_time ...
>  0: 1     2013  1     1     517      515            2         830      ...
>  1: 2     2013  1     1     533      529            4         850      ...
>  2: 3     2013  1     1     542      540            2         923      ...
>  3: 4     2013  1     1     544      545            -1        1004     ...
>  4: 5     2013  1     1     554      600            -6        812      ...
>  5: 6     2013  1     1     554      558            -4        740      ...
>  6: 7     2013  1     1     555      600            -5        913      ...
>  7: 8     2013  1     1     557      600            -3        709      ...
>  8: 9     2013  1     1     557      600            -3        838      ...
>  9: 10    2013  1     1     558      600            -2        753      ...
>     ...   ...   ...   ...   ...      ...            ...       ...      ...
>     <int> <int> <int> <int> <int>    <int>          <int>     <int>    ...
> 
> Not Showing: sched_arr_time <int>, arr_delay <int>, carrier <string>, flight <int>,
> tailnum <string>, origin <string>, dest <string>, air_time <int>, distance <int>, hour <int>,
> minute <int>, time_hour <string>
```

#### Interfacing with gonum

A `gonum/mat64.Matrix` can be loaded as a `DataFrame` by using the
`LoadMatrix()` function and back to `mat64.Matrix` via
`DataFrame.Matrix()`:

```
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "true"},
			[]string{"1", "4", "6.0", "true"},
			[]string{"2", "3", "6.0", "false"},
			[]string{"2", "2", "7.1", "false"},
		},
	)
	m := a.Matrix()
	sum := mat64.Sum(m)
```

Series
------

Series are essentially vectors of elements of the same type with
support for missing values. Series are the building blocks for
DataFrame columns.

Four types are currently supported:

```
Int
Float
String
Bool
```

For more information about the API, make sure to check:

- [dataframe godoc][3]
- [series godoc][4]

License
-------
Copyright 2016 Alejandro Sanchez Brotons

Licensed under the Apache License, Version 2.0 (the "License"); you
may not use this file except in compliance with the License.  You may
obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

[1]: https://github.com/gonum
[2]: https://github.com/kniren/gota
[3]: https://godoc.org/github.com/kniren/gota/dataframe
[4]: https://godoc.org/github.com/kniren/gota/series
