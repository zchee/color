// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	colorable "github.com/mattn/go-colorable"

	"github.com/zchee/color"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())

	var code int
	defer os.Exit(code)
	code = m.Run()
}

// Testing colors is kinda different. First we test for given colors and their
// escaped formatted results. Next we create some visual tests to be tested.
// Each visual test includes the color name to be compared.
func TestColor(t *testing.T) {
	t.Parallel()

	b := new(bytes.Buffer)
	color.Output = b
	color.NoColor = false

	tests := []struct {
		text string
		attr color.Attribute
	}{
		{
			text: "black",
			attr: color.FgBlack,
		},
		{
			text: "red",
			attr: color.FgRed,
		},
		{
			text: "green",
			attr: color.FgGreen,
		},
		{
			text: "yellow",
			attr: color.FgYellow,
		},
		{
			text: "blue",
			attr: color.FgBlue,
		},
		{
			text: "magent",
			attr: color.FgMagenta,
		},
		{
			text: "cyan",
			attr: color.FgCyan,
		},
		{
			text: "white",
			attr: color.FgWhite,
		},
		{
			text: "hiblack",
			attr: color.FgHiBlack,
		},
		{
			text: "hired",
			attr: color.FgHiRed,
		},
		{
			text: "higreen",
			attr: color.FgHiGreen,
		},
		{
			text: "hiyellow",
			attr: color.FgHiYellow,
		},
		{
			text: "hiblue",
			attr: color.FgHiBlue,
		},
		{
			text: "himagent",
			attr: color.FgHiMagenta,
		},
		{
			text: "hicyan",
			attr: color.FgHiCyan,
		},
		{
			text: "hiwhite",
			attr: color.FgHiWhite,
		},
	}

	for _, tt := range tests {
		t.Run("New.Print("+tt.text+")", func(t *testing.T) {
			color.New(tt.attr).Print(tt.text)

			line, _ := b.ReadString('\n')
			scannedLine := fmt.Sprintf("%q", line)
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", tt.attr, tt.text)
			escapedForm := fmt.Sprintf("%q", colored)

			fmt.Printf("%s: %s\n", tt.text, line)

			if scannedLine != escapedForm {
				t.Errorf("got %q, want %q", scannedLine, escapedForm)
			}
		})
	}

	for _, tt := range tests {
		t.Run("New.Sprintf("+tt.text+")", func(t *testing.T) {
			line := color.New(tt.attr).Sprintf("%s", tt.text)
			scannedLine := fmt.Sprintf("%q", line)
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", tt.attr, tt.text)
			escapedForm := fmt.Sprintf("%q", colored)

			fmt.Printf("%s: %s\n", tt.text, line)

			if scannedLine != escapedForm {
				t.Errorf("got %q, want %q", scannedLine, escapedForm)
			}
		})
	}
}

func TestColorEquals(t *testing.T) {
	tests := []struct {
		name  string
		c     *color.Color
		wantC *color.Color
		want  bool
	}{
		{
			name:  "Two black colors are equal",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.FgBlack),
			want:  true,
		},
		{
			name:  "Fg and bg black colors are not equal",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.BgBlack),
			want:  false,
		},
		{
			name:  "Fg black not equals fg/bg black color",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.FgBlack, color.BgBlack),
			want:  false,
		},
		{
			name:  "Fg black not equals Fg red",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.FgRed),
			want:  false,
		},
		{
			name:  "Fg black not equals Bg red",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.BgRed),
			want:  false,
		},
		{
			name:  "Fg black not equals fg black bg red",
			c:     color.New(color.FgBlack),
			wantC: color.New(color.FgBlack, color.BgRed),
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.c.Equals(tt.wantC); got != tt.want {
				t.Errorf(strings.Replace(tt.name, " not", "", 1))
			}
		})
	}
}

func TestNoColor(t *testing.T) {
	b := new(bytes.Buffer)
	color.Output = b

	tests := []struct {
		text string
		attr color.Attribute
	}{
		{
			text: "black",
			attr: color.FgBlack,
		},
		{
			text: "red",
			attr: color.FgRed,
		},
		{
			text: "green",
			attr: color.FgGreen,
		},
		{
			text: "yellow",
			attr: color.FgYellow,
		},
		{
			text: "blue",
			attr: color.FgBlue,
		},
		{
			text: "magent",
			attr: color.FgMagenta,
		},
		{
			text: "cyan",
			attr: color.FgCyan,
		},
		{
			text: "white",
			attr: color.FgWhite,
		},
		{
			text: "hiblack",
			attr: color.FgHiBlack,
		},
		{
			text: "hired",
			attr: color.FgHiRed,
		},
		{
			text: "higreen",
			attr: color.FgHiGreen,
		},
		{
			text: "hiyellow",
			attr: color.FgHiYellow,
		},
		{
			text: "hiblue",
			attr: color.FgHiBlue,
		},
		{
			text: "himagent",
			attr: color.FgHiMagenta,
		},
		{
			text: "hicyan",
			attr: color.FgHiCyan,
		},
		{
			text: "hiwhite",
			attr: color.FgHiWhite,
		},
	}

	for _, tt := range tests {
		t.Run("DisableColor", func(t *testing.T) {
			p := color.New(tt.attr)
			p.DisableColor()
			p.Print(tt.text)

			line, _ := b.ReadString('\n')
			if line != tt.text {
				t.Errorf("got %q, want %q", line, tt.text)
			}
		})
	}

	for _, tt := range tests {
		t.Run("color.NoColor", func(t *testing.T) {
			// global check
			color.NoColor = true
			defer func() {
				color.NoColor = false
			}()

			p := color.New(tt.attr)
			p.Print(tt.text)

			line, _ := b.ReadString('\n')
			if line != tt.text {
				t.Errorf("got %q, want %q", line, tt.text)
			}
		})
	}

}

