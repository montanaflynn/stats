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

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	ErrEmptyInput = statsError{"Input must not be empty."}
	ErrNaN        = statsError{"Not a number."}
	ErrNegative   = statsError{"Must not contain negative values."}
	ErrZero       = statsError{"Must not contain zero values."}
	ErrBounds     = statsError{"Input is outside of range."}
	ErrSize       = statsError{"Must be the same length."}
	ErrInfValue   = statsError{"Value is infinite."}
	ErrYCoord     = statsError{"Y Value must be greater than zero."}
)
