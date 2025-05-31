package shell

import (
	"context"
	"os/exec"
)

type DirectShell struct {
	output string
}

func NewDirectShell() Shell {
	return &DirectShell{}
}

func (s *DirectShell) Run(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	s.output = string(out[:])

	return nil
}

func (s *DirectShell) Output() string {
	return s.output
}
