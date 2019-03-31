// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark_fatih

package benchmarks_test

import (
	"io/ioutil"
	"math/rand"
	"testing"

	fatihcolor "github.com/fatih/color"
	"github.com/zchee/color/benchmarks"
)

const length = int64(1024)

func BenchmarkNewPrint(b *testing.B) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	benchmarkNewPrint(b, fatihcolor.New(fatihcolor.FgGreen), length)
}

func BenchmarkColorPrint(b *testing.B) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	var fns = printFuncs{
		fatihcolor.Black,
		fatihcolor.Red,
		fatihcolor.Green,
		fatihcolor.Yellow,
		fatihcolor.Blue,
		fatihcolor.Magenta,
		fatihcolor.Cyan,
		fatihcolor.White,
	}

	benchmarkColorPrint(b, fns, length)
}

func BenchmarkColorString(b *testing.B) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	var fns = stringFuncs{
		fatihcolor.BlackString,
		fatihcolor.RedString,
		fatihcolor.GreenString,
		fatihcolor.YellowString,
		fatihcolor.BlueString,
		fatihcolor.MagentaString,
		fatihcolor.CyanString,
		fatihcolor.WhiteString,
	}

	benchmarkColorString(b, fns, length)
}

func benchmark_getCacheColor(b *testing.B, i int) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := fatihcolor.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			_ = benchmarks.GetCachedColor(p)
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
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := fatihcolor.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			benchmarks.ColorPrint(format, p, buf)
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

func benchmark_colorString(b *testing.B, i int) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := fatihcolor.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			_ = benchmarks.ColorString(format, p, buf)
		}
	})
}

func BenchmarkColorStringFg(b *testing.B) {
	benchmark_colorString(b, 30)
}

func BenchmarkColorStringFgHi(b *testing.B) {
	benchmark_colorString(b, 90)
}

func BenchmarColorStringBg(b *testing.B) {
	benchmark_colorString(b, 40)
}

func BenchmarkColorStringBgHi(b *testing.B) {
	benchmark_colorString(b, 100)
}
