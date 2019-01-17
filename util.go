package stats

import (
	"math"
	"sort"
	"time"
)

// Float64ToInt rounds a float64 to an int
func Float64ToInt(input float64) (output int) {
	r, _ := Round(input, 0)
	return int(r)
}

// Counts returns a map with the times an item
// is included in the provided slice or array
func Counts(input Float64Data) map[float64]int {
	c := make(map[float64]int, len(input))
	for _, value := range input {
		if math.IsNaN(value) {
			continue
		}
		c[value]++
	}
	return c
}

// unixnano returns nanoseconds from UTC epoch
func unixnano() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func copyslice(input Float64Data) Float64Data {
	s := make(Float64Data, input.Len())
	copy(s, input)
	return s
}

// sortedCopy returns a sorted copy of float64s
func sortedCopy(input Float64Data) (copy Float64Data) {
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func sortedCopyDif(input Float64Data) (copy Float64Data) {
	if sort.Float64sAreSorted(input) {
		return input
	}
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}
