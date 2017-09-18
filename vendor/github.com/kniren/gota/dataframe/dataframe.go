// Package dataframe provides an implementation of data frames and methods to
// subset, join, mutate, set, arrange, summarize, etc.
package dataframe

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gonum/matrix/mat64"
	"github.com/kniren/gota/series"
)

// DataFrame is a data structure designed for operating on table like data (Such
// as Excel, CSV files, SQL table results...) where every column have to keep type
// integrity. As a general rule of thumb, variables are stored on columns where
// every row of a DataFrame represents an observation for each variable.
//
// On the real world, data is very messy and sometimes there are non measurements
// or missing data. For this reason, DataFrame has support for NaN elements and
// allows the most common data cleaning and mungling operations such as
// subsetting, filtering, type transformations, etc. In addition to this, this
// library provides the necessary functions to concatenate DataFrames (By rows or
// columns), different Join operations (Inner, Outer, Left, Right, Cross) and the
// ability to read and write from different formats (CSV/JSON).
type DataFrame struct {
	columns []series.Series
	ncols   int
	nrows   int
	Err     error
}

// New is the generic DataFrame constructor
func New(se ...series.Series) DataFrame {
	if se == nil || len(se) == 0 {
		return DataFrame{Err: fmt.Errorf("empty DataFrame")}
	}

	nrows := -1
	columns := make([]series.Series, len(se))
	for k, s := range se {
		if s.Err != nil {
			err := fmt.Errorf("error on series %v: %v", k, s.Err)
			return DataFrame{Err: err}
		}
		columns[k] = s.Copy()
		if nrows == -1 {
			nrows = s.Len()
		}
		if nrows != s.Len() {
			return DataFrame{Err: fmt.Errorf("arguments have different dimensions")}
		}
	}

	// Fill DataFrame base structure
	df := DataFrame{
		columns: columns,
		ncols:   len(se),
		nrows:   nrows,
	}
	colnames := df.Names()
	fixColnames(colnames)
	for i, colname := range colnames {
		df.columns[i].Name = colname
	}
	return df
}

// Copy returns a copy of the DataFrame
func (df DataFrame) Copy() DataFrame {
	copy := New(df.columns...)
	if df.Err != nil {
		copy.Err = df.Err
	}
	return copy
}

// String implements the Stringer interface for DataFrame
func (df DataFrame) String() (str string) {
	return df.print(true, true, true, true, 10, 70, "DataFrame")
}

func (df DataFrame) print(
	shortRows, shortCols, showDims, showTypes bool,
	maxRows int,
	maxCharsTotal int,
	class string) (str string) {

	addRightPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return s + strings.Repeat(" ", nchar-utf8.RuneCountInString(s))
		}
		return s
	}

	addLeftPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return strings.Repeat(" ", nchar-utf8.RuneCountInString(s)) + s
		}
		return s
	}

	if df.Err != nil {
		str = fmt.Sprintf("%v error: %v", class, df.Err)
		return
	}
	nrows, ncols := df.Dims()
	if nrows == 0 || ncols == 0 {
		str = fmt.Sprintf("Empty %v", class)
		return
	}
	idx := make([]int, maxRows)
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}
	var records [][]string
	shortening := false
	if shortRows && nrows > maxRows {
		shortening = true
		df = df.Subset(idx)
		records = df.Records()
	} else {
		records = df.Records()
	}

	if showDims {
		str += fmt.Sprintf("[%vx%v] %v\n\n", nrows, ncols, class)
	}

	// Add the row numbers
	for i := 0; i < df.nrows+1; i++ {
		add := ""
		if i != 0 {
			add = strconv.Itoa(i-1) + ":"
		}
		records[i] = append([]string{add}, records[i]...)
	}
	if shortening {
		dots := make([]string, ncols+1)
		for i := 1; i < ncols+1; i++ {
			dots[i] = "..."
		}
		records = append(records, dots)
	}
	types := df.Types()
	typesrow := make([]string, ncols)
	for i := 0; i < ncols; i++ {
		typesrow[i] = fmt.Sprintf("<%v>", types[i])
	}
	typesrow = append([]string{""}, typesrow...)

	if showTypes {
		records = append(records, typesrow)
	}

	maxChars := make([]int, df.ncols+1)
	for i := 0; i < len(records); i++ {
		for j := 0; j < df.ncols+1; j++ {
			// Escape special characters
			records[i][j] = strconv.Quote(records[i][j])
			records[i][j] = records[i][j][1 : len(records[i][j])-1]

			// Detect maximum number of characters per column
			if len(records[i][j]) > maxChars[j] {
				maxChars[j] = utf8.RuneCountInString(records[i][j])
			}
		}
	}
	maxCols := len(records[0])
	var notShowing []string
	if shortCols {
		maxCharsCum := 0
		for colnum, m := range maxChars {
			maxCharsCum += m
			if maxCharsCum > maxCharsTotal {
				maxCols = colnum
				break
			}
		}
		notShowingNames := records[0][maxCols:]
		notShowingTypes := typesrow[maxCols:]
		notShowing = make([]string, len(notShowingNames))
		for i := 0; i < len(notShowingNames); i++ {
			notShowing[i] = fmt.Sprintf("%v %v", notShowingNames[i], notShowingTypes[i])
		}
	}
	for i := 0; i < len(records); i++ {
		// Add right padding to all elements
		records[i][0] = addLeftPadding(records[i][0], maxChars[0]+1)
		for j := 1; j < df.ncols+1; j++ {
			records[i][j] = addRightPadding(records[i][j], maxChars[j])
		}
		records[i] = records[i][0:maxCols]
		if shortCols && len(notShowing) != 0 {
			records[i] = append(records[i], "...")
		}
		// Create the final string
		str += strings.Join(records[i], " ")
		str += "\n"
	}
	if shortCols && len(notShowing) != 0 {
		var notShown string
		var notShownArr [][]string
		cum := 0
		i := 0
		for n, ns := range notShowing {
			cum += len(ns)
			if cum > maxCharsTotal {
				notShownArr = append(notShownArr, notShowing[i:n])
				cum = 0
				i = n
			}
		}
		if i < len(notShowing) {
			notShownArr = append(notShownArr, notShowing[i:len(notShowing)])
		}
		for k, ns := range notShownArr {
			notShown += fmt.Sprintf("%v", strings.Join(ns, ", "))
			if k != len(notShownArr)-1 {
				notShown += ","
			}
			notShown += "\n"
		}
		str += fmt.Sprintf("\nNot Showing: %v", notShown)
	}
	return str
}

