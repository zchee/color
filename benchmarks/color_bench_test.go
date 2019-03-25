// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark

package benchmarks_test

import (
	crand "crypto/rand"
	"io/ioutil"
	"testing"

	fatihcolor "github.com/fatih/color"
	"github.com/zchee/color"
)

func genRandomBytes(tb testing.TB, length int64) []byte {
	tb.Helper()

	b := make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		tb.Fatal(err)
	}

	return b
}

type printFunc func(...interface{}) (int, error)

func benchmarkNewPrint(b *testing.B, fn printFunc, length int64) {
	b.Helper()

	buf := genRandomBytes(b, length)
	b.SetBytes(length)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn(string(buf))
		}
	})
}

func BenchmarkZcheeColor(b *testing.B) {
	color.Output = ioutil.Discard
	color.NoColor = false
	benchmarkNewPrint(b, color.New(color.FgGreen).Print, int64(1024))
}

func BenchmarkFatihColor(b *testing.B) {
	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false
	benchmarkNewPrint(b, fatihcolor.New(fatihcolor.FgGreen).Print, int64(1024))
}
