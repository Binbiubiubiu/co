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

## Bench

```sh
$ go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/Binbiubiubiu/co
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
Benchmark_co-8                   1986559               606.8 ns/op
Benchmark_fatih_color-8           534586              2065 ns/op
Benchmark_gookit_color-8         1463462               804.2 ns/op
PASS
ok      github.com/Binbiubiubiu/co      6.429s
```

## Thanks

[colorette](https://github.com/jorgebucaran/colorette)  🌈Easily set your terminal text color & styles.