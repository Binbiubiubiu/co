# co

The library for terminal output formatting with ANSI colors

## Install

``` sh
$ go get -u github.com/Binbiubiubiu/co
```

## Example

``` go
package main

import (
	. "github.com/Binbiubiubiu/co"
)

func main() {
	println(GreenBright("Hello"), BgCyanBright(RedBright("Binbiubiubiu")))
}

```

## Thanks

[colorette](https://github.com/jorgebucaran/colorette)  🌈Easily set your terminal text color & styles.