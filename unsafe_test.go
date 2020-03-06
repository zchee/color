// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

import (
	"math/rand"
	"reflect"
	"testing"
	"unsafe"
)

func Test_unsafeToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []byte
	}{
		{
			name: "hello",
			s:    "Hello world",
			want: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
		},
		{
			name: "empty",
			s:    "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unsafeByteSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unsafeToSlice(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func Benchmark_unsafeToSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		s := randomString(rand.Intn(65526) + 10)
		b.SetBytes(int64(len(s)))
		for pb.Next() {
			_ = unsafeByteSlice(s)
		}
	})
}
