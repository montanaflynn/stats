package stats_test

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
)

func ExampleSigmoid() {
	s, _ := stats.Sigmoid([]float64{3.0, 1.0, 2.1})
	fmt.Println(s)
	// Output: [0.9525741268224334 0.7310585786300049 0.8909031788043871]
}

func TestSigmoidEmptyInput(t *testing.T) {
	_, err := stats.Sigmoid([]float64{})
	if err != stats.EmptyInputErr {
		t.Errorf("Should have returned empty input error")
	}
}

func TestSigmoid(t *testing.T) {
	sm, err := stats.Sigmoid([]float64{-0.54761371, 17.04850603, 4.86054302})
	if err != nil {
		t.Error(err)
	}

	a := 0.3664182235138545
	if sm[0] != a {
		t.Errorf("%v != %v", sm[0], a)
	}

	a = 0.9999999605608187
	if sm[1] != a {
		t.Errorf("%v != %v", sm[1], a)
	}

	a = 0.9923132671908277
	if sm[2] != a {
		t.Errorf("%v != %v", sm[2], a)
	}
}
