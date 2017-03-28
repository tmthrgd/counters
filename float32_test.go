// Code generated by go run generate-tests.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"testing"
	"testing/quick"
)

func TestFloat32Default(t *testing.T) {
	var c Float32
	if c.Load() != 0 {
		t.Fatal("invalid default value")
	}
}

func TestNewFloat32(t *testing.T) {
	if NewFloat32(0) == nil {
		t.Fatal("NewFloat32 returned nil")
	}
}

func TestFloat32UnsafeRaw(t *testing.T) {
	var c Float32
	if c.UnsafeRaw() == nil {
		t.Fatal("UnsafeRaw returned nil")
	}
}

func TestFloat32Load(t *testing.T) {
	if err := quick.Check(func(v float32) bool {
		return NewFloat32(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Store(t *testing.T) {
	if err := quick.Check(func(v float32) bool {
		var c Float32
		c.Store(v)
		return c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Swap(t *testing.T) {
	if err := quick.Check(func(old, new float32) bool {
		c := NewFloat32(old)
		return c.Swap(new) == old && c.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new float32) bool {
		c := NewFloat32(old)
		return !c.CompareAndSwap(-old, new) &&
			c.Load() == old &&
			c.CompareAndSwap(old, new) &&
			c.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Add(t *testing.T) {
	if err := quick.Check(func(v, delta float32) bool {
		c := NewFloat32(v)
		v += delta
		return c.Add(delta) == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Increment(t *testing.T) {
	if err := quick.Check(func(v float32) bool {
		c := NewFloat32(v)
		v++
		return c.Increment() == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta float32) bool {
		c := NewFloat32(v)
		v -= delta
		return c.Subtract(delta) == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Decrement(t *testing.T) {
	if err := quick.Check(func(v float32) bool {
		c := NewFloat32(v)
		v--
		return c.Decrement() == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat32Reset(t *testing.T) {
	if err := quick.Check(func(v float32) bool {
		c := NewFloat32(v)
		return c.Reset() == v && c.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNewFloat32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewFloat32(0)
	}
}

func BenchmarkFloat32Load(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func BenchmarkFloat32Store(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Store(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat32Swap(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Swap(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat32CompareAndSwap(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(0, 0)
	}
}

func BenchmarkFloat32Add(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Add(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat32Increment(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Increment()
	}
}

func BenchmarkFloat32Subtract(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Subtract(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat32Decrement(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Decrement()
	}
}

func BenchmarkFloat32Reset(b *testing.B) {
	var v Float32

	for n := 0; n < b.N; n++ {
		v.Reset()
	}
}
