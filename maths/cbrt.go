package maths

import "math"

// Cbrt returns the cube root of x.
func Cbrt[T RationalNumber](n T) float64 {
   return math.Cbrt(float64(n))
}

// TODO: Add additional functions that return the same type given as the input.
//       They could be named e.g. CbrtFloor, CbrtCeil, CbrtRound, etc.

