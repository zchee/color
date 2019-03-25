// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color_test

import (
	crand "crypto/rand"
	"io/ioutil"
	"testing"

	fatihcolor "github.com/fatih/color"
	"github.com/zchee/color"
)

func BenchmarkZcheeColor(b *testing.B) {
	const length = int64(1024)

	color.Output = ioutil.Discard
	color.NoColor = false

	buf := genRandomBytes(b, length)
	b.SetBytes(length)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			color.New(color.FgGreen).Print(string(buf))
		}
	})
}

func BenchmarkFatihColor(b *testing.B) {
	const length = int64(1024)

	fatihcolor.Output = ioutil.Discard
	fatihcolor.NoColor = false

	buf := genRandomBytes(b, length)
	b.SetBytes(length)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fatihcolor.New(fatihcolor.FgGreen).Print(string(buf))
		}
	})
}

func genRandomBytes(tb testing.TB, length int64) []byte {
	tb.Helper()

	b := make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		tb.Fatal(err)
	}

	return b
}
