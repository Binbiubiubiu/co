package co

type Style struct {
	_reset           StyleFunc
	_bold            StyleFunc
	_dim             StyleFunc
	_italic          StyleFunc
	_underline       StyleFunc
	_inverse         StyleFunc
	_hidden          StyleFunc
	_strikethrough   StyleFunc
	_black           StyleFunc
	_red             StyleFunc
	_green           StyleFunc
	_yellow          StyleFunc
	_blue            StyleFunc
	_magenta         StyleFunc
	_cyan            StyleFunc
	_white           StyleFunc
	_gray            StyleFunc
	_bgBlack         StyleFunc
	_bgRed           StyleFunc
	_bgGreen         StyleFunc
	_bgYellow        StyleFunc
	_bgBlue          StyleFunc
	_bgMagenta       StyleFunc
	_bgCyan          StyleFunc
	_bgWhite         StyleFunc
	_blackBright     StyleFunc
	_redBright       StyleFunc
	_greenBright     StyleFunc
	_yellowBright    StyleFunc
	_blueBright      StyleFunc
	_magentaBright   StyleFunc
	_cyanBright      StyleFunc
	_whiteBright     StyleFunc
	_bgBlackBright   StyleFunc
	_bgRedBright     StyleFunc
	_bgGreenBright   StyleFunc
	_bgYellowBright  StyleFunc
	_bgBlueBright    StyleFunc
	_bgMagentaBright StyleFunc
	_bgCyanBright    StyleFunc
	_bgWhiteBright   StyleFunc
}

func (s *Style) Reset(input string) string {
	return s._reset(input)
}

func (s *Style) Bold(input string) string {
	return s._bold(input)
}

func (s *Style) Dim(input string) string {
	return s._dim(input)
}

func (s *Style) Italic(input string) string {
	return s._italic(input)
}

func (s *Style) Underline(input string) string {
	return s._underline(input)
}

func (s *Style) Inverse(input string) string {
	return s._inverse(input)
}

func (s *Style) Hidden(input string) string {
	return s._hidden(input)
}

func (s *Style) Strikethrough(input string) string {
	return s._strikethrough(input)
}

func (s *Style) Black(input string) string {
	return s._black(input)
}

func (s *Style) Red(input string) string {
	return s._red(input)
}

func (s *Style) Green(input string) string {
	return s._green(input)
}

func (s *Style) Yellow(input string) string {
	return s._yellow(input)
}

func (s *Style) Blue(input string) string {
	return s._blue(input)
}

func (s *Style) Magenta(input string) string {
	return s._magenta(input)
}

func (s *Style) Cyan(input string) string {
	return s._cyan(input)
}

func (s *Style) White(input string) string {
	return s._white(input)
}

func (s *Style) Gray(input string) string {
	return s._gray(input)
}

func (s *Style) BgBlack(input string) string {
	return s._bgBlack(input)
}

func (s *Style) BgRed(input string) string {
	return s._bgRed(input)
}

func (s *Style) BgGreen(input string) string {
	return s._bgGreen(input)
}

func (s *Style) BgYellow(input string) string {
	return s._bgYellow(input)
}

func (s *Style) BgBlue(input string) string {
	return s._bgBlue(input)
}

func (s *Style) BgMagenta(input string) string {
	return s._bgMagenta(input)
}

func (s *Style) BgCyan(input string) string {
	return s._bgCyan(input)
}

func (s *Style) BgWhite(input string) string {
	return s._bgWhite(input)
}

func (s *Style) BlackBright(input string) string {
	return s._blackBright(input)
}

func (s *Style) RedBright(input string) string {
	return s._redBright(input)
}

func (s *Style) GreenBright(input string) string {
	return s._greenBright(input)
}

func (s *Style) YellowBright(input string) string {
	return s._yellowBright(input)
}

func (s *Style) BlueBright(input string) string {
	return s._blueBright(input)
}

func (s *Style) MagentaBright(input string) string {
	return s._magentaBright(input)
}

func (s *Style) CyanBright(input string) string {
	return s._cyanBright(input)
}

func (s *Style) WhiteBright(input string) string {
	return s._whiteBright(input)
}

func (s *Style) BgBlackBright(input string) string {
	return s._bgBlackBright(input)
}

func (s *Style) BgRedBright(input string) string {
	return s._bgRedBright(input)
}

func (s *Style) BgGreenBright(input string) string {
	return s._bgGreenBright(input)
}

func (s *Style) BgYellowBright(input string) string {
	return s._bgYellowBright(input)
}

func (s *Style) BgBlueBright(input string) string {
	return s._bgBlueBright(input)
}

func (s *Style) BgMagentaBright(input string) string {
	return s._bgMagentaBright(input)
}

func (s *Style) BgCyanBright(input string) string {
	return s._bgCyanBright(input)
}

func (s *Style) BgWhiteBright(input string) string {
	return s._bgWhiteBright(input)
}

