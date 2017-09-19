package stats

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

var sf = []float64{1.1, 2, 3, 4, 5}

func makeLargeFloatSlice(c int) []float64 {
	lf := []float64{}
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func TestMin(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1.1, 2, 3, 4, 5}, 1.1},
		{[]float64{10.534, 3, 5, 7, 9}, 3.0},
		{[]float64{-5, 1, 5}, -5.0},
		{[]float64{5}, 5},
	} {
		got, err := Min(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Min(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Min([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func BenchmarkMinSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(sf)
	}
}

func BenchmarkMinLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(lf)
	}
}

func TestMax(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 5.0},
		{[]float64{10.5, 3, 5, 7, 9}, 10.5},
		{[]float64{-20, -1, -5.5}, -1.0},
		{[]float64{-1.0}, -1.0},
	} {
		got, err := Max(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Max(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Max([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func BenchmarkMaxSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(sf)
	}
}

func BenchmarkMaxLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(lf)
	}
}

func TestMean(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := Mean(c.in)
		if got != c.out {
			t.Errorf("Mean(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Mean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkMeanSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(sf)
	}
}

func BenchmarkMeanLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mean(lf)
	}
}

func TestMedian(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{5, 3, 4, 2, 1}, 3.0},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := Median(c.in)
		if got != c.out {
			t.Errorf("Median(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Median([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkMedianSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Median(sf)
	}
}

func BenchmarkMedianLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Median(lf)
	}
}

func TestMedianSortSideEffects(t *testing.T) {
	s := []float64{0.1, 0.3, 0.2, 0.4, 0.5}
	a := []float64{0.1, 0.3, 0.2, 0.4, 0.5}
	Median(s)
	if !reflect.DeepEqual(s, a) {
		t.Errorf("%.1f != %.1f", s, a)
	}
}

