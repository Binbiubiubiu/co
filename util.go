package co

import (
	"os"
	"unicode/utf8"
)

func contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

func isEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

func substringNoEnd(str string, s int) string {
	if s <= 0 {
		return str
	}
	return string([]rune(str)[s:])
}

func substring(str string, s int, e int) string {
	n := lenRune(str)
	if s <= 0 && e >= n {
		return str
	}

	if s > e {
		s, e = e, s
	}

	if s < 0 {
		s = 0
	}
	if e > n {
		e = n
	}

	l := len(str)
	if l == n {
		return str[s:e]
	}

	return string([]rune(str)[s:e])
}

func hasEnv(vars ...string) bool {
	for _, v := range vars {
		_, exist := os.LookupEnv(v)
		if exist {
			return true
		}
	}
	return false
}

func lenRune(s string) int {
	return utf8.RuneCountInString(s)
}

func indexOf(source string, substr string, args ...int) int {
	at := 0
	if len(args) > 0 {
		at = args[0]
	}

	input := []rune(source)
	iLen := len(input)
	if at >= iLen {
		return -1
	}

	input = input[at:]
	iLen = len(input)

	search := []rune(substr)
	sLen := len(search)

pin:
	for i := 0; i+sLen <= iLen; i++ {
		for j := 0; j < sLen; j++ {
			if input[i+j] != search[j] {
				continue pin
			}
		}
		return i + at
	}
	return -1
}