// Subsetting, mutating and transforming DataFrame methods
// =======================================================

// Set will updated the values of a DataFrame for the rows selected via indexes.
func (df DataFrame) Set(indexes series.Indexes, newvalues DataFrame) DataFrame {
	if df.Err != nil {
		return df
	}
	if newvalues.Err != nil {
		return DataFrame{Err: fmt.Errorf("argument has errors: %v", newvalues.Err)}
	}
	if df.ncols != newvalues.ncols {
		return DataFrame{Err: fmt.Errorf("different number of columns")}
	}
	columns := make([]series.Series, df.ncols)
	for i, s := range df.columns {
		columns[i] = s.Set(indexes, newvalues.columns[i])
		if columns[i].Err != nil {
			df = DataFrame{Err: fmt.Errorf("setting error on column %v: %v", i, columns[i].Err)}
			return df
		}
	}
	return df
}

// Subset returns a subset of the rows of the original DataFrame based on the
// Series subsetting indexes.
func (df DataFrame) Subset(indexes series.Indexes) DataFrame {
	if df.Err != nil {
		return df
	}
	columns := make([]series.Series, df.ncols)
	for i, column := range df.columns {
		sub := column.Subset(indexes)
		if sub.Err != nil {
			return DataFrame{Err: fmt.Errorf("can't subset: %v", sub.Err)}
		}
		columns[i] = sub
	}
	return New(columns...)
}

// SelectIndexes are the supported indexes used for the DataFrame.Select method. Currently supported are:
//
//     int              // Matches the given index number
//     []int            // Matches all given index numbers
//     []bool           // Matches all columns marked as true
//     string           // Matches the column with the matching column name
//     []string         // Matches all columns with the matching column names
//     Series [Int]     // Same as []int
//     Series [Bool]    // Same as []bool
//     Series [String]  // Same as []string
type SelectIndexes interface{}

// Select the given DataFrame columns
func (df DataFrame) Select(indexes SelectIndexes) DataFrame {
	if df.Err != nil {
		return df
	}
	idx, err := parseSelectIndexes(df.ncols, indexes, df.Names())
	if err != nil {
		return DataFrame{Err: fmt.Errorf("can't select columns: %v", err)}
	}
	columns := make([]series.Series, len(idx))
	for k, i := range idx {
		if i < 0 || i >= df.ncols {
			return DataFrame{Err: fmt.Errorf("can't select columns: index out of range")}
		}
		columns[k] = df.columns[i]
	}
	return New(columns...)
}

// Rename changes the name of one of the columns of a DataFrame
func (df DataFrame) Rename(newname, oldname string) DataFrame {
	if df.Err != nil {
		return df
	}
	// Check that colname exist on dataframe
	var copy DataFrame
	colnames := df.Names()
	if idx := findInStringSlice(oldname, colnames); idx >= 0 {
		copy = df.Copy()
		copy.columns[idx].Name = newname
	} else {
		return DataFrame{
			Err: fmt.Errorf("rename: can't find column name"),
		}
	}
	return copy
}

// CBind combines the columns of two DataFrames
func (df DataFrame) CBind(dfb DataFrame) DataFrame {
	if df.Err != nil {
		return df
	}
	if dfb.Err != nil {
		return dfb
	}
	cols := append(df.columns, dfb.columns...)
	return New(cols...)
}

// RBind matches the column names of two DataFrames and returns the combination of
// the rows of both of them.
func (df DataFrame) RBind(dfb DataFrame) DataFrame {
	if df.Err != nil {
		return df
	}
	if dfb.Err != nil {
		return dfb
	}
	expandedSeries := make([]series.Series, df.ncols)
	for k, v := range df.Names() {
		idx := findInStringSlice(v, dfb.Names())
		if idx < 0 {
			return DataFrame{Err: fmt.Errorf("rbind: column names are not compatible")}
		}
		originalSeries := df.columns[k]
		addedSeries := dfb.columns[idx]
		newSeries := originalSeries.Concat(addedSeries)
		if err := newSeries.Err; err != nil {
			return DataFrame{Err: fmt.Errorf("rbind: %v", err)}
		}
		expandedSeries[k] = newSeries
	}
	return New(expandedSeries...)
}

