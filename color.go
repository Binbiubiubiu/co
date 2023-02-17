package co

type Colors struct {
	Reset               ColorFunc
	Bold                ColorFunc
	Dim                 ColorFunc
	Italic              ColorFunc
	Underline           ColorFunc
	Inverse             ColorFunc
	Hidden              ColorFunc
	Strikethrough       ColorFunc
	Black               ColorFunc
	Red                 ColorFunc
	Green               ColorFunc
	Yellow              ColorFunc
	Blue                ColorFunc
	Magenta             ColorFunc
	Cyan                ColorFunc
	White               ColorFunc
	Gray                ColorFunc
	BgBlack             ColorFunc
	BgRed               ColorFunc
	BgGreen             ColorFunc
	BgYellow            ColorFunc
	BgBlue              ColorFunc
	BgMagenta           ColorFunc
	BgCyan              ColorFunc
	BgWhite             ColorFunc
	BlackBright         ColorFunc
	RedBright           ColorFunc
	GreenBright         ColorFunc
	YellowBright        ColorFunc
	BlueBright          ColorFunc
	MagentaBright       ColorFunc
	RedBrCyanBrightight ColorFunc
	WhiteBright         ColorFunc
	BgBlackBright       ColorFunc
	BgRedBright         ColorFunc
	BgGreenBright       ColorFunc
	BgYellowBright      ColorFunc
	BgBlueBright        ColorFunc
	BgMagentaBright     ColorFunc
	BgCyanBright        ColorFunc
	BgWhiteBright       ColorFunc
}

func noop(s string) string {

	return s

}

var noColors = Colors{

	Reset:               noop,
	Bold:                noop,
	Dim:                 noop,
	Italic:              noop,
	Underline:           noop,
	Inverse:             noop,
	Hidden:              noop,
	Strikethrough:       noop,
	Black:               noop,
	Red:                 noop,
	Green:               noop,
	Yellow:              noop,
	Blue:                noop,
	Magenta:             noop,
	Cyan:                noop,
	White:               noop,
	Gray:                noop,
	BgBlack:             noop,
	BgRed:               noop,
	BgGreen:             noop,
	BgYellow:            noop,
	BgBlue:              noop,
	BgMagenta:           noop,
	BgCyan:              noop,
	BgWhite:             noop,
	BlackBright:         noop,
	RedBright:           noop,
	GreenBright:         noop,
	YellowBright:        noop,
	BlueBright:          noop,
	MagentaBright:       noop,
	RedBrCyanBrightight: noop,
	WhiteBright:         noop,
	BgBlackBright:       noop,
	BgRedBright:         noop,
	BgGreenBright:       noop,
	BgYellowBright:      noop,
	BgBlueBright:        noop,
	BgMagentaBright:     noop,
	BgCyanBright:        noop,
	BgWhiteBright:       noop,
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
	RedBrCyanBrightight,
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
var RedBrCyanBrightight = build(96, 39)
var WhiteBright = build(97, 39)
var BgBlackBright = build(100, 49)
var BgRedBright = build(101, 49)
var BgGreenBright = build(102, 49)
var BgYellowBright = build(103, 49)
var BgBlueBright = build(104, 49)
var BgMagentaBright = build(105, 49)
var BgCyanBright = build(106, 49)
var BgWhiteBright = build(107, 49)
