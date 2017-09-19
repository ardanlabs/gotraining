package stats

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Min finds the lowest number in a set of data
func Min(input Float64Data) (min float64, err error) {

	// Get the count of numbers in the slice
	l := input.Len()

	// Return an error if there are no numbers
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the first value as the starting point
	min = input.Get(0)

	// Iterate until done checking for a lower value
	for i := 1; i < l; i++ {
		if input.Get(i) < min {
			min = input.Get(i)
		}
	}
	return min, nil
}

// Max finds the highest number in a slice
func Max(input Float64Data) (max float64, err error) {

	// Return an error if there are no numbers
	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the first value as the starting point
	max = input.Get(0)

	// Loop and replace higher values
	for i := 1; i < input.Len(); i++ {
		if input.Get(i) > max {
			max = input.Get(i)
		}
	}

	return max, nil
}

// Sum adds all the numbers of a slice together
func Sum(input Float64Data) (sum float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Add em up
	for _, n := range input {
		sum += n
	}

	return sum, nil
}

// Mean gets the average of a slice of numbers
func Mean(input Float64Data) (float64, error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	sum, _ := input.Sum()

	return sum / float64(input.Len()), nil
}

// GeometricMean gets the geometric mean for a slice of numbers
func GeometricMean(input Float64Data) (float64, error) {

	l := input.Len()
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the product of all the numbers
	var p float64
	for _, n := range input {
		if p == 0 {
			p = n
		} else {
			p *= n
		}
	}

	// Calculate the geometric mean
	return math.Pow(p, 1/float64(l)), nil
}

// HarmonicMean gets the harmonic mean for a slice of numbers
func HarmonicMean(input Float64Data) (float64, error) {

	l := input.Len()
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the sum of all the numbers reciprocals and return an
	// error for values that cannot be included in harmonic mean
	var p float64
	for _, n := range input {
		if n < 0 {
			return 0, errors.New("Input must not contain a negative number")
		} else if n == 0 {
			return 0, errors.New("Input must not contain a zero value")
		}
		p += (1 / n)
	}

	return float64(l) / p, nil
}

// Median gets the median number in a slice of numbers
func Median(input Float64Data) (median float64, err error) {

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	l := len(c)
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	} else if l%2 == 0 {
		median, _ = Mean(c[l/2-1 : l/2+1])
	} else {
		median = float64(c[l/2])
	}

	return median, nil
}

// Mode gets the mode of a slice of numbers
func Mode(input Float64Data) (mode []float64, err error) {

	// Return the input if there's only one number
	l := input.Len()
	if l == 1 {
		return input, nil
	} else if l == 0 {
		return nil, errors.New("Input must not be empty")
	}

	// Create a map with the counts for each number
	m := make(map[float64]int)
	for _, v := range input {
		m[v]++
	}

	// Find the highest counts to return as a slice
	// of ints to accomodate duplicate counts
	var current int
	for k, v := range m {

		// Depending if the count is lower, higher
		// or equal to the current numbers count
		// we return nothing, start a new mode or
		// append to the current mode
		switch {
		case v < current:
		case v > current:
			current = v
			mode = append(mode[:0], k)
		default:
			mode = append(mode, k)
		}
	}

	// Finally we check to see if there actually was
	// a mode by checking the length of the input and
	// mode against eachother
	lm := len(mode)
	if l == lm {
		return Float64Data{}, nil
	}

	return mode, nil
}

// _variance finds the variance for both population and sample data
func _variance(input Float64Data, sample int) (variance float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Sum the square of the mean subtracted from each number
	m, _ := Mean(input)

	for _, n := range input {
		variance += (float64(n) - m) * (float64(n) - m)
	}

	// When getting the mean of the squared differences
	// "sample" will allow us to know if it's a sample
	// or population and wether to subtract by one or not
	return variance / float64((input.Len() - (1 * sample))), nil
}

// Variance the amount of variation in the dataset
func Variance(input Float64Data) (sdev float64, err error) {
	return PopulationVariance(input)
}

// PopulationVariance finds the amount of variance within a population
func PopulationVariance(input Float64Data) (pvar float64, err error) {

	v, err := _variance(input, 0)
	if err != nil {
		return 0, err
	}

	return v, nil
}

// SampleVariance finds the amount of variance within a sample
func SampleVariance(input Float64Data) (svar float64, err error) {

	v, err := _variance(input, 1)
	if err != nil {
		return 0, err
	}

	return v, nil
}

// StandardDeviation the amount of variation in the dataset
func StandardDeviation(input Float64Data) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

// StandardDeviationPopulation finds the amount of variation from the population
func StandardDeviationPopulation(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the population variance
	vp, _ := PopulationVariance(input)

	// Return the population standard deviation
	return math.Pow(vp, 0.5), nil
}

// StandardDeviationSample finds the amount of variation from a sample
func StandardDeviationSample(input Float64Data) (sdev float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Get the sample variance
	vs, _ := SampleVariance(input)

	// Return the sample standard deviation
	return math.Pow(vs, 0.5), nil
}