// Mutate changes a column of the DataFrame with the given Series or adds it as
// a new column if the column name does not exist.
func (df DataFrame) Mutate(s series.Series) DataFrame {
	if df.Err != nil {
		return df
	}
	colname := s.Name
	if s.Len() != df.nrows {
		return DataFrame{
			Err: fmt.Errorf("mutate: wrong dimensions"),
		}
	}
	df = df.Copy()
	// Check that colname exist on dataframe
	newSeries := df.columns
	if idx := findInStringSlice(colname, df.Names()); idx >= 0 {
		newSeries[idx] = s
	} else {
		s.Name = colname
		newSeries = append(newSeries, s)
	}
	return New(newSeries...)
}

// F is the filtering structure
type F struct {
	Colname    string
	Comparator series.Comparator
	Comparando interface{}
}

// Filter will filter the rows of a DataFrame based on the given filters. All
// filters on the argument of a Filter call are aggregated as an OR operation
// whereas if we chain Filter calls, every filter will act as an AND operation
// with regards to the rest.
func (df DataFrame) Filter(filters ...F) DataFrame {
	if df.Err != nil {
		return df
	}
	compResults := make([]series.Series, len(filters))
	for i, f := range filters {
		idx := findInStringSlice(f.Colname, df.Names())
		if idx < 0 {
			return DataFrame{Err: fmt.Errorf("filter: can't find column name")}
		}
		res := df.columns[idx].Compare(f.Comparator, f.Comparando)
		if err := res.Err; err != nil {
			return DataFrame{Err: fmt.Errorf("filter: %v", err)}
		}
		compResults[i] = res
	}
	// Join compResults via "OR"
	if len(compResults) == 0 {
		return df.Copy()
	}
	res, err := compResults[0].Bool()
	if err != nil {
		return DataFrame{Err: fmt.Errorf("filter: %v", err)}
	}
	for i := 1; i < len(compResults); i++ {
		nextRes, err := compResults[i].Bool()
		if err != nil {
			return DataFrame{Err: fmt.Errorf("filter: %v", err)}
		}
		for j := 0; j < len(res); j++ {
			res[j] = res[j] || nextRes[j]
		}
	}
	return df.Subset(res)
}

// Order is the ordering structure
type Order struct {
	Colname string
	Reverse bool
}

// Sort return an ordering structure for regular column sorting sort.
func Sort(colname string) Order {
	return Order{colname, false}
}

// RevSort return an ordering structure for reverse column sorting.
func RevSort(colname string) Order {
	return Order{colname, true}
}

// Arrange sort the rows of a DataFrame according to the given Order
func (df DataFrame) Arrange(order ...Order) DataFrame {
	if df.Err != nil {
		return df
	}
	if order == nil || len(order) == 0 {
		return DataFrame{
			Err: fmt.Errorf("rename: no arguments"),
		}
	}

	// Check that all colnames exist before starting to sort
	for i := 0; i < len(order); i++ {
		colname := order[i].Colname
		if df.colIndex(colname) == -1 {
			return DataFrame{Err: fmt.Errorf("colname %v doesn't exist", colname)}
		}
	}

	// Initialize the index that will be used to store temporary and final order
	// results.
	origIdx := make([]int, df.nrows)
	for i := 0; i < df.nrows; i++ {
		origIdx[i] = i
	}

	swapOrigIdx := func(newidx []int) {
		newOrigIdx := make([]int, len(newidx))
		for k, i := range newidx {
			newOrigIdx[k] = origIdx[i]
		}
		origIdx = newOrigIdx
	}

	suborder := origIdx
	for i := len(order) - 1; i >= 0; i-- {
		colname := order[i].Colname
		idx := df.colIndex(colname)
		nextSeries := df.columns[idx].Subset(suborder)
		suborder = nextSeries.Order(order[i].Reverse)
		swapOrigIdx(suborder)
	}
	return df.Subset(origIdx)
}

// Capply applies the given function to the columns of a DataFrame
func (df DataFrame) Capply(f func(series.Series) series.Series) DataFrame {
	if df.Err != nil {
		return df
	}
	columns := make([]series.Series, df.ncols)
	for i, s := range df.columns {
		applied := f(s)
		applied.Name = s.Name
		columns[i] = applied
	}
	return New(columns...)
}

// Rapply applies the given function to the rows of a DataFrame. Prior to applying
// the function the elements of each row are casted to a Series of a specific
// type. In order of priority: String -> Float -> Int -> Bool. This casting also
// takes place after the function application to equalize the type of the columns.
func (df DataFrame) Rapply(f func(series.Series) series.Series) DataFrame {
	if df.Err != nil {
		return df
	}

	detectType := func(types []series.Type) series.Type {
		hasFloats := false
		hasInts := false
		hasBools := false
		hasStrings := false
		for _, t := range types {
			switch t {
			case series.Int:
				hasInts = true
			case series.Float:
				hasFloats = true
			case series.Bool:
				hasBools = true
			case series.String:
				hasStrings = true
			}
		}
		if hasStrings {
			return series.String
		} else if hasFloats {
			return series.Float
		} else if hasInts {
			return series.Int
		} else if hasBools {
			return series.Bool
		}
		panic("type not supported")
	}

	// Detect row type prior to function application
	types := df.Types()
	rowType := detectType(types)

	// Create Element matrix
	elements := make([][]series.Element, df.nrows)
	rowlen := -1
	for i := 0; i < df.nrows; i++ {
		row := series.New(nil, rowType, "").Empty()
		for _, col := range df.columns {
			row.Append(col.Elem(i))
		}
		row = f(row)
		if row.Err != nil {
			return DataFrame{
				Err: fmt.Errorf("error applying function on row %v: %v", i, row.Err),
			}
		}

		if rowlen != -1 && rowlen != row.Len() {
			return DataFrame{
				Err: fmt.Errorf("error applying function: rows have different lengths"),
			}
		}
		rowlen = row.Len()

		rowElems := make([]series.Element, rowlen)
		for j := 0; j < rowlen; j++ {
			rowElems[j] = row.Elem(j)
		}
		elements[i] = rowElems
	}

	// Cast columns if necessary
	columns := make([]series.Series, rowlen)
	for j := 0; j < rowlen; j++ {
		types := make([]series.Type, df.nrows)
		for i := 0; i < df.nrows; i++ {
			types[i] = elements[i][j].Type()
		}
		colType := detectType(types)
		s := series.New(nil, colType, "").Empty()
		for i := 0; i < df.nrows; i++ {
			s.Append(elements[i][j])
		}
		columns[j] = s
	}

	return New(columns...)
}

