package stats

import (
	"fmt"
	"testing"
)

func ExampleSoftMax() {
	sm, _ := SoftMax([]float64{3.0, 1.0, 0.2})
	fmt.Println(sm)
	// Output: [0.8360188027814407 0.11314284146556013 0.05083835575299916]
}

func TestSoftMaxEmptyInput(t *testing.T) {
	_, err := SoftMax([]float64{})
	if err != EmptyInputErr {
		t.Errorf("Should have returned empty input error")
	}
}

func TestSoftMax(t *testing.T) {
	sm, err := SoftMax([]float64{3.0, 1.0, 0.2})
	if err != nil {
		t.Error(err)
	}

	a := 0.8360188027814407
	if sm[0] != a {
		t.Errorf("%v != %v", sm[0], a)
	}

	a = 0.11314284146556013
	if sm[1] != a {
		t.Errorf("%v != %v", sm[1], a)
	}

	a = 0.05083835575299916
	if sm[2] != a {
		t.Errorf("%v != %v", sm[1], a)
	}
}
