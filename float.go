// Copyright (c) 2016, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package float

import "math"

var minNormalFloat32 = math.Float32frombits(0x00800000)

func Next64(a float64) float64 {
	if math.IsInf(a, 0) || math.IsNaN(a) {
		return a
	}
	return math.Float64frombits(math.Float64bits(a) + 1)
}

func Next32(a float32) float32 {
	if math.IsInf(float64(a), 0) || math.IsNaN(float64(a)) {
		return a
	}
	return math.Float32frombits(math.Float32bits(a) + 1)
}

func Equals64(a, b, epsilon float64) bool {
	diff := math.Abs(a - b)
	if a == b {
		// Shortcuts and handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < minNormalFloat32 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < epsilon*minNormalFloat32
	} else {
		// Use relative error.
		return diff/math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64) < epsilon
	}
}

func Equals32(a, b, epsilon float32) bool {
	diff := math.Abs(a - b)
	if a == b {
		// Shortcuts and handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < minNormalFloat32 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < epsilon*minNormalFloat32
	} else {
		// Use relative error.
		return diff/math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64) < epsilon
	}
}
