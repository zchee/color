// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark

package benchmarks_test

import (
	"io/ioutil"
	"testing"

	"github.com/zchee/color"
)

const length = int64(1024)

func BenchmarkNewPrint(b *testing.B) {
	color.Output = ioutil.Discard
	color.NoColor = false

	benchmarkNewPrint(b, color.New(color.FgGreen), length)
}

func BenchmarkColorPrint(b *testing.B) {
	color.Output = ioutil.Discard
	color.NoColor = false

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
	color.Output = ioutil.Discard
	color.NoColor = false

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
