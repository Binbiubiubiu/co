package co

import (
	"fmt"
	"os"
	"runtime"

	tty "github.com/mattn/go-isatty"
)

const IS_WINDOWS = runtime.GOOS == "windows"

var (
	isDisabled           = hasEnv("NO_COLOR") || contains(os.Args, "--no-color")
	isForced             = hasEnv("FORCE_COLOR") || contains(os.Args, "--color")
	isDumbTerminal       = os.Getenv("TERM") == "dumb"
	isCompatibleTerminal = tty.IsTerminal(os.Stdin.Fd()) && !isDumbTerminal
	isCI                 = hasEnv("CI", "GITHUB_ACTIONS", "GITLAB_CI", "CIRCLECI")
	IsColorSupported     = !isDisabled &&
		(isForced || (IS_WINDOWS && !isDumbTerminal) || isCompatibleTerminal || isCI)
)

func replaceClose(i int, source string, close string, replace string) string {
	head := substring(source, 0, i) + replace
	tail := substring(source, i+lenRune(close))
	next := indexOf(tail, close)
	if next < 0 {
		return head + tail
	}
	return head + replaceClose(next, tail, close, replace)
}

func clearBleed(i int, source string, open string, close string, replace string) string {
	if i < 0 {
		return open + source + close
	}
	return open + replaceClose(i, source, close, replace) + close
}

func filterEmpty(open string, close string, replace string) StyleFunc {
	return func(input string) string {
		if isEmpty(input) {
			return ""
		}
		if isEmpty(replace) {
			replace = open
		}
		return clearBleed(indexOf(input, close, lenRune(open)), input, open, close, replace)
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
func CreateStyles() Style {
	return UseStyles(IsColorSupported)
}
