// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

// +build ignore

package main

import (
	"os"
	"text/template"
)

func main() {
	intTmpl := template.Must(template.New("").Parse(intTemplate))
	floatTmpl := template.Must(template.New("").Parse(floatTemplate))

	for _, typ := range []struct {
		File, Name, Type, Atomic string
		Unsigned                 bool
	}{
		{"int32.go", "Int32", "int32", "Int32", false},
		{"int64.go", "Int64", "int64", "Int64", false},
		{"uint32.go", "Uint32", "uint32", "Uint32", true},
		{"uint64.go", "Uint64", "uint64", "Uint64", true},
	} {
		f, err := os.Create(typ.File)
		if err != nil {
			panic(err)
		}

		if err = intTmpl.Execute(f, typ); err != nil {
			panic(err)
		}

		f.Close()
	}

	for _, typ := range []struct{ File, Name, Type, Atomic, AtomicType, MathName string }{
		{"float32.go", "Float32", "float32", "Uint32", "uint32", "Float32"},
		{"float64.go", "Float64", "float64", "Uint64", "uint64", "Float64"},
	} {
		f, err := os.Create(typ.File)
		if err != nil {
			panic(err)
		}

		if err = floatTmpl.Execute(f, typ); err != nil {
			panic(err)
		}

		f.Close()
	}
}

const intTemplate = `// Code generated by go run generate.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import (
	"sync/atomic"

	"github.com/golang/sync/syncmap"
)

// {{.Name}} provides a map of atomic counters of type {{.Type}}.
type {{.Name}} struct {
	m syncmap.Map // map[interface{}]*{{.Type}}
}

// UnsafeLoad returns a pointer to the counter key.
//
// It is only safe to access the return value with
// methods from the sync/atomic package. It must
// not be manually dereferenced.
func (c *{{.Name}}) UnsafeLoad(key interface{}) *{{.Type}} {
	v, _ := c.m.LoadOrStore(key, new({{.Type}}))
	return v.(*{{.Type}})
}

// Load returns the value of the counter key.
func (c *{{.Name}}) Load(key interface{}) (val {{.Type}}) {
	return atomic.Load{{.Atomic}}(c.UnsafeLoad(key))
}

// Store sets the value of the counter key.
func (c *{{.Name}}) Store(key interface{}, val {{.Type}}) {
	atomic.Store{{.Atomic}}(c.UnsafeLoad(key), val)
}

// Swap sets the value of the counter key and returns the
// old value.
func (c *{{.Name}}) Swap(key interface{}, new {{.Type}}) (old {{.Type}}) {
	return atomic.Swap{{.Atomic}}(c.UnsafeLoad(key), new)
}

// CompareAndSwap sets the value of the counter key to new
// but only if it currently has the value old.
func (c *{{.Name}}) CompareAndSwap(key interface{}, old, new {{.Type}}) (swapped bool) {
	return atomic.CompareAndSwap{{.Atomic}}(c.UnsafeLoad(key), old, new)
}

// Add adds delta to the counter key.
func (c *{{.Name}}) Add(key interface{}, delta {{.Type}}) (new {{.Type}}) {
	return atomic.Add{{.Atomic}}(c.UnsafeLoad(key), delta)
}

// Increment is a wrapper for Add(key, 1).
func (c *{{.Name}}) Increment(key interface{}) (new {{.Type}}) {
	return c.Add(key, 1)
}

{{if .Unsigned -}}

// Subtract subtracts delta from the counter key.
func (c *{{.Name}}) Subtract(key interface{}, delta {{.Type}}) (new {{.Type}}) {
	return atomic.Add{{.Atomic}}(c.UnsafeLoad(key), ^(delta - 1))
}

// Decrement is a wrapper for Subtract(key, 1).
func (c *{{.Name}}) Decrement(key interface{}) (new {{.Type}}) {
	return c.Subtract(key, 1)
}

{{- else -}}

// Decrement is a wrapper for Add(key, -1).
func (c *{{.Name}}) Decrement(key interface{}) (new {{.Type}}) {
	return c.Add(key, -1)
}

{{- end}}

// Reset is a wrapper for Swap(key, 0).
func (c *{{.Name}}) Reset(key interface{}) (old {{.Type}}) {
	return c.Swap(key, 0)
}

// Delete removes the counter key from the map.
func (c *{{.Name}}) Delete(key interface{}) {
	c.m.Delete(key)
}

// Keys returns the list of all counters.
func (c *{{.Name}}) Keys() []interface{} {
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
func (c *{{.Name}}) UnsafeRange(f func(key interface{}, val *{{.Type}}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, val.(*{{.Type}}))
	})
}

// RangeKeys calls f with the key of each counter.
func (c *{{.Name}}) RangeKeys(f func(key interface{}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key)
	})
}

// RangeLoad calls f with the value of each counter.
func (c *{{.Name}}) RangeLoad(f func(key interface{}, val {{.Type}}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.Load{{.Atomic}}(val.(*{{.Type}})))
	})
}

// RangeStore sets each counter to the return value of f.
func (c *{{.Name}}) RangeStore(f func(key interface{}) (val {{.Type}}, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		v, ok := f(key)
		atomic.Store{{.Atomic}}(val.(*{{.Type}}), v)
		return ok
	})
}

// RangeAdd adds the return value of f to each counter.
func (c *{{.Name}}) RangeAdd(f func(key interface{}) (delta {{.Type}}, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		atomic.Add{{.Atomic}}(val.(*{{.Type}}), delta)
		return ok
	})
}

{{- if .Unsigned}}

// RangeSubtract subtracts the return value of f from
// each counter.
func (c *{{.Name}}) RangeSubtract(f func(key interface{}) (delta {{.Type}}, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		atomic.Add{{.Atomic}}(val.(*{{.Type}}), ^(delta - 1))
		return ok
	})
}

{{- end}}

// RangeReset resets each counter and calls f with the
// old value.
func (c *{{.Name}}) RangeReset(f func(key interface{}, old {{.Type}}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.Swap{{.Atomic}}(val.(*{{.Type}}), 0))
	})
}
`

