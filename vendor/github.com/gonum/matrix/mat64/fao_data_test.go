// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64_test

import "github.com/gonum/matrix/mat64"

// FAO is a dataset extracted from Food and Agriculture Organization of the
// United Nations "FAO Statistical Pocketbook: World Food and Agriculture 2015".
// pp49-52.
var FAO = struct {
	Africa                *mat64.Dense
	Asia                  *mat64.Dense
	LatinAmericaCaribbean *mat64.Dense
	Oceania               *mat64.Dense
}{
	Africa: mat64.NewDense(21, 3, []float64{
		// 1990, 2000, 2014
		35.3, 38, 30.7, // Employment in agriculture (%)
		9.2, 20.3, 25.2, // Employment in agriculture, female (%)
		3163, 14718, 20667, // Energy consump, power irrigation (mln kWh)

		2597, 2717, 2903, // Dietary energy supply (kcal/pc/day)
		113, 116, 123, // Average dietary energy supply adequacy (%)
		58, 55, 52, // Dietary en supp, cereals/roots/tubers (%)
		18.6, 15, 10.8, // Prevalence of undernourishment (%)
		8832, 10241, 13915, // GDP per capita (US$, PPP)
		-0.4, -0.2, 50.7, // Cereal import dependency ratio (%)
		78.5, 83, 88.7, // Improved water source (% pop)

		// Production indices (2004-06=100)
		73, 90, 121, // Net food
		72, 89, 123, // Net crops
		82, 92, 123, // Cereals
		51, 77, 141, // Vegetable oils
		74, 94, 119, // Roots and tubers
		58, 86, 127, // Fruit and vegetables
		86, 93, 132, // Sugar
		76, 92, 115, // Livestock
		83, 89, 114, // Milk
		74, 91, 118, // Meat
		72, 92, 119, // Fish
	}),

	Asia: mat64.NewDense(21, 3, []float64{
		// 1990, 2000, 2014
		30.9, 24.5, 27.6, // Employment in agriculture (%)
		40.9, 29.4, 31.1, // Employment in agriculture, female (%)
		7614, 38316, 82411, // Energy consump, power irrigation (mln kWh)

		2320, 2402, 2581, // Dietary energy supply (kcal/pc/day)
		107, 110, 117, // Average dietary energy supply adequacy (%)
		66, 65, 63, // Dietary en supp, cereals/roots/tubers (%)
		27.6, 25.7, 19.8, // Prevalence of undernourishment (%)
		3315, 3421, 4575, // GDP per capita (US$, PPP)
		25.9, 28.1, 42, // Cereal import dependency ratio (%)
		55.5, 61.1, 68.7, // Improved water source (% pop)

		// Production indices (2004-06=100)
		60, 82, 129, // Net food
		59, 82, 127, // Net crops
		66, 79, 131, // Cereals
		58, 79, 128, // Vegetable oils
		50, 80, 133, // Roots and tubers
		58, 82, 124, // Fruit and vegetables
		76, 94, 114, // Sugar
		65, 84, 126, // Livestock
		59, 77, 125, // Milk
		67, 87, 127, // Meat
		65, 90, 119, // Fish
	}),

	LatinAmericaCaribbean: mat64.NewDense(14, 3, []float64{
		// 1990, 2000, 2014
		19.5, 14.2, 15.8, // Employment in agriculture (%)
		13.7, 6.2, 7.6, // Employment in agriculture, female (%)

		2669, 2787, 3069, // Dietary energy supply (kcal/pc/day)
		117, 120, 129, // Average dietary energy supply adequacy (%)
		42, 41, 40, // Dietary en supp, cereals/roots/tubers (%)
		14.7, 12.1, 5.5, // Prevalence of undernourishment (%)
		9837, 10976, 13915, // GDP per capita (US$, PPP)
		13, 12, 49.7, // Cereal import dependency ratio (%)
		85.1, 89.8, 94, // Improved water source (% pop)

		// Production indices (2004-06=100)
		60, 83, 129, // Net food
		64, 83, 131, // Net crops
		62, 88, 139, // Cereals
		58, 84, 123, // Livestock
		82, 107, 71, // Fish
	}),

	Oceania: mat64.NewDense(21, 3, []float64{
		// 1990, 2000, 2014
		6.2, 17.1, 3.8, // Employment in agriculture (%)
		4.5, 3.9, 4.4, // Employment in agriculture, female (%)
		415, 1028, 8667, // Energy consump, power irrigation (mln kWh)

		2454, 2436, 2542, // Dietary energy supply (kcal/pc/day)
		113, 112, 114, // Average dietary energy supply adequacy (%)
		49, 50, 48, // Dietary en supp, cereals/roots/tubers (%)
		15.7, 16.1, 14.2, // Prevalence of undernourishment (%)
		2269, 2536, 3110, // GDP per capita (US$, PPP)
		95.2, 95.9, 95.4, // Cereal import dependency ratio (%)
		49.7, 53.2, 55.5, // Improved water source (% pop)

		// Production indices (2004-06=100)
		72, 99, 116, // Net food
		69, 105, 126, // Net crops
		77, 113, 117, // Cereals
		41, 122, 215, // Vegetable oils
		80, 90, 110, // Roots and tubers
		66, 88, 104, // Fruit and vegetables
		70, 104, 71, // Sugar
		79, 97, 107, // Livestock
		56, 92, 113, // Milk
		79, 96, 105, // Meat
		51, 78, 85, // Fish
	}),
}
