package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachCommand(t *testing.T) {
	withTestPath(t, func(t *testing.T) {
		// Test that any valid dc args pass through.
		t.Run("all args pass through", func(t *testing.T) {
			stdout, stderr, err := runTestCommand(nil, []string{
				"attach",
				"--detach-keys", "alt-q",
				"--no-stdin",
				"--sig-proxy",
				"svc",
			})

			// sorted and normalized
			expOut := `docker
attach
--detach-keys=alt-q
--no-stdin
--sig-proxy
some:svc:cid
`

			assert.Nil(t, err)
			assert.Equal(t, "", stderr)
			assert.Equal(t, expOut, stdout)
		})

		t.Run("no flags", func(t *testing.T) {
			stdout, stderr, err := runTestCommand(nil, []string{"attach", "foo"})

			expOut := `docker
attach
--detach-keys=ctrl-c
--no-stdin=false
--sig-proxy=false
some:foo:cid
`

			assert.Nil(t, err)
			assert.Equal(t, "", stderr)
			assert.Equal(t, expOut, stdout)
		})

		t.Run("index", func(t *testing.T) {
			os.Setenv("MUSS_TEST_PS_SCALE", "1")
			defer os.Unsetenv("MUSS_TEST_PS_SCALE")

			stdout, stderr, err := runTestCommand(nil, []string{"attach", "foo", "--index=1"})

			expOut := `docker
attach
--detach-keys=ctrl-c
--no-stdin=false
--sig-proxy=false
some:foo:cid
`

			assert.Nil(t, err)
			assert.Equal(t, "", stderr)
			assert.Equal(t, expOut, stdout)
		})

		t.Run("index=2", func(t *testing.T) {
			os.Setenv("MUSS_TEST_PS_SCALE", "1")
			defer os.Unsetenv("MUSS_TEST_PS_SCALE")

			stdout, stderr, err := runTestCommand(nil, []string{"attach", "foo", "--index", "2"})

			expOut := `docker
attach
--detach-keys=ctrl-c
--no-stdin=false
--sig-proxy=false
second:foo:cid
`

			assert.Nil(t, err)
			assert.Equal(t, "", stderr)
			assert.Equal(t, expOut, stdout)
		})
	})
}
