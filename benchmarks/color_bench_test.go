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

	benchmarkColorPrint(b, color.Magenta, length)
}

func BenchmarkColorString(b *testing.B) {
	color.Output = ioutil.Discard
	color.NoColor = false

	benchmarkColorString(b, color.RedString, length)
}
