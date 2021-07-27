package math

import "math"

const DefaultFloat64EqualityEpsilon = 1e-9

func Float64Equal(f1, f2, epsilon float64) bool {
	return math.Abs(f1 - f2) <= epsilon
}
