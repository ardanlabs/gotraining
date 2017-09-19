// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// a set of benchmarks to evaluate the performance of the various
// moment statistics: Mean, Variance, StdDev, MeanVariance, MeanStdDev,
// Covariance, Correlation, Skew, ExKurtosis, Moment, MomentAbout, ...
//
// It tests both weighted and unweighted versions by using a slice of
// all ones.

package stat

import (
	"math/rand"
	"testing"
)

const (
	small  = 10
	medium = 1000
	large  = 100000
	huge   = 10000000
)

// tests for unweighted versions

func RandomSlice(l int) []float64 {
	s := make([]float64, l)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}

func benchmarkMean(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mean(s, wts)
	}
}

func BenchmarkMeanSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkMean(b, s, nil)
}

func BenchmarkMeanMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkMean(b, s, nil)
}

func BenchmarkMeanLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkMean(b, s, nil)
}

func BenchmarkMeanHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkMean(b, s, nil)
}

func BenchmarkMeanSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkMean(b, s, wts)
}

func BenchmarkMeanMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkMean(b, s, wts)
}

func BenchmarkMeanLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkMean(b, s, wts)
}

func BenchmarkMeanHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkMean(b, s, wts)
}

func benchmarkVariance(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Variance(s, wts)
	}
}

func BenchmarkVarianceSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkVariance(b, s, nil)
}

func BenchmarkVarianceMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkVariance(b, s, nil)
}

func BenchmarkVarianceLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkVariance(b, s, nil)
}

func BenchmarkVarianceHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkVariance(b, s, nil)
}

func BenchmarkVarianceSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkVariance(b, s, wts)
}

func BenchmarkVarianceMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkVariance(b, s, wts)
}

func BenchmarkVarianceLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkVariance(b, s, wts)
}

func BenchmarkVarianceHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkVariance(b, s, wts)
}

func benchmarkStdDev(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StdDev(s, wts)
	}
}

func BenchmarkStdDevSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkStdDev(b, s, nil)
}

func BenchmarkStdDevMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkStdDev(b, s, nil)
}

func BenchmarkStdDevLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkStdDev(b, s, nil)
}

func BenchmarkStdDevHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkStdDev(b, s, nil)
}

func BenchmarkStdDevSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkStdDev(b, s, wts)
}

func BenchmarkStdDevMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkStdDev(b, s, wts)
}

func BenchmarkStdDevLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkStdDev(b, s, wts)
}

func BenchmarkStdDevHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkStdDev(b, s, wts)
}

func benchmarkMeanVariance(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MeanVariance(s, wts)
	}
}

func BenchmarkMeanVarianceSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkMeanVariance(b, s, nil)
}

func BenchmarkMeanVarianceMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkMeanVariance(b, s, nil)
}

func BenchmarkMeanVarianceLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkMeanVariance(b, s, nil)
}

func BenchmarkMeanVarianceHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkMeanVariance(b, s, nil)
}

func BenchmarkMeanVarianceSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkMeanVariance(b, s, wts)
}

func BenchmarkMeanVarianceMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkMeanVariance(b, s, wts)
}

func BenchmarkMeanVarianceLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkMeanVariance(b, s, wts)
}

func BenchmarkMeanVarianceHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkMeanVariance(b, s, wts)
}

func benchmarkMeanStdDev(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MeanStdDev(s, wts)
	}
}

func BenchmarkMeanStdDevSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkMeanStdDev(b, s, nil)
}

func BenchmarkMeanStdDevMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkMeanStdDev(b, s, nil)
}

func BenchmarkMeanStdDevLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkMeanStdDev(b, s, nil)
}

func BenchmarkMeanStdDevHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkMeanStdDev(b, s, nil)
}

func BenchmarkMeanStdDevSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkMeanStdDev(b, s, wts)
}

func BenchmarkMeanStdDevMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkMeanStdDev(b, s, wts)
}

func BenchmarkMeanStdDevLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkMeanStdDev(b, s, wts)
}

func BenchmarkMeanStdDevHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkMeanStdDev(b, s, wts)
}

func benchmarkCovariance(b *testing.B, s1, s2, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Covariance(s1, s2, wts)
	}
}

func BenchmarkCovarianceSmall(b *testing.B) {
	s1 := RandomSlice(small)
	s2 := RandomSlice(small)
	benchmarkCovariance(b, s1, s2, nil)
}

func BenchmarkCovarianceMedium(b *testing.B) {
	s1 := RandomSlice(medium)
	s2 := RandomSlice(medium)
	benchmarkCovariance(b, s1, s2, nil)
}

func BenchmarkCovarianceLarge(b *testing.B) {
	s1 := RandomSlice(large)
	s2 := RandomSlice(large)
	benchmarkCovariance(b, s1, s2, nil)
}

func BenchmarkCovarianceHuge(b *testing.B) {
	s1 := RandomSlice(huge)
	s2 := RandomSlice(huge)
	benchmarkCovariance(b, s1, s2, nil)
}

func BenchmarkCovarianceSmallWeighted(b *testing.B) {
	s1 := RandomSlice(small)
	s2 := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkCovariance(b, s1, s2, wts)
}

func BenchmarkCovarianceMediumWeighted(b *testing.B) {
	s1 := RandomSlice(medium)
	s2 := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkCovariance(b, s1, s2, wts)
}

func BenchmarkCovarianceLargeWeighted(b *testing.B) {
	s1 := RandomSlice(large)
	s2 := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkCovariance(b, s1, s2, wts)
}

