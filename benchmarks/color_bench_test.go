// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark

package benchmarks_test

import (
	"math/rand"
	"testing"

	"github.com/zchee/color"
	"github.com/zchee/color/benchmarks"
)

const length = int64(1024)

func BenchmarkNewPrint(b *testing.B) {
	benchmarkNewPrint(b, color.New(color.FgGreen), length)
}

func BenchmarkColorPrint(b *testing.B) {
	var fns = printFuncs{
		color.Black,
		color.Red,
		color.Green,
		color.Yellow,
		color.Blue,
		color.Magenta,
		color.Cyan,
		color.White,
	}

	benchmarkColorPrint(b, fns, length)
}

func BenchmarkColorString(b *testing.B) {
	var fns = stringFuncs{
		color.BlackString,
		color.RedString,
		color.GreenString,
		color.YellowString,
		color.BlueString,
		color.MagentaString,
		color.CyanString,
		color.WhiteString,
	}

	benchmarkColorString(b, fns, length)
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

func benchmark_colorString(b *testing.B, i int) {
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		p := color.Attribute(rand.Intn(7) + i)
		for pb.Next() {
			_ = benchmarks.ColorString(p, format, buf)
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
