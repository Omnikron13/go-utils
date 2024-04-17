// maths contains genericised versions of as many of the functions found in the std math package as possible.
// Due ro quite obvious differeces ub precision between the Int types and float64, additional care should be taken
// with the most basic versiojs of provided function to enure they return what you want and epect.
// Wheverer deemed necessary (and possible) additioional functions will be provided that reuire qdditional arguments,
// return additional data, or error out in one way or another.
package maths


// UnsignedInt is a type constraint that represents a whole number that is greater than or equal to zero.
type UnsignedInt interface { uint8 | uint16 | uint32 | uint64 | uint }

// SignedInt is a type constraint that represents a number that can be positive, negative, or zero.
type SignedInt interface { int8 | int16 | int32 | int64 | int }

// Int is a type constraint that represents a whole number that can be positive, negative, or zero.
type Int interface { UnsignedInt | SignedInt }

// Float is a type constraint that represents a floating-point number.
type Float interface { float32 | float64 }

// RationalNumber is a type constraint that represents a number that can be positive, negative, or zero.
type RationalNumber interface { Int | Float }

