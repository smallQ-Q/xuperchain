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

// Int64 provides an atomic int64.
type Int64 struct {
	noCopy noCopy
	val    int64
}

// NewInt64 returns an atomic int64 with a given value.
func NewInt64(val int64) *Int64 {
	return &Int64{val: val}
}

// Raw returns a pointer to the int64.
//
// It is only safe to access the pointer with methods from the
// sync/atomic package. Use caution if manually dereferencing.
func (v *Int64) Raw() *int64 {
	return &v.val
}

// Load returns the value of the int64.
func (v *Int64) Load() (val int64) {
	return atomic.LoadInt64(&v.val)
}

// Store sets the value of the int64.
func (v *Int64) Store(val int64) {
	atomic.StoreInt64(&v.val, val)
}

// Swap sets the value of the int64 and returns the old value.
func (v *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(&v.val, new)
}

// CompareAndSwap sets the value of the int64 to new but only
// if it currently has the value old. It returns true if the swap
// succeeded.
func (v *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.val, old, new)
}

// Add adds delta to the int64.
func (v *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&v.val, delta)
}

// Increment is a wrapper for Add(1).
func (v *Int64) Increment() (new int64) {
	return v.Add(1)
}

// Subtract is a wrapper for Add(-delta)
func (v *Int64) Subtract(delta int64) (new int64) {
	return v.Add(-delta)
}

// Decrement is a wrapper for Add(-1).
func (v *Int64) Decrement() (new int64) {
	return v.Add(-1)
}

// Reset is a wrapper for Swap(0).
func (v *Int64) Reset() (old int64) {
	return v.Swap(0)
}

// String implements fmt.Stringer.
func (v *Int64) String() string {
	return strconv.FormatInt(int64(v.Load()), 10)
}
