package regression

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	r := new(Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")
	r.Train(
		DataPoint(11.2, []float64{587000, 16.5, 6.2}),
		DataPoint(13.4, []float64{643000, 20.5, 6.4}),
		DataPoint(40.7, []float64{635000, 26.3, 9.3}),
		DataPoint(5.3, []float64{692000, 16.5, 5.3}),
		DataPoint(24.8, []float64{1248000, 19.2, 7.3}),
		DataPoint(12.7, []float64{643000, 16.5, 5.9}),
		DataPoint(20.9, []float64{1964000, 20.2, 6.4}),
		DataPoint(35.7, []float64{1531000, 21.3, 7.6}),
		DataPoint(8.7, []float64{713000, 17.2, 4.9}),
		DataPoint(9.6, []float64{749000, 14.3, 6.4}),
		DataPoint(14.5, []float64{7895000, 18.1, 6}),
		DataPoint(26.9, []float64{762000, 23.1, 7.4}),
		DataPoint(15.7, []float64{2793000, 19.1, 5.8}),
		DataPoint(36.2, []float64{741000, 24.7, 8.6}),
		DataPoint(18.1, []float64{625000, 18.6, 6.5}),
		DataPoint(28.9, []float64{854000, 24.9, 8.3}),
		DataPoint(14.9, []float64{716000, 17.9, 6.7}),
		DataPoint(25.8, []float64{921000, 22.4, 8.6}),
		DataPoint(21.7, []float64{595000, 20.2, 8.4}),
		DataPoint(25.7, []float64{3353000, 16.9, 6.7}),
	)
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

	// All vars are known to positively correlate with the murder rate
	for i, c := range r.coeff {
		if i == 0 {
			// This is the offset and not a coeff
			continue
		}
		if c < 0 {
			t.Errorf("Coefficient is negative, but shouldn't be: %.2f", c)
		}
	}

	//  We know this set has an R^2 above 80
	if r.R2 < 0.8 {
		t.Errorf("R^2 was %.2f, but we expected > 80", r.R2)
	}
}

func TestCrossApply(t *testing.T) {
	r := new(Regression)
	r.SetObserved("Input-Squared plus Input")
	r.SetVar(0, "Input")
	r.Train(
		DataPoint(6, []float64{2}),
		DataPoint(20, []float64{4}),
		DataPoint(30, []float64{5}),
		DataPoint(72, []float64{8}),
		DataPoint(156, []float64{12}),
	)
	r.AddCross(PowCross(0, 2))
	r.AddCross(PowCross(0, 7))
	err := r.Run()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
	if r.names.vars[1] != "(Input)^2" {
		t.Error("Name incorrect")
	}

	for i, c := range r.coeff {
		if i == 0 {
			// This is the offset and not a coeff
			continue
		}
		if c < 0 {
			t.Errorf("Coefficient is negative, but shouldn't be: %.2f", c)
		}
	}

	//  We know this set has an R^2 above 80
	if r.R2 < 0.8 {
		t.Errorf("R^2 was %.2f, but we expected > 80", r.R2)
	}

	// Test that predict uses the cross as well
	val, err := r.Predict([]float64{6})
	if err != nil {
		t.Error(err)
	}
	if val <= 41.999 && val >= 42.001 {
		t.Errorf("Expected 42, got %.2f", val)
	}
}
