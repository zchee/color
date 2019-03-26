// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks_test

import (
	crand "crypto/rand"
	"testing"
)

func genRandomBytes(tb testing.TB, length int64) (b []byte) {
	tb.Helper()

	b = make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		tb.Fatal(err)
	}

	return b
}

type printFunc func(...interface{}) (int, error)

func benchmarkNewPrint(b *testing.B, fn printFunc, length int64) {
	buf := genRandomBytes(b, length)
	b.SetBytes(length)
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn(buf)
		}
	})
}

type colorPrintFunc func(format string, a ...interface{})

func benchmarkColorPrint(b *testing.B, fn colorPrintFunc, length int64) {
	const format = "buf: %x"
	buf := genRandomBytes(b, length)
	b.SetBytes(length)
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn(format, buf)
		}
	})
}

type colorStringFunc func(format string, a ...interface{}) string

func benchmarkColorString(b *testing.B, fn colorStringFunc, length int64) {
	const format = "buf: %x"
	buf := genRandomBytes(b, length)
	b.SetBytes(length)
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fn(format, buf)
		}
	})
}
