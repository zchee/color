// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

func init() {
	for i := 0; i < 96; i++ { // init pooled 64(default)+32(num of *Color cache)
		colorPool.Put(&Color{params: make([]Attribute, allocMinSize, allocMaxSize)})
	}

	m := make(map[Attribute]*Color, 32) // Total for loop is 8(Black~White) * 4({F,B}g{,Hi}) = 32

	for _, attrs := range [4][2]Attribute{
		{
			FgBlack,
			FgWhite,
		},
		{
			FgHiBlack,
			FgHiWhite,
		},
		{
			BgBlack,
			BgWhite,
		},
		{
			BgHiBlack,
			BgHiWhite,
		},
	} {
		start := attrs[0]
		end := attrs[1]
		for attr := start; attr < end; attr++ {
			m[attr] = New(attr)
		}
	}

	colorsCache.Store(m)
}
