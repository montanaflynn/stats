package stats

import "math"

// RemoveDuplicates removes duplicate numbers
// in a slice which also means it removes NaN
func RemoveDuplicates(input []float64) []float64 {
	seen := make(map[float64]bool)
	pos := 0
	for _, f := range input {
		if math.IsNaN(f) {
			continue
		}
		if _, ok := seen[f]; !ok {
			seen[f] = true
			input[pos] = f
			pos++
		}
	}
	input = input[:pos]
	return input
}

// FindUniques finds any unique numbers in a
// slice which are not repeated while ignoring
// NaN since they are always unique
func FindUniques(input []float64) []float64 {
	size := len(input)
	freq := []int{}
	for i := 0; i < size; i++ {
		freq = append(freq, -1)
	}
	output := []float64{}
	for i := 0; i < size; i++ {
		if math.IsNaN(input[i]) {
			continue
		}
		count := 1
		for j := i + 1; j < size; j++ {
			if input[i] == input[j] {
				count++
				freq[j] = 0
			}
		}
		if freq[i] != 0 {
			freq[i] = count
		}
	}
	for i := 0; i < size; i++ {
		if freq[i] == 1 {
			output = append(output, input[i])
		}
	}
	return output
}