// Read/Write Methods
// =================

// LoadOption is the type used to configure the load of elements
type LoadOption func(*loadOptions)

type loadOptions struct {
	// If set, the type of each column will be automatically detected unless
	// otherwise specified.
	detectTypes bool

	// If set, the first row of the tabular structure will be used as column
	// names.
	hasHeader bool

	// The types of specific columns can be specified via column name.
	types map[string]series.Type

	// Specifies which is the default type in case detectTypes is disabled.
	defaultType series.Type

	// Defines which valeus are going to be considered as NaN when parsing from string
	nanValues []string
}

// DefaultType set the defaultType option for loadOptions.
func DefaultType(t series.Type) LoadOption {
	return func(c *loadOptions) {
		c.defaultType = t
	}
}

// DetectTypes set the detectTypes option for loadOptions.
func DetectTypes(b bool) LoadOption {
	return func(c *loadOptions) {
		c.detectTypes = b
	}
}

// HasHeader set the hasHeader option for loadOptions.
func HasHeader(b bool) LoadOption {
	return func(c *loadOptions) {
		c.hasHeader = b
	}
}

// NaNValues set which values are to be parsed as NaN
func NaNValues(nanValues []string) LoadOption {
	return func(c *loadOptions) {
		c.nanValues = nanValues
	}
}

// WithTypes set the types option for loadOptions.
func WithTypes(coltypes map[string]series.Type) LoadOption {
	return func(c *loadOptions) {
		c.types = coltypes
	}
}

// LoadRecords creates a new DataFrame based on the given records.
func LoadRecords(records [][]string, options ...LoadOption) DataFrame {
	// Load the options
	cfg := loadOptions{
		types:       make(map[string]series.Type),
		detectTypes: true,
		defaultType: series.String,
		hasHeader:   true,
		nanValues:   []string{"NA", "NaN", "<nil>"},
	}
	for _, option := range options {
		option(&cfg)
	}

	if len(records) == 0 {
		return DataFrame{Err: fmt.Errorf("load records: empty DataFrame")}
	}
	if cfg.hasHeader && len(records) <= 1 {
		return DataFrame{Err: fmt.Errorf("load records: empty DataFrame")}
	}

	// Extract headers
	headers := make([]string, len(records[0]))
	if cfg.hasHeader {
		headers = records[0]
		records = records[1:]
	} else {
		fixColnames(headers)
	}
	types := make([]series.Type, len(headers))
	rawcols := make([][]string, len(headers))
	for i, colname := range headers {
		rawcol := make([]string, len(records))
		for j := 0; j < len(records); j++ {
			rawcol[j] = records[j][i]
			if findInStringSlice(rawcol[j], cfg.nanValues) != -1 {
				rawcol[j] = "NaN"
			}
		}
		rawcols[i] = rawcol

		t, ok := cfg.types[colname]
		if !ok {
			t = cfg.defaultType
			if cfg.detectTypes {
				t = findType(rawcol)
			}
		}
		types[i] = t
	}

	columns := make([]series.Series, len(headers))
	for i, colname := range headers {
		col := series.New(rawcols[i], types[i], colname)
		if col.Err != nil {
			return DataFrame{Err: col.Err}
		}
		columns[i] = col
	}
	return New(columns...)
}

// LoadMaps creates a new DataFrame based on the given maps. This function assumes
// that every map on the array represents a row of observations.
func LoadMaps(maps []map[string]interface{}, options ...LoadOption) DataFrame {
	if len(maps) == 0 {
		return DataFrame{
			Err: fmt.Errorf("load maps: empty array"),
		}
	}
	inStrSlice := func(i string, s []string) bool {
		for _, v := range s {
			if v == i {
				return true
			}
		}
		return false
	}
	// Detect all colnames
	var colnames []string
	for _, v := range maps {
		for k := range v {
			if exists := inStrSlice(k, colnames); !exists {
				colnames = append(colnames, k)
			}
		}
	}
	sort.Strings(colnames)
	records := make([][]string, len(maps)+1)
	records[0] = colnames
	for k, m := range maps {
		row := make([]string, len(colnames))
		for i, colname := range colnames {
			element := ""
			val, ok := m[colname]
			if ok {
				element = fmt.Sprint(val)
			}
			row[i] = element
		}
		records[k+1] = row
	}
	return LoadRecords(records, options...)
}

// LoadMatrix loads the  given mat64.Matrix as a DataFrame
func LoadMatrix(mat mat64.Matrix) DataFrame {
	nrows, ncols := mat.Dims()
	columns := make([]series.Series, ncols)
	for i := 0; i < ncols; i++ {
		floats := make([]float64, nrows)
		for j := 0; j < nrows; j++ {
			floats[j] = mat.At(j, i)
		}
		columns[i] = series.Floats(floats)
	}
	return New(columns...)
}

