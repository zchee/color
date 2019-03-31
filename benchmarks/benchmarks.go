// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build benchmark benchmark_fatih

package benchmarks

import (
	_ "unsafe" // required for go:linkname

	"github.com/zchee/color"
)

//go:linkname getCacheColor github.com/zchee/color.getCacheColor
func getCacheColor(p color.Attribute) (c *color.Color)

func GetCacheColor(p color.Attribute) (c *color.Color) {
	return getCacheColor(p)
}

//go:linkname colorPrint github.com/zchee/color.colorPrint
func colorPrint(p color.Attribute, format string, a ...interface{})

func ColorPrint(p color.Attribute, format string, a ...interface{}) {
	colorPrint(p, format, a...)
}

//go:linkname colorString github.com/zchee/color.colorString
func colorString(p color.Attribute, format string, a ...interface{}) string

func ColorString(p color.Attribute, format string, a ...interface{}) string {
	return colorString(p, format, a...)
}
