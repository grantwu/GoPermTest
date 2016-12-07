package cli

import (
	"github.com/palantir/go-palantir/cli/flag"
)

type Command struct {
	Name         string
	Alias        string
	Usage        string
	Description  string // prose
	Flags        []flag.Flag
	Subcommands  []Command
	DecisionFlag string
	Action       func(ctx Context) error
}

const DefaultDecision string = ""

func (cmd Command) Names() []string {
	names := []string{}
	if cmd.Name != DefaultDecision {
		names = append(names, cmd.Name)
	}
	if cmd.Alias != "" {
		names = append(names, cmd.Alias)
	}
	return names
}
