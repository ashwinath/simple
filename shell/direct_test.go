package shell

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectShellOK(t *testing.T) {
	exe := NewDirectShell()
	err := exe.Run(context.Background(), "echo", "hello world")
	assert.Nil(t, err)
	assert.Equal(t, "hello world\n", exe.Output())
}

func TestDirectShellError(t *testing.T) {
	exe := NewDirectShell()
	err := exe.Run(context.Background(), "nosuchcommand")
	assert.NotNil(t, err)
}
