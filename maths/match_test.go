package maths

import (
   "math"
   "testing"

   "github.com/stretchr/testify/assert"
)

// Test matrix for each of the numeric types
var tests map[string]any = map[string]any {
   "uint8":   []uint8{+0, -0, 1, 127,                  128,                  129,                 254,                  255},
   // TODO: set specific middle values
   "uint16": []uint16{+0, -0, 1, 127,                  128,                  129,                 254,                65535},
   "uint32": []uint32{+0, -0, 1, 127,                  128,                  129,                 254,           4294967295},
   "uint64": []uint64{+0, -0, 1, 127,                  128,                  129,                 254, 18446744073709551615},
   "int8":     []int8{+0, -0, -1,  1,                 -127,                 -128,                 126,                  127},
   "int16":   []int16{+0, -0, -1,  1,               -32767,               -32768,               32766,                32767},
   "int32":   []int32{+0, -0, -1,  1,          -2147483647,          -2147483648,          2147483646,           2147483647},
   "int64":   []int64{+0, -0, -1,  1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},
   // TODO: set bounds properly
   "float32": []float32{+0, -0, -1, 1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},
   "float64": []float64{+0, -0, -1, 1, -9223372036854775807, -9223372036854775808, 9223372036854775806,  9223372036854775807},
}

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

