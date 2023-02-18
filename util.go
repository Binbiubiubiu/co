package co

import (
	"os"
	"strings"
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

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start > end {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end > length {
		end = length
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
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

func indexOf(input string, substr string, at int) int {
	if at >= len(input) {
		return -1
	}
	input = substring(input, at, len(input))
	i := strings.Index(input, substr)
	if i < 0 {
		return i
	}
	return i + at
}
