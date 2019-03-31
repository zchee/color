// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"

	colorable "github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// Output defines the standard output of the print functions. By default
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()

	// Error defines a color supporting writer for os.Stderr.
	Error = colorable.NewColorableStderr()
)

// colorCache is used to reduce the count of created Color objects and
// allows to reuse already created objects with required Attribute using intern sync.Pool pattern.
var colorCache = sync.Pool{
	New: func() interface{} {
		return make(map[Attribute]*Color)
	},
}

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

const (
	allocMinSize = 0
	allocMaxSize = 2
)

type pool struct {
	c chan *Color
}

func newColorPool(size int) (p *pool) {
	return &pool{
		c: make(chan *Color, size),
	}
}

func (p *pool) Get() (c *Color) {
	select {
	case c = <-p.c:
		// reuse existing *Color
	default:
		c = &Color{params: make([]Attribute, allocMinSize, allocMaxSize)}
	}

	return
}

func (p *pool) Put(c *Color) {
	c.Reset()

	select {
	case p.c <- c:
	default:
		// Discard the buffer if the pool is full.
	}
}

var colorPool *pool

const (
	escapePrefix = "\x1b["
	escapeSuffix = "m"
)

// Attribute defines a single SGR Code.
type Attribute int

// Base attributes.
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors.
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors.
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors.
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors.
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New returns a newly created color object.
func New(value ...Attribute) (c *Color) {
	c = colorPool.Get()
	c.Add(value...)

	runtime.SetFinalizer(c, (*Color).Put)

	return c
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()

	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if NoColor {
		return
	}

	Output.Write(unsafeToSlice(escapePrefix + Reset.String() + escapeSuffix))
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)

	return c
}

// Put resets c.params and puts colorPool.
func (c *Color) Put() {
	c.Reset()
	colorPool.Put(c)
}

// Prepend prepends value Attribute to c.
func (c *Color) Prepend(value Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Reset resets the c.params slice.
func (c *Color) Reset() {
	c.params = c.params[:0]
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}
	fmt.Fprintf(Output, c.format())

	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() (s string) {
	var b strings.Builder

	for _, attr := range c.params {
		b.Write(unsafeToSlice(attr.String()))
		b.WriteByte(';')
	}

	s = b.String()[:b.Len()-1] // trim last ';'

	return
}

func (c *Color) format() string {
	return escapePrefix + c.sequence() + escapeSuffix
}

func (c *Color) unformat() string {
	return escapePrefix + Reset.String() + escapeSuffix
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func (c *Color) setWriter(w io.Writer) *Color {
	if c.isNoColorSet() {
		return c
	}
	fmt.Fprintf(w, c.format())

	return c
}

func (c *Color) unsetWriter(w io.Writer) {
	if c.isNoColorSet() || NoColor {
		return
	}

	w.Write(unsafeToSlice(escapePrefix + Reset.String() + escapeSuffix))
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	n, err = fmt.Fprint(w, a...)
	c.unsetWriter(w)

	return n, err
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	n, err = fmt.Fprint(Output, a...)
	c.unset()

	return n, err
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	n, err = fmt.Fprintf(w, format, a...)
	c.unsetWriter(w)

	return n, err
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	n, err = fmt.Fprintf(Output, format, a...)
	c.unset()

	return n, err
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	n, err = fmt.Fprintln(w, a...)
	c.unsetWriter(w)

	return n, err
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	n, err = fmt.Fprintln(Output, a...)
	c.unset()

	return n, err
}

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// FprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprint(w, a...)
	}
}

// PrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Print(a...)
	}
}

// FprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		c.Fprintf(w, format, a...)
	}
}

// PrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		c.Printf(format, a...)
	}
}

// FprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprintln(w, a...)
	}
}

// PrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Println(a...)
	}
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