const floatTemplate = `// Code generated by go run generate.go.

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

// {{.Name}} provides a map of atomic counters of type {{.Type}}.
type {{.Name}} struct {
	m syncmap.Map // map[interface{}]*{{.AtomicType}}
}

func (c *{{.Name}}) unsafeLoad(key interface{}) *{{.AtomicType}} {
	v, _ := c.m.LoadOrStore(key, new({{.AtomicType}}))
	return v.(*{{.AtomicType}})
}

// Load returns the value of the counter key.
func (c *{{.Name}}) Load(key interface{}) (val {{.Type}}) {
	return math.{{.MathName}}frombits(atomic.Load{{.Atomic}}(c.unsafeLoad(key)))
}

// Store sets the value of the counter key.
func (c *{{.Name}}) Store(key interface{}, val {{.Type}}) {
	atomic.Store{{.Atomic}}(c.unsafeLoad(key), math.{{.MathName}}bits(val))
}

// Swap sets the value of the counter key and returns the
// old value.
func (c *{{.Name}}) Swap(key interface{}, new {{.Type}}) (old {{.Type}}) {
	return math.{{.MathName}}frombits(atomic.Swap{{.Atomic}}(c.unsafeLoad(key), math.{{.MathName}}bits(new)))
}

// CompareAndSwap sets the value of the counter key to new
// but only if it currently has the value old.
func (c *{{.Name}}) CompareAndSwap(key interface{}, old, new {{.Type}}) (swapped bool) {
	return atomic.CompareAndSwap{{.Atomic}}(c.unsafeLoad(key), math.{{.MathName}}bits(old), math.{{.MathName}}bits(new))
}

// Reset is a wrapper for Swap(key, 0).
func (c *{{.Name}}) Reset(key interface{}) (old {{.Type}}) {
	return c.Swap(key, 0)
}

// Delete removes the counter key from the map.
func (c *{{.Name}}) Delete(key interface{}) {
	c.m.Delete(key)
}

// Keys returns the list of all counters.
func (c *{{.Name}}) Keys() []interface{} {
	var keys []interface{}
	c.m.Range(func(key, val interface{}) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// RangeKeys calls f with the key of each counter.
func (c *{{.Name}}) RangeKeys(f func(key interface{}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key)
	})
}

// RangeLoad calls f with the value of each counter.
func (c *{{.Name}}) RangeLoad(f func(key interface{}, val {{.Type}}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, math.{{.MathName}}frombits(atomic.Load{{.Atomic}}(val.(*{{.AtomicType}}))))
	})
}

// RangeStore sets each counter to the return value of f.
func (c *{{.Name}}) RangeStore(f func(key interface{}) (val {{.Type}}, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		v, ok := f(key)
		atomic.Store{{.Atomic}}(val.(*{{.AtomicType}}), math.{{.MathName}}bits(v))
		return ok
	})
}

// RangeReset resets each counter and calls f with the
// old value.
func (c *{{.Name}}) RangeReset(f func(key interface{}, old {{.Type}}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, math.{{.MathName}}frombits(atomic.Swap{{.Atomic}}(val.(*{{.AtomicType}}), 0)))
	})
}
`
