package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/google/go-github/github"
	client "github.com/jit-y/greco/github"
)

type diff struct {
	name     string
	out, err io.Writer
	owner    string
	repo     string
	token    string
	from     string
	to       string
	onlyName bool
}

func newDiffCmd(out, err io.Writer) *cobra.Command {
	d := diff{
		out: out,
		err: err,
	}

	cmd := &cobra.Command{
		Use:   "diff [flags] <owner name> <repo name> <from version> <to version>",
		Short: "diff versions",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 4 {
				return errors.New("command `tags` requires <owner> <repo>")
			}

			d.owner = args[0]
			d.repo = args[1]
			d.from = args[2]
			d.to = args[3]

			if err := d.run(args); err != nil {
				return err
			}

			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&d.token, "token", "t", "", "github token")
	f.BoolVar(&d.onlyName, "only-name", false, "out only file names")

	return cmd
}

// Run shows diff of comparison with to and from.
func (d *diff) run(args []string) error {
	gh, err := client.NewClient(d.owner, d.repo, d.token)
	if err != nil {
		return err
	}

	comparison, err := gh.Compare(d.from, d.to)
	if err != nil {
		return err
	}

	files := comparison.Files

	var fn func(file github.CommitFile) string

	if d.onlyName {
		fn = func(file github.CommitFile) string { return file.GetFilename() }
	} else {
		fn = func(file github.CommitFile) string { return file.GetPatch() }
	}

	output(&d.out, &files, fn)

	return nil
}

func output(out *io.Writer, files *[]github.CommitFile, fn func(file github.CommitFile) string) {
	for _, file := range *files {
		fmt.Fprintln(*out, fn(file))
	}
}
