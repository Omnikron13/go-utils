package maths

import (
	"math"
   "math/rand"
	"testing"
)


func Test_Abs(t *testing.T) {
   if got, want := Abs(uint8(8)), uint8(8); got != want {
      t.Errorf("Abs(8) = %v, want %v", got, want)
   }
   if got, want := Abs(uint16(16)), uint16(16); got != want {
      t.Errorf("Abs(16) = %v, want %v", got, want)
   }
   if got, want := Abs(uint32(32)), uint32(32); got != want {
      t.Errorf("Abs(32) = %v, want %v", got, want)
   }
   if got, want := Abs(uint64(64)), uint64(64); got != want {
      t.Errorf("Abs(64) = %v, want %v", got, want)
   }
   if got, want := Abs(uint(0)), uint(0); got != want {
      t.Errorf("Abs(0) = %v, want %v", got, want)
   }
   if got, want := Abs(int8(8)), int8(8); got != want {
      t.Errorf("Abs(8) = %v, want %v", got, want)
   }
   if got, want := Abs(int16(16)), int16(16); got != want {
      t.Errorf("Abs(16) = %v, want %v", got, want)
   }
   if got, want := Abs(int32(32)), int32(32); got != want {
      t.Errorf("Abs(32) = %v, want %v", got, want)
   }
   if got, want := Abs(int64(64)), int64(64); got != want {
      t.Errorf("Abs(64) = %v, want %v", got, want)
   }
   if got, want := Abs(int(0)), int(0); got != want {
      t.Errorf("Abs(0) = %v, want %v", got, want)
   }
   if got, want := Abs(int8(-8)), int8(8); got != want {
      t.Errorf("Abs(8) = %v, want %v", got, want)
   }
   if got, want := Abs(int16(-16)), int16(16); got != want {
      t.Errorf("Abs(16) = %v, want %v", got, want)
   }
   if got, want := Abs(int32(-32)), int32(32); got != want {
      t.Errorf("Abs(32) = %v, want %v", got, want)
   }
   if got, want := Abs(int64(-64)), int64(64); got != want {
      t.Errorf("Abs(64) = %v, want %v", got, want)
   }
   if got, want := Abs(int(-0)), int(0); got != want {
      t.Errorf("Abs(-0) = %v, want %v", got, want)
   }

   if got, want := Abs(float32(3.2)), float32(3.2); got != want {
      t.Errorf("Abs(3.2) = %v, want %v", got, want)
   }
   if got, want := Abs(float32(-3.2)), float32(3.2); got != want {
      t.Errorf("Abs(-3.2) = %v, want %v", got, want)
   }
   if got, want := Abs(float64(6.4)), float64(6.4); got != want {
      t.Errorf("Abs(6.4) = %v, want %v", got, want)
   }
   if got, want := Abs(float64(-6.4)), float64(6.4); got != want {
      t.Errorf("Abs(-6.4) = %v, want %v", got, want)
   }

   if got := Abs(math.NaN()); !math.IsNaN(got) {
      t.Errorf("Abs(NaN) = %v, want NaN", got)
   }

   if got := Abs(math.Inf(1)); !math.IsInf(got, 1) {
      t.Errorf("Abs(+Inf) = %v, want +Inf", got)
   }
   if got := Abs(math.Inf(-1)); math.IsInf(got, -1) {
      t.Errorf("Abs(-Inf) = %v, want +Inf", got)
   }
}


// Bench_Abs is a benchmark for Abs()
// It tests against the tedious casting bach and forth between float64 and the type that math.Abs() requires,
// and against the current (naive?) implementation of maths.Abs(), and potentially alternatives.
func Benchmark_Abs(b *testing.B) {
   tests8 := []int8{}
   for i := 0; i < 8; i++ {
      tests8 = append(tests8, -(1<<i), 1<<i - 1, int8(rand.Intn(256) - 128))
   }
   tests16 := []int16{}
   for i := 0; i < 16; i++ {
      tests16 = append(tests16, -(1<<i), 1<<i - 1, int16(rand.Intn(65536) - 32768))
   }
   tests32 := []int32{}
   for i := 0; i < 32; i++ {
      tests32 = append(tests32, -(1<<i), 1<<i - 1, int32(rand.Uint32() - 2147483648))
   }
   tests64 := []int64{}
   for i := 0; i < 64; i++ {
      tests64 = append(tests64, -(1<<i), 1<<i - 1, int64(rand.Uint64() - 9223372036854775808))
   }

   b.Run("math.Abs", func(b *testing.B) {
      b.Run("int8", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests8 {
               a := int8(math.Abs(float64(tests8[x])))
               _ = a
            }
         }
      })
      b.Run("int16", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests16 {
               a := int16(math.Abs(float64(tests16[x])))
               _ = a
            }
         }
      })
      b.Run("int32", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests32 {
               a := int32(math.Abs(float64(tests32[x])))
               _ = a
            }
         }
      })
      b.Run("int64", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests64 {
               a := int64(math.Abs(float64(tests64[x])))
               _ = a
            }
         }
      })
   })

   b.Run("maths.Abs", func(b *testing.B) {
      b.Run("int8", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests8 {
               a := Abs(tests8[x])
               _ = a
            }
         }
      })
      b.Run("int16", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests16 {
               a := Abs(tests16[x])
               _ = a
            }
         }
      })
      b.Run("int32", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests32 {
               a := Abs(tests32[x])
               _ = a
            }
         }
      })
      b.Run("int64", func(b *testing.B) {
         for i := 0; i < b.N; i++ {
            for x := range tests64 {
               a := Abs(tests64[x])
               _ = a
            }
         }
      })
   })
}

