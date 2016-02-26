// Copyright (c) 2016, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package float

import "math"

// Because of compiler guarantees, these cannot be consts.
var (
	NaN32 = math.Float32frombits(0x7f800001)
	NaN64 = math.Float64frombits(0x7ff8000000000001)

	PosInf32 = math.Float32frombits(0x7f800000)
	NegInf32 = math.Float32frombits(0xff800000)
	PosInf64 = math.Float64frombits(0x7ff0000000000000)
	NegInf64 = math.Float64frombits(0xfff0000000000000)
)

const (
	MinNormal32 = math.Float32frombits(0x00800000)
	MinNormal64 = math.Float64frombits(0x0010000000000000)
)

func Next32(a float32) float32 {
	if a != a || math.IsInf(float64(a), 0) {
		return a
	}
	return math.Float32frombits(math.Float32bits(a) + 1)
}

func Next(a float64) float64 {
	if a != a || math.IsInf(a, 0) {
		return a
	}
	return math.Float64frombits(math.Float64bits(a) + 1)
}

/* this probably doesn't work at all */
func Equals32(af, bf, epsilonf float32) bool {
	a, b, epsilon := float64(af), float64(bf), float64(epsilonf)
	diff := math.Abs(a - b)
	if a == b {
		// Shortcuts and handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < float64(minNormalFloat32) {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < epsilon*float64(minNormalFloat32)
	} else {
		// Use relative error.
		return diff/math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat32) < epsilon
	}
}

func Equals(a, b, epsilon float64) bool {
	diff := math.Abs(a - b)
	if a == b {
		// Shortcuts and handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < minNormalFloat64 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < epsilon*minNormalFloat64
	} else {
		// Use relative error.
		return diff/math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64) < epsilon
	}
}
