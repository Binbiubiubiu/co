package co

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_environment(t *testing.T) {
	type vars struct {
		env  map[string]string
		args []string
	}
	tests := []struct {
		name     string
		vars     vars
		expected string
	}{
		{
			name: "`FORCE_COLOR` forces color",
			vars: vars{
				env: map[string]string{
					"FORCE_COLOR": "1",
				},
				args: []string{},
			},
			expected: "\x1b[34mfoo\x1b[39m\n",
		},
		{
			name: "`NO_COLOR` disables color",
			vars: vars{
				env: map[string]string{
					"NO_COLOR": "1",
				},
				args: []string{"--color"},
			},
			expected: "foo\n",
		},
		{
			name: "`--no-color` disables color",
			vars: vars{
				env: map[string]string{
					"FORCE_COLOR": "1",
				},
				args: []string{"--no-color"},
			},
			expected: "foo\n",
		},
		{
			name: "`--color` enables color",
			vars: vars{
				env:  map[string]string{},
				args: []string{"--color"},
			},
			expected: "\x1b[34mfoo\x1b[39m\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.vars.env {
				os.Setenv(k, v)
			}
			args := append([]string{"run", "./fixtures/foo.go"}, tt.vars.args...)
			cmd := exec.Command("go", args...)
			output, _ := cmd.CombinedOutput()
			assert.Equal(t, tt.expected, string(output))
			for k := range tt.vars.env {
				os.Unsetenv(k)
			}
		})
	}

}

func TestUseStyles(t *testing.T) {
	yes := UseStyles(true)
	no := UseStyles(false)
	assert.Equal(t, "\x1b[34mblue\x1b[39m", yes.Blue("blue"))
	assert.Equal(t, "nope", no.Blue("nope"))
	assert.Equal(t, "", no.Blue(""))
}

func TestCreateColors(t *testing.T) {
	style := CreateStyles()
	if IsColorSupported {
		assert.Equal(t, "\x1b[34mblue\x1b[39m", style.Blue("blue"))
	} else {
		assert.Equal(t, "blue", style.Blue("blue"))
	}
}
