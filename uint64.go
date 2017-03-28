// Code generated by go run generate.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import (
	"sync/atomic"

	"github.com/golang/sync/syncmap"
)

// Uint64 provides a map of atomic counters of type uint64.
type Uint64 struct {
	m syncmap.Map // map[interface{}]*uint64
}

// UnsafeLoad returns a pointer to the counter key.
//
// It is only safe to access the return value with
// methods from the sync/atomic package. It must
// not be manually dereferenced.
func (c *Uint64) UnsafeLoad(key interface{}) *uint64 {
	v, _ := c.m.LoadOrStore(key, new(uint64))
	return v.(*uint64)
}

// Load returns the value of the counter key.
func (c *Uint64) Load(key interface{}) (val uint64) {
	return atomic.LoadUint64(c.UnsafeLoad(key))
}

// Store sets the value of the counter key.
func (c *Uint64) Store(key interface{}, val uint64) {
	atomic.StoreUint64(c.UnsafeLoad(key), val)
}

// Swap sets the value of the counter key and returns the
// old value.
func (c *Uint64) Swap(key interface{}, new uint64) (old uint64) {
	return atomic.SwapUint64(c.UnsafeLoad(key), new)
}

// CompareAndSwap sets the value of the counter key to new
// but only if it currently has the value old.
func (c *Uint64) CompareAndSwap(key interface{}, old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(c.UnsafeLoad(key), old, new)
}

// Add adds delta to the counter key.
func (c *Uint64) Add(key interface{}, delta uint64) (new uint64) {
	return atomic.AddUint64(c.UnsafeLoad(key), delta)
}

// Increment is a wrapper for Add(key, 1).
func (c *Uint64) Increment(key interface{}) (new uint64) {
	return c.Add(key, 1)
}

// Subtract subtracts delta from the counter key.
func (c *Uint64) Subtract(key interface{}, delta uint64) (new uint64) {
	return atomic.AddUint64(c.UnsafeLoad(key), ^(delta - 1))
}

// Decrement is a wrapper for Subtract(key, 1).
func (c *Uint64) Decrement(key interface{}) (new uint64) {
	return c.Subtract(key, 1)
}

// Reset is a wrapper for Swap(key, 0).
func (c *Uint64) Reset(key interface{}) (old uint64) {
	return c.Swap(key, 0)
}

// Delete removes the counter key from the map.
func (c *Uint64) Delete(key interface{}) {
	c.m.Delete(key)
}

// Keys returns the list of all counters.
func (c *Uint64) Keys() []interface{} {
	var keys []interface{}
	c.m.Range(func(key, val interface{}) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// UnsafeRange calls f with a pointer to each
// counter.
//
// It is only safe to access val with methods from
// the sync/atomic package. It must not be manually
// dereferenced.
func (c *Uint64) UnsafeRange(f func(key interface{}, val *uint64) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, val.(*uint64))
	})
}

// RangeKeys calls f with the key of each counter.
func (c *Uint64) RangeKeys(f func(key interface{}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key)
	})
}

// RangeLoad calls f with the value of each counter.
func (c *Uint64) RangeLoad(f func(key interface{}, val uint64) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.LoadUint64(val.(*uint64)))
	})
}

// RangeStore sets each counter to the return value of f.
func (c *Uint64) RangeStore(f func(key interface{}) (val uint64, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		v, ok := f(key)
		atomic.StoreUint64(val.(*uint64), v)
		return ok
	})
}

// RangeAdd adds the return value of f to each counter.
func (c *Uint64) RangeAdd(f func(key interface{}) (delta uint64, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		atomic.AddUint64(val.(*uint64), delta)
		return ok
	})
}

// RangeSubtract subtracts the return value of f from
// each counter.
func (c *Uint64) RangeSubtract(f func(key interface{}) (delta uint64, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		atomic.AddUint64(val.(*uint64), ^(delta - 1))
		return ok
	})
}

// RangeReset resets each counter and calls f with the
// old value.
func (c *Uint64) RangeReset(f func(key interface{}, old uint64) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.SwapUint64(val.(*uint64), 0))
	})
}