// Create a style utility function with isSupportColor
func UseStyles(isColorSupported bool) Style {
	if isColorSupported {
		return Style{
			_reset:           build(0, 0),
			_bold:            buildWithReplace(1, 22, "\x1b[22m\x1b[1m"),
			_dim:             buildWithReplace(2, 22, "\x1b[22m\x1b[2m"),
			_italic:          build(3, 23),
			_underline:       build(4, 24),
			_inverse:         build(7, 27),
			_hidden:          build(8, 28),
			_strikethrough:   build(9, 29),
			_black:           build(30, 39),
			_red:             build(31, 39),
			_green:           build(32, 39),
			_yellow:          build(33, 39),
			_blue:            build(34, 39),
			_magenta:         build(35, 39),
			_cyan:            build(36, 39),
			_white:           build(37, 39),
			_gray:            build(90, 39),
			_bgBlack:         build(40, 49),
			_bgRed:           build(41, 49),
			_bgGreen:         build(42, 49),
			_bgYellow:        build(43, 49),
			_bgBlue:          build(44, 49),
			_bgMagenta:       build(45, 49),
			_bgCyan:          build(46, 49),
			_bgWhite:         build(47, 49),
			_blackBright:     build(90, 39),
			_redBright:       build(91, 39),
			_greenBright:     build(92, 39),
			_yellowBright:    build(93, 39),
			_blueBright:      build(94, 39),
			_magentaBright:   build(95, 39),
			_cyanBright:      build(96, 39),
			_whiteBright:     build(97, 39),
			_bgBlackBright:   build(100, 49),
			_bgRedBright:     build(101, 49),
			_bgGreenBright:   build(102, 49),
			_bgYellowBright:  build(103, 49),
			_bgBlueBright:    build(104, 49),
			_bgMagentaBright: build(105, 49),
			_bgCyanBright:    build(106, 49),
			_bgWhiteBright:   build(107, 49),
		}
	} else {
		return Style{
			_reset:           noop,
			_bold:            noop,
			_dim:             noop,
			_italic:          noop,
			_underline:       noop,
			_inverse:         noop,
			_hidden:          noop,
			_strikethrough:   noop,
			_black:           noop,
			_red:             noop,
			_green:           noop,
			_yellow:          noop,
			_blue:            noop,
			_magenta:         noop,
			_cyan:            noop,
			_white:           noop,
			_gray:            noop,
			_bgBlack:         noop,
			_bgRed:           noop,
			_bgGreen:         noop,
			_bgYellow:        noop,
			_bgBlue:          noop,
			_bgMagenta:       noop,
			_bgCyan:          noop,
			_bgWhite:         noop,
			_blackBright:     noop,
			_redBright:       noop,
			_greenBright:     noop,
			_yellowBright:    noop,
			_blueBright:      noop,
			_magentaBright:   noop,
			_cyanBright:      noop,
			_whiteBright:     noop,
			_bgBlackBright:   noop,
			_bgRedBright:     noop,
			_bgGreenBright:   noop,
			_bgYellowBright:  noop,
			_bgBlueBright:    noop,
			_bgMagentaBright: noop,
			_bgCyanBright:    noop,
			_bgWhiteBright:   noop,
		}
	}
}

var _styles = CreateStyles()

func Reset(input string) string {
	return _styles.Reset(input)
}

func Bold(input string) string {
	return _styles.Bold(input)
}

func Dim(input string) string {
	return _styles.Dim(input)
}

func Italic(input string) string {
	return _styles.Italic(input)
}

func Underline(input string) string {
	return _styles.Underline(input)
}

func Inverse(input string) string {
	return _styles.Inverse(input)
}

func Hidden(input string) string {
	return _styles.Hidden(input)
}

func Strikethrough(input string) string {
	return _styles.Strikethrough(input)
}

func Black(input string) string {
	return _styles.Black(input)
}

func Red(input string) string {
	return _styles.Red(input)
}

func Green(input string) string {
	return _styles.Green(input)
}

func Yellow(input string) string {
	return _styles.Yellow(input)
}

func Blue(input string) string {
	return _styles.Blue(input)
}

func Magenta(input string) string {
	return _styles.Magenta(input)
}

func Cyan(input string) string {
	return _styles.Cyan(input)
}

func White(input string) string {
	return _styles.White(input)
}

func Gray(input string) string {
	return _styles.Gray(input)
}

func BgBlack(input string) string {
	return _styles.BgBlack(input)
}

func BgRed(input string) string {
	return _styles.BgRed(input)
}

func BgGreen(input string) string {
	return _styles.BgGreen(input)
}

func BgYellow(input string) string {
	return _styles.BgYellow(input)
}

func BgBlue(input string) string {
	return _styles.BgBlue(input)
}

func BgMagenta(input string) string {
	return _styles.BgMagenta(input)
}

func BgCyan(input string) string {
	return _styles.BgCyan(input)
}

func BgWhite(input string) string {
	return _styles.BgWhite(input)
}

func BlackBright(input string) string {
	return _styles.BlackBright(input)
}

func RedBright(input string) string {
	return _styles.RedBright(input)
}

func GreenBright(input string) string {
	return _styles.GreenBright(input)
}

func YellowBright(input string) string {
	return _styles.YellowBright(input)
}

func BlueBright(input string) string {
	return _styles.BlueBright(input)
}

func MagentaBright(input string) string {
	return _styles.MagentaBright(input)
}

func CyanBright(input string) string {
	return _styles.CyanBright(input)
}

func WhiteBright(input string) string {
	return _styles.WhiteBright(input)
}

func BgBlackBright(input string) string {
	return _styles.BgBlackBright(input)
}

func BgRedBright(input string) string {
	return _styles.BgRedBright(input)
}

func BgGreenBright(input string) string {
	return _styles.BgGreenBright(input)
}

func BgYellowBright(input string) string {
	return _styles.BgYellowBright(input)
}

func BgBlueBright(input string) string {
	return _styles.BgBlueBright(input)
}

func BgMagentaBright(input string) string {
	return _styles.BgMagentaBright(input)
}

func BgCyanBright(input string) string {
	return _styles.BgCyanBright(input)
}

func BgWhiteBright(input string) string {
	return _styles.BgWhiteBright(input)
}
