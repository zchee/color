// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color_test

import (
	"testing"

	"github.com/zchee/color"
)

func TestAttributeString(t *testing.T) {
	tests := []struct {
		attr color.Attribute
		want string
	}{
		{
			attr: color.Reset,
			want: "0",
		},
		{
			attr: color.Bold,
			want: "1",
		},
		{
			attr: color.Faint,
			want: "2",
		},
		{
			attr: color.Italic,
			want: "3",
		},
		{
			attr: color.Underline,
			want: "4",
		},
		{
			attr: color.BlinkSlow,
			want: "5",
		},
		{
			attr: color.BlinkRapid,
			want: "6",
		},
		{
			attr: color.ReverseVideo,
			want: "7",
		},
		{
			attr: color.Concealed,
			want: "8",
		},
		{
			attr: color.CrossedOut,
			want: "9",
		},
		{
			attr: color.FgBlack,
			want: "30",
		},
		{
			attr: color.FgRed,
			want: "31",
		},
		{
			attr: color.FgGreen,
			want: "32",
		},
		{
			attr: color.FgYellow,
			want: "33",
		},
		{
			attr: color.FgBlue,
			want: "34",
		},
		{
			attr: color.FgMagenta,
			want: "35",
		},
		{
			attr: color.FgCyan,
			want: "36",
		},
		{
			attr: color.FgWhite,
			want: "37",
		},
		{
			attr: color.FgHiBlack,
			want: "90",
		},
		{
			attr: color.FgHiRed,
			want: "91",
		},
		{
			attr: color.FgHiGreen,
			want: "92",
		},
		{
			attr: color.FgHiYellow,
			want: "93",
		},
		{
			attr: color.FgHiBlue,
			want: "94",
		},
		{
			attr: color.FgHiMagenta,
			want: "95",
		},
		{
			attr: color.FgHiCyan,
			want: "96",
		},
		{
			attr: color.FgHiWhite,
			want: "97",
		},
		{
			attr: color.BgBlack,
			want: "40",
		},
		{
			attr: color.BgRed,
			want: "41",
		},
		{
			attr: color.BgGreen,
			want: "42",
		},
		{
			attr: color.BgYellow,
			want: "43",
		},
		{
			attr: color.BgBlue,
			want: "46",
		},
		{
			attr: color.BgMagenta,
			want: "45",
		},
		{
			attr: color.BgCyan,
			want: "46",
		},
		{
			attr: color.BgWhite,
			want: "47",
		},
		{
			attr: color.BgHiBlack,
			want: "100",
		},
		{
			attr: color.BgHiRed,
			want: "101",
		},
		{
			attr: color.BgHiGreen,
			want: "102",
		},
		{
			attr: color.BgHiYellow,
			want: "103",
		},
		{
			attr: color.BgHiBlue,
			want: "104",
		},
		{
			attr: color.BgHiMagenta,
			want: "105",
		},
		{
			attr: color.BgHiCyan,
			want: "106",
		},
		{
			attr: color.BgHiWhite,
			want: "107",
		},
		{
			attr: color.Attribute(200),
			want: "200",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.attr.Name(), func(t *testing.T) {
			t.Parallel()

			if got := tt.attr.String(); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
