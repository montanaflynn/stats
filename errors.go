// Package stats error values, all exported APIs
// should use these errors as returned values
// https://github.com/golang/go/wiki/Errors#naming
package stats

type statsError struct {
	err string
}

func (s statsError) Error() string {
	return s.err
}

func (s statsError) String() string {
	return s.err
}

// ErrEmptyInput Input must not be empty
var ErrEmptyInput = statsError{"Input must not be empty."}

// ErrNaN Not a number
var ErrNaN = statsError{"Not a number."}

// ErrNegative Must not contain negative values
var ErrNegative = statsError{"Must not contain negative values."}

// ErrZero Must not contain zero values
var ErrZero = statsError{"Must not contain zero values."}

// ErrBounds Input is outside of range
var ErrBounds = statsError{"Input is outside of range."}

// ErrSize Must be the same length
var ErrSize = statsError{"Must be the same length."}

// ErrInfValue Value is infinite
var ErrInfValue = statsError{"Value is infinite."}

// ErrYCoord Y Value must be greater than zero
var ErrYCoord = statsError{"Y Value must be greater than zero."}