// ReadCSV reads a CSV file from a io.Reader and builds a DataFrame with the
// resulting records.
func ReadCSV(r io.Reader, options ...LoadOption) DataFrame {
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		return DataFrame{Err: err}
	}
	return LoadRecords(records, options...)
}

// ReadJSON reads a JSON array from a io.Reader and builds a DataFrame with the
// resulting records.
func ReadJSON(r io.Reader, options ...LoadOption) DataFrame {
	var m []map[string]interface{}
	err := json.NewDecoder(r).Decode(&m)
	if err != nil {
		return DataFrame{Err: err}
	}
	return LoadMaps(m, options...)
}

// WriteCSV writes the DataFrame to the given io.Writer as a CSV file.
func (df DataFrame) WriteCSV(w io.Writer) error {
	if df.Err != nil {
		return df.Err
	}
	records := df.Records()
	return csv.NewWriter(w).WriteAll(records)
}

// WriteJSON writes the DataFrame to the given io.Writer as a JSON array.
func (df DataFrame) WriteJSON(w io.Writer) error {
	if df.Err != nil {
		return df.Err
	}
	m := df.Maps()
	return json.NewEncoder(w).Encode(m)
}

// Getters/Setters for DataFrame fields
// ====================================

// Names returns the name of the columns on a DataFrame.
func (df DataFrame) Names() []string {
	colnames := make([]string, df.ncols)
	for i, s := range df.columns {
		colnames[i] = s.Name
	}
	return colnames
}

// Types returns the types of the columns on a DataFrame.
func (df DataFrame) Types() []series.Type {
	coltypes := make([]series.Type, df.ncols)
	for i, s := range df.columns {
		coltypes[i] = s.Type()
	}
	return coltypes
}

// SetNames changes the column names of a DataFrame to the ones passed as an
// argument.
func (df DataFrame) SetNames(colnames []string) error {
	if df.Err != nil {
		return df.Err
	}
	if len(colnames) != df.ncols {
		err := fmt.Errorf("setting names: wrong dimensions")
		return err
	}
	for k, s := range colnames {
		df.columns[k].Name = s
	}
	return nil
}

// Dims retrieves the dimensions of a DataFrame.
func (df DataFrame) Dims() (r, c int) {
	r, c = df.Nrow(), df.Ncol()
	return
}

// Nrow returns the number of rows on a DataFrame.
func (df DataFrame) Nrow() int {
	return df.nrows
}

// Ncol returns the number of columns on a DataFrame.
func (df DataFrame) Ncol() int {
	return df.ncols
}

// Col returns the Series with the given column name contained in the DataFrame.
func (df DataFrame) Col(colname string) series.Series {
	if df.Err != nil {
		return series.Series{Err: df.Err}
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, df.Names())
	if idx < 0 {
		return series.Series{
			Err: fmt.Errorf("unknown column name"),
		}
	}
	return df.columns[idx].Copy()
}

// InnerJoin returns a DataFrame containing the inner join of two DataFrames.
func (df DataFrame) InnerJoin(b DataFrame, keys ...string) DataFrame {
	if len(keys) == 0 {
		return DataFrame{Err: fmt.Errorf("join keys not specified")}
	}
	// Check that we have all given keys in both DataFrames
	errorArr := []string{}
	var iKeysA []int
	var iKeysB []int
	for _, key := range keys {
		i := df.colIndex(key)
		if i < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on left DataFrame"))
		}
		iKeysA = append(iKeysA, i)
		j := b.colIndex(key)
		if j < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on right DataFrame"))
		}
		iKeysB = append(iKeysB, j)
	}
	if len(errorArr) != 0 {
		return DataFrame{Err: fmt.Errorf("%v", strings.Join(errorArr, "\n"))}
	}

	aCols := df.columns
	bCols := b.columns
	// Initialize newCols
	var newCols []series.Series
	for _, i := range iKeysA {
		newCols = append(newCols, aCols[i].Empty())
	}
	var iNotKeysA []int
	for i := 0; i < df.ncols; i++ {
		if !inIntSlice(i, iKeysA) {
			iNotKeysA = append(iNotKeysA, i)
			newCols = append(newCols, aCols[i].Empty())
		}
	}
	var iNotKeysB []int
	for i := 0; i < b.ncols; i++ {
		if !inIntSlice(i, iKeysB) {
			iNotKeysB = append(iNotKeysB, i)
			newCols = append(newCols, bCols[i].Empty())
		}
	}

	// Fill newCols
	for i := 0; i < df.nrows; i++ {
		for j := 0; j < b.nrows; j++ {
			match := true
			for k := range keys {
				aElem := aCols[iKeysA[k]].Elem(i)
				bElem := bCols[iKeysB[k]].Elem(j)
				match = match && aElem.Eq(bElem)
			}
			if match {
				ii := 0
				for _, k := range iKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysB {
					elem := bCols[k].Elem(j)
					newCols[ii].Append(elem)
					ii++
				}
			}
		}
	}
	return New(newCols...)
}

