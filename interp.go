package stats

import "sort"

// Interp calculates the one-dimensional piecewise-linear interpolant to a
// function with given discrete data points (xp, fp), evaluated at each x.
// Values of x below xp[0] return fp[0] and values above xp[len(xp)-1] return
// fp[len(xp)-1], so no extrapolation is performed. Unlike numpy's interp,
// which silently returns nonsense for unsorted coordinates, xp must be
// strictly increasing or ErrBounds is returned. An empty x or xp returns
// ErrEmptyInput and xp and fp of different lengths return ErrSize.
func Interp(x, xp, fp Float64Data) ([]float64, error) {

	if x.Len() == 0 || xp.Len() == 0 {
		return nil, ErrEmptyInput
	}

	if xp.Len() != fp.Len() {
		return nil, ErrSize
	}

	for i := 1; i < xp.Len(); i++ {
		if xp[i] <= xp[i-1] {
			return nil, ErrBounds
		}
	}

	output := make([]float64, x.Len())

	for i, xv := range x {
		switch {
		case xv <= xp[0]:
			output[i] = fp[0]
		case xv >= xp[xp.Len()-1]:
			output[i] = fp[fp.Len()-1]
		default:
			// The first index with xp[j] >= xv, which the clamping
			// above guarantees is within [1, len(xp)-1]
			j := sort.SearchFloat64s(xp, xv)
			t := (xv - xp[j-1]) / (xp[j] - xp[j-1])
			output[i] = fp[j-1] + t*(fp[j]-fp[j-1])
		}
	}

	return output, nil
}
