package cli

import (
	"io"
	"os"

	"github.com/palantir/go-palantir/cli/completion"
)

type App struct {
	Command
	Before       func(ctx Context) error
	ErrorHandler func(ctx Context, err error) int
	Version      string
	Stdout       io.Writer
	Stderr       io.Writer
	Completion   map[string]completion.Provider
	Manpage      *Manpage
	Backcompat   []Backcompat
}

func NewApp() *App {
	return &App{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
