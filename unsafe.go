// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

import (
	"unsafe"
)

// sliceHeader is the same as reflect.SliceHeader but with unsafe.Pointers to
// guarantee they don't get collected by the GC.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// stringHeader is the same as reflect.StringHeader but with unsafe.Pointers to
// guarantee they don't get collected by the GC.
type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// unsafeByteSlice returns a byte array that points to the given string without a heap allocation.
// The string must be preserved until the byte array is disposed.
func unsafeByteSlice(s string) (p []byte) {
	if s == "" {
		return nil
	}

	sh := (*stringHeader)(unsafe.Pointer(&s))
	p = *(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}))

	return
}
