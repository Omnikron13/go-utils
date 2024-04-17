package maths

import "math"

// Sqrt returns the square root of a number.
func Sqrt[T Rational](n T) float64 {
   return math.Sqrt(float64(n))
}

// TODO: Add additional functions that return the same type given as the input.
//       They could be named e.g. SqrtFloor, SqrtCeil, SqrtRound, etc.