func TestColorVisual(t *testing.T) {
	// First Visual Test
	color.Output = colorable.NewColorableStdout()

	color.New(color.FgRed).Printf("red\t")
	color.New(color.BgRed).Print("         ")
	color.New(color.FgRed, color.Bold).Println(" red")

	color.New(color.FgGreen).Printf("green\t")
	color.New(color.BgGreen).Print("         ")
	color.New(color.FgGreen, color.Bold).Println(" green")

	color.New(color.FgYellow).Printf("yellow\t")
	color.New(color.BgYellow).Print("         ")
	color.New(color.FgYellow, color.Bold).Println(" yellow")

	color.New(color.FgBlue).Printf("blue\t")
	color.New(color.BgBlue).Print("         ")
	color.New(color.FgBlue, color.Bold).Println(" blue")

	color.New(color.FgMagenta).Printf("magenta\t")
	color.New(color.BgMagenta).Print("         ")
	color.New(color.FgMagenta, color.Bold).Println(" magenta")

	color.New(color.FgCyan).Printf("cyan\t")
	color.New(color.BgCyan).Print("         ")
	color.New(color.FgCyan, color.Bold).Println(" cyan")

	color.New(color.FgWhite).Printf("white\t")
	color.New(color.BgWhite).Print("         ")
	color.New(color.FgWhite, color.Bold).Println(" white")
	fmt.Println("")

	// Second Visual test
	color.Black("black")
	color.Red("red")
	color.Green("green")
	color.Yellow("yellow")
	color.Blue("blue")
	color.Magenta("magenta")
	color.Cyan("cyan")
	color.White("white")
	color.HiBlack("hblack")
	color.HiRed("hred")
	color.HiGreen("hgreen")
	color.HiYellow("hyellow")
	color.HiBlue("hblue")
	color.HiMagenta("hmagenta")
	color.HiCyan("hcyan")
	color.HiWhite("hwhite")

	// Third visual test
	fmt.Println()
	color.Set(color.FgBlue)
	fmt.Println("is this blue?")
	color.Unset()

	color.Set(color.FgMagenta)
	fmt.Println("and this magenta?")
	color.Unset()

	// Fourth Visual test
	fmt.Println()
	blue := color.New(color.FgBlue).PrintlnFunc()
	blue("blue text with custom print func")

	red := color.New(color.FgRed).PrintfFunc()
	red("red text with a printf func: %d\n", 123)

	put := color.New(color.FgYellow).SprintFunc()
	warn := color.New(color.FgRed).SprintFunc()

	fmt.Fprintf(color.Output, "this is a %s and this is %s.\n", put("warning"), warn("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Fprintf(color.Output, "this %s rocks!\n", info("package"))

	notice := color.New(color.FgBlue).FprintFunc()
	notice(os.Stderr, "just a blue notice to stderr")

	// Fifth Visual Test
	fmt.Println()

	fmt.Fprintln(color.Output, color.BlackString("black"))
	fmt.Fprintln(color.Output, color.RedString("red"))
	fmt.Fprintln(color.Output, color.GreenString("green"))
	fmt.Fprintln(color.Output, color.YellowString("yellow"))
	fmt.Fprintln(color.Output, color.BlueString("blue"))
	fmt.Fprintln(color.Output, color.MagentaString("magenta"))
	fmt.Fprintln(color.Output, color.CyanString("cyan"))
	fmt.Fprintln(color.Output, color.WhiteString("white"))
	fmt.Fprintln(color.Output, color.HiBlackString("hblack"))
	fmt.Fprintln(color.Output, color.HiRedString("hred"))
	fmt.Fprintln(color.Output, color.HiGreenString("hgreen"))
	fmt.Fprintln(color.Output, color.HiYellowString("hyellow"))
	fmt.Fprintln(color.Output, color.HiBlueString("hblue"))
	fmt.Fprintln(color.Output, color.HiMagentaString("hmagenta"))
	fmt.Fprintln(color.Output, color.HiCyanString("hcyan"))
	fmt.Fprintln(color.Output, color.HiWhiteString("hwhite"))
}

func TestNoFormat(t *testing.T) {
	fmt.Printf("%s   %%s = ", color.BlackString("Black"))
	color.Black("%s")

	fmt.Printf("%s     %%s = ", color.RedString("Red"))
	color.Red("%s")

	fmt.Printf("%s   %%s = ", color.GreenString("Green"))
	color.Green("%s")

	fmt.Printf("%s  %%s = ", color.YellowString("Yellow"))
	color.Yellow("%s")

	fmt.Printf("%s    %%s = ", color.BlueString("Blue"))
	color.Blue("%s")

	fmt.Printf("%s %%s = ", color.MagentaString("Magenta"))
	color.Magenta("%s")

	fmt.Printf("%s    %%s = ", color.CyanString("Cyan"))
	color.Cyan("%s")

	fmt.Printf("%s   %%s = ", color.WhiteString("White"))
	color.White("%s")

	fmt.Printf("%s   %%s = ", color.HiBlackString("HiBlack"))
	color.HiBlack("%s")

	fmt.Printf("%s     %%s = ", color.HiRedString("HiRed"))
	color.HiRed("%s")

	fmt.Printf("%s   %%s = ", color.HiGreenString("HiGreen"))
	color.HiGreen("%s")

	fmt.Printf("%s  %%s = ", color.HiYellowString("HiYellow"))
	color.HiYellow("%s")

	fmt.Printf("%s    %%s = ", color.HiBlueString("HiBlue"))
	color.HiBlue("%s")

	fmt.Printf("%s %%s = ", color.HiMagentaString("HiMagenta"))
	color.HiMagenta("%s")

	fmt.Printf("%s    %%s = ", color.HiCyanString("HiCyan"))
	color.HiCyan("%s")

	fmt.Printf("%s   %%s = ", color.HiWhiteString("HiWhite"))
	color.HiWhite("%s")
}

func TestNoFormatString(t *testing.T) {
	tests := []struct {
		fn     func(string, ...interface{}) string
		format string
		args   []interface{}
		want   string
	}{
		{
			fn:     color.BlackString,
			format: "%s",
			args:   nil,
			want:   "\x1b[30m%s\x1b[0m",
		},
		{
			fn:     color.RedString,
			format: "%s",
			args:   nil,
			want:   "\x1b[31m%s\x1b[0m",
		},
		{
			fn:     color.GreenString,
			format: "%s",
			args:   nil,
			want:   "\x1b[32m%s\x1b[0m",
		},
		{
			fn:     color.YellowString,
			format: "%s",
			args:   nil,
			want:   "\x1b[33m%s\x1b[0m",
		},
		{
			fn:     color.BlueString,
			format: "%s",
			args:   nil,
			want:   "\x1b[34m%s\x1b[0m",
		},
		{
			fn:     color.MagentaString,
			format: "%s",
			args:   nil,
			want:   "\x1b[35m%s\x1b[0m",
		},
		{
			fn:     color.CyanString,
			format: "%s",
			args:   nil,
			want:   "\x1b[36m%s\x1b[0m",
		},
		{
			fn:     color.WhiteString,
			format: "%s",
			args:   nil,
			want:   "\x1b[37m%s\x1b[0m",
		},
		{
			fn:     color.HiBlackString,
			format: "%s",
			args:   nil,
			want:   "\x1b[90m%s\x1b[0m",
		},
		{
			fn:     color.HiRedString,
			format: "%s",
			args:   nil,
			want:   "\x1b[91m%s\x1b[0m",
		},
		{
			fn:     color.HiGreenString,
			format: "%s",
			args:   nil,
			want:   "\x1b[92m%s\x1b[0m",
		},
		{
			fn:     color.HiYellowString,
			format: "%s",
			args:   nil,
			want:   "\x1b[93m%s\x1b[0m",
		},
		{
			fn:     color.HiBlueString,
			format: "%s",
			args:   nil,
			want:   "\x1b[94m%s\x1b[0m",
		},
		{
			fn:     color.HiMagentaString,
			format: "%s",
			args:   nil,
			want:   "\x1b[95m%s\x1b[0m",
		},
		{
			fn:     color.HiCyanString,
			format: "%s",
			args:   nil,
			want:   "\x1b[96m%s\x1b[0m",
		},
		{
			fn:     color.HiWhiteString,
			format: "%s",
			args:   nil,
			want:   "\x1b[97m%s\x1b[0m",
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			s := fmt.Sprintf("%s", tt.fn(tt.format, tt.args...))
			if s != tt.want {
				t.Errorf("[%d] got: %q, want: %q", i, s, tt.want)
			}
		})
	}
}
