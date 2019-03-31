// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark_fatih

package benchmarks

import (
	_ "unsafe" // required for go:linkname

	fatihcolor "github.com/fatih/color"
)

//go:linkname getCachedColor github.com/fatih/color.getCachedColor
func getCachedColor(p fatihcolor.Attribute) (c *fatihcolor.Color)

func GetCachedColor(p fatihcolor.Attribute) (c *fatihcolor.Color) {
	return getCachedColor(p)
}

//go:linkname colorPrint github.com/fatih/color.colorPrint
func colorPrint(format string, p fatihcolor.Attribute, a ...interface{})

func ColorPrint(format string, p fatihcolor.Attribute, a ...interface{}) {
	colorPrint(format, p, a...)
}

//go:linkname colorString github.com/fatih/color.colorString
func colorString(format string, p fatihcolor.Attribute, a ...interface{}) string

func ColorString(format string, p fatihcolor.Attribute, a ...interface{}) string {
	return colorString(format, p, a...)
}
