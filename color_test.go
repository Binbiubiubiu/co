package co

import (
	"fmt"
	"reflect"
	"testing"

	fatihColor "github.com/fatih/color"
	"github.com/gkampitakis/go-snaps/snaps"
	gookitColor "github.com/gookit/color"
)

type expectedColor struct {
	f     StyleFunc
	open  string
	close string
}

var eColor = UseStyles(true)

func decode(p any) string {
	return fmt.Sprintf("%q", p)
}

func Test_simple(t *testing.T) {
	mock := "test"
	typeStyle := reflect.TypeOf(&eColor)
	fLen := typeStyle.NumMethod()
	for i := 0; i < fLen; i++ {
		method := typeStyle.Method(i)
		if method.IsExported() {
			input := []reflect.Value{reflect.ValueOf(&eColor), reflect.ValueOf(mock)}
			vs := method.Func.Call(input)
			snaps.MatchSnapshot(t, decode(vs[0]))
		}
	}
}

func Test_nesting(t *testing.T) {
	snaps.MatchSnapshot(t, decode(eColor.Bold("bold "+eColor.Red("你好 "+eColor.Dim("🌛")+" 123")+" bold")))
	snaps.MatchSnapshot(t, decode(eColor.Magenta("你好 "+eColor.Yellow("黄色 "+eColor.Cyan("123")+" "+eColor.Red("red")+" "+eColor.Green("green")+" 黄色")+" magenta")))
}

func Test_empty(t *testing.T) {
	snaps.MatchSnapshot(t, Bold(""))
}

func Benchmark_co(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", eColor.Yellow("yellow"), eColor.Bold(eColor.Cyan("cyan")), eColor.Red("red"))
	}
}

func Benchmark_fatih_color(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", fatihColor.YellowString("yellow"), fatihColor.New(fatihColor.FgCyan, fatihColor.Bold).Sprint("cyan"), fatihColor.RedString("red"))
	}
}

func Benchmark_gookit_color(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", gookitColor.Yellow.Render("yellow"), gookitColor.Style{gookitColor.FgCyan, gookitColor.OpBold}.Render("cyan"), gookitColor.Red.Render("red"))
	}
}