// LeftJoin returns a DataFrame containing the left join of two DataFrames.
func (df DataFrame) LeftJoin(b DataFrame, keys ...string) DataFrame {
	if len(keys) == 0 {
		return DataFrame{Err: fmt.Errorf("join keys not specified")}
	}
	// Check that we have all given keys in both DataFrames
	errorArr := []string{}
	var iKeysA []int
	var iKeysB []int
	for _, key := range keys {
		i := df.colIndex(key)
		if i < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on left DataFrame"))
		}
		iKeysA = append(iKeysA, i)
		j := b.colIndex(key)
		if j < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on right DataFrame"))
		}
		iKeysB = append(iKeysB, j)
	}
	if len(errorArr) != 0 {
		return DataFrame{Err: fmt.Errorf(strings.Join(errorArr, "\n"))}
	}

	aCols := df.columns
	bCols := b.columns
	// Initialize newCols
	var newCols []series.Series
	for _, i := range iKeysA {
		newCols = append(newCols, aCols[i].Empty())
	}
	var iNotKeysA []int
	for i := 0; i < df.ncols; i++ {
		if !inIntSlice(i, iKeysA) {
			iNotKeysA = append(iNotKeysA, i)
			newCols = append(newCols, aCols[i].Empty())
		}
	}
	var iNotKeysB []int
	for i := 0; i < b.ncols; i++ {
		if !inIntSlice(i, iKeysB) {
			iNotKeysB = append(iNotKeysB, i)
			newCols = append(newCols, bCols[i].Empty())
		}
	}

	// Fill newCols
	for i := 0; i < df.nrows; i++ {
		matched := false
		for j := 0; j < b.nrows; j++ {
			match := true
			for k := range keys {
				aElem := aCols[iKeysA[k]].Elem(i)
				bElem := bCols[iKeysB[k]].Elem(j)
				match = match && aElem.Eq(bElem)
			}
			if match {
				matched = true
				ii := 0
				for _, k := range iKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysB {
					elem := bCols[k].Elem(j)
					newCols[ii].Append(elem)
					ii++
				}
			}
		}
		if !matched {
			ii := 0
			for _, k := range iKeysA {
				elem := aCols[k].Elem(i)
				newCols[ii].Append(elem)
				ii++
			}
			for _, k := range iNotKeysA {
				elem := aCols[k].Elem(i)
				newCols[ii].Append(elem)
				ii++
			}
			for _ = range iNotKeysB {
				newCols[ii].Append(nil)
				ii++
			}
		}
	}
	return New(newCols...)
}

// RightJoin returns a DataFrame containing the right join of two DataFrames.
func (df DataFrame) RightJoin(b DataFrame, keys ...string) DataFrame {
	if len(keys) == 0 {
		return DataFrame{Err: fmt.Errorf("join keys not specified")}
	}
	// Check that we have all given keys in both DataFrames
	errorArr := []string{}
	var iKeysA []int
	var iKeysB []int
	for _, key := range keys {
		i := df.colIndex(key)
		if i < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on left DataFrame"))
		}
		iKeysA = append(iKeysA, i)
		j := b.colIndex(key)
		if j < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on right DataFrame"))
		}
		iKeysB = append(iKeysB, j)
	}
	if len(errorArr) != 0 {
		return DataFrame{Err: fmt.Errorf(strings.Join(errorArr, "\n"))}
	}

	aCols := df.columns
	bCols := b.columns
	// Initialize newCols
	var newCols []series.Series
	for _, i := range iKeysA {
		newCols = append(newCols, aCols[i].Empty())
	}
	var iNotKeysA []int
	for i := 0; i < df.ncols; i++ {
		if !inIntSlice(i, iKeysA) {
			iNotKeysA = append(iNotKeysA, i)
			newCols = append(newCols, aCols[i].Empty())
		}
	}
	var iNotKeysB []int
	for i := 0; i < b.ncols; i++ {
		if !inIntSlice(i, iKeysB) {
			iNotKeysB = append(iNotKeysB, i)
			newCols = append(newCols, bCols[i].Empty())
		}
	}

	// Fill newCols
	var yesmatched []struct{ i, j int }
	var nonmatched []int
	for j := 0; j < b.nrows; j++ {
		matched := false
		for i := 0; i < df.nrows; i++ {
			match := true
			for k := range keys {
				aElem := aCols[iKeysA[k]].Elem(i)
				bElem := bCols[iKeysB[k]].Elem(j)
				match = match && aElem.Eq(bElem)
			}
			if match {
				matched = true
				yesmatched = append(yesmatched, struct{ i, j int }{i, j})
			}
		}
		if !matched {
			nonmatched = append(nonmatched, j)
		}
	}
	for _, v := range yesmatched {
		i := v.i
		j := v.j
		ii := 0
		for _, k := range iKeysA {
			elem := aCols[k].Elem(i)
			newCols[ii].Append(elem)
			ii++
		}
		for _, k := range iNotKeysA {
			elem := aCols[k].Elem(i)
			newCols[ii].Append(elem)
			ii++
		}
		for _, k := range iNotKeysB {
			elem := bCols[k].Elem(j)
			newCols[ii].Append(elem)
			ii++
		}
	}
	for _, j := range nonmatched {
		ii := 0
		for _, k := range iKeysB {
			elem := bCols[k].Elem(j)
			newCols[ii].Append(elem)
			ii++
		}
		for _ = range iNotKeysA {
			newCols[ii].Append(nil)
			ii++
		}
		for _, k := range iNotKeysB {
			elem := bCols[k].Elem(j)
			newCols[ii].Append(elem)
			ii++
		}
	}
	return New(newCols...)
}

