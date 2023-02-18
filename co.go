package co

import (
	"fmt"
	"os"
	"runtime"
	"strings"

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

func replaceClose(i int, str string, close string, replace string) string {
	head := substring(str, 0, i) + replace
	tail := substring(str, i+len(close), len(str))
	next := strings.Index(tail, close)
	if next < 0 {
		return head + tail
	}
	return head + replaceClose(next, tail, close, replace)
}

func clearBleed(i int, str string, open string, close string, replace string) string {
	if i < 0 {
		return open + str + close
	}
	return open + replaceClose(i, str, close, replace) + close
}

func filterEmpty(open string, close string, replace string) ColorFunc {
	return func(input string) string {
		if isEmpty(input) {
			return ""
		}
		if isEmpty(replace) {
			replace = open
		}
		return clearBleed(indexOf(input, close, len(open)), input, open, close, replace)
	}
}

// Color formatting functions
type ColorFunc func(string) string

func buildWithReplace(open int, close int, replace string) ColorFunc {
	return filterEmpty(fmt.Sprintf("\x1b[%vm", open), fmt.Sprintf("\x1b[%vm", close), replace)
}

func build(open int, close int) ColorFunc {
	return buildWithReplace(open, close, "")
}

// Create a color tool function
func CreateColors() Colors {
	return UseColors(IsColorSupported)
}