// Round a float to a specific decimal place or precision
func Round(input float64, places int) (rounded float64, err error) {

	// If the float is not a number
	if math.IsNaN(input) {
		return 0.0, errors.New("Not a number")
	}

	// Find out the actual sign and correct the input for later
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	// Use the places arg to get the amount of precision wanted
	precision := math.Pow(10, float64(places))

	// Find the decimal place we are looking to round
	digit := input * precision

	// Get the actual decimal number as a fraction to be compared
	_, decimal := math.Modf(digit)

	// If the decimal is less than .5 we round down otherwise up
	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	// Finally we do the math to actually create a rounded number
	return rounded / precision * sign, nil
}

// Percentile finds the relative standing in a slice of floats
func Percentile(input Float64Data, percent float64) (percentile float64, err error) {

	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// Multiple percent by length of input
	index := (percent / 100) * float64(len(c))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := float64ToInt(index)

		// Find the average of the index and following values
		percentile, _ = Mean(Float64Data{c[i-1], c[i]})

	} else {

		// Convert float to int
		i := float64ToInt(index)

		// Find the value at the index
		percentile = c[i-1]

	}

	return percentile, nil

}

// PercentileNearestRank finds the relative standing in a slice of floats using the Nearest Rank method
func PercentileNearestRank(input Float64Data, percent float64) (percentile float64, err error) {

	// Find the length of items in the slice
	il := input.Len()

	// Return an error for empty slices
	if il == 0 {
		return 0, errors.New("Input must not be empty")
	}

	// Return error for less than 0 percentages
	if percent <= 0 {
		return 0, errors.New("Percentage must be above 0")
	}

	// Return error for greater than 100 percentages
	if percent > 100 {
		return 0, errors.New("Percentage must not be above 100")
	}

	// Start by sorting a copy of the slice
	c := sortedCopy(input)

	// Return the last item
	if percent == 100.0 {
		return c[il-1], nil
	}

	// Find ordinal ranking
	or := int(math.Ceil(float64(il) * percent / 100))

	// Return the item that is in the place of the ordinal rank
	return c[or-1], nil

}

// Series is a container for a series of data
type Series []Coordinate

// Coordinate holds the data in a series
type Coordinate struct {
	X, Y float64
}

// LinearRegression finds the least squares linear regression on data series
func LinearRegression(s Series) (regressions Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Input must not be empty")
	}

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)

	// Create the new regression series
	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: s[j].X*gradient + intercept,
		})
	}

	return regressions, nil

}

// ExponentialRegression returns an exponential regression on data series
func ExponentialRegression(s Series) (regressions Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Input must not be empty")
	}

	var sum [6]float64

	for i := 0; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X * s[i].Y
		sum[3] += s[i].Y * math.Log(s[i].Y)
		sum[4] += s[i].X * s[i].Y * math.Log(s[i].Y)
		sum[5] += s[i].X * s[i].Y
	}

	denominator := (sum[1]*sum[2] - sum[5]*sum[5])
	a := math.Pow(math.E, (sum[2]*sum[3]-sum[5]*sum[4])/denominator)
	b := (sum[1]*sum[4] - sum[5]*sum[3]) / denominator

	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: a * math.Pow(2.718281828459045, b*s[j].X),
		})
	}

	return regressions, nil

}

// LogarithmicRegression returns an logarithmic regression on data series
func LogarithmicRegression(s Series) (regressions Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Input must not be empty")
	}

	var sum [4]float64

	i := 0
	for ; i < len(s); i++ {
		sum[0] += math.Log(s[i].X)
		sum[1] += s[i].Y * math.Log(s[i].X)
		sum[2] += s[i].Y
		sum[3] += math.Pow(math.Log(s[i].X), 2)
	}

	f := float64(i)
	a := (f*sum[1] - sum[2]*sum[0]) / (f*sum[3] - sum[0]*sum[0])
	b := (sum[2] - a*sum[0]) / f

	for j := 0; j < len(s); j++ {
		regressions = append(regressions, Coordinate{
			X: s[j].X,
			Y: b + a*math.Log(s[j].X),
		})
	}

	return regressions, nil

}

// Sample returns sample from input with replacement or without
func Sample(input Float64Data, takenum int, replacement bool) ([]float64, error) {

	if input.Len() == 0 {
		return nil, errors.New("Input must not be empty")
	}

	length := input.Len()
	if replacement {

		result := Float64Data{}
		rand.Seed(unixnano())

		// In every step, randomly take the num for
		for i := 0; i < takenum; i++ {
			idx := rand.Intn(length)
			result = append(result, input[idx])
		}

		return result, nil

	} else if !replacement && takenum <= length {

		rand.Seed(unixnano())

		// Get permutation of number of indexies
		perm := rand.Perm(length)
		result := Float64Data{}

		// Get element of input by permutated index
		for _, idx := range perm[0:takenum] {
			result = append(result, input[idx])
		}

		return result, nil

	}

	return nil, errors.New("Number of taken elements must be less than length of input")
}

