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

// Run is main interfacen
func (r *Runner) Run(args []string) int {
	c := cli.NewCLI(Name, Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"tags": func() (cli.Command, error) {
			return &command.Tags{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}