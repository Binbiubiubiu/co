package main

import (
	"fmt"

	co "github.com/Binbiubiubiu/co"
)

func main() {
	println(co.GreenBright("Hello"), co.BgCyanBright(co.RedBright("Binbiubiubiu")))
	var a = co.UseColors(true)
	fmt.Printf("%v\n", a.BgBlueBright("323"))
}
