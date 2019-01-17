package stats

// VarP is a shortcut to PopulationVariance
func VarP(input Float64Data) (sdev float64, err error) {
	return PopulationVariance(input)
}

// VarS is a shortcut to SampleVariance
func VarS(input Float64Data) (sdev float64, err error) {
	return SampleVariance(input)
}

// StdDevP is a shortcut to StandardDeviationPopulation
func StdDevP(input Float64Data) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

// StdDevS is a shortcut to StandardDeviationSample
func StdDevS(input Float64Data) (sdev float64, err error) {
	return StandardDeviationSample(input)
}

// LinReg is a shortcut to LinearRegression
func LinReg(s []Coordinate) (regressions []Coordinate, err error) {
	return LinearRegression(s)
}

// ExpReg is a shortcut to ExponentialRegression
func ExpReg(s []Coordinate) (regressions []Coordinate, err error) {
	return ExponentialRegression(s)
}

// LogReg is a shortcut to LogarithmicRegression
func LogReg(s []Coordinate) (regressions []Coordinate, err error) {
	return LogarithmicRegression(s)
}

var (
	// EmptyInputErr legacy error name didn't start with Err
	EmptyInputErr = ErrEmptyInput
	// NaNErr legacy error name didn't start with Err
	NaNErr = ErrNaN
	// NegativeErr legacy error name didn't start with Err
	NegativeErr = ErrNegative
	// ZeroErr legacy error name didn't start with Err
	ZeroErr = ErrZero
	// BoundsErr legacy error name didn't start with Err
	BoundsErr = ErrBounds
	// SizeErr legacy error name didn't start with Err
	SizeErr = ErrSize
	// InfValue legacy error name didn't start with Err
	InfValue = ErrInfValue
	// YCoordErr legacy error name didn't start with Err
	YCoordErr = ErrYCoord
	// EmptyInput legacy error name didn't end with Err
	EmptyInput = ErrEmptyInput
)
