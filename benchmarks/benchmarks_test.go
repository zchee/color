// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks_test

import (
	crand "crypto/rand"
	"io"
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

type newPrintFunc interface {
	Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Print(a ...interface{}) (n int, err error)
	Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	Sprint(a ...interface{}) string
	Sprintln(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
}

func benchmarkNewPrint(b *testing.B, fn newPrintFunc, length int64) {
	buf := genRandomBytes(b, length)
	b.SetBytes(length)
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn.Print(buf)
		}
	})
}

type printFunc func(format string, a ...interface{})

func benchmarkColorPrint(b *testing.B, fn printFunc, length int64) {
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

type stringFunc func(format string, a ...interface{}) string

func benchmarkColorString(b *testing.B, fn stringFunc, length int64) {
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
