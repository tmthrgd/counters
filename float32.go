// Code generated by go run generate.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import (
	"math"
	"sync/atomic"

	"github.com/golang/sync/syncmap"
)

// Float32 provides a map of atomic counters of type float32.
type Float32 struct {
	m syncmap.Map // map[interface{}]*uint32
}

func (c *Float32) unsafeLoad(key interface{}) *uint32 {
	v, _ := c.m.LoadOrStore(key, new(uint32))
	return v.(*uint32)
}

// Load returns the value of the counter key.
func (c *Float32) Load(key interface{}) (val float32) {
	return math.Float32frombits(atomic.LoadUint32(c.unsafeLoad(key)))
}

// Store sets the value of the counter key.
func (c *Float32) Store(key interface{}, val float32) {
	atomic.StoreUint32(c.unsafeLoad(key), math.Float32bits(val))
}

// Swap sets the value of the counter key and returns the
// old value.
func (c *Float32) Swap(key interface{}, new float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(c.unsafeLoad(key), math.Float32bits(new)))
}

// CompareAndSwap sets the value of the counter key to new
// but only if it currently has the value old.
func (c *Float32) CompareAndSwap(key interface{}, old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(c.unsafeLoad(key), math.Float32bits(old), math.Float32bits(new))
}

func addFloat32(ptr *uint32, delta float32) (new float32) {
	for {
		old := atomic.LoadUint32(ptr)
		new := math.Float32frombits(old) + delta

		if atomic.CompareAndSwapUint32(ptr, old, math.Float32bits(new)) {
			return new
		}
	}
}

// Add adds delta to the counter key.
func (c *Float32) Add(key interface{}, delta float32) (new float32) {
	return addFloat32(c.unsafeLoad(key), delta)
}

// Subtract is a wrapper for Add(key, -delta)
func (c *Float32) Subtract(key interface{}, delta float32) (new float32) {
	return c.Add(key, -delta)
}

// Reset is a wrapper for Swap(key, 0).
func (c *Float32) Reset(key interface{}) (old float32) {
	return c.Swap(key, 0)
}

// Delete removes the counter key from the map.
func (c *Float32) Delete(key interface{}) {
	c.m.Delete(key)
}

// Keys returns the list of all counters.
func (c *Float32) Keys() []interface{} {
	var keys []interface{}
	c.m.Range(func(key, val interface{}) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// RangeKeys calls f with the key of each counter.
func (c *Float32) RangeKeys(f func(key interface{}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key)
	})
}

// RangeLoad calls f with the value of each counter.
func (c *Float32) RangeLoad(f func(key interface{}, val float32) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, math.Float32frombits(atomic.LoadUint32(val.(*uint32))))
	})
}

// RangeStore sets each counter to the return value of f.
func (c *Float32) RangeStore(f func(key interface{}) (val float32, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		v, ok := f(key)
		atomic.StoreUint32(val.(*uint32), math.Float32bits(v))
		return ok
	})
}

// RangeSubtract subtracts the return value of f from
// each counter.
func (c *Float32) RangeSubtract(f func(key interface{}) (delta float32, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		addFloat32(val.(*uint32), -delta)
		return ok
	})
}

// RangeAdd adds the return value of f to each counter.
func (c *Float32) RangeAdd(f func(key interface{}) (delta float32, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		addFloat32(val.(*uint32), delta)
		return ok
	})
}

// RangeReset resets each counter and calls f with the
// old value.
func (c *Float32) RangeReset(f func(key interface{}, old float32) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, math.Float32frombits(atomic.SwapUint32(val.(*uint32), 0)))
	})
}
