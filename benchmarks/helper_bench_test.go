// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks_test

import (
	crand "crypto/rand"
	"testing"
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
	buf := genRandomBytes(b, length)
	b.SetBytes(length)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn(string(buf))
		}
	})
}
