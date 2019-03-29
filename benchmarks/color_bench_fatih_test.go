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

	benchmarkColorPrint(b, fatihcolor.Magenta, length)
}

func BenchmarkColorString(b *testing.B) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	benchmarkColorString(b, fatihcolor.RedString, length)
}
