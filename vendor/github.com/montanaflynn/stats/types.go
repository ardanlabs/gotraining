package stats

// Float64Data is a named type for []float64 with helper methods
type Float64Data []float64

// Get item in slice
func (f Float64Data) Get(i int) float64 { return f[i] }

// Len returns length of slice
func (f Float64Data) Len() int { return len(f) }

// Less returns if one number is less than another
func (f Float64Data) Less(i, j int) bool { return f[i] < f[j] }

// Swap switches out two numbers in slice
func (f Float64Data) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// Min returns the minimum number in the data
func (f Float64Data) Min() (float64, error) { return Min(f) }

// Max returns the maximum number in the data
func (f Float64Data) Max() (float64, error) { return Max(f) }

// Sum returns the total of all the numbers in the data
func (f Float64Data) Sum() (float64, error) { return Sum(f) }

// Mean returns the mean of the data
func (f Float64Data) Mean() (float64, error) { return Mean(f) }

// Median returns the median of the data
func (f Float64Data) Median() (float64, error) { return Median(f) }

// Mode returns the mode of the data
func (f Float64Data) Mode() ([]float64, error) { return Mode(f) }
