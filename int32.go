// Code generated by go run generate-int.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"strconv"
	"sync/atomic"
)

// Int32 provides an atomic int32.
type Int32 struct {
	noCopy noCopy
	val    int32
}

// NewInt32 returns an atomic int32 with a given value.
func NewInt32(val int32) *Int32 {
	return &Int32{val: val}
}

// UnsafeRaw returns a pointer to the int32.
//
// It is only safe to access the pointer with methods from the
// sync/atomic package. It must not be manually dereferenced.
func (v *Int32) UnsafeRaw() *int32 {
	return &v.val
}

// Load returns the value of the int32.
func (v *Int32) Load() (val int32) {
	return atomic.LoadInt32(&v.val)
}

// Store sets the value of the int32.
func (v *Int32) Store(val int32) {
	atomic.StoreInt32(&v.val, val)
}

// Swap sets the value of the int32 and returns the old value.
func (v *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&v.val, new)
}

// CompareAndSwap sets the value of the int32 to new but only
// if it currently has the value old. It returns true if the swap
// succeeded.
func (v *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.val, old, new)
}

// Add adds delta to the int32.
func (v *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.val, delta)
}

// Increment is a wrapper for Add(1).
func (v *Int32) Increment() (new int32) {
	return v.Add(1)
}

// Subtract is a wrapper for Add(-delta)
func (v *Int32) Subtract(delta int32) (new int32) {
	return v.Add(-delta)
}

// Decrement is a wrapper for Add(-1).
func (v *Int32) Decrement() (new int32) {
	return v.Add(-1)
}

// Reset is a wrapper for Swap(0).
func (v *Int32) Reset() (old int32) {
	return v.Swap(0)
}

// String implements fmt.Stringer.
func (v *Int32) String() string {
	return strconv.FormatInt(int64(v.Load()), 10)
}
