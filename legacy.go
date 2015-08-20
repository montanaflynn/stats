package stats

func VarP(input []float64) (sdev float64, err error) {
	return PopulationVariance(input)
}

func VarS(input []float64) (sdev float64, err error) {
	return SampleVariance(input)
}

func StdDevP(input []float64) (sdev float64, err error) {
	return StandardDeviationPopulation(input)
}

func StdDevS(input []float64) (sdev float64, err error) {
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