func TestMode(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{5, 3, 4, 2, 1}, []float64{}},
		{[]float64{5, 5, 3, 4, 2, 1}, []float64{5}},
		{[]float64{5, 5, 3, 3, 4, 2, 1}, []float64{3, 5}},
		{[]float64{1}, []float64{1}},
	} {
		got, err := Mode(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		sort.Float64s(got)
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Mode(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Mode([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkModeSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mode(sf)
	}
}

func BenchmarkModeLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mode(lf)
	}
}

func TestSum(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3}, 6},
		{[]float64{1.0, 1.1, 1.2, 2.2}, 5.5},
		{[]float64{1, -1, 2, -3}, -1},
	} {
		got, err := Sum(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Sum(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Sum([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkSumSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(sf)
	}
}

func BenchmarkSumLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(lf)
	}
}

func TestVariance(t *testing.T) {
	_, err := Variance([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestPopulationVariance(t *testing.T) {
	e, _ := PopulationVariance([]float64{})
	if e != 0.0 {
		t.Errorf("%.1f != %.1f", e, 0.0)
	}
	pv, _ := PopulationVariance([]float64{1, 2, 3})
	a, err := Round(pv, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if a != 0.7 {
		t.Errorf("%.1f != %.1f", a, 0.7)
	}
}

func TestSampleVariance(t *testing.T) {
	m, _ := SampleVariance([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
	m, _ = SampleVariance([]float64{1, 2, 3})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestStandardDeviation(t *testing.T) {
	_, err := StandardDeviation([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestStandardDeviationPopulation(t *testing.T) {
	s, _ := StandardDeviationPopulation([]float64{1, 2, 3})
	m, err := Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.82 {
		t.Errorf("%.10f != %.10f", m, 0.82)
	}
	s, _ = StandardDeviationPopulation([]float64{-1, -2, -3.3})
	m, err = Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.94 {
		t.Errorf("%.10f != %.10f", m, 0.94)
	}

	m, _ = StandardDeviationPopulation([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestStandardDeviationSample(t *testing.T) {
	s, _ := StandardDeviationSample([]float64{1, 2, 3})
	m, err := Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.0 {
		t.Errorf("%.10f != %.10f", m, 1.0)
	}
	s, _ = StandardDeviationSample([]float64{-1, -2, -3.3})
	m, err = Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.15 {
		t.Errorf("%.10f != %.10f", m, 1.15)
	}

	m, _ = StandardDeviationSample([]float64{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestRound(t *testing.T) {
	m, err := Round(0.1111, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.1 {
		t.Errorf("%.1f != %.1f", m, 0.1)
	}

	m, err = Round(-0.1111, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != -0.11 {
		t.Errorf("%.1f != %.1f", m, -0.11)
	}

	m, err = Round(5.3253, 3)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 5.325 {
		t.Errorf("%.1f != %.1f", m, 5.325)
	}

	m, err = Round(5.3253, 0)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 5.0 {
		t.Errorf("%.1f != %.1f", m, 5.0)
	}

	m, err = Round(math.NaN(), 2)
	if err == nil {
		t.Errorf("Round should error on NaN")
	}
}

func BenchmarkRoundSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(0.1111, 1)
	}
}

func TestPercentile(t *testing.T) {
	m, _ := Percentile([]float64{43, 54, 56, 61, 62, 66}, 90)
	if m != 62.0 {
		t.Errorf("%.1f != %.1f", m, 62.0)
	}
	m, _ = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50)
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
	m, _ = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 99.9)
	if m != 10.0 {
		t.Errorf("%.1f != %.1f", m, 10.0)
	}
	_, err := Percentile([]float64{}, 99.9)
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func TestPercentileSortSideEffects(t *testing.T) {
	s := []float64{43, 54, 56, 44, 62, 66}
	a := []float64{43, 54, 56, 44, 62, 66}
	Percentile(s, 90)
	if !reflect.DeepEqual(s, a) {
		t.Errorf("%.1f != %.1f", s, a)
	}
}

func BenchmarkPercentileSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Percentile(sf, 50)
	}
}

func BenchmarkPercentileLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Percentile(lf, 50)
	}
}

func TestPercentileNearestRank(t *testing.T) {
	f1 := []float64{35, 20, 15, 40, 50}
	f2 := []float64{20, 6, 7, 8, 8, 10, 13, 15, 16, 3}
	f3 := makeLargeFloatSlice(101)

	for _, c := range []struct {
		sample  []float64
		percent float64
		result  float64
	}{
		{f1, 30, 20},
		{f1, 40, 20},
		{f1, 50, 35},
		{f1, 75, 40},
		{f1, 95, 50},
		{f1, 99, 50},
		{f1, 99.9, 50},
		{f1, 100, 50},
		{f2, 25, 7},
		{f2, 50, 8},
		{f2, 75, 15},
		{f2, 100, 20},
		{f3, 1, 100},
		{f3, 99, 9900},
		{f3, 100, 10000},
	} {
		got, err := PercentileNearestRank(c.sample, c.percent)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}
		if got != c.result {
			t.Errorf("%v != %v", got, c.result)
		}
	}

	_, err := PercentileNearestRank([]float64{}, 50)
	if err == nil {
		t.Errorf("Should have returned an empty slice error")
	}

	_, err = PercentileNearestRank([]float64{1, 2, 3, 4, 5}, 0)
	if err == nil {
		t.Errorf("Should have returned an percentage must be above 0 error")
	}

	_, err = PercentileNearestRank([]float64{1, 2, 3, 4, 5}, 110)
	if err == nil {
		t.Errorf("Should have returned an percentage must not be above 100 error")
	}

}

func BenchmarkPercentileNearestRankSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PercentileNearestRank(sf, 50)
	}
}

func BenchmarkPercentileNearestRankLargeFloatSlice(b *testing.B) {
	lf := makeLargeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PercentileNearestRank(lf, 50)
	}
}

func TestFloat64ToInt(t *testing.T) {
	m := float64ToInt(234.0234)
	if m != 234 {
		t.Errorf("%x != %x", m, 234)
	}
	m = float64ToInt(-234.0234)
	if m != -234 {
		t.Errorf("%x != %x", m, -234)
	}
	m = float64ToInt(1)
	if m != 1 {
		t.Errorf("%x != %x", m, 1)
	}
}

func TestLinearRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := LinearRegression(data)
	a := 2.3800000000000026
	if r[0].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 3.0800000000000014
	if r[1].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 3.7800000000000002
	if r[2].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 4.479999999999999
	if r[3].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 5.179999999999998
	if r[4].Y != a {
		t.Errorf("%v != %v", r, a)
	}

	_, err := LinearRegression([]Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestExponentialRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := ExponentialRegression(data)
	a, _ := Round(r[0].Y, 3)
	if a != 2.515 {
		t.Errorf("%v != %v", r, 2.515)
	}
	a, _ = Round(r[1].Y, 3)
	if a != 3.032 {
		t.Errorf("%v != %v", r, 3.032)
	}
	a, _ = Round(r[2].Y, 3)
	if a != 3.655 {
		t.Errorf("%v != %v", r, 3.655)
	}
	a, _ = Round(r[3].Y, 3)
	if a != 4.407 {
		t.Errorf("%v != %v", r, 4.407)
	}
	a, _ = Round(r[4].Y, 3)
	if a != 5.313 {
		t.Errorf("%v != %v", r, 5.313)
	}

	_, err := ExponentialRegression([]Coordinate{})
	if err == nil {

		t.Errorf("Empty slice should have returned an error")
	}
}

func TestLogarithmicRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := LogarithmicRegression(data)
	a := 2.1520822363811702
	if r[0].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 3.3305559222492214
	if r[1].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 4.019918836568674
	if r[2].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 4.509029608117273
	if r[3].Y != a {
		t.Errorf("%v != %v", r, a)
	}
	a = 4.888413396683663
	if r[4].Y != a {
		t.Errorf("%v != %v", r, a)
	}

	_, err := LogarithmicRegression([]Coordinate{})
	if err == nil {

		t.Errorf("Empty slice should have returned an error")
	}
}

func TestSample(t *testing.T) {
	_, err := Sample([]float64{}, 10, false)
	if err == nil {
		t.Errorf("Returned an error")
	}

	_, err2 := Sample([]float64{0.1, 0.2}, 10, false)
	if err2 == nil {
		t.Errorf("Returned an error")
	}
}

func TestSampleWithoutReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result, _ := Sample(arr, 5, false)
	checks := map[float64]bool{}
	for _, res := range result {
		_, ok := checks[res]
		if ok {
			t.Errorf("%v already seen", res)
		}
		checks[res] = true
	}
}

func TestSampleWithReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	numsamples := 100
	result, _ := Sample(arr, numsamples, true)
	if len(result) != numsamples {
		t.Errorf("%v != %v", len(result), numsamples)
	}
}

func TestQuartile(t *testing.T) {
	s1 := []float64{6, 7, 15, 36, 39, 40, 41, 42, 43, 47, 49}
	s2 := []float64{7, 15, 36, 39, 40, 41}

	for _, c := range []struct {
		in []float64
		Q1 float64
		Q2 float64
		Q3 float64
	}{
		{s1, 15, 40, 43},
		{s2, 15, 37.5, 40},
	} {
		quartiles, err := Quartile(c.in)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}

		if quartiles.Q1 != c.Q1 {
			t.Errorf("Q1 %v != %v", quartiles.Q1, c.Q1)
		}
		if quartiles.Q2 != c.Q2 {
			t.Errorf("Q2 %v != %v", quartiles.Q2, c.Q2)
		}
		if quartiles.Q3 != c.Q3 {
			t.Errorf("Q3 %v != %v", quartiles.Q3, c.Q3)
		}
	}

	_, err := Quartile([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestInterQuartileRange(t *testing.T) {
	s1 := []float64{102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118}
	iqr, _ := InterQuartileRange(s1)

	if iqr != 10 {
		t.Errorf("IQR %v != 10", iqr)
	}

	_, err := InterQuartileRange([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestMidhinge(t *testing.T) {
	s1 := []float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13}
	mh, _ := Midhinge(s1)

	if mh != 7.5 {
		t.Errorf("Midhinge %v != 7.5", mh)
	}

	_, err := Midhinge([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestTrimean(t *testing.T) {
	s1 := []float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13}
	tr, _ := Trimean(s1)

	if tr != 7.25 {
		t.Errorf("Trimean %v != 7.25", tr)
	}

	_, err := Trimean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestQuartileOutliers(t *testing.T) {
	s1 := []float64{-1000, 1, 3, 4, 4, 6, 6, 6, 6, 7, 8, 15, 18, 100}
	o, _ := QuartileOutliers(s1)

	if o.Mild[0] != 15 {
		t.Errorf("First Mild Outlier %v != 15", o.Mild[0])
	}

	if o.Mild[1] != 18 {
		t.Errorf("Second Mild Outlier %v != 18", o.Mild[1])
	}

	if o.Extreme[0] != -1000 {
		t.Errorf("First Extreme Outlier %v != -1000", o.Extreme[0])
	}

	if o.Extreme[1] != 100 {
		t.Errorf("Second Extreme Outlier %v != 100", o.Extreme[1])
	}

	_, err := QuartileOutliers([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestGeometricMean(t *testing.T) {
	s1 := []float64{2, 18}
	s2 := []float64{10, 51.2, 8}
	s3 := []float64{1, 3, 9, 27, 81}

	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{s1, 6},
		{s2, 16},
		{s3, 9},
	} {
		gm, err := GeometricMean(c.in)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}

		gm, _ = Round(gm, 0)
		if gm != c.out {
			t.Errorf("Geometric Mean %v != %v", gm, c.out)
		}
	}

	_, err := GeometricMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestHarmonicMean(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 0, 9, 27, 81}

	hm, err := HarmonicMean(s1)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	hm, _ = Round(hm, 2)
	if hm != 2.19 {
		t.Errorf("Geometric Mean %v != %v", hm, 2.19)
	}

	hm, err = HarmonicMean(s2)
	if err == nil {
		t.Errorf("Should have returned a negative number error")
	}

	hm, err = HarmonicMean(s3)
	if err == nil {
		t.Errorf("Should have returned a zero number error")
	}

	_, err = HarmonicMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestCovariance(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := Covariance(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := Covariance(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 3.2499999999999996 {
		t.Errorf("Covariance %v != %v", a, 3.2499999999999996)
	}

	_, err = Covariance(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestCorrelation(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := Correlation(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := Correlation(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 0.9912407071619301 {
		t.Errorf("Correlation %v != %v", a, 0.9912407071619301)
	}

	_, err = Correlation(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}

}
