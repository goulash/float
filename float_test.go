// Copyright (c) 2016, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package float

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext32(z *testing.T) {
	list := float32{+0.0}
}

func TestEquals(z *testing.T) {
	assert := assert.New(z)

	var (
		nan    = math.NaN()
		posinf = math.Inf(1)
		neginf = math.Inf(-1)
		max    = math.MaxFloat64
		min    = math.SmallestNonzeroFloat64
	)

	epsilon := 1e-5
	equals := func(a, b float64) bool { return Equals(a, b, epsilon) }

	// Regular large numbers - generally not problematic
	assert.True(equals(1000000, 1000001))
	assert.True(equals(1000001, 1000000))
	assert.False(equals(10000, 10001))
	assert.False(equals(10001, 10000))

	// Negative large numbers
	assert.True(equals(-1000000, -1000001))
	assert.True(equals(-1000001, -1000000))
	assert.False(equals(-10000, -10001))
	assert.False(equals(-10001, -10000))

	// Numbers around 1
	assert.True(equals(1.0000001, 1.0000002))
	assert.True(equals(1.0000002, 1.0000001))
	assert.False(equals(1.0002, 1.0001))
	assert.False(equals(1.0001, 1.0002))

	// Numbers around -1
	assert.True(equals(-1.000001, -1.000002))
	assert.True(equals(-1.000002, -1.000001))
	assert.False(equals(-1.0001, -1.0002))
	assert.False(equals(-1.0002, -1.0001))

	// Numbers between 1 and 0
	assert.True(equals(0.000000001000001, 0.000000001000002))
	assert.True(equals(0.000000001000002, 0.000000001000001))
	assert.False(equals(0.000000000001002, 0.000000000001001))
	assert.False(equals(0.000000000001001, 0.000000000001002))

	// Numbers between -1 and 0
	assert.True(equals(-0.000000001000001, -0.000000001000002))
	assert.True(equals(-0.000000001000002, -0.000000001000001))
	assert.False(equals(-0.000000000001002, -0.000000000001001))
	assert.False(equals(-0.000000000001001, -0.000000000001002))

	// Comparisons involving zero
	assert.True(equals(0.0, 0.0))
	assert.True(equals(0.0, -0.0))
	assert.True(equals(-0.0, -0.0))
	assert.False(equals(0.00000001, 0.0))
	assert.False(equals(0.0, 0.00000001))
	assert.False(equals(-0.00000001, 0.0))
	assert.False(equals(0.0, -0.00000001))

	assert.True(Equals(0.0, 1e-40, 0.01))
	assert.True(Equals(1e-40, 0.0, 0.01))
	assert.False(Equals(1e-40, 0.0, 0.000001))
	assert.False(Equals(0.0, 1e-40, 0.000001))

	assert.True(Equals(0.0, -1e-40, 0.1))
	assert.True(Equals(-1e-40, 0.0, 0.1))
	assert.False(Equals(-1e-40, 0.0, 0.00000001))
	assert.False(Equals(0.0, -1e-40, 0.00000001))

	// Comparisons involving extreme values (overlow potential)
	assert.True(equals(max, max))
	assert.False(equals(max, -max))
	assert.False(equals(-max, max))
	assert.False(equals(max, max/2))
	assert.False(equals(max, -max/2))
	assert.False(equals(-max, max/2))

	// Comparisons involving ininities
	assert.True(equals(posinf, posinf))
	assert.True(equals(neginf, neginf))
	assert.False(equals(neginf, posinf))
	assert.False(equals(posinf, max))
	assert.False(equals(neginf, -max))

	// Comparisons involving NaN values
	assert.False(equals(nan, nan))
	assert.False(equals(nan, 0.0))
	assert.False(equals(-0.0, nan))
	assert.False(equals(nan, -0.0))
	assert.False(equals(0.0, nan))
	assert.False(equals(nan, posinf))
	assert.False(equals(posinf, nan))
	assert.False(equals(nan, neginf))
	assert.False(equals(neginf, nan))
	assert.False(equals(nan, max))
	assert.False(equals(max, nan))
	assert.False(equals(nan, -max))
	assert.False(equals(-max, nan))
	assert.False(equals(nan, min))
	assert.False(equals(min, nan))
	assert.False(equals(nan, -min))
	assert.False(equals(-min, nan))

	// Comparisons of numbers on opposite sides of 0
	assert.False(equals(1.000000001, -1.0))
	assert.False(equals(-1.0, 1.000000001))
	assert.False(equals(-1.000000001, 1.0))
	assert.False(equals(1.0, -1.000000001))
	assert.True(equals(10*min, 10*-min))
	assert.False(equals(10000*min, 10000*-min))

	// The really tricky part - comparisons of numbers very close to zero.
	assert.True(equals(min, min))
	assert.True(equals(min, -min))
	assert.True(equals(-min, min))
	assert.True(equals(min, 0))
	assert.True(equals(0, min))
	assert.True(equals(-min, 0))
	assert.True(equals(0, -min))

	assert.False(equals(0.000000001, -min))
	assert.False(equals(0.000000001, min))
	assert.False(equals(min, 0.000000001))
	assert.False(equals(-min, 0.000000001))
}
