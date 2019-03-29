// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

func init() {
	for i := 0; i < 107; i++ { // init pooled 64(default)+44(num of *Color cache)
		colorPool.Put(&Color{params: make([]Attribute, 0, defaultAllocSize)})
	}

	m := make(map[Attribute]*Color, 44) // Total for loop is 8+3(bold,italic,underline)*4 = 44

	for fg := Attribute(30); fg < Attribute(38); fg++ {
		m[fg] = New(fg)

		if fg == 38 {
			for _, j := range []Attribute{1, 3, 4} {
				m[fg+j] = New(fg + j)
			}
		}
	}

	for fghi := Attribute(90); fghi < Attribute(98); fghi++ {
		m[fghi] = New(fghi)

		if fghi == 98 {
			for _, j := range []Attribute{1, 3, 4} {
				m[fghi+j] = New(fghi + j)
			}
		}
	}

	for bg := Attribute(40); bg < Attribute(48); bg++ {
		m[bg] = New(bg)

		if bg == 48 {
			for _, j := range []Attribute{1, 3, 4} {
				m[bg+j] = New(bg + j)
			}
		}
	}

	for bghi := Attribute(100); bghi < Attribute(108); bghi++ {
		m[bghi] = New(bghi)

		if bghi == 108 {
			for _, j := range []Attribute{1, 3, 4} {
				m[bghi+j] = New(bghi + j)
			}
		}
	}

	colorsCache.Store(m)
}
