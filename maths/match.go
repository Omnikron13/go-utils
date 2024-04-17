package maths

import "math"

// TODO: Expand this to include additional float types (float16, and perhaps float128/float8 of some kind..?)
// Mod returns the remainder of x/y.
func Mod[T RationalNumber](x, y T) T {
   return T(math.Mod(float64(x), float64(y)))
}

