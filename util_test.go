package co

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_contains(t *testing.T) {
	assert.True(t, contains([]string{"--no-color"}, "--no-color"))
	assert.False(t, contains([]string{}, "--no-color"))
}

func Test_isEmpty(t *testing.T) {
	assert.True(t, isEmpty(""))
	assert.True(t, isEmpty(0))
	assert.True(t, isEmpty(0.0))
}

func Test_substring(t *testing.T) {
	assert.Equal(t, []rune("45"), substring([]rune("123456"), 3, 5))
	assert.Equal(t, []rune("456"), substring([]rune("123456"), 3, 6))
	assert.Equal(t, []rune("456"), substring([]rune("123456"), 3, 7))
	assert.Equal(t, []rune("我的"), substring([]rune("你好！我的太阳"), 3, 5))
	assert.Equal(t, []rune("23"), substring([]rune("123456"), 3, 1))
}

func Test_hasEnv(t *testing.T) {
	os.Setenv("COLOR", "1")
	assert.True(t, hasEnv("COLOR"))
	assert.False(t, hasEnv("NO_COLOR"))
}

func Test_indexOf(t *testing.T) {
	assert.Equal(t, 8, indexOf([]rune("321 123 321"), []rune("321"), 3))
	assert.Equal(t, 0, indexOf([]rune("321 123 321"), []rune("321"), 0))
	assert.Equal(t, 5, indexOf([]rune("你好！我的太阳"), []rune("太阳"), 2))
	assert.Equal(t, -1, indexOf([]rune("123"), []rune("321"), 0))
	assert.Equal(t, 0, indexOf([]rune("123"), []rune("123"), 0))
	assert.Equal(t, 4, indexOf([]rune("我的太阳123"), []rune("123"), 0))
}
