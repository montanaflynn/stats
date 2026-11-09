package stats

import "math"

// Entropy provides calculation of the entropy
func Entropy(input Float64Data) (float64, error) {
	input, err := normalize(input)
	if err != nil {
		return math.NaN(), err
	}
	var result float64
	for i := 0; i < input.Len(); i++ {
		v := input.Get(i)
		if v == 0 {
			continue
		}
		result += (v * math.Log(v))
	}
	return -result, nil
}

// normalize divides each value by the sum of all values,
// leaving the input itself unchanged.
func normalize(input Float64Data) (Float64Data, error) {
	if input.Len() == 0 {
		return nil, ErrEmptyInput
	}
	if hasNaN(input) {
		return Float64Data{math.NaN()}, nil
	}

	var sum float64
	f := math.Ldexp(1, -unitExponent(input))
	c := make(Float64Data, len(input))
	for i, v := range input {
		if v < 0 {
			return nil, ErrNegative
		}
		scaled := v * f
		c[i] = scaled
		sum += scaled
	}
	if sum == 0 {
		return nil, ErrZero
	}
	for i := range c {
		c[i] /= sum
	}
	return c, nil
}
