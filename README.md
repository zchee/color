# Color

[![CircleCI][circleci-badge]][circleci] [![godoc.org][godoc-badge]][godoc] [![codecov.io][codecov-badge]][codecov] [![Releases][tag-badge]][tag] [![GA][ga-badge]][ga]

Color lets you use colorized outputs in terms of [ANSI Escape
Codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors) in Go (Golang). It
has support for Windows too! The API can be used in several ways, pick one that
suits you.


![Color](https://i.imgur.com/c1JI0lA.png)


## Install

```bash
go get github.com/zchee/color
```

Note that the `vendor` folder is here for stability. Remove the folder if you
already have the dependencies in your GOPATH.

## Examples

### Standard colors

```go
// Print with default helper functions
color.Cyan("Prints text in cyan.")

// A newline will be appended automatically
color.Blue("Prints %s in blue.", "text")

// These are using the default foreground colors
color.Red("We have red")
color.Magenta("And many others ..")

```

### Mix and reuse colors

```go
// Create a new color object
c := color.New(color.FgCyan).Add(color.Underline)
c.Println("Prints cyan text with an underline.")

// Or just add them to New()
d := color.New(color.FgCyan, color.Bold)
d.Printf("This prints bold cyan %s\n", "too!.")

// Mix up foreground and background colors, create new mixes!
red := color.New(color.FgRed)

boldRed := red.Add(color.Bold)
boldRed.Println("This will print text in bold red.")

whiteBackground := red.Add(color.BgWhite)
whiteBackground.Println("Red text with white background.")
```

### Use your own output (io.Writer)

```go
// Use your own io.Writer output
color.New(color.FgBlue).Fprintln(myWriter, "blue color!")

blue := color.New(color.FgBlue)
blue.Fprint(writer, "This will print text in blue.")
```

### Custom print functions (PrintFunc)

```go
// Create a custom print function for convenience
red := color.New(color.FgRed).PrintfFunc()
red("Warning")
red("Error: %s", err)

// Mix up multiple attributes
notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
notice("Don't forget this...")
```

### Custom fprint functions (FprintFunc)

```go
blue := color.New(FgBlue).FprintfFunc()
blue(myWriter, "important notice: %s", stars)

// Mix up with multiple attributes
success := color.New(color.Bold, color.FgGreen).FprintlnFunc()
success(myWriter, "Don't forget this...")
```

### Insert into noncolor strings (SprintFunc)

```go
// Create SprintXxx functions to mix strings with other non-colorized strings:
yellow := color.New(color.FgYellow).SprintFunc()
red := color.New(color.FgRed).SprintFunc()
fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
fmt.Printf("This %s rocks!\n", info("package"))

// Use helper functions
fmt.Println("This", color.RedString("warning"), "should be not neglected.")
fmt.Printf("%v %v\n", color.GreenString("Info:"), "an important message.")

// Windows supported too! Just don't forget to change the output to color.Output
fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))
```

### Plug into existing code

```go
// Use handy standard colors
color.Set(color.FgYellow)

fmt.Println("Existing text will now be in yellow")
fmt.Printf("This one %s\n", "too")

color.Unset() // Don't forget to unset

// You can mix up parameters
color.Set(color.FgMagenta, color.Bold)
defer color.Unset() // Use it in your function

fmt.Println("All text will now be bold magenta.")
```

### Disable/Enable color
 
There might be a case where you want to explicitly disable/enable color output. the 
`go-isatty` package will automatically disable color output for non-tty output streams 
(for example if the output were piped directly to `less`)

`Color` has support to disable/enable colors both globally and for single color 
definitions. For example suppose you have a CLI app and a `--no-color` bool flag. You 
can easily disable the color output with:

```go

var flagNoColor = flag.Bool("no-color", false, "Disable color output")

if *flagNoColor {
	color.NoColor = true // disables colorized output
}
```

It also has support for single color definitions (local). You can
disable/enable color output on the fly:

```go
c := color.New(color.FgCyan)
c.Println("Prints cyan text")

c.DisableColor()
c.Println("This is printed without any color")

c.EnableColor()
c.Println("This prints again cyan...")
```

## Benchmark

### Run benchmark

```sh
$ cd ./benchmarks
$ go test -v -tags=benchmark_fatih -cpu 1,4,12 -count 10 -run='^$' -bench=. -benchtime=2s . | tee old.txt
$ go test -v -tags=benchmark -cpu 1,4,12 -count 10 -run='^$' -bench=. -benchtime=2s . | tee new.txt
$ benchstat old.txt new.txt
```

### Benchmark result

On my Macbook Pro.

- `list_cpu_features` command
  - [google/cpu_features](https://github.com/google/cpu_features)
- `lscpu` command (on macOS)
  - [NanXiao/lscpu](https://github.com/NanXiao/lscpu)

```console
$ system_profiler SPHardwareDataType
# (omitted)
Model Name: MacBook Pro
Model Identifier: MacBookPro15,1
Processor Name: Intel Core i9
Processor Speed: 2.9 GHz
Number of Processors: 1
Total Number of Cores: 6
L2 Cache (per Core): 256 KB
L3 Cache: 12 MB
Memory: 32 GB

$ list_cpu_features
arch            : x86
brand           : Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
family          :   6 (0x06)
model           : 158 (0x9E)
stepping        :  10 (0x0A)
uarch           : INTEL_KBL
flags           : aes,avx,avx2,bmi1,bmi2,cx16,erms,f16c,fma3,movbe,popcnt,rdrnd,sgx,sse4_1,sse4_2,ssse3

$ lscpu
Architecture:            x86_64
Byte Order:              Little Endian
Total CPU(s):            12
Thread(s) per core:      2
Core(s) per socket:      6
Socket(s):               1
Vendor:                  GenuineIntel
CPU family:              6
Model:                   158
Model name:              MacBookPro15,1
Stepping:                10
L1d cache:               32K
L1i cache:               32K
L2 cache:                256K
L3 cache:                12M
Flags:                   fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 cflsh ds acpi mmx fxsr sse sse2 ss htt tm pbe sse3 pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline aes xsave osxsave avx f16c rdrnd syscall nx pdpe1gb rdtscp lm lahf_lm lzcnt
```

```
name            old time/op    new time/op    delta
NewPrint          18.9µs ± 2%    19.7µs ± 3%   +3.88%  (p=0.000 n=10+10)
NewPrint-4        6.26µs ± 2%    6.34µs ±13%     ~     (p=0.965 n=8+10)
NewPrint-12       4.09µs ± 3%    4.73µs ± 7%  +15.76%  (p=0.000 n=10+10)
ColorPrint        2.66µs ± 5%    2.25µs ± 7%  -15.66%  (p=0.000 n=10+10)
ColorPrint-4       871ns ± 8%     694ns ± 5%  -20.39%  (p=0.000 n=10+10)
ColorPrint-12      639ns ± 2%     527ns ± 3%  -17.56%  (p=0.000 n=10+10)
ColorString       3.60µs ± 4%    2.97µs ± 2%  -17.65%  (p=0.000 n=10+9)
ColorString-4     1.37µs ±11%    1.16µs ± 8%  -15.70%  (p=0.000 n=10+10)
ColorString-12    1.56µs ±19%    1.20µs ± 3%  -23.04%  (p=0.000 n=10+10)

name            old alloc/op   new alloc/op   delta
NewPrint           85.0B ± 0%     68.0B ± 0%  -20.00%  (p=0.000 n=10+10)
NewPrint-4         85.5B ± 1%     68.0B ± 0%  -20.47%  (p=0.000 n=10+10)
NewPrint-12        86.0B ± 0%     69.0B ± 0%  -19.77%  (p=0.000 n=10+9)
ColorPrint         96.0B ± 0%     68.0B ± 0%  -29.17%  (p=0.000 n=10+10)
ColorPrint-4       96.0B ± 0%     68.0B ± 0%  -29.17%  (p=0.000 n=10+10)
ColorPrint-12      96.0B ± 0%     68.0B ± 0%  -29.17%  (p=0.000 n=10+10)
ColorString       4.72kB ± 0%    4.69kB ± 0%   -0.62%  (p=0.000 n=8+10)
ColorString-4     4.74kB ± 0%    4.71kB ± 0%   -0.60%  (p=0.000 n=9+10)
ColorString-12    4.76kB ± 0%    4.74kB ± 0%   -0.52%  (p=0.000 n=10+9)

name            old allocs/op  new allocs/op  delta
NewPrint            5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.000 n=10+10)
NewPrint-4          5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.000 n=10+10)
NewPrint-12         5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.000 n=10+10)
ColorPrint          6.00 ± 0%      4.00 ± 0%  -33.33%  (p=0.000 n=10+10)
ColorPrint-4        6.00 ± 0%      4.00 ± 0%  -33.33%  (p=0.000 n=10+10)
ColorPrint-12       6.00 ± 0%      4.00 ± 0%  -33.33%  (p=0.000 n=10+10)
ColorString         9.00 ± 0%      6.00 ± 0%  -33.33%  (p=0.000 n=10+10)
ColorString-4       9.00 ± 0%      6.00 ± 0%  -33.33%  (p=0.000 n=10+10)
ColorString-12      9.00 ± 0%      6.00 ± 0%  -33.33%  (p=0.000 n=10+10)
```

## Todo

- [ ] Save/Return previous values
- [ ] Evaluate fmt.Formatter interface


## Credits

- [Fatih Arslan](https://github.com/fatih)
- Windows support via @mattn: [colorable](https://github.com/mattn/go-colorable)
- 2018- The color Authors.

## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/zchee/color/blob/master/LICENSE.md) for more details


<!-- badge links -->
[circleci]: https://circleci.com/gh/zchee/workflows/color
[codecov]: https://codecov.io/gh/zchee/color
[godoc]: https://godoc.org/github.com/zchee/color
[tag]: https://github.com/zchee/color/releases
[ga]: https://github.com/zchee/color

[circleci-badge]: https://img.shields.io/circleci/project/github/zchee/color/master.svg?style=for-the-badge&label=CIRCLECI&logo=circleci
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=for-the-badge&label=GODOC.ORG&logoWidth=25&logo=data%3Aimage%2Fsvg%2Bxml%3Bcharset%3Dutf-8%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgdmlld0JveD0iODUgNTUgMTIwIDEyMCI+PHBhdGggZmlsbD0iIzJEQkNBRiIgZD0iTTQwLjIgMTAxLjFjLS40IDAtLjUtLjItLjMtLjVsMi4xLTIuN2MuMi0uMy43LS41IDEuMS0uNWgzNS43Yy40IDAgLjUuMy4zLjZsLTEuNyAyLjZjLS4yLjMtLjcuNi0xIC42bC0zNi4yLS4xek0yNS4xIDExMC4zYy0uNCAwLS41LS4yLS4zLS41bDIuMS0yLjdjLjItLjMuNy0uNSAxLjEtLjVoNDUuNmMuNCAwIC42LjMuNS42bC0uOCAyLjRjLS4xLjQtLjUuNi0uOS42bC00Ny4zLjF6TTQ5LjMgMTE5LjVjLS40IDAtLjUtLjMtLjMtLjZsMS40LTIuNWMuMi0uMy42LS42IDEtLjZoMjBjLjQgMCAuNi4zLjYuN2wtLjIgMi40YzAgLjQtLjQuNy0uNy43bC0yMS44LS4xek0xNTMuMSA5OS4zYy02LjMgMS42LTEwLjYgMi44LTE2LjggNC40LTEuNS40LTEuNi41LTIuOS0xLTEuNS0xLjctMi42LTIuOC00LjctMy44LTYuMy0zLjEtMTIuNC0yLjItMTguMSAxLjUtNi44IDQuNC0xMC4zIDEwLjktMTAuMiAxOSAuMSA4IDUuNiAxNC42IDEzLjUgMTUuNyA2LjguOSAxMi41LTEuNSAxNy02LjYuOS0xLjEgMS43LTIuMyAyLjctMy43aC0xOS4zYy0yLjEgMC0yLjYtMS4zLTEuOS0zIDEuMy0zLjEgMy43LTguMyA1LjEtMTAuOS4zLS42IDEtMS42IDIuNS0xLjZoMzYuNGMtLjIgMi43LS4yIDUuNC0uNiA4LjEtMS4xIDcuMi0zLjggMTMuOC04LjIgMTkuNi03LjIgOS41LTE2LjYgMTUuNC0yOC41IDE3LTkuOCAxLjMtMTguOS0uNi0yNi45LTYuNi03LjQtNS42LTExLjYtMTMtMTIuNy0yMi4yLTEuMy0xMC45IDEuOS0yMC43IDguNS0yOS4zIDcuMS05LjMgMTYuNS0xNS4yIDI4LTE3LjMgOS40LTEuNyAxOC40LS42IDI2LjUgNC45IDUuMyAzLjUgOS4xIDguMyAxMS42IDE0LjEuNi45LjIgMS40LTEgMS43eiIvPjxwYXRoIGZpbGw9IiMyREJDQUYiIGQ9Ik0xODYuMiAxNTQuNmMtOS4xLS4yLTE3LjQtMi44LTI0LjQtOC44LTUuOS01LjEtOS42LTExLjYtMTAuOC0xOS4zLTEuOC0xMS4zIDEuMy0yMS4zIDguMS0zMC4yIDcuMy05LjYgMTYuMS0xNC42IDI4LTE2LjcgMTAuMi0xLjggMTkuOC0uOCAyOC41IDUuMSA3LjkgNS40IDEyLjggMTIuNyAxNC4xIDIyLjMgMS43IDEzLjUtMi4yIDI0LjUtMTEuNSAzMy45LTYuNiA2LjctMTQuNyAxMC45LTI0IDEyLjgtMi43LjUtNS40LjYtOCAuOXptMjMuOC00MC40Yy0uMS0xLjMtLjEtMi4zLS4zLTMuMy0xLjgtOS45LTEwLjktMTUuNS0yMC40LTEzLjMtOS4zIDIuMS0xNS4zIDgtMTcuNSAxNy40LTEuOCA3LjggMiAxNS43IDkuMiAxOC45IDUuNSAyLjQgMTEgMi4xIDE2LjMtLjYgNy45LTQuMSAxMi4yLTEwLjUgMTIuNy0xOS4xeiIvPjwvc3ZnPg==
[codecov-badge]: https://img.shields.io/codecov/c/github/zchee/color/master.svg?style=for-the-badge&logo=codecov&cacheSeconds=600
[tag-badge]: https://img.shields.io/github/tag/zchee/color.svg?style=for-the-badge&logo=github?cacheSeconds=600
[ga-badge]: https://gh-ga-beacon.appspot.com/UA-89201129-1/zchee/color?useReferer&pixel