func boolPtr(v bool) *bool {
	return &v
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

func (c *Color) attrExists(a Attribute) bool {
	for _, attr := range c.params {
		if attr == a {
			return true
		}
	}

	return false
}

// Equals returns a boolean value indicating whether two colors are equal.
func (c *Color) Equals(c2 *Color) bool {
	if len(c.params) != len(c2.params) {
		return false
	}

	for _, attr := range c.params {
		if !c2.attrExists(attr) {
			return false
		}
	}

	return true
}

func getCacheColor(p Attribute) (c *Color) {
	m := colorCache.Get().(map[Attribute]*Color)
	c, ok := m[p]
	if ok {
		colorCache.Put(m)
		return c
	}

	c = New(p)
	m[p] = c
	colorCache.Put(m)

	return c
}

func colorPrint(p Attribute, format string, a ...interface{}) {
	c := getCacheColor(p)

	if len(a) == 0 {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		c.Print(format)
		return
	}

	c.Printf(format, a...)
}

func colorString(p Attribute, format string, a ...interface{}) string {
	c := getCacheColor(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

// Black is a convenient helper function to print with black foreground. A
// newline is appended to format by default.
func Black(format string, a ...interface{}) { colorPrint(FgBlack, format, a...) }

// Red is a convenient helper function to print with red foreground. A
// newline is appended to format by default.
func Red(format string, a ...interface{}) { colorPrint(FgRed, format, a...) }

// Green is a convenient helper function to print with green foreground. A
// newline is appended to format by default.
func Green(format string, a ...interface{}) { colorPrint(FgGreen, format, a...) }

// Yellow is a convenient helper function to print with yellow foreground.
// A newline is appended to format by default.
func Yellow(format string, a ...interface{}) { colorPrint(FgYellow, format, a...) }

// Blue is a convenient helper function to print with blue foreground. A
// newline is appended to format by default.
func Blue(format string, a ...interface{}) { colorPrint(FgBlue, format, a...) }

// Magenta is a convenient helper function to print with magenta foreground.
// A newline is appended to format by default.
func Magenta(format string, a ...interface{}) { colorPrint(FgMagenta, format, a...) }

// Cyan is a convenient helper function to print with cyan foreground. A
// newline is appended to format by default.
func Cyan(format string, a ...interface{}) { colorPrint(FgCyan, format, a...) }

// White is a convenient helper function to print with white foreground. A
// newline is appended to format by default.
func White(format string, a ...interface{}) { colorPrint(FgWhite, format, a...) }

// BlackString is a convenient helper function to return a string with black
// foreground.
func BlackString(format string, a ...interface{}) string { return colorString(FgBlack, format, a...) }

// RedString is a convenient helper function to return a string with red
// foreground.
func RedString(format string, a ...interface{}) string { return colorString(FgRed, format, a...) }

// GreenString is a convenient helper function to return a string with green
// foreground.
func GreenString(format string, a ...interface{}) string { return colorString(FgGreen, format, a...) }

// YellowString is a convenient helper function to return a string with yellow
// foreground.
func YellowString(format string, a ...interface{}) string { return colorString(FgYellow, format, a...) }

// BlueString is a convenient helper function to return a string with blue
// foreground.
func BlueString(format string, a ...interface{}) string { return colorString(FgBlue, format, a...) }

// MagentaString is a convenient helper function to return a string with magenta
// foreground.
func MagentaString(format string, a ...interface{}) string {
	return colorString(FgMagenta, format, a...)
}

// CyanString is a convenient helper function to return a string with cyan
// foreground.
func CyanString(format string, a ...interface{}) string { return colorString(FgCyan, format, a...) }

// WhiteString is a convenient helper function to return a string with white
// foreground.
func WhiteString(format string, a ...interface{}) string { return colorString(FgWhite, format, a...) }

// HiBlack is a convenient helper function to print with hi-intensity black foreground. A
// newline is appended to format by default.
func HiBlack(format string, a ...interface{}) { colorPrint(FgHiBlack, format, a...) }

// HiRed is a convenient helper function to print with hi-intensity red foreground. A
// newline is appended to format by default.
func HiRed(format string, a ...interface{}) { colorPrint(FgHiRed, format, a...) }

// HiGreen is a convenient helper function to print with hi-intensity green foreground. A
// newline is appended to format by default.
func HiGreen(format string, a ...interface{}) { colorPrint(FgHiGreen, format, a...) }

// HiYellow is a convenient helper function to print with hi-intensity yellow foreground.
// A newline is appended to format by default.
func HiYellow(format string, a ...interface{}) { colorPrint(FgHiYellow, format, a...) }

// HiBlue is a convenient helper function to print with hi-intensity blue foreground. A
// newline is appended to format by default.
func HiBlue(format string, a ...interface{}) { colorPrint(FgHiBlue, format, a...) }

// HiMagenta is a convenient helper function to print with hi-intensity magenta foreground.
// A newline is appended to format by default.
func HiMagenta(format string, a ...interface{}) { colorPrint(FgHiMagenta, format, a...) }

// HiCyan is a convenient helper function to print with hi-intensity cyan foreground. A
// newline is appended to format by default.
func HiCyan(format string, a ...interface{}) { colorPrint(FgHiCyan, format, a...) }

// HiWhite is a convenient helper function to print with hi-intensity white foreground. A
// newline is appended to format by default.
func HiWhite(format string, a ...interface{}) { colorPrint(FgHiWhite, format, a...) }

// HiBlackString is a convenient helper function to return a string with hi-intensity black
// foreground.
func HiBlackString(format string, a ...interface{}) string {
	return colorString(FgHiBlack, format, a...)
}

// HiRedString is a convenient helper function to return a string with hi-intensity red
// foreground.
func HiRedString(format string, a ...interface{}) string { return colorString(FgHiRed, format, a...) }

// HiGreenString is a convenient helper function to return a string with hi-intensity green
// foreground.
func HiGreenString(format string, a ...interface{}) string {
	return colorString(FgHiGreen, format, a...)
}

// HiYellowString is a convenient helper function to return a string with hi-intensity yellow
// foreground.
func HiYellowString(format string, a ...interface{}) string {
	return colorString(FgHiYellow, format, a...)
}

// HiBlueString is a convenient helper function to return a string with hi-intensity blue
// foreground.
func HiBlueString(format string, a ...interface{}) string { return colorString(FgHiBlue, format, a...) }

// HiMagentaString is a convenient helper function to return a string with hi-intensity magenta
// foreground.
func HiMagentaString(format string, a ...interface{}) string {
	return colorString(FgHiMagenta, format, a...)
}

// HiCyanString is a convenient helper function to return a string with hi-intensity cyan
// foreground.
func HiCyanString(format string, a ...interface{}) string { return colorString(FgHiCyan, format, a...) }

// HiWhiteString is a convenient helper function to return a string with hi-intensity white
// foreground.
func HiWhiteString(format string, a ...interface{}) string {
	return colorString(FgHiWhite, format, a...)
}
