package maths

import (
   "math"
   "testing"

   "github.com/stretchr/testify/assert"
)

// Test matrix for each of the numeric types
var tests map[string]any = map[string]any {
   "uint8":   []uint8{+0, -0, 1,                 127,                 128,                 129,                  254,                  255},
   "uint16": []uint16{+0, -0, 1,               32767,               32768,               32769,                65534,                65535},
   "uint32": []uint32{+0, -0, 1,          2147483647,          2147483648,          2147483649,           4294967294,           4294967295},
   "uint64": []uint64{+0, -0, 1, 9223372036854775807, 9223372036854775808, 9223372036854775809, 18446744073709551614, 18446744073709551615},

   "int8":   []int8{+0, -0, -1, 1,                 -127,                 -128,                 126,                  127},
   "int16": []int16{+0, -0, -1, 1,               -32767,               -32768,               32766,                32767},
   "int32": []int32{+0, -0, -1, 1,          -2147483647,          -2147483648,          2147483646,           2147483647},
   "int64": []int64{+0, -0, -1, 1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},

   // TODO: set bounds properly
   "float32": []float32{+0, -0, -1, 1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},
   "float64": []float64{+0, -0, -1, 1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},
}

// TODO: document/comment this
// TODO: refactor this to not require the test string - should be able to infer it from the type constraint T
func testMatrix[T RationalNumber](t *testing.T, test string, f func(t *testing.T, v T)) {
   if v, ok := tests[test].([]T); ok {
      t.Run(test, func(t *testing.T) {
         for _, x := range v {
            f(t, T(x))
         }
      })
   }
}

// TODO: add benchmarkMatrix[T RationalNumber](b *testing.B, test string, f func(b *testing.B, v T))

// Test_Abs tests `Abs()` outputs against `math.Abs()` outputs.
func Test_Abs(t *testing.T) {
   testMatrix(t, "uint8", func(t *testing.T, x uint8) {
      assert.Equal(t, Abs(x), uint8(math.Abs(float64(x))))
   })
   testMatrix(t, "uint16", func(t *testing.T, x uint16) {
      assert.Equal(t, Abs(x), uint16(math.Abs(float64(x))))
   })
   testMatrix(t, "uint32", func(t *testing.T, x uint32) {
      assert.Equal(t, Abs(x), uint32(math.Abs(float64(x))))
   })
   testMatrix(t, "uint64", func(t *testing.T, x uint64) {
      // NOTE: uint64 is always positive, so Abs() should return the same value.
      //       However math.Abs() requires casting to and from float64, which isn't capable of representing all uint64 values.
      //       So we have to ignore what math.Abs() would do here, as our Abs() is more accurate.
      assert.Equal(t, Abs(x), x)
   })
   testMatrix(t, "int8", func(t *testing.T, x int8) {
      assert.Equal(t, Abs(x), int8(math.Abs(float64(x))))
   })
   testMatrix(t, "int16", func(t *testing.T, x int16) {
      assert.Equal(t, Abs(x), int16(math.Abs(float64(x))))
   })
   testMatrix(t, "int32", func(t *testing.T, x int32) {
      assert.Equal(t, Abs(x), int32(math.Abs(float64(x))))
   })
   testMatrix(t, "int64", func(t *testing.T, x int64) {
      // NOTE: int64 is not always positive, so Abs() should invert negatives.
      //       Again math.Abs() requires casting to and from float64, which isn't capable of representing all uint64 values.
      //       So we have to ignore what math.Abs() would do here, as our Abs() is more accurate.
      // TODO: handle signed integer overflows after Abs() is updated to handle them.
      if x < 0 {
         assert.Equal(t, Abs(x), -x)
      } else {
         assert.Equal(t, Abs(x), x)
      }
   })
   testMatrix(t, "float32", func(t *testing.T, x float32) {
      assert.Equal(t, Abs(x), float32(math.Abs(float64(x))))
   })
   testMatrix(t, "float64", func(t *testing.T, x float64) {
      assert.Equal(t, Abs(x), math.Abs(x))
   })
}

// TODO: refactor to use testMatrix() - after the refactor of testMatrix() is done!
// Test_Mod tests `Mod()` outputs against `math.Mod()` outputs.
func Test_Mod(t *testing.T) {
   for k, v := range tests {
      switch v.(type) {
         case []uint8:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]uint8) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []uint16:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]uint16) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []uint32:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]uint32) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []uint64:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]uint64) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []int8:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]int8) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []int16:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]int16) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []int32:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]int32) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []int64:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]int64) {
                  assert.Equal(t, float64(Mod(x, 2)), math.Mod(float64(x), 2))
               }
            })
         case []float32:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]float32) {
                  assert.Equal(t, Mod(x, 1.2), float32(math.Mod(float64(x), float64(float32(1.2)))))
               }
               // TODO: check special cases
            })
         case []float64:
            t.Run(k, func(t *testing.T) {
               for _, x := range v.([]float64) {
                  assert.Equal(t, Mod(x, 1.2), float64(math.Mod(float64(x), 1.2)))
               }
               // TODO: check special cases
            })
      }
   }
}

