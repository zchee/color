// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark

package benchmarks_test

import (
	crand "crypto/rand"
	"flag"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
	_ "unsafe" // required for go:linkname

	fatihcolor "github.com/fatih/color"
	color "github.com/zchee/color/v2"
)

const length = 1024 * 1024 // 1 MB

// benchFatih is for taking a switching between "zchee/color" and "fatih/color" benchmarking result.
var benchFatih = flag.Bool("fatih", false, "benchmark the fatih/color package")

func TestMain(m *testing.M) {
	// force coloring output and discard output to /dev/null.
	color.Output = ioutil.Discard
	color.NoColor = false
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	rand.Seed(time.Now().UTC().UnixNano())

	var status int
	status = m.Run()
	defer func() { os.Exit(status) }()
}

// Color provides the both of this and fatih package's Color method.
//
// This interface is for the take a same function name benchmark results using test flag.
type Color interface {
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

func genRandomBytes(tb testing.TB, length int) (b []byte) {
	tb.Helper()

	b = make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		tb.Fatal(err)
	}

	return b
}

var testAttributes = []color.Attribute{
	color.Bold,
	color.Italic,
	color.Underline,
	color.BlinkRapid,
	color.FgRed,
	color.FgCyan,
	color.FgHiGreen,
	color.FgHiBlue,
	color.BgRed,
	color.BgCyan,
	color.BgHiGreen,
	color.BgHiBlue,
}

var testFatihcAttributes = []fatihcolor.Attribute{
	fatihcolor.Bold,
	fatihcolor.Italic,
	fatihcolor.Underline,
	fatihcolor.BlinkRapid,
	fatihcolor.FgRed,
	fatihcolor.FgCyan,
	fatihcolor.FgHiGreen,
	fatihcolor.FgHiBlue,
	fatihcolor.BgRed,
	fatihcolor.BgCyan,
	fatihcolor.BgHiGreen,
	fatihcolor.BgHiBlue,
}

func BenchmarkNew(b *testing.B) {
	n := rand.Intn(11) + 1

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !*benchFatih {
			attrs := make([]color.Attribute, n)
			for i := 1; i < n; i++ {
				attrs[i] = testAttributes[rand.Intn(n)]
			}
			_ = color.New(attrs...)
		} else {
			attrs := make([]fatihcolor.Attribute, n)
			for i := 1; i < n; i++ {
				attrs[i] = testFatihcAttributes[rand.Intn(n)]
			}
			_ = fatihcolor.New(attrs...)
		}
	}
}

func getNewFnuc() Color {
	if *benchFatih {
		return fatihcolor.New(fatihcolor.FgGreen)
	}

	return color.New(color.FgGreen)
}

func benchmarkNewPrint(b *testing.B, fn Color, length int) {
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn.Print(buf)
	}
}

func BenchmarkNewPrint(b *testing.B) {
	benchmarkNewPrint(b, getNewFnuc(), length)
}

const numPrintFunc = 8

type printFuncs [numPrintFunc]func(format string, a ...interface{})

func getPrintFuncs() printFuncs {
	if *benchFatih {
		return printFuncs{
			fatihcolor.Black,
			fatihcolor.Red,
			fatihcolor.Green,
			fatihcolor.Yellow,
			fatihcolor.Blue,
			fatihcolor.Magenta,
			fatihcolor.Cyan,
			fatihcolor.White,
		}
	}

	return printFuncs{
		color.Black,
		color.Red,
		color.Green,
		color.Yellow,
		color.Blue,
		color.Magenta,
		color.Cyan,
		color.White,
	}
}

func benchmarkColorPrint(b *testing.B, fn printFuncs, length int) {
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(numPrintFunc)
		fn[n](format, buf)
	}
	b.SetBytes(int64(len(buf)))
}

func BenchmarkColorPrint(b *testing.B) {
	benchmarkColorPrint(b, getPrintFuncs(), length)
}

const numstringFunc = 8

type stringFuncs [numstringFunc]func(format string, a ...interface{}) string

func getStringFuncs() stringFuncs {
	if *benchFatih {
		return stringFuncs{
			fatihcolor.BlackString,
			fatihcolor.RedString,
			fatihcolor.GreenString,
			fatihcolor.YellowString,
			fatihcolor.BlueString,
			fatihcolor.MagentaString,
			fatihcolor.CyanString,
			fatihcolor.WhiteString,
		}
	}

	return stringFuncs{
		color.BlackString,
		color.RedString,
		color.GreenString,
		color.YellowString,
		color.BlueString,
		color.MagentaString,
		color.CyanString,
		color.WhiteString,
	}
}

func benchmarkColorString(b *testing.B, fn stringFuncs, length int) {
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(numstringFunc)
		_ = fn[n](format, buf)
	}
}

func BenchmarkColorString(b *testing.B) {
	benchmarkColorString(b, getStringFuncs(), length)
}

//go:linkname getCacheColor github.com/zchee/color.getCacheColor
func getCacheColor(p ...color.Attribute) (c *color.Color)

//go:linkname getCachedColorFatih github.com/fatih/color.getCachedColor
func getCachedColorFatih(p fatihcolor.Attribute) (c *fatihcolor.Color)

type attribute int

func GetCacheColor(p attribute) Color {
	if !*benchFatih {
		return getCacheColor(color.Attribute(p))
	}

	return getCachedColorFatih(fatihcolor.Attribute(p))
}

func genAttribute(i int) attribute {
	if !*benchFatih {
		return attribute(color.Attribute(rand.Intn(7) + i))
	}

	return attribute(fatihcolor.Attribute(rand.Intn(7) + i))
}

func benchmark_getCacheColor(b *testing.B, i int) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := genAttribute(i)
		_ = GetCacheColor(p)
	}
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

//go:linkname colorPrint github.com/zchee/color.colorPrint
func colorPrint(p color.Attribute, format string, a ...interface{})

//go:linkname colorPrintFatih github.com/fatih/color.colorPrint
func colorPrintFatih(format string, p fatihcolor.Attribute, a ...interface{})

func ColorPrint(p attribute, format string, a ...interface{}) {
	if !*benchFatih {
		colorPrint(color.Attribute(p), format, a...)
	} else {
		colorPrintFatih(format, fatihcolor.Attribute(p), a...)
	}
}

func benchmark_colorPrint(b *testing.B, i int) {
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := genAttribute(i)
		ColorPrint(p, format, buf)
	}
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

//go:linkname colorString github.com/zchee/color.colorString
func colorString(p color.Attribute, format string, a ...interface{}) string

//go:linkname colorStringFatih github.com/fatih/color.colorString
func colorStringFatih(format string, p fatihcolor.Attribute, a ...interface{}) string

func ColorString(format string, p attribute, a ...interface{}) string {
	if !*benchFatih {
		return colorString(color.Attribute(p), format, a...)
	}

	return colorStringFatih(format, fatihcolor.Attribute(p), a...)
}

func benchmark_colorString(b *testing.B, i int) {
	const format = "buf: %x\n"
	buf := genRandomBytes(b, length)

	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := genAttribute(i)
		_ = ColorString(format, p, buf)
	}
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
