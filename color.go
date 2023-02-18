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

func Reset(source string) string {
	return colors.Reset(source)
}

func Bold(source string) string {
	return colors.Bold(source)
}

func Dim(source string) string {
	return colors.Dim(source)
}

func Italic(source string) string {
	return colors.Italic(source)
}

func Underline(source string) string {
	return colors.Underline(source)
}

func Inverse(source string) string {
	return colors.Inverse(source)
}

func Hidden(source string) string {
	return colors.Hidden(source)
}

func Strikethrough(source string) string {
	return colors.Strikethrough(source)
}

func Black(source string) string {
	return colors.Black(source)
}

func Red(source string) string {
	return colors.Red(source)
}

func Green(source string) string {
	return colors.Green(source)
}

func Yellow(source string) string {
	return colors.Yellow(source)
}

func Blue(source string) string {
	return colors.Blue(source)
}

func Magenta(source string) string {
	return colors.Magenta(source)
}

func Cyan(source string) string {
	return colors.Cyan(source)
}

func White(source string) string {
	return colors.White(source)
}

func Gray(source string) string {
	return colors.Gray(source)
}

func BgBlack(source string) string {
	return colors.BgBlack(source)
}

func BgRed(source string) string {
	return colors.BgRed(source)
}

func BgGreen(source string) string {
	return colors.BgGreen(source)
}

func BgYellow(source string) string {
	return colors.BgYellow(source)
}

func BgBlue(source string) string {
	return colors.BgBlue(source)
}

func BgMagenta(source string) string {
	return colors.BgMagenta(source)
}

func BgCyan(source string) string {
	return colors.BgCyan(source)
}

func BgWhite(source string) string {
	return colors.BgWhite(source)
}

func BlackBright(source string) string {
	return colors.BlackBright(source)
}

func RedBright(source string) string {
	return colors.RedBright(source)
}

func GreenBright(source string) string {
	return colors.GreenBright(source)
}

func YellowBright(source string) string {
	return colors.YellowBright(source)
}

func BlueBright(source string) string {
	return colors.BlueBright(source)
}

func MagentaBright(source string) string {
	return colors.MagentaBright(source)
}

func CyanBright(source string) string {
	return colors.CyanBright(source)
}

func WhiteBright(source string) string {
	return colors.WhiteBright(source)
}

func BgBlackBright(source string) string {
	return colors.BgBlackBright(source)
}

func BgRedBright(source string) string {
	return colors.BgRedBright(source)
}

func BgGreenBright(source string) string {
	return colors.BgGreenBright(source)
}

func BgYellowBright(source string) string {
	return colors.BgYellowBright(source)
}

func BgBlueBright(source string) string {
	return colors.BgBlueBright(source)
}

func BgMagentaBright(source string) string {
	return colors.BgMagentaBright(source)
}

func BgCyanBright(source string) string {
	return colors.BgCyanBright(source)
}

func BgWhiteBright(source string) string {
	return colors.BgWhiteBright(source)
}
