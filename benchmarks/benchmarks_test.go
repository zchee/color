// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks_test

import (
	crand "crypto/rand"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/zchee/color"
	"github.com/zchee/color/benchmarks"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())

	os.Exit(m.Run())
}

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

const numPrintFunc = 8

type printFuncs [numPrintFunc]func(format string, a ...interface{})

func benchmarkColorPrint(b *testing.B, fn printFuncs, length int64) {
	const format = "buf: %x"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := rand.Intn(numPrintFunc)
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

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		n := rand.Intn(numstringFunc)
		for pb.Next() {
			_ = fn[n](format, buf)
		}
	})
}

func benchmark_getCacheColor(b *testing.B, i int) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := color.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			_ = benchmarks.GetCacheColor(p)
		}
	})
}

func BenchmarkGetCacheColorFg(b *testing.B) {
	benchmark_getCacheColor(b, 30)
}

func BenchmarkGetCacheColorFgHi(b *testing.B) {
	benchmark_getCacheColor(b, 90)
}

func BenchmarkGetCacheColorBg(b *testing.B) {
	benchmark_getCacheColor(b, 40)
}

func BenchmarkGetCacheColorBgHi(b *testing.B) {
	benchmark_getCacheColor(b, 100)
}

func benchmark_colorPrint(b *testing.B, i int) {
	color.Output = ioutil.Discard
	color.NoColor = false
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := color.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			benchmarks.ColorPrint(p, format, buf)
		}
	})
}

func BenchmarkColorPrintFg(b *testing.B) {
	benchmark_colorPrint(b, 30)
}

func BenchmarkColorPrintFgHi(b *testing.B) {
	benchmark_colorPrint(b, 90)
}

func BenchmarColorPrintBg(b *testing.B) {
	benchmark_colorPrint(b, 40)
}

func BenchmarkColorPrintBgHi(b *testing.B) {
	benchmark_colorPrint(b, 100)
}
