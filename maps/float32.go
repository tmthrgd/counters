// Code generated by go run generate-int.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package maps

import (
	"sync"

	"github.com/tmthrgd/atomics"
)

// Float32 provides a map of atomic float32s.
type Float32 struct {
	m sync.Map // map[interface{}]*atomics.Float32
}

// Retrieve returns the atomics.Float32 associated with
// the given key or nil if it does not exist in the map.
func (m *Float32) Retrieve(key interface{}) *atomics.Float32 {
	v, ok := m.m.Load(key)
	if !ok {
		return nil
	}

	return v.(*atomics.Float32)
}

// Insert inserts the atomics.Float32 into the map for
// the given key.
func (m *Float32) Insert(key interface{}, val *atomics.Float32) {
	m.m.Store(key, val)
}

// Value returns the atomics.Float32 associated with the
// given key or atomically inserts a new atomics.Float32
// into the map if an entry did not exist in the map
// for the given key.
func (m *Float32) Value(key interface{}) *atomics.Float32 {
	v, ok := m.m.Load(key)
	if !ok {
		v, _ = m.m.LoadOrStore(key, new(atomics.Float32))
	}

	return v.(*atomics.Float32)
}

// Delete removes an atomics.Float32 from the map.
func (m *Float32) Delete(key interface{}) {
	m.m.Delete(key)
}

// Range calls f for each entry in the map. If f
// returns false Range stops iterating over the map.
func (m *Float32) Range(f func(key interface{}, val *atomics.Float32) bool) {
	m.m.Range(func(key, val interface{}) bool {
		return f(key, val.(*atomics.Float32))
	})
}

// Load is a wrapper for Value(key).Load().
func (m *Float32) Load(key interface{}) (val float32) {
	return m.Value(key).Load()
}

// Store is a wrapper for Value(key).Store(val).
func (m *Float32) Store(key interface{}, val float32) {
	m.Value(key).Store(val)
}

// Swap is a wrapper for Value(key).Swap(new).
func (m *Float32) Swap(key interface{}, new float32) (old float32) {
	return m.Value(key).Swap(new)
}

// CompareAndSwap is a wrapper for
// Value(key).CompareAndSwap(old, new).
func (m *Float32) CompareAndSwap(key interface{}, old, new float32) (swapped bool) {
	return m.Value(key).CompareAndSwap(old, new)
}

// Add is a wrapper for Value(key).Add(delta).
func (m *Float32) Add(key interface{}, delta float32) (new float32) {
	return m.Value(key).Add(delta)
}

// Increment is a wrapper for Value(key).Increment().
func (m *Float32) Increment(key interface{}) (new float32) {
	return m.Value(key).Increment()
}

// Subtract is a wrapper for Value(key).Subtract(delta).
func (m *Float32) Subtract(key interface{}, delta float32) (new float32) {
	return m.Value(key).Subtract(delta)
}

// Decrement is a wrapper for Value(key).Decrement().
func (m *Float32) Decrement(key interface{}) (new float32) {
	return m.Value(key).Decrement()
}

// Reset is a wrapper for Value(key).Reset().
func (m *Float32) Reset(key interface{}) (old float32) {
	return m.Value(key).Reset()
}
