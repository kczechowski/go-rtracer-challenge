package rtracer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	assert.True(t, p.IsPoint())
	assert.False(t, p.IsVector())

	assert.Equal(t, 4.3, p[0])
	assert.Equal(t, -4.2, p[1])
	assert.Equal(t, 3.1, p[2])
}

func TestTupleIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	assert.False(t, v.IsPoint())
	assert.True(t, v.IsVector())

	assert.Equal(t, 4.3, v[0])
	assert.Equal(t, -4.2, v[1])
	assert.Equal(t, 3.1, v[2])
}

func TestTupleEquality(t *testing.T) {
	assert.True(t, TupleEqual(NewPoint(0, 0, 0), NewPoint(0, 0, 0)))
	assert.False(t, TupleEqual(NewPoint(1, 0, 0), NewPoint(0, 0, 0)))
	assert.False(t, TupleEqual(NewPoint(1, 0, 0), NewVector(0, 0, 0)))
	assert.True(t, TupleXYZEqual(NewPoint(0, 0, 0), NewVector(0, 0, 0)))
}

func TestTupleAdd(t *testing.T) {
	t1 := NewPoint(3.0,-2.0,5.0)
	t2 := NewVector(-2.0,3.0,1.0)
	t3 := TupleAdd(t1, t2)

	assert.Equal(t, 1.0, t3[0])
	assert.Equal(t, 1.0, t3[1])
	assert.Equal(t, 6.0, t3[2])
	assert.Equal(t, 1.0, t3[3])
}

func TestTupleSub(t *testing.T) {
	t1 := NewPoint(3.0,2.0,1.0)
	t2 := NewPoint(5.0,6.0,7.0)
	t3 := TupleSub(t1, t2)

	assert.Equal(t, -2.0, t3[0])
	assert.Equal(t, -4.0, t3[1])
	assert.Equal(t, -6.0, t3[2])
	assert.True(t, t3.IsVector())
}

/*
​ 	​Scenario​: Negating a tuple
​ 	  ​Given​ a ← tuple(1, -2, 3, -4)
​ 	  ​Then​ -a = tuple(-1, 2, -3, 4)
*/
func TestTupleNegate(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := TupleNegate(t1)

	assert.Equal(t, -1.0, t2[0])
	assert.Equal(t, 2.0, t2[1])
	assert.Equal(t, -3.0, t2[2])
	assert.Equal(t, 4.0, t2[3])
}

/*
​ 	​Scenario​: Multiplying a tuple by a scalar
​ 	  ​Given​ a ← tuple(1, -2, 3, -4)
​ 	  ​Then​ a * 3.5 = tuple(3.5, -7, 10.5, -14)
​ 
​ 	​Scenario​: Multiplying a tuple by a fraction
​ 	  ​Given​ a ← tuple(1, -2, 3, -4)
​ 	  ​Then​ a * 0.5 = tuple(0.5, -1, 1.5, -2)
*/
func TestTupleMultiplyByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := TupleMultiplyByScalar(t1, 3.5)

	assert.Equal(t, 3.5, t2[0])
	assert.Equal(t, -7.0, t2[1])
	assert.Equal(t, 10.5, t2[2])
	assert.Equal(t, -14.0, t2[3])

	t3 := TupleMultiplyByScalar(t1, 0.5)
	assert.Equal(t, 0.5, t3[0])
	assert.Equal(t, -1.0, t3[1])
	assert.Equal(t, 1.5, t3[2])
	assert.Equal(t, -2.0, t3[3])
}

/*
​ 	​Scenario​: Dividing a tuple by a scalar
​ 	  ​Given​ a ← tuple(1, -2, 3, -4)
​ 	  ​Then​ a / 2 = tuple(0.5, -1, 1.5, -2)
*/
func TestTupleDivideByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := TupleDivideByScalar(t1,2.0)

	assert.Equal(t, 0.5, t2[0])
	assert.Equal(t, -1.0, t2[1])
	assert.Equal(t, 1.5, t2[2])
	assert.Equal(t, -2.0, t2[3])
}

/*
​ 	​Scenario​: Computing the magnitude of vector(1, 0, 0)
​ 	  ​Given​ v ← vector(1, 0, 0)
​ 	  ​Then​ magnitude(v) = 1
​ 	​Scenario​: Computing the magnitude of vector(0, 1, 0)
​ 	  ​Given​ v ← vector(0, 1, 0)
​ 	  ​Then​ magnitude(v) = 1
​ 	​Scenario​: Computing the magnitude of vector(0, 0, 1)
​ 	  ​Given​ v ← vector(0, 0, 1)
​ 	  ​Then​ magnitude(v) = 1
*/
func TestTupleMagnitude(t *testing.T) {
	assert.Equal(t, 1.0, TupleMagnitude(NewVector(0,1,0)))
	assert.Equal(t, 1.0, TupleMagnitude(NewVector(1,0,0)))
	assert.Equal(t, 1.0, TupleMagnitude(NewVector(0,0,1)))
}

/*
​ 	​Scenario​: Normalizing vector(4, 0, 0) gives (1, 0, 0)
​ 	  ​Given​ v ← vector(4, 0, 0)
​ 	  ​Then​ normalize(v) = vector(1, 0, 0)
​ 	​Scenario​: The magnitude of a normalized vector
​ 	  ​Given​ v ← vector(1, 2, 3)
​ 	  ​When​ norm ← normalize(v)
​ 	  ​Then​ magnitude(norm) = 1
*/
func TestTupleNormalize(t *testing.T)  {
	t1 := NewVector(4,0,0)
	nt1 := TupleNormalize(t1)
	assert.Equal(t, 1.0, nt1[0])
	assert.Equal(t, 0.0, nt1[1])
	assert.Equal(t, 0.0, nt1[2])

	t2 := NewVector(1,2,3)
	nt2 := TupleNormalize(t2)
	assert.Equal(t, 1.0, nt2.Magnitude())
}

/*
​ 	​Scenario​: The dot product of two tuples
​ 	  ​Given​ a ← vector(1, 2, 3)
​ 	    ​And​ b ← vector(2, 3, 4)
​ 	  ​Then​ dot(a, b) = 20
*/
func TestTupleDot(t *testing.T) {
	t1 := NewVector(1,2,3)
	t2 := NewVector(2,3,4)
	assert.Equal(t, 20.0, TupleDot(t1, t2))
}

/*
​ 	​Scenario​: The cross product of two vectors
​ 	  ​Given​ a ← vector(1, 2, 3)
​ 	    ​And​ b ← vector(2, 3, 4)
​ 	  ​Then​ cross(a, b) = vector(-1, 2, -1)
​ 	    ​And​ cross(b, a) = vector(1, -2, 1)
*/
func TestTupleCross(t *testing.T) {
	t1 := NewVector(1,2,3)
	t2 := NewVector(2,3,4)

	c1 := TupleCross(t1, t2)
	assert.True(t, TupleEqual(c1, NewVector(-1, 2, -1)))

	c2 := TupleCross(t2, t1)
	assert.True(t, TupleEqual(c2, NewVector(1, -2, 1)))
}