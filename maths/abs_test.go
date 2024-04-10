package maths

import (
	"math"
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

