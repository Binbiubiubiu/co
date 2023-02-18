package co

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

func noop(s string) string {
	return s
}

// Create a color utility function with a useColor property
func UseColors(useColor bool) Colors {
	if useColor {
		return Colors{
			Reset:           build(0, 0),
			Bold:            buildWithReplace(1, 22, "\x1b[22m\x1b[1m"),
			Dim:             buildWithReplace(2, 22, "\x1b[22m\x1b[2m"),
			Italic:          build(3, 23),
			Underline:       build(4, 24),
			Inverse:         build(7, 27),
			Hidden:          build(8, 28),
			Strikethrough:   build(9, 29),
			Black:           build(30, 39),
			Red:             build(31, 39),
			Green:           build(32, 39),
			Yellow:          build(33, 39),
			Blue:            build(34, 39),
			Magenta:         build(35, 39),
			Cyan:            build(36, 39),
			White:           build(37, 39),
			Gray:            build(90, 39),
			BgBlack:         build(40, 49),
			BgRed:           build(41, 49),
			BgGreen:         build(42, 49),
			BgYellow:        build(43, 49),
			BgBlue:          build(44, 49),
			BgMagenta:       build(45, 49),
			BgCyan:          build(46, 49),
			BgWhite:         build(47, 49),
			BlackBright:     build(90, 39),
			RedBright:       build(91, 39),
			GreenBright:     build(92, 39),
			YellowBright:    build(93, 39),
			BlueBright:      build(94, 39),
			MagentaBright:   build(95, 39),
			CyanBright:      build(96, 39),
			WhiteBright:     build(97, 39),
			BgBlackBright:   build(100, 49),
			BgRedBright:     build(101, 49),
			BgGreenBright:   build(102, 49),
			BgYellowBright:  build(103, 49),
			BgBlueBright:    build(104, 49),
			BgMagentaBright: build(105, 49),
			BgCyanBright:    build(106, 49),
			BgWhiteBright:   build(107, 49),
		}
	} else {
		return Colors{
			Reset:           noop,
			Bold:            noop,
			Dim:             noop,
			Italic:          noop,
			Underline:       noop,
			Inverse:         noop,
			Hidden:          noop,
			Strikethrough:   noop,
			Black:           noop,
			Red:             noop,
			Green:           noop,
			Yellow:          noop,
			Blue:            noop,
			Magenta:         noop,
			Cyan:            noop,
			White:           noop,
			Gray:            noop,
			BgBlack:         noop,
			BgRed:           noop,
			BgGreen:         noop,
			BgYellow:        noop,
			BgBlue:          noop,
			BgMagenta:       noop,
			BgCyan:          noop,
			BgWhite:         noop,
			BlackBright:     noop,
			RedBright:       noop,
			GreenBright:     noop,
			YellowBright:    noop,
			BlueBright:      noop,
			MagentaBright:   noop,
			CyanBright:      noop,
			WhiteBright:     noop,
			BgBlackBright:   noop,
			BgRedBright:     noop,
			BgGreenBright:   noop,
			BgYellowBright:  noop,
			BgBlueBright:    noop,
			BgMagentaBright: noop,
			BgCyanBright:    noop,
			BgWhiteBright:   noop,
		}
	}
}

var colors = CreateColors()

var Reset = colors.Reset
var Bold = colors.Bold
var Dim = colors.Dim
var Italic = colors.Italic
var Underline = colors.Underline
var Inverse = colors.Inverse
var Hidden = colors.Hidden
var Strikethrough = colors.Strikethrough
var Black = colors.Black
var Red = colors.Red
var Green = colors.Green
var Yellow = colors.Yellow
var Blue = colors.Blue
var Magenta = colors.Magenta
var Cyan = colors.Cyan
var White = colors.White
var Gray = colors.Gray
var BgBlack = colors.BgBlack
var BgRed = colors.BgRed
var BgGreen = colors.BgGreen
var BgYellow = colors.BgYellow
var BgBlue = colors.BgBlue
var BgMagenta = colors.BgMagenta
var BgCyan = colors.BgCyan
var BgWhite = colors.BgWhite
var BlackBright = colors.BlackBright
var RedBright = colors.RedBright
var GreenBright = colors.GreenBright
var YellowBright = colors.YellowBright
var BlueBright = colors.BlueBright
var MagentaBright = colors.MagentaBright
var CyanBright = colors.CyanBright
var WhiteBright = colors.WhiteBright
var BgBlackBright = colors.BgBlackBright
var BgRedBright = colors.BgRedBright
var BgGreenBright = colors.BgGreenBright
var BgYellowBright = colors.BgYellowBright
var BgBlueBright = colors.BgBlueBright
var BgMagentaBright = colors.BgMagentaBright
var BgCyanBright = colors.BgCyanBright
var BgWhiteBright = colors.BgWhiteBright
