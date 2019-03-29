// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

func init() {
	for i := 0; i < 107; i++ { // init pooled 64(default)+44(num of *Color cache)
		colorPool.Put(&Color{params: make([]Attribute, allocMinSize, allocMaxSize)})
	}

	m := make(map[Attribute]*Color, 44) // Total for loop is 8+3(bold,italic,underline)*4 = 44

	const (
		fgStart   = 30
		fgEnd     = 38
		fgHiStart = 90
		fgHiEnd   = 98
		bgStart   = 40
		bgEnd     = 48
		bgHiStart = 100
		bgHiEnd   = 108
	)

	var attrs = []Attribute{1, 3, 4}

	for fg := Attribute(fgStart); fg < Attribute(fgEnd); fg++ {
		m[fg] = New(fg)

		if fg == fgEnd {
			for _, j := range attrs {
				m[fg+j] = New(fg + j)
			}
		}
	}

	for fghi := Attribute(fgHiStart); fghi < Attribute(fgHiEnd); fghi++ {
		m[fghi] = New(fghi)

		if fghi == fgHiEnd {
			for _, j := range attrs {
				m[fghi+j] = New(fghi + j)
			}
		}
	}

	for bg := Attribute(bgStart); bg < Attribute(bgEnd); bg++ {
		m[bg] = New(bg)

		if bg == bgEnd {
			for _, j := range attrs {
				m[bg+j] = New(bg + j)
			}
		}
	}

	for bghi := Attribute(bgHiStart); bghi < Attribute(bgHiEnd); bghi++ {
		m[bghi] = New(bghi)

		if bghi == bgHiEnd {
			for _, j := range attrs {
				m[bghi+j] = New(bghi + j)
			}
		}
	}

	colorsCache.Store(m)
}
