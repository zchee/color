// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark_fatih

package benchmarks_test

import (
	"io/ioutil"
	"testing"

	fatihcolor "github.com/fatih/color"
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
