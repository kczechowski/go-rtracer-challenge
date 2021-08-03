package rtracer

import (
	"github.com/kczechowski/go-rtracer-challenge/rtracer/math"
	math2 "math"
)

const TupleTypeVector = 0.0
const TupleTypePoint = 1.0

type Tuple [4]float64

func NewPoint(x, y, z float64) Tuple {
	t := Tuple{x, y, z, TupleTypePoint}

	return t
}

func NewVector(x, y, z float64) Tuple {
	t := Tuple{x, y, z, TupleTypeVector}

	return t
}

func (t Tuple) IsPoint() bool {
	return t[3] == TupleTypePoint
}

func (t Tuple) IsVector() bool {
	return t[3] == TupleTypeVector
}

func TupleEqual(t1, t2 Tuple) bool {
	return math.Float64Equal(t1[0], t2[0], math.DefaultFloat64EqualityEpsilon) &&
		math.Float64Equal(t1[1], t2[1], math.DefaultFloat64EqualityEpsilon) &&
		math.Float64Equal(t1[2], t2[2], math.DefaultFloat64EqualityEpsilon) &&
		math.Float64Equal(t1[3], t2[3], math.DefaultFloat64EqualityEpsilon)
}

func (t Tuple) Equals(t2 Tuple) bool {
	return TupleEqual(t, t2)
}

func TupleXYZEqual(t1, t2 Tuple) bool {
	return math.Float64Equal(t1[0], t2[0], math.DefaultFloat64EqualityEpsilon) &&
		math.Float64Equal(t1[1], t2[2], math.DefaultFloat64EqualityEpsilon) &&
		math.Float64Equal(t1[2], t2[2], math.DefaultFloat64EqualityEpsilon)
}

func (t Tuple) XYZEquals(t2 Tuple) bool {
	return TupleXYZEqual(t, t2)
}

func TupleAdd(t1, t2 Tuple) Tuple {
	rt := Tuple{}
	for i := 0; i < 4; i++ {
		rt[i] = t1[i] + t2[i]
	}

	return rt
}

func (t Tuple) Add(t2 Tuple) Tuple {
	for i := 0; i < 4; i++ {
		t[i] = t[i] + t2[i]
	}

	return t
}

func TupleSub(t1, t2 Tuple) Tuple {
	rt := Tuple{}
	for i := 0; i < 4; i++ {
		rt[i] = t1[i] - t2[i]
	}

	return rt
}

func (t Tuple) Sub(t2 Tuple) Tuple {
	for i := 0; i < 4; i++ {
		t[i] = t[i] - t2[i]
	}

	return t
}

func TupleNegate(t Tuple) Tuple {
	rt := Tuple{}
	for i := 0; i < 4; i++ {
		rt[i] = 0 - t[i]
	}

	return rt
}

func TupleMultiplyByScalar(t1 Tuple, scalar float64) Tuple {
	rt := Tuple{}
	for i := 0; i < 4; i++ {
		rt[i] = t1[i] * scalar
	}

	return rt
}

func (t Tuple) MultiplyByScalar(scalar float64) Tuple {
	for i := 0; i < 4; i++ {
		t[i] = t[i] * scalar
	}

	return t
}

func TupleDivideByScalar(t1 Tuple, scalar float64) Tuple {
	rt := Tuple{}
	for i := 0; i < 4; i++ {
		rt[i] = t1[i] / scalar
	}

	return rt
}

func (t Tuple) DivideByScalar(scalar float64) Tuple {
	for i := 0; i < 4; i++ {
		t[i] = t[i] / scalar
	}

	return t
}

func TupleMagnitude(t Tuple) float64 {
	return math2.Sqrt(math2.Pow(t[0], 2) + math2.Pow(t[1], 2) + math2.Pow(t[2], 2))
}

func (t Tuple) Magnitude() float64 {
	return math2.Sqrt(math2.Pow(t[0], 2) + math2.Pow(t[1], 2) + math2.Pow(t[2], 2))
}

func TupleNormalize(t Tuple) Tuple {
	magnitude := TupleMagnitude(t)
	return Tuple{t[0] / magnitude, t[1] / magnitude, t[2] / magnitude, t[3] / magnitude}
}

func (t Tuple) Normalize() Tuple {
	magnitude := TupleMagnitude(t)
	t[0] = t[0] / magnitude
	t[1] = t[1] / magnitude
	t[2] = t[2] / magnitude
	t[3] = t[3] / magnitude

	return t
}

func TupleDot(t1, t2 Tuple) float64 {
	return t1[0]*t2[0] + t1[1]*t2[1] + t1[2]*t2[2] + t1[3]*t2[3]
}

func TupleCross(t1, t2 Tuple) Tuple {
	return NewVector(t1[1]*t2[2]-t1[2]*t2[1], t1[2]*t2[0]-t1[0]*t2[2], t1[0]*t2[1]-t1[1]*t2[0])
}
