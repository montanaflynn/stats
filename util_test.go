package stats

import (
	"testing"
)

func TestFloat64ToInt(t *testing.T) {
	m := float64ToInt(234.0234)
	if m != 234 {
		t.Errorf("%x != %x", m, 234)
	}
	m = float64ToInt(-234.0234)
	if m != -234 {
		t.Errorf("%x != %x", m, -234)
	}
	m = float64ToInt(1)
	if m != 1 {
		t.Errorf("%x != %x", m, 1)
	}
}
