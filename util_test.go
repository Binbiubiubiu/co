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
	assert.Equal(t, "45", substring("123456", 3, 5))
	assert.Equal(t, "456", substring("123456", 3, 6))
	assert.Equal(t, "456", substring("123456", 3, 7))
	assert.Equal(t, "", substring("123456", 3, 1))
}

func Test_hasEnv(t *testing.T) {
	os.Setenv("COLOR", "1")
	assert.True(t, hasEnv("COLOR"))
	assert.False(t, hasEnv("NO_COLOR"))
}

func Test_indexOf(t *testing.T) {
	assert.Equal(t, 8, indexOf("321 123 321", "321", 3))
	assert.Equal(t, 0, indexOf("321 123 321", "321", 0))
	assert.Equal(t, -1, indexOf("123", "321", 0))
}
