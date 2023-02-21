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

func substringNoEnd(str []rune, s int) []rune {
	if s <= 0 {
		return str
	}
	return str[s:]
}

func substring(str []rune, s int, e int) []rune {
	n := len(str)
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

	return str[s:e]
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

func indexOf(input []rune, search []rune, at uint) int {

	iLen, sLen := len(input), len(search)
	if int(at) >= iLen {
		return -1
	}

	input = input[at:]
	iLen = len(input)

pin:
	for i := 0; i+sLen <= iLen; i++ {
		for j := 0; j < sLen; j++ {
			if input[i+j] != search[j] {
				continue pin
			}
		}
		return i + int(at)
	}
	return -1
}
