package stats

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := statsError{"test error"}
	if err.Error() != "test error" {
		t.Errorf("Error method message didn't match")
	}
	if err.String() != "test error" {
		t.Errorf("String method message didn't match")
	}
}

func TestErrTooFew(t *testing.T) {
	if !errors.Is(ErrTooFew, ErrEmptyInput) {
		t.Errorf("errors.Is(ErrTooFew, ErrEmptyInput) should be true")
	}
	if !errors.Is(ErrTooFew, ErrTooFew) {
		t.Errorf("errors.Is(ErrTooFew, ErrTooFew) should be true")
	}
	if errors.Is(ErrTooFew, ErrZero) {
		t.Errorf("errors.Is(ErrTooFew, ErrZero) should be false")
	}
	if errors.Is(ErrEmptyInput, ErrTooFew) {
		t.Errorf("errors.Is(ErrEmptyInput, ErrTooFew) should be false")
	}
	if error(ErrTooFew) == error(ErrEmptyInput) {
		t.Errorf("ErrTooFew should not equal ErrEmptyInput")
	}
	if ErrTooFew.Error() != "Input has too few elements." {
		t.Errorf("ErrTooFew message didn't match")
	}
}
