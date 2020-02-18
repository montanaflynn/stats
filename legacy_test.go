package stats_test

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// Create working sample data to test if the legacy
// functions cause a runtime crash or return an error
func TestLegacy(t *testing.T) {

	// Slice of data
	s := []float64{-10, -10.001, 5, 1.1, 2, 3, 4.20, 5}

	// Slice of coordinates
	d := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	// VarP rename compatibility
	_, err := stats.VarP(s)
	if err != nil {
		t.Errorf("VarP not successfully returning PopulationVariance.")
	}

	// VarS rename compatibility
	_, err = stats.VarS(s)
	if err != nil {
		t.Errorf("VarS not successfully returning SampleVariance.")
	}

	// StdDevP rename compatibility
	_, err = stats.StdDevP(s)
	if err != nil {
		t.Errorf("StdDevP not successfully returning StandardDeviationPopulation.")
	}

	// StdDevS rename compatibility
	_, err = stats.StdDevS(s)
	if err != nil {
		t.Errorf("StdDevS not successfully returning StandardDeviationSample.")
	}

	// LinReg rename compatibility
	_, err = stats.LinReg(d)
	if err != nil {
		t.Errorf("LinReg not successfully returning LinearRegression.")
	}

	// ExpReg rename compatibility
	_, err = stats.ExpReg(d)
	if err != nil {
		t.Errorf("ExpReg not successfully returning ExponentialRegression.")
	}

	// LogReg rename compatibility
	_, err = stats.LogReg(d)
	if err != nil {
		t.Errorf("LogReg not successfully returning LogarithmicRegression.")
	}
}
