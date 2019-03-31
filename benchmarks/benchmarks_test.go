// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks_test

import (
	crand "crypto/rand"
	"io"
	"math/rand"
	"testing"
	"time"
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

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn.Print(buf)
		}
	})
}

var randSrc = rand.NewSource(time.Now().UTC().UnixNano())

const numPrintFunc = 8

type printFuncs [numPrintFunc]func(format string, a ...interface{})

func benchmarkColorPrint(b *testing.B, fn printFuncs, length int64) {
	const format = "buf: %x"
	buf := genRandomBytes(b, length)
	r := rand.New(randSrc)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := r.Intn(numPrintFunc)
		for pb.Next() {
			fn[n](format, buf)
		}
	})
}

const numstringFunc = 8

type stringFuncs [numstringFunc]func(format string, a ...interface{}) string

func benchmarkColorString(b *testing.B, fn stringFuncs, length int64) {
	const format = "buf: %x"
	buf := genRandomBytes(b, length)
	r := rand.New(randSrc)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := r.Intn(numPrintFunc)
		for pb.Next() {
			_ = fn[n](format, buf)
		}
	})
}
