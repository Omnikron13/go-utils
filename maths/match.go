package maths

import "math"

// Abs returns the absolute value of x.
func Abs[T RationalNumber](x T) T {
   // TODO: handle overflow for min value of signed integers
   if x >= 0 {
      return x
   }
   return -x
}

// TODO: Expand this to include additional float types (float16, and perhaps float128/float8 of some kind..?)
// Mod returns the remainder of x/y.
func Mod[T RationalNumber](x, y T) T {
   return T(math.Mod(float64(x), float64(y)))
}

// Modf returns integer and fractional floating-point numbers that sum to f. Both values have the same sign as f.
func Modf[T RationalNumber](f T) (T, T) {
   int, frac := math.Modf(float64(f))
   return T(int), T(frac)
}

