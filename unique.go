package stats

import "errors"

// Unique returns any unique numbers in the data
func Unique(input Float64Data) (result []float64, err error) {

	// Return an error if there are no numbers
	if input.Len() == 0 {
		return nil, errors.New("Input must not be empty")
	}

	for i := range input {
		// fmt.Println(result, input[i], in(result, input[i]))
		if in(input, input[i]) == 1 {
			result = append(result, input[i])
		}
	}

	return result, nil
}