func BenchmarkCovarianceHugeWeighted(b *testing.B) {
	s1 := RandomSlice(huge)
	s2 := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkCovariance(b, s1, s2, wts)
}

func benchmarkCorrelation(b *testing.B, s1, s2, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Correlation(s1, s2, wts)
	}
}

func BenchmarkCorrelationSmall(b *testing.B) {
	s1 := RandomSlice(small)
	s2 := RandomSlice(small)
	benchmarkCorrelation(b, s1, s2, nil)
}

func BenchmarkCorrelationMedium(b *testing.B) {
	s1 := RandomSlice(medium)
	s2 := RandomSlice(medium)
	benchmarkCorrelation(b, s1, s2, nil)
}

func BenchmarkCorrelationLarge(b *testing.B) {
	s1 := RandomSlice(large)
	s2 := RandomSlice(large)
	benchmarkCorrelation(b, s1, s2, nil)
}

func BenchmarkCorrelationHuge(b *testing.B) {
	s1 := RandomSlice(huge)
	s2 := RandomSlice(huge)
	benchmarkCorrelation(b, s1, s2, nil)
}

func BenchmarkCorrelationSmallWeighted(b *testing.B) {
	s1 := RandomSlice(small)
	s2 := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkCorrelation(b, s1, s2, wts)
}

func BenchmarkCorrelationMediumWeighted(b *testing.B) {
	s1 := RandomSlice(medium)
	s2 := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkCorrelation(b, s1, s2, wts)
}

func BenchmarkCorrelationLargeWeighted(b *testing.B) {
	s1 := RandomSlice(large)
	s2 := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkCorrelation(b, s1, s2, wts)
}

func BenchmarkCorrelationHugeWeighted(b *testing.B) {
	s1 := RandomSlice(huge)
	s2 := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkCorrelation(b, s1, s2, wts)
}

func benchmarkSkew(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Skew(s, wts)
	}
}

func BenchmarkSkewSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkSkew(b, s, nil)
}

func BenchmarkSkewMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkSkew(b, s, nil)
}

func BenchmarkSkewLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkSkew(b, s, nil)
}

func BenchmarkSkewHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkSkew(b, s, nil)
}

func BenchmarkSkewSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkSkew(b, s, wts)
}

func BenchmarkSkewMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkSkew(b, s, wts)
}

func BenchmarkSkewLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkSkew(b, s, wts)
}

func BenchmarkSkewHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkSkew(b, s, wts)
}

func benchmarkExKurtosis(b *testing.B, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExKurtosis(s, wts)
	}
}

func BenchmarkExKurtosisSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkExKurtosis(b, s, nil)
}

func BenchmarkExKurtosisMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkExKurtosis(b, s, nil)
}

func BenchmarkExKurtosisLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkExKurtosis(b, s, nil)
}

func BenchmarkExKurtosisHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkExKurtosis(b, s, nil)
}

func BenchmarkExKurtosisSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkExKurtosis(b, s, wts)
}

func BenchmarkExKurtosisMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkExKurtosis(b, s, wts)
}

func BenchmarkExKurtosisLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkExKurtosis(b, s, wts)
}

func BenchmarkExKurtosisHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkExKurtosis(b, s, wts)
}

func benchmarkMoment(b *testing.B, n float64, s, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Moment(n, s, wts)
	}
}

func BenchmarkMomentSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkMoment(b, 5, s, nil)
}

func BenchmarkMomentMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkMoment(b, 5, s, nil)
}

func BenchmarkMomentLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkMoment(b, 5, s, nil)
}

func BenchmarkMomentHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkMoment(b, 5, s, nil)
}

func BenchmarkMomentSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkMoment(b, 5, s, wts)
}

func BenchmarkMomentMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkMoment(b, 5, s, wts)
}

func BenchmarkMomentLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkMoment(b, 5, s, wts)
}

func BenchmarkMomentHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkMoment(b, 5, s, wts)
}

func benchmarkMomentAbout(b *testing.B, n float64, s []float64, mean float64, wts []float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MomentAbout(n, s, mean, wts)
	}
}

func BenchmarkMomentAboutSmall(b *testing.B) {
	s := RandomSlice(small)
	benchmarkMomentAbout(b, 5, s, 0, nil)
}

func BenchmarkMomentAboutMedium(b *testing.B) {
	s := RandomSlice(medium)
	benchmarkMomentAbout(b, 5, s, 0, nil)
}

func BenchmarkMomentAboutLarge(b *testing.B) {
	s := RandomSlice(large)
	benchmarkMomentAbout(b, 5, s, 0, nil)
}

func BenchmarkMomentAboutHuge(b *testing.B) {
	s := RandomSlice(huge)
	benchmarkMomentAbout(b, 5, s, 0, nil)
}

func BenchmarkMomentAboutSmallWeighted(b *testing.B) {
	s := RandomSlice(small)
	wts := RandomSlice(small)
	benchmarkMomentAbout(b, 5, s, 0, wts)
}

func BenchmarkMomentAboutMediumWeighted(b *testing.B) {
	s := RandomSlice(medium)
	wts := RandomSlice(medium)
	benchmarkMomentAbout(b, 5, s, 0, wts)
}

func BenchmarkMomentAboutLargeWeighted(b *testing.B) {
	s := RandomSlice(large)
	wts := RandomSlice(large)
	benchmarkMomentAbout(b, 5, s, 0, wts)
}

func BenchmarkMomentAboutHugeWeighted(b *testing.B) {
	s := RandomSlice(huge)
	wts := RandomSlice(huge)
	benchmarkMomentAbout(b, 5, s, 0, wts)
}
