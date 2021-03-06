// Code generated by go run generate-tests.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestUint32Default(t *testing.T) {
	var v Uint32

	if v.Load() != 0 {
		t.Fatal("invalid default value")
	}
}

func TestNewUint32(t *testing.T) {
	if NewUint32(0) == nil {
		t.Fatal("NewUint32 returned nil")
	}
}

func TestUint32Raw(t *testing.T) {
	var v Uint32

	if v.Raw() == nil {
		t.Fatal("Raw returned nil")
	}

	if err := quick.Check(func(v uint32) bool {
		return *NewUint32(v).Raw() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Load(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		return NewUint32(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Store(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		var a Uint32
		a.Store(v)
		return a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Swap(t *testing.T) {
	if err := quick.Check(func(old, new uint32) bool {
		a := NewUint32(old)
		return a.Swap(new) == old && a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new uint32) bool {
		a := NewUint32(old)
		return !a.CompareAndSwap(-old, new) &&
			a.Load() == old &&
			a.CompareAndSwap(old, new) &&
			a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Add(t *testing.T) {
	if err := quick.Check(func(v, delta uint32) bool {
		a := NewUint32(v)
		v += delta
		return a.Add(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Increment(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		a := NewUint32(v)
		v++
		return a.Increment() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta uint32) bool {
		a := NewUint32(v)
		v -= delta
		return a.Subtract(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Decrement(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		a := NewUint32(v)
		v--
		return a.Decrement() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32Reset(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		a := NewUint32(v)
		return a.Reset() == v && a.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestUint32String(t *testing.T) {
	if err := quick.Check(func(v uint32) bool {
		return NewUint32(v).String() == fmt.Sprint(v)
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNewUint32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewUint32(0)
	}
}

func BenchmarkUint32Load(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func BenchmarkUint32Store(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Store(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkUint32Swap(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Swap(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkUint32CompareAndSwap(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(0, 0)
	}
}

func BenchmarkUint32Add(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Add(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkUint32Increment(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Increment()
	}
}

func BenchmarkUint32Subtract(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Subtract(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkUint32Decrement(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Decrement()
	}
}

func BenchmarkUint32Reset(b *testing.B) {
	var v Uint32

	for n := 0; n < b.N; n++ {
		v.Reset()
	}
}
