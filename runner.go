package main

import (
	"io"
	"log"
	"os"

	"github.com/jit-y/greco/command"
	"github.com/mitchellh/cli"
)

// Runner is interface of this tool
type Runner struct {
	out, err io.Writer
}

// Run is main interface.
func (r *Runner) Run(args []string) int {
	c := cli.NewCLI(Name, Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"tags": func() (cli.Command, error) {
			return command.NewTagsCommand("tags", r.out, r.err)
		},
		"browse": func() (cli.Command, error) {
			return command.NewBrowseCommand("browse", r.out, r.err)
		},
		"diff": func() (cli.Command, error) {
			return command.NewDiffCommand("diff", r.out, r.err)
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
