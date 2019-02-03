// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	colorable "github.com/mattn/go-colorable"
	"github.com/zchee/color"
)

// Testing colors is kinda different. First we test for given colors and their
// escaped formatted results. Next we create some visual tests to be tested.
// Each visual test includes the color name to be compared.
func TestColor(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb

	color.NoColor = false

	testColors := []struct {
		text string
		code color.Attribute
	}{
		{text: "black", code: color.FgBlack},
		{text: "red", code: color.FgRed},
		{text: "green", code: color.FgGreen},
		{text: "yellow", code: color.FgYellow},
		{text: "blue", code: color.FgBlue},
		{text: "magent", code: color.FgMagenta},
		{text: "cyan", code: color.FgCyan},
		{text: "white", code: color.FgWhite},
		{text: "hblack", code: color.FgHiBlack},
		{text: "hred", code: color.FgHiRed},
		{text: "hgreen", code: color.FgHiGreen},
		{text: "hyellow", code: color.FgHiYellow},
		{text: "hblue", code: color.FgHiBlue},
		{text: "hmagent", code: color.FgHiMagenta},
		{text: "hcyan", code: color.FgHiCyan},
		{text: "hwhite", code: color.FgHiWhite},
	}

	for _, c := range testColors {
		color.New(c.code).Print(c.text)

		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)

		fmt.Printf("%s\t: %s\n", c.text, line)

		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}

	for _, c := range testColors {
		line := color.New(c.code).Sprintf("%s", c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)

		fmt.Printf("%s\t: %s\n", c.text, line)

		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestColorEquals(t *testing.T) {
	fgblack1 := color.New(color.FgBlack)
	fgblack2 := color.New(color.FgBlack)
	bgblack := color.New(color.BgBlack)
	fgbgblack := color.New(color.FgBlack, color.BgBlack)
	fgblackbgred := color.New(color.FgBlack, color.BgRed)
	fgred := color.New(color.FgRed)
	bgred := color.New(color.BgRed)

	if !fgblack1.Equals(fgblack2) {
		t.Error("Two black colors are not equal")
	}

	if fgblack1.Equals(bgblack) {
		t.Error("Fg and bg black colors are equal")
	}

	if fgblack1.Equals(fgbgblack) {
		t.Error("Fg black equals fg/bg black color")
	}

	if fgblack1.Equals(fgred) {
		t.Error("Fg black equals Fg red")
	}

	if fgblack1.Equals(bgred) {
		t.Error("Fg black equals Bg red")
	}

	if fgblack1.Equals(fgblackbgred) {
		t.Error("Fg black equals fg black bg red")
	}
}

func TestNoColor(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb

	testColors := []struct {
		text string
		code color.Attribute
	}{
		{text: "black", code: color.FgBlack},
		{text: "red", code: color.FgRed},
		{text: "green", code: color.FgGreen},
		{text: "yellow", code: color.FgYellow},
		{text: "blue", code: color.FgBlue},
		{text: "magent", code: color.FgMagenta},
		{text: "cyan", code: color.FgCyan},
		{text: "white", code: color.FgWhite},
		{text: "hblack", code: color.FgHiBlack},
		{text: "hred", code: color.FgHiRed},
		{text: "hgreen", code: color.FgHiGreen},
		{text: "hyellow", code: color.FgHiYellow},
		{text: "hblue", code: color.FgHiBlue},
		{text: "hmagent", code: color.FgHiMagenta},
		{text: "hcyan", code: color.FgHiCyan},
		{text: "hwhite", code: color.FgHiWhite},
	}

	for _, c := range testColors {
		p := color.New(c.code)
		p.DisableColor()
		p.Print(c.text)

		line, _ := rb.ReadString('\n')
		if line != c.text {
			t.Errorf("Expecting %s, got '%s'\n", c.text, line)
		}
	}

	// global check
	color.NoColor = true
	defer func() {
		color.NoColor = false
	}()
	for _, c := range testColors {
		p := color.New(c.code)
		p.Print(c.text)

		line, _ := rb.ReadString('\n')
		if line != c.text {
			t.Errorf("Expecting %s, got '%s'\n", c.text, line)
		}
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
		f      func(string, ...interface{}) string
		format string
		args   []interface{}
		want   string
	}{
		{color.BlackString, "%s", nil, "\x1b[30m%s\x1b[0m"},
		{color.RedString, "%s", nil, "\x1b[31m%s\x1b[0m"},
		{color.GreenString, "%s", nil, "\x1b[32m%s\x1b[0m"},
		{color.YellowString, "%s", nil, "\x1b[33m%s\x1b[0m"},
		{color.BlueString, "%s", nil, "\x1b[34m%s\x1b[0m"},
		{color.MagentaString, "%s", nil, "\x1b[35m%s\x1b[0m"},
		{color.CyanString, "%s", nil, "\x1b[36m%s\x1b[0m"},
		{color.WhiteString, "%s", nil, "\x1b[37m%s\x1b[0m"},
		{color.HiBlackString, "%s", nil, "\x1b[90m%s\x1b[0m"},
		{color.HiRedString, "%s", nil, "\x1b[91m%s\x1b[0m"},
		{color.HiGreenString, "%s", nil, "\x1b[92m%s\x1b[0m"},
		{color.HiYellowString, "%s", nil, "\x1b[93m%s\x1b[0m"},
		{color.HiBlueString, "%s", nil, "\x1b[94m%s\x1b[0m"},
		{color.HiMagentaString, "%s", nil, "\x1b[95m%s\x1b[0m"},
		{color.HiCyanString, "%s", nil, "\x1b[96m%s\x1b[0m"},
		{color.HiWhiteString, "%s", nil, "\x1b[97m%s\x1b[0m"},
	}

	for i, test := range tests {
		s := fmt.Sprintf("%s", test.f(test.format, test.args...))
		if s != test.want {
			t.Errorf("[%d] want: %q, got: %q", i, test.want, s)
		}
	}
}
