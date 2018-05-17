package cmd

import (
	"errors"
	"fmt"
	"io"

	client "github.com/jit-y/greco/github"
	"github.com/spf13/cobra"
)

// Tags struct
type Tags struct {
	out, err io.Writer
	token    string
	per      int
	page     int
	owner    string
	repo     string
}

func newTagsCmd(out, err io.Writer) *cobra.Command {
	t := Tags{
		out: out,
		err: err,
	}

	cmd := &cobra.Command{
		Use:     "tags [flags] <owner name> <repo name>",
		Aliases: []string{"list"},
		Short:   "list tags of owner/repo",
		Long:    "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("command `tags` requires <owner> <repo>")
			}

			t.owner = args[0]
			t.repo = args[1]

			if err := t.run(args); err != nil {
				return err
			}

			return nil
		},
	}

	f := cmd.Flags()
	f.IntVar(&t.page, "page", 1, "page")
	f.IntVarP(&t.per, "per", "p", 10, "per page")
	f.StringVarP(&t.token, "token", "t", "", "github token")

	return cmd
}

func (t *Tags) run(args []string) error {
	github, err := client.NewClient(t.owner, t.repo, t.token)
	if err != nil {
		return err
	}

	tags, err := github.Tags(t.per, t.page)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		fmt.Fprintln(t.out, tag.GetName())
	}

	return nil
}
