package shell

import "context"

type Shell interface {
	Run(context.Context, string, ...string) error
	Output() string
}
