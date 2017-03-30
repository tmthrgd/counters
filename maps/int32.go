// Code generated by go run generate-int.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package maps

import (
	"github.com/tmthrgd/atomics"
	"golang.org/x/sync/syncmap"
)

// Int32 provides a map of atomic int32s.
type Int32 struct {
	m syncmap.Map // map[interface{}]*atomics.Int32
}

// Retrieve returns the atomics.Int32 associated with
// the given key or nil if it does not exist in the map.
func (m *Int32) Retrieve(key interface{}) *atomics.Int32 {
	v, ok := m.m.Load(key)
	if !ok {
		return nil
	}

	return v.(*atomics.Int32)
}

// Insert inserts the atomics.Int32 into the map for
// the given key.
func (m *Int32) Insert(key interface{}, val *atomics.Int32) {
	m.m.Store(key, val)
}

// Value returns the atomics.Int32 associated with the
// given key or atomically inserts a new atomics.Int32
// into the map if an entry did not exist in the map
// for the given key.
func (m *Int32) Value(key interface{}) *atomics.Int32 {
	v, ok := m.m.Load(key)
	if !ok {
		v, _ = m.m.LoadOrStore(key, new(atomics.Int32))
	}

	return v.(*atomics.Int32)
}

// Delete removes an atomics.Int32 from the map.
func (m *Int32) Delete(key interface{}) {
	m.m.Delete(key)
}

// Range calls f for each entry in the map. If f
// returns false Range stops iterating over the map.
func (m *Int32) Range(f func(key interface{}, val *atomics.Int32) bool) {
	m.m.Range(func(key, val interface{}) bool {
		return f(key, val.(*atomics.Int32))
	})
}

// Load is a wrapper for Value(key).Load().
func (m *Int32) Load(key interface{}) (val int32) {
	return m.Value(key).Load()
}

// Store is a wrapper for Value(key).Store(val).
func (m *Int32) Store(key interface{}, val int32) {
	m.Value(key).Store(val)
}

// Swap is a wrapper for Value(key).Swap(new).
func (m *Int32) Swap(key interface{}, new int32) (old int32) {
	return m.Value(key).Swap(new)
}

// CompareAndSwap is a wrapper for
// Value(key).CompareAndSwap(old, new).
func (m *Int32) CompareAndSwap(key interface{}, old, new int32) (swapped bool) {
	return m.Value(key).CompareAndSwap(old, new)
}

// Add is a wrapper for Value(key).Add(delta).
func (m *Int32) Add(key interface{}, delta int32) (new int32) {
	return m.Value(key).Add(delta)
}

// Increment is a wrapper for Value(key).Increment().
func (m *Int32) Increment(key interface{}) (new int32) {
	return m.Value(key).Increment()
}

// Subtract is a wrapper for Value(key).Subtract(delta).
func (m *Int32) Subtract(key interface{}, delta int32) (new int32) {
	return m.Value(key).Subtract(delta)
}

// Decrement is a wrapper for Value(key).Decrement().
func (m *Int32) Decrement(key interface{}) (new int32) {
	return m.Value(key).Decrement()
}

// Reset is a wrapper for Value(key).Reset().
func (m *Int32) Reset(key interface{}) (old int32) {
	return m.Value(key).Reset()
}