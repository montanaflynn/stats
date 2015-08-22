package stats

// Float64Data is a named type for []float64 with helper methods
type Float64Data []float64

// These methods implement standard interfaces
func (f Float64Data) Get(i int) float64  { return f[i] }
func (f Float64Data) Len() int           { return len(f) }
func (f Float64Data) Less(i, j int) bool { return f[i] < f[j] }
func (f Float64Data) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// Helper methods for common statistics functions
func (f Float64Data) Min() (float64, error)    { return Min(f) }
func (f Float64Data) Max() (float64, error)    { return Max(f) }
func (f Float64Data) Sum() (float64, error)    { return Sum(f) }
func (f Float64Data) Mean() (float64, error)   { return Mean(f) }
func (f Float64Data) Median() (float64, error) { return Median(f) }
func (f Float64Data) Mode() ([]float64, error) { return Mode(f) }
