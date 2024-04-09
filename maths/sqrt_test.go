package maths

import "testing"

func Test_Sqrt(t *testing.T) {
   // TODO: Pretty urgently need to work out a way to automaticqally run these tests over all the types.
   t.Run("uint8", func(t *testing.T) {
      if got, want := Sqrt(uint8(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("uint16", func(t *testing.T) {
      if got, want := Sqrt(uint16(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("uint32", func(t *testing.T) {
      if got, want := Sqrt(uint32(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("uint64", func(t *testing.T) {
      if got, want := Sqrt(uint64(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("uint", func(t *testing.T) {
      if got, want := Sqrt(uint(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("int8", func(t *testing.T) {
      if got, want := Sqrt(int8(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("int16", func(t *testing.T) {
      if got, want := Sqrt(int16(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("int32", func(t *testing.T) {
      if got, want := Sqrt(int32(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("int64", func(t *testing.T) {
      if got, want := Sqrt(int64(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("int", func(t *testing.T) {
      if got, want := Sqrt(int(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
   t.Run("float32", func(t *testing.T) {
      if got, want := Sqrt(float32(9)), 3.0; got != want {
         t.Errorf("Sqrt(%v) = %v, want %v", 9, got, want)
      }
   })
}

