package co

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	tty "github.com/mattn/go-isatty"
)

const isWindows = runtime.GOOS == "windows"

var (
	isDisabled           = hasEnv("NO_COLOR") || contains(os.Args, "--no-color")
	isForced             = hasEnv("FORCE_COLOR") || contains(os.Args, "--color")
	isDumbTerminal       = os.Getenv("TERM") == "dumb"
	isCompatibleTerminal = tty.IsTerminal(os.Stdin.Fd()) && !isDumbTerminal
	isCI                 = hasEnv("CI", "GITHUB_ACTIONS", "GITLAB_CI", "CIRCLECI")
	IsColorSupported     = !isDisabled &&
		(isForced || (isWindows && !isDumbTerminal) || isCompatibleTerminal || isCI)
)

func replaceClose(i int, tail []rune, close []rune, replace string) string {
	lClose := len(close)
	var sb strings.Builder
	for i >= 0 {
		for _, c := range substring(tail, 0, i) {
			sb.WriteRune(c)
		}
		sb.WriteString(replace)
		tail = substringNoEnd(tail, i+lClose)
		i = indexOf(tail, close, 0)
	}
	for _, c := range tail {
		sb.WriteRune(c)
	}
	return sb.String()
}

func clearBleed(i int, source string, open string, close string, replace string) string {
	if i < 0 {
		return open + source + close
	}
	return open + replaceClose(i, []rune(source), []rune(close), replace) + close
}

func filterEmpty(open string, close string, replace string) StyleFunc {
	return func(input string) string {
		if isEmpty(input) {
			return ""
		}
		if isEmpty(replace) {
			replace = open
		}
		return clearBleed(indexOf([]rune(input), []rune(close), uint(lenRune(open))), input, open, close, replace)
	}
}

// Style formatting functions
type StyleFunc func(string) string

func noop(input string) string {
	return input
}

func buildWithReplace(open int, close int, replace string) StyleFunc {
	return filterEmpty(fmt.Sprintf("\x1b[%vm", open), fmt.Sprintf("\x1b[%vm", close), replace)
}

func build(open int, close int) StyleFunc {
	return buildWithReplace(open, close, "")
}

// Create a ansi-style tool function
func CreateStyles() *Style {
	return UseStyles(IsColorSupported)
}

func Compose(styl ...StyleFunc) func(str string) string {
	return func(str string) string {
		for _, c := range styl {
			str = c(str)
		}
		return str
	}
}
