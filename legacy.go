package stats

func VarP(input Float64Data) (sdev float64, err error) {
	return PopulationVariance(input)
}

func VarS(input Float64Data) (sdev float64, err error) {
	return SampleVariance(input)
}

func StdDevP(input Float64Data) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

func StdDevS(input Float64Data) (sdev float64, err error) {
	return StandardDeviationSample(input)
}

func LinReg(s []Coordinate) (regressions []Coordinate, err error) {
	return LinearRegression(s)
}

func ExpReg(s []Coordinate) (regressions []Coordinate, err error) {
	return ExponentialRegression(s)
}

func LogReg(s []Coordinate) (regressions []Coordinate, err error) {
	return LogarithmicRegression(s)
}
