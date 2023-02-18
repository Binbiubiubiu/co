package co

import (
	"fmt"
	"testing"

	fatihColor "github.com/fatih/color"
	gookitColor "github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

type expectedColor struct {
	f     ColorFunc
	open  string
	close string
}

var eColor = UseColors(true)
var eColors = []expectedColor{
	{f: eColor.Reset, open: "\x1b[0m", close: "\x1b[0m"},
	{f: eColor.Bold, open: "\x1b[1m", close: "\x1b[22m"},
	{f: eColor.Dim, open: "\x1b[2m", close: "\x1b[22m"},
	{f: eColor.Italic, open: "\x1b[3m", close: "\x1b[23m"},
	{f: eColor.Underline, open: "\x1b[4m", close: "\x1b[24m"},
	{f: eColor.Inverse, open: "\x1b[7m", close: "\x1b[27m"},
	{f: eColor.Hidden, open: "\x1b[8m", close: "\x1b[28m"},
	{f: eColor.Strikethrough, open: "\x1b[9m", close: "\x1b[29m"},
	{f: eColor.Black, open: "\x1b[30m", close: "\x1b[39m"},
	{f: eColor.Red, open: "\x1b[31m", close: "\x1b[39m"},
	{f: eColor.Green, open: "\x1b[32m", close: "\x1b[39m"},
	{f: eColor.Yellow, open: "\x1b[33m", close: "\x1b[39m"},
	{f: eColor.Blue, open: "\x1b[34m", close: "\x1b[39m"},
	{f: eColor.Magenta, open: "\x1b[35m", close: "\x1b[39m"},
	{f: eColor.Cyan, open: "\x1b[36m", close: "\x1b[39m"},
	{f: eColor.White, open: "\x1b[37m", close: "\x1b[39m"},
	{f: eColor.Gray, open: "\x1b[90m", close: "\x1b[39m"},
	{f: eColor.BgBlack, open: "\x1b[40m", close: "\x1b[49m"},
	{f: eColor.BgRed, open: "\x1b[41m", close: "\x1b[49m"},
	{f: eColor.BgGreen, open: "\x1b[42m", close: "\x1b[49m"},
	{f: eColor.BgYellow, open: "\x1b[43m", close: "\x1b[49m"},
	{f: eColor.BgBlue, open: "\x1b[44m", close: "\x1b[49m"},
	{f: eColor.BgMagenta, open: "\x1b[45m", close: "\x1b[49m"},
	{f: eColor.BgCyan, open: "\x1b[46m", close: "\x1b[49m"},
	{f: eColor.BgWhite, open: "\x1b[47m", close: "\x1b[49m"},
	{f: eColor.BlackBright, open: "\x1b[90m", close: "\x1b[39m"},
	{f: eColor.RedBright, open: "\x1b[91m", close: "\x1b[39m"},
	{f: eColor.GreenBright, open: "\x1b[92m", close: "\x1b[39m"},
	{f: eColor.YellowBright, open: "\x1b[93m", close: "\x1b[39m"},
	{f: eColor.BlueBright, open: "\x1b[94m", close: "\x1b[39m"},
	{f: eColor.MagentaBright, open: "\x1b[95m", close: "\x1b[39m"},
	{f: eColor.CyanBright, open: "\x1b[96m", close: "\x1b[39m"},
	{f: eColor.WhiteBright, open: "\x1b[97m", close: "\x1b[39m"},
	{f: eColor.BgBlackBright, open: "\x1b[100m", close: "\x1b[49m"},
	{f: eColor.BgRedBright, open: "\x1b[101m", close: "\x1b[49m"},
	{f: eColor.BgGreenBright, open: "\x1b[102m", close: "\x1b[49m"},
	{f: eColor.BgYellowBright, open: "\x1b[103m", close: "\x1b[49m"},
	{f: eColor.BgBlueBright, open: "\x1b[104m", close: "\x1b[49m"},
	{f: eColor.BgMagentaBright, open: "\x1b[105m", close: "\x1b[49m"},
	{f: eColor.BgCyanBright, open: "\x1b[106m", close: "\x1b[49m"},
	{f: eColor.BgWhiteBright, open: "\x1b[107m", close: "\x1b[49m"},
}

func Test_simple(t *testing.T) {
	s := "test"
	for _, co := range eColors {
		assert.Equal(t, co.f(s), co.open+s+co.close)
	}
}

func Test_nesting(t *testing.T) {
	assert.Equal(t, eColor.Bold("bold "+eColor.Red("red "+eColor.Dim("dim")+" red")+" bold"), "\x1B[1mbold \x1B[31mred \x1B[2mdim\x1B[22m\x1B[1m red\x1B[39m bold\x1B[22m")
	assert.Equal(t, eColor.Magenta("magenta "+eColor.Yellow("yellow "+eColor.Cyan("cyan")+" "+eColor.Red("red")+" "+eColor.Green("green")+" yellow")+" magenta"), "\x1B[35mmagenta \x1B[33myellow \x1B[36mcyan\x1B[33m \x1B[31mred\x1B[33m \x1B[32mgreen\x1B[33m yellow\x1B[35m magenta\x1B[39m")
}

func Test_empty(t *testing.T) {
	assert.Equal(t, "", Bold(""))
}

func Benchmark_co(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s%s%s", eColor.Yellow("yellow"), eColor.Bold(eColor.Cyan("cyan")), eColor.Red("red"))
	}
}

func Benchmark_fatih_color(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s%s%s", fatihColor.YellowString("yellow"), fatihColor.New(fatihColor.FgCyan, fatihColor.Bold).Sprint("cyan"), fatihColor.RedString("red"))
	}
}

func Benchmark_gookit_color(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s%s%s", gookitColor.Yellow.Render("yellow"), gookitColor.Style{gookitColor.FgCyan, gookitColor.OpBold}.Render("cyan"), gookitColor.Red.Render("red"))
	}
}
