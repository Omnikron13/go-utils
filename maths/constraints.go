// Description: This file contains type constraints for the maths module.

package maths


// UnsignedInt is a type constraint that represents a whole number that is greater than or equal to zero.
type UnsignedInt interface { ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint }

// SignedInt is a type constraint that represents a number that can be positive, negative, or zero.
type SignedInt interface { ~int8 | ~int16 | ~int32 | ~int64 | ~int }

// Int is a type constraint that represents a whole number that can be positive, negative, or zero.
type Int interface { UnsignedInt | SignedInt }

// Float is a type constraint that represents a floating-point number.
type Float interface { ~float32 | ~float64 }

// Unsigned combines `UnsignedInt` with `uintptr`.
type Unsigned interface { UnsignedInt | ~uintptr }

// Signed is a type constraint that represents a number that can be positive, negative, or zero.
type Signed interface { SignedInt | Float }

// Rational is a type constraint that represents a number that can be positive, negative, or zero.
type Rational interface { Int | Float }