// OuterJoin returns a DataFrame containing the outer join of two DataFrames.
func (df DataFrame) OuterJoin(b DataFrame, keys ...string) DataFrame {
	if len(keys) == 0 {
		return DataFrame{Err: fmt.Errorf("join keys not specified")}
	}
	// Check that we have all given keys in both DataFrames
	errorArr := []string{}
	var iKeysA []int
	var iKeysB []int
	for _, key := range keys {
		i := df.colIndex(key)
		if i < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on left DataFrame"))
		}
		iKeysA = append(iKeysA, i)
		j := b.colIndex(key)
		if j < 0 {
			errorArr = append(errorArr, fmt.Sprint("can't find key \"", key, "\" on right DataFrame"))
		}
		iKeysB = append(iKeysB, j)
	}
	if len(errorArr) != 0 {
		return DataFrame{Err: fmt.Errorf(strings.Join(errorArr, "\n"))}
	}

	aCols := df.columns
	bCols := b.columns
	// Initialize newCols
	var newCols []series.Series
	for _, i := range iKeysA {
		newCols = append(newCols, aCols[i].Empty())
	}
	var iNotKeysA []int
	for i := 0; i < df.ncols; i++ {
		if !inIntSlice(i, iKeysA) {
			iNotKeysA = append(iNotKeysA, i)
			newCols = append(newCols, aCols[i].Empty())
		}
	}
	var iNotKeysB []int
	for i := 0; i < b.ncols; i++ {
		if !inIntSlice(i, iKeysB) {
			iNotKeysB = append(iNotKeysB, i)
			newCols = append(newCols, bCols[i].Empty())
		}
	}

	// Fill newCols
	for i := 0; i < df.nrows; i++ {
		matched := false
		for j := 0; j < b.nrows; j++ {
			match := true
			for k := range keys {
				aElem := aCols[iKeysA[k]].Elem(i)
				bElem := bCols[iKeysB[k]].Elem(j)
				match = match && aElem.Eq(bElem)
			}
			if match {
				matched = true
				ii := 0
				for _, k := range iKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysA {
					elem := aCols[k].Elem(i)
					newCols[ii].Append(elem)
					ii++
				}
				for _, k := range iNotKeysB {
					elem := bCols[k].Elem(j)
					newCols[ii].Append(elem)
					ii++
				}
			}
		}
		if !matched {
			ii := 0
			for _, k := range iKeysA {
				elem := aCols[k].Elem(i)
				newCols[ii].Append(elem)
				ii++
			}
			for _, k := range iNotKeysA {
				elem := aCols[k].Elem(i)
				newCols[ii].Append(elem)
				ii++
			}
			for _ = range iNotKeysB {
				newCols[ii].Append(nil)
				ii++
			}
		}
	}
	for j := 0; j < b.nrows; j++ {
		matched := false
		for i := 0; i < df.nrows; i++ {
			match := true
			for k := range keys {
				aElem := aCols[iKeysA[k]].Elem(i)
				bElem := bCols[iKeysB[k]].Elem(j)
				match = match && aElem.Eq(bElem)
			}
			if match {
				matched = true
			}
		}
		if !matched {
			ii := 0
			for _, k := range iKeysB {
				elem := bCols[k].Elem(j)
				newCols[ii].Append(elem)
				ii++
			}
			for _ = range iNotKeysA {
				newCols[ii].Append(nil)
				ii++
			}
			for _, k := range iNotKeysB {
				elem := bCols[k].Elem(j)
				newCols[ii].Append(elem)
				ii++
			}
		}
	}
	return New(newCols...)
}

// CrossJoin returns a DataFrame containing the cross join of two DataFrames.
func (df DataFrame) CrossJoin(b DataFrame) DataFrame {
	aCols := df.columns
	bCols := b.columns
	// Initialize newCols
	var newCols []series.Series
	for i := 0; i < df.ncols; i++ {
		newCols = append(newCols, aCols[i].Empty())
	}
	for i := 0; i < b.ncols; i++ {
		newCols = append(newCols, bCols[i].Empty())
	}
	// Fill newCols
	for i := 0; i < df.nrows; i++ {
		for j := 0; j < b.nrows; j++ {
			for ii := 0; ii < df.ncols; ii++ {
				elem := aCols[ii].Elem(i)
				newCols[ii].Append(elem)
			}
			for ii := 0; ii < b.ncols; ii++ {
				jj := ii + df.ncols
				elem := bCols[ii].Elem(j)
				newCols[jj].Append(elem)
			}
		}
	}
	return New(newCols...)
}

// colIndex returns the index of the column with name `s`. If it fails to find the
// column it returns -1 instead.
func (df DataFrame) colIndex(s string) int {
	for k, v := range df.Names() {
		if v == s {
			return k
		}
	}
	return -1
}

// Records return the string record representation of a DataFrame.
func (df DataFrame) Records() [][]string {
	var records [][]string
	records = append(records, df.Names())
	if df.ncols == 0 || df.nrows == 0 {
		return records
	}
	var tRecords [][]string
	for _, col := range df.columns {
		tRecords = append(tRecords, col.Records())
	}
	records = append(records, transposeRecords(tRecords)...)
	return records
}

// Maps return the array of maps representation of a DataFrame.
func (df DataFrame) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, df.nrows)
	colnames := df.Names()
	for i := 0; i < df.nrows; i++ {
		m := make(map[string]interface{})
		for k, v := range colnames {
			val, _ := df.columns[k].Val(i) // Ignoring the error as the index should not be out of bounds
			m[v] = val
		}
		maps[i] = m
	}
	return maps
}

