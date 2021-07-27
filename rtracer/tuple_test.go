package rtracer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	assert.True(t, IsPoint(p))
	assert.False(t, isVector(p))

	assert.Equal(t, 4.3, p[0])
	assert.Equal(t, -4.2, p[1])
	assert.Equal(t, 3.1, p[2])
}

func TestTupleIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	assert.False(t, IsPoint(v))
	assert.True(t, isVector(v))

	assert.Equal(t, 4.3, v[0])
	assert.Equal(t, -4.2, v[1])
	assert.Equal(t, 3.1, v[2])
}
