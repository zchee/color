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

func BenchmarkNewPrint(b *testing.B) {
	const length = int64(1024)
	color.Output = ioutil.Discard
	color.NoColor = false

	benchmarkNewPrint(b, color.New(color.FgGreen).Print, length)
}

func BenchmarkColorPrint(b *testing.B) {
	const length = int64(1024)
	color.Output = ioutil.Discard
	color.NoColor = false

	benchmarkColorPrint(b, color.Magenta, length)
}