// Quartiles holds the three quartile points
type Quartiles struct {
	Q1 float64
	Q2 float64
	Q3 float64
}

// Quartile returns the three quartile points from a slice of data
func Quartile(input Float64Data) (Quartiles, error) {

	il := input.Len()
	if il == 0 {
		return Quartiles{}, errors.New("Input must not be empty")
	}

	// Start by sorting a copy of the slice
	copy := sortedCopy(input)

	// Find the cutoff places depeding on if
	// the input slice length is even or odd
	var c1 int
	var c2 int
	if il%2 == 0 {
		c1 = il / 2
		c2 = il / 2
	} else {
		c1 = (il - 1) / 2
		c2 = c1 + 1
	}

	// Find the Medians with the cutoff points
	Q1, _ := Median(copy[:c1])
	Q2, _ := Median(copy)
	Q3, _ := Median(copy[c2:])

	return Quartiles{Q1, Q2, Q3}, nil

}

// InterQuartileRange finds the range between Q1 and Q3
func InterQuartileRange(input Float64Data) (float64, error) {
	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}
	qs, _ := Quartile(input)
	iqr := qs.Q3 - qs.Q1
	return iqr, nil
}

// Midhinge finds the average of the first and third quartiles
func Midhinge(input Float64Data) (float64, error) {
	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}
	qs, _ := Quartile(input)
	mh := (qs.Q1 + qs.Q3) / 2
	return mh, nil
}

// Trimean finds the average of the median and the midhinge
func Trimean(input Float64Data) (float64, error) {
	if input.Len() == 0 {
		return 0, errors.New("Input must not be empty")
	}

	c := sortedCopy(input)
	q, _ := Quartile(c)

	return (q.Q1 + (q.Q2 * 2) + q.Q3) / 4, nil
}

// Outliers holds mild and extreme outliers found in data
type Outliers struct {
	Mild    Float64Data
	Extreme Float64Data
}

// QuartileOutliers finds the mild and extreme outliers
func QuartileOutliers(input Float64Data) (Outliers, error) {
	if input.Len() == 0 {
		return Outliers{}, errors.New("Input must not be empty")
	}

	// Start by sorting a copy of the slice
	copy := sortedCopy(input)

	// Calculate the quartiles and interquartile range
	qs, _ := Quartile(copy)
	iqr, _ := InterQuartileRange(copy)

	// Calculate the lower and upper inner and outer fences
	lif := qs.Q1 - (1.5 * iqr)
	uif := qs.Q3 + (1.5 * iqr)
	lof := qs.Q1 - (3 * iqr)
	uof := qs.Q3 + (3 * iqr)

	// Find the data points that are outside of the
	// inner and upper fences and add them to mild
	// and extreme outlier slices
	var mild Float64Data
	var extreme Float64Data
	for _, v := range copy {

		if v < lof || v > uof {
			extreme = append(extreme, v)
		} else if v < lif || v > uif {
			mild = append(mild, v)
		}
	}

	// Wrap them into our struct
	return Outliers{mild, extreme}, nil
}

// Covariance is a measure of how much two sets of data change
func Covariance(data1, data2 Float64Data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return 0, errors.New("Input data must not be empty")
	}

	if l1 != l2 {
		return 0, errors.New("Input data must be same length")
	}

	m1, _ := Mean(data1)
	m2, _ := Mean(data2)

	// Calculate sum of squares
	var ss float64
	for i := 0; i < l1; i++ {
		delta1 := (data1.Get(i) - m1)
		delta2 := (data2.Get(i) - m2)
		ss += (delta1*delta2 - ss) / float64(i+1)
	}

	return ss * float64(l1) / float64(l1-1), nil
}

// Correlation describes the degree of relationship between two sets of data
func Correlation(data1, data2 Float64Data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return 0, errors.New("Input data must not be empty")
	}

	if l1 != l2 {
		return 0, errors.New("Input data must be same length")
	}

	var sumX, sumY, sumCross float64

	meanX := data1.Get(0)
	meanY := data2.Get(0)

	for i := 1; i < l1; i++ {
		ratio := float64(i) / float64(i+1)
		deltaX := data1.Get(i) - meanX
		deltaY := data2.Get(i) - meanY
		sumX += deltaX * deltaX * ratio
		sumY += deltaY * deltaY * ratio
		sumCross += deltaX * deltaY * ratio
		meanX += deltaX / float64(i+1)
		meanY += deltaY / float64(i+1)
	}

	return sumCross / (math.Sqrt(sumX) * math.Sqrt(sumY)), nil
}

// float64ToInt rounds a float64 to an int
func float64ToInt(input float64) (output int) {
	r, _ := Round(input, 0)
	return int(r)
}

// unixnano returns nanoseconds from UTC epoch
func unixnano() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func copyslice(input Float64Data) Float64Data {
	s := make(Float64Data, input.Len())
	copy(s, input)
	return s
}

// sortedCopy returns a sorted copy of float64s
func sortedCopy(input Float64Data) (copy Float64Data) {
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}
