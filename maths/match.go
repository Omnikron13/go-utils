package maths

import "math"

// Abs returns the absolute value of x.
func Abs[T Rational](x T) T {
   // TODO: handle overflow for min value of signed integers
   if x >= 0 {
      return x
   }
   return -x
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil[T Float](x T) T {
   return T(math.Ceil(float64(x)))
}

// Copysign returns a value with the magnitude of x and the sign of sign.
func Copysign[T Signed](x, sign T) T {
   return T(math.Copysign(float64(x), float64(sign)))
}

// Floor returns the greatest integer value less than or equal to x.
func Floor[T Float](x T) T {
   return T(math.Floor(float64(x)))
}

// Log returns the natural logarithm of x.
func Log[T Rational](x T) T {
   return T(math.Log(float64(x)))
}

// Log10 returns the decimal logarithm of x.
func Log10[T Rational](x T) T {
   return T(math.Log10(float64(x)))
}

// Log1p returns the natural logarithm of 1 plus its argument.
func Log1p[T Rational](x T) T {
   return T(math.Log1p(float64(x)))
}

// Log2 returns the binary logarithm of x.
func Log2[T Rational](x T) T {
   return T(math.Log2(float64(x)))
}

// Logb returns the binary exponent of x.
func Logb[T Rational](x T) T {
   return T(math.Logb(float64(x)))
}

// Remainder returns the IEEE 754 floating-point remainder of x/y.
func Remainder[T Float](x, y T) T {
   return T(math.Remainder(float64(x), float64(y)))
}

// Trunc returns the integer value of x.
func Trunc[T Rational](x T) T {
   return T(math.Trunc(float64(x)))
}

// Round returns the nearest integer, rounding half away from zero.
func Round[T Float](x T) T {
   return T(math.Round(float64(x)))
}

// RoundToEven returns the nearest integer, rounding ties to even.
func RoundToEven[T Float](x T) T {
   return T(math.RoundToEven(float64(x)))
}

// Max returns the larger of x or y.
func Max[T Rational](x, y T) T {
   return max(x, y)
}

// Min returns the smaller of x or y.
func Min[T Rational](x, y T) T {
   return min(x, y)
}

// TODO: Expand this to include additional float types (float16, and perhaps float128/float8 of some kind..?)
// Mod returns the remainder of x/y.
func Mod[T Rational](x, y T) T {
   return T(math.Mod(float64(x), float64(y)))
}

// Modf returns integer and fractional floating-point numbers that sum to f. Both values have the same sign as f.
func Modf[T Rational](f T) (T, T) {
   int, frac := math.Modf(float64(f))
   return T(int), T(frac)
}
