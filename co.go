package co

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	tty "github.com/mattn/go-isatty"
)

const isWindows = runtime.GOOS == "windows"

var (
	isDisabled           = hasEnv("NO_COLOR") || contains(os.Args, "--no-color")
	isForced             = hasEnv("FORCE_COLOR") || contains(os.Args, "--color")
	isDumbTerminal       = os.Getenv("TERM") == "dumb"
	isCompatibleTerminal = tty.IsTerminal(os.Stdin.Fd()) && !isDumbTerminal
	isCI                 = hasEnv("CI", "GITHUB_ACTIONS", "GITLAB_CI", "CIRCLECI")
	isColorSupported     = !isDisabled &&
		(isForced || (isWindows && !isDumbTerminal) || isCompatibleTerminal || isCI)
)

func replaceClose(i int, str string, close string, replace string) string {
	head := substring(str, 0, i) + replace
	tail := substring(str, i+len(close), len(str)-1)
	next := strings.Index(tail, close)
	if next < 0 {
		return head + tail
	}
	return head + replaceClose(next, tail, close, replace)
}

func clearBleed(i int, str string, open string, close string, replace string) string {
	if i < 0 {
		return open + str + close
	}
	return open + replaceClose(i, str, close, replace) + close
}

func filterEmpty(open string, close string, replace string) ColorFunc {
	return func(input string) string {
		if isEmpty(input) {
			return ""
		}
		if isEmpty(replace) {
			replace = open
		}
		return clearBleed(indexOf(input, close, len(open)), input, open, close, replace)
	}
}

func buildWithReplace(open int, close int, replace string) ColorFunc {
	return filterEmpty(fmt.Sprintf("\x1b[%vm", open), fmt.Sprintf("\x1b[%vm", close), replace)
}

func build(open int, close int) ColorFunc {
	return buildWithReplace(open, close, "")
}

type CreateColorsOption struct {
	useColor bool
}

type ColorFunc func(string) string

func CreateColors(option CreateColorsOption) Colors {
	if option.useColor {
		return colors
	} else {
		return Colors{}
	}
}

type Colors struct {
	Reset           ColorFunc
	Bold            ColorFunc
	Dim             ColorFunc
	Italic          ColorFunc
	Underline       ColorFunc
	Inverse         ColorFunc
	Hidden          ColorFunc
	Strikethrough   ColorFunc
	Black           ColorFunc
	Red             ColorFunc
	Green           ColorFunc
	Yellow          ColorFunc
	Blue            ColorFunc
	Magenta         ColorFunc
	Cyan            ColorFunc
	White           ColorFunc
	Gray            ColorFunc
	BgBlack         ColorFunc
	BgRed           ColorFunc
	BgGreen         ColorFunc
	BgYellow        ColorFunc
	BgBlue          ColorFunc
	BgMagenta       ColorFunc
	BgCyan          ColorFunc
	BgWhite         ColorFunc
	BlackBright     ColorFunc
	RedBright       ColorFunc
	GreenBright     ColorFunc
	YellowBright    ColorFunc
	BlueBright      ColorFunc
	MagentaBright   ColorFunc
	CyanBright      ColorFunc
	WhiteBright     ColorFunc
	BgBlackBright   ColorFunc
	BgRedBright     ColorFunc
	BgGreenBright   ColorFunc
	BgYellowBright  ColorFunc
	BgBlueBright    ColorFunc
	BgMagentaBright ColorFunc
	BgCyanBright    ColorFunc
	BgWhiteBright   ColorFunc
}

var colors = Colors{
	Reset,
	Bold,
	Dim,
	Italic,
	Underline,
	Inverse,
	Hidden,
	Strikethrough,
	Black,
	Red,
	Green,
	Yellow,
	Blue,
	Magenta,
	Cyan,
	White,
	Gray,
	BgBlack,
	BgRed,
	BgGreen,
	BgYellow,
	BgBlue,
	BgMagenta,
	BgCyan,
	BgWhite,
	BlackBright,
	RedBright,
	GreenBright,
	YellowBright,
	BlueBright,
	MagentaBright,
	CyanBright,
	WhiteBright,
	BgBlackBright,
	BgRedBright,
	BgGreenBright,
	BgYellowBright,
	BgBlueBright,
	BgMagentaBright,
	BgCyanBright,
	BgWhiteBright,
}

var Reset = build(0, 0)
var Bold = buildWithReplace(1, 22, "\x1b[22m\x1b[1m")
var Dim = buildWithReplace(2, 22, "\x1b[22m\x1b[2m")
var Italic = build(3, 23)
var Underline = build(4, 24)
var Inverse = build(7, 27)
var Hidden = build(8, 28)
var Strikethrough = build(9, 29)
var Black = build(30, 39)
var Red = build(31, 39)
var Green = build(32, 39)
var Yellow = build(33, 39)
var Blue = build(34, 39)
var Magenta = build(35, 39)
var Cyan = build(36, 39)
var White = build(37, 39)
var Gray = build(90, 39)
var BgBlack = build(40, 49)
var BgRed = build(41, 49)
var BgGreen = build(42, 49)
var BgYellow = build(43, 49)
var BgBlue = build(44, 49)
var BgMagenta = build(45, 49)
var BgCyan = build(46, 49)
var BgWhite = build(47, 49)
var BlackBright = build(90, 39)
var RedBright = build(91, 39)
var GreenBright = build(92, 39)
var YellowBright = build(93, 39)
var BlueBright = build(94, 39)
var MagentaBright = build(95, 39)
var CyanBright = build(96, 39)
var WhiteBright = build(97, 39)
var BgBlackBright = build(100, 49)
var BgRedBright = build(101, 49)
var BgGreenBright = build(102, 49)
var BgYellowBright = build(103, 49)
var BgBlueBright = build(104, 49)
var BgMagentaBright = build(105, 49)
var BgCyanBright = build(106, 49)
var BgWhiteBright = build(107, 49)
