package rtracer

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

func IsPoint(t Tuple) bool {
	return t[3] == TupleTypePoint
}

func isVector(t Tuple) bool  {
	return t[3] == TupleTypeVector
}