// Matrix returns the mat64.Matrix representation of a DataFrame
func (df DataFrame) Matrix() mat64.Matrix {
	columns := make([]series.Series, df.ncols)
	for i := 0; i < df.ncols; i++ {
		columns[i] = series.New(df.columns[i], series.Float, df.columns[i].Name)
	}
	m := DataFrame{
		columns: columns,
		ncols:   df.ncols,
		nrows:   df.nrows,
		Err:     df.Err,
	}
	return matrix{m}
}

// fixColnames assigns a name to the missing column names and makes it so that the
// column names are unique.
func fixColnames(colnames []string) {
	// Find duplicated colnames
	dupnamesidx := make(map[string][]int)
	var missingnames []int
	for i := 0; i < len(colnames); i++ {
		a := colnames[i]
		if a == "" {
			missingnames = append(missingnames, i)
			continue
		}
		for j := 0; j < len(colnames); j++ {
			b := colnames[j]
			if i != j && a == b {
				temp := dupnamesidx[a]
				if !inIntSlice(i, temp) {
					dupnamesidx[a] = append(temp, i)
				}
			}
		}
	}

	// Autofill missing column names
	counter := 0
	for _, i := range missingnames {
		proposedName := fmt.Sprintf("X%v", counter)
		for findInStringSlice(proposedName, colnames) >= 0 {
			counter++
			proposedName = fmt.Sprintf("X%v", counter)
		}
		colnames[i] = proposedName
		counter++
	}

	// Sort map keys to make sure it always follows the same order
	var keys []string
	for k := range dupnamesidx {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Add a suffix to the duplicated colnames
	for _, name := range keys {
		idx := dupnamesidx[name]
		if name == "" {
			name = "X"
		}
		counter := 0
		for _, i := range idx {
			proposedName := fmt.Sprintf("%v_%v", name, counter)
			for findInStringSlice(proposedName, colnames) >= 0 {
				counter++
				proposedName = fmt.Sprintf("%v_%v", name, counter)
			}
			colnames[i] = proposedName
			counter++
		}
	}
}

func findInStringSlice(str string, s []string) int {
	for i, e := range s {
		if e == str {
			return i
		}
	}
	return -1
}

func parseSelectIndexes(l int, indexes SelectIndexes, colnames []string) ([]int, error) {
	var idx []int
	switch indexes.(type) {
	case []int:
		idx = indexes.([]int)
	case int:
		idx = []int{indexes.(int)}
	case []bool:
		bools := indexes.([]bool)
		if len(bools) != l {
			return nil, fmt.Errorf("indexing error: index dimensions mismatch")
		}
		for i, b := range bools {
			if b {
				idx = append(idx, i)
			}
		}
	case string:
		s := indexes.(string)
		i := findInStringSlice(s, colnames)
		if i < 0 {
			return nil, fmt.Errorf("can't select columns: column name \"%v\" not found", s)
		}
		idx = append(idx, i)
	case []string:
		xs := indexes.([]string)
		for _, s := range xs {
			i := findInStringSlice(s, colnames)
			if i < 0 {
				return nil, fmt.Errorf("can't select columns: column name \"%v\" not found", s)
			}
			idx = append(idx, i)
		}
	case series.Series:
		s := indexes.(series.Series)
		if err := s.Err; err != nil {
			return nil, fmt.Errorf("indexing error: new values has errors: %v", err)
		}
		if s.HasNaN() {
			return nil, fmt.Errorf("indexing error: indexes contain NaN")
		}
		switch s.Type() {
		case series.Int:
			return s.Int()
		case series.Bool:
			bools, err := s.Bool()
			if err != nil {
				return nil, fmt.Errorf("indexing error: %v", err)
			}
			return parseSelectIndexes(l, bools, colnames)
		case series.String:
			xs := indexes.(series.Series).Records()
			return parseSelectIndexes(l, xs, colnames)
		default:
			return nil, fmt.Errorf("indexing error: unknown indexing mode")
		}
	default:
		return nil, fmt.Errorf("indexing error: unknown indexing mode")
	}
	return idx, nil
}

func findType(arr []string) series.Type {
	hasFloats := false
	hasInts := false
	hasBools := false
	hasStrings := false
	for _, str := range arr {
		if str == "" || str == "NaN" {
			continue
		}
		if _, err := strconv.Atoi(str); err == nil {
			hasInts = true
			continue
		}
		if _, err := strconv.ParseFloat(str, 64); err == nil {
			hasFloats = true
			continue
		}
		if str == "true" || str == "false" {
			hasBools = true
			continue
		}
		hasStrings = true
	}
	if hasFloats && !hasBools && !hasStrings {
		return series.Float
	}
	if hasInts && !hasFloats && !hasBools && !hasStrings {
		return series.Int
	}
	if !hasInts && !hasFloats && hasBools && !hasStrings {
		return series.Bool
	}
	return series.String
}

func transposeRecords(x [][]string) [][]string {
	n := len(x)
	if n == 0 {
		return x
	}
	m := len(x[0])
	y := make([][]string, m)
	for i := 0; i < m; i++ {
		z := make([]string, n)
		for j := 0; j < n; j++ {
			z[j] = x[j][i]
		}
		y[i] = z
	}
	return y
}

func inIntSlice(i int, is []int) bool {
	for _, v := range is {
		if v == i {
			return true
		}
	}
	return false
}

type matrix struct {
	DataFrame
}

func (m matrix) String() string {
	return m.print(true, true, true, true, 10, 70, "Matrix")
}

func (m matrix) At(i, j int) float64 {
	return m.columns[j].Elem(i).Float()
}

func (m matrix) T() mat64.Matrix {
	return mat64.Transpose{Matrix: m}
}
