package main

import (
	co "github.com/Binbiubiubiu/co"
)

func main() {
	println(co.GreenBright("Hello"), co.BgCyanBright(co.RedBright("Binbiubiubiu")))
	println(co.Compose(co.BgGreen, co.Blue)("123"))
}
