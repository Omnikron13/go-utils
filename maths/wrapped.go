// Description: This file contains the wrapped functions from the `math` package.
// The functions are wrapped to accept any numeric type in place of `float64`.
// The arguments are simply passed into `float64()` before being passed to the original function.
package maths

import "math"

// Acos returns the arccosine, in radians, of x.
func Acos[T RationalNumber](x T) float64 {
   return math.Acos(float64(x))
}

// Acosh returns the inverse hyperbolic cosine of x.
func Acosh[T RationalNumber](x T) float64 {
   return math.Acosh(float64(x))
}

// Asin returns the arcsine, in radians, of x.
func Asin[T RationalNumber](x T) float64 {
   return math.Asin(float64(x))
}

// Asinh returns the inverse hyperbolic sine of x.
func Asinh[T RationalNumber](x T) float64 {
   return math.Asinh(float64(x))
}

// Atan returns the arctangent, in radians, of x.
func Atan[T RationalNumber](x T) float64 {
   return math.Atan(float64(x))
}

// Atan2 returns the arctangent of y/x, in radians, using the signs of the two to determine the quadrant of the result.
func Atan2[T RationalNumber](y, x T) float64 {
   return math.Atan2(float64(y), float64(x))
}

// Atanh returns the inverse hyperbolic tangent of x.
func Atanh[T RationalNumber](x T) float64 {
   return math.Atanh(float64(x))
}

// Cos returns the cosine of the radian argument x.
func Cos[T RationalNumber](x T) float64 {
   return math.Cos(float64(x))
}

// Cosh returns the hyperbolic cosine of x.
func Cosh[T RationalNumber](x T) float64 {
   return math.Cosh(float64(x))
}

// Erf returns the error function of x.
func Erf[T RationalNumber](x T) float64 {
   return math.Erf(float64(x))
}

// Erfc returns the complementary error function of x.
func Erfc[T RationalNumber](x T) float64 {
   return math.Erfc(float64(x))
}

// Erfcinv returns the inverse of Erfc(x).
func Erfcinv[T RationalNumber](x T) float64 {
   return math.Erfcinv(float64(x))
}

// Erfinv returns the inverse error function of x.
func Erfinv[T RationalNumber](x T) float64 {
   return math.Erfinv(float64(x))
}

// Exp returns e**x, the base-e exponential of x.
func Exp[T RationalNumber](x T) float64 {
   return math.Exp(float64(x))
}

// Exp2 returns 2**x, the base-2 exponential of x.
func Exp2[T RationalNumber](x T) float64 {
   return math.Exp2(float64(x))
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1.
func Expm1[T RationalNumber](x T) float64 {
   return math.Expm1(float64(x))
}

// TODO: test if R can be inferred
// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
func Hypot[T1 RationalNumber, T2 RationalNumber, R Float](p T1, q T2) R {
   return R(math.Hypot(float64(p), float64(q)))
}

// TODO: test if R can be inferred
// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf[T RationalNumber, R Float](sign T) R {
   return R(math.Inf(int(sign)))
}

// IsInf reports whether f is an infinity, according to sign.
func IsInf[T1 RationalNumber, T2 RationalNumber](f T1, sign T2) bool {
   return math.IsInf(float64(f), int(sign))
}

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN[T RationalNumber](f T) bool {
   return math.IsNaN(float64(f))
}

// J0 returns the order-zero Bessel function of the first kind.
func J0[T RationalNumber](x T) float64 {
   return math.J0(float64(x))
}

// J1 returns the order-one Bessel function of the first kind.
func J1[T RationalNumber](x T) float64 {
   return math.J1(float64(x))
}

// Jn returns the order-n Bessel function of the first kind.
func Jn[T1 RationalNumber, T2 RationalNumber](n T1, x T2) float64 {
   return math.Jn(int(n), float64(x))
}

// TODO: test if R can be inferred
// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN[R Float]() R {
   return R(math.NaN())
}

// TODO: additional options? generic return? Next/Prev using min/max values?
// Nextafter returns the next representable float64 value after x towards y.
func Nextafter[T1 RationalNumber, T2 RationalNumber](x T1, y T2) float64 {
   return math.Nextafter(float64(x), float64(y))
}

// Nextafter32 returns the next representable float32 value after x towards y.
func Nextafter32[T1 RationalNumber, T2 RationalNumber](x T1, y T2) float32 {
   return float32(math.Nextafter32(float32(x), float32(y)))
}

// Pow returns x**y, the base-x exponential of y.
func Pow[T1 RationalNumber, T2 RationalNumber](x T1, y T2) float64 {
   return math.Pow(float64(x), float64(y))
}

// Pow10 returns 10**e, the base-10 exponential of n.
func Pow10[T RationalNumber](n T) float64 {
   return math.Pow10(int(n))
}

// Signbit reports whether x is negative or negative zero.
func Signbit[T RationalNumber](x T) bool {
   return math.Signbit(float64(x))
}

// Sin returns the sine of the radian argument x.
func Sin[T RationalNumber](x T) float64 {
   return math.Sin(float64(x))
}

// Sincos returns returns Sin(x), Cos(x).
func Sincos[T RationalNumber](x T) (float64, float64) {
   return math.Sincos(float64(x))
}

// Sinh returns the hyperbolic sine of x.
func Sinh[T RationalNumber](x T) float64 {
   return math.Sinh(float64(x))
}

// Tan returns the tangent of the radian argument x.
func Tan[T RationalNumber](x T) float64 {
   return math.Tan(float64(x))
}

// Tanh returns the hyperbolic tangent of x.
func Tanh[T RationalNumber](x T) float64 {
   return math.Tanh(float64(x))
}

// Y0 returns the order-zero Bessel function of the second kind.
func Y0[T RationalNumber](x T) float64 {
   return math.Y0(float64(x))
}

// Y1 returns the order-one Bessel function of the second kind.
func Y1[T RationalNumber](x T) float64 {
   return math.Y1(float64(x))
}

// Yn returns the order-n Bessel function of the second kind.
func Yn[T1 RationalNumber, T2 RationalNumber](n T1, x T2) float64 {
   return math.Yn(int(n), float64(x))
}

// NOTE: The following are out of order as they purely pass through to `math`, with no attempt to genericise them.

// Float32bits returns the IEEE 754 binary representation of f.
func Float32bits(f float32) uint32 {
   return math.Float32bits(f)
}

// Float32frombits returns the floating point number corresponding the IEEE 754 binary representation b.
func Float32frombits(b uint32) float32 {
   return math.Float32frombits(b)
}

// Float64bits returns the IEEE 754 binary representation of f.
func Float64bits(f float64) uint64 {
   return math.Float64bits(f)
}

// Float64frombits returns the floating point number corresponding the IEEE 754 binary representation b.
func Float64frombits(b uint64) float64 {
   return math.Float64frombits(b)
}

