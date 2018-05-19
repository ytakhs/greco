package cmd

import (
	"errors"
	"fmt"
	"io"

	client "github.com/jit-y/greco/github"
	"github.com/spf13/cobra"
)

type tags struct {
	out, err io.Writer
	token    string
	per      int
	page     int
	client   *client.Client
}

func newTagsCmd(out, err io.Writer) *cobra.Command {
	t := tags{
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

			owner := args[0]
			repo := args[1]

			client, err := client.NewClient(owner, repo, t.token)
			if err != nil {
				return err
			}

			t.client = client

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

func (t *tags) run(args []string) error {
	tags, err := t.client.Tags(t.per, t.page)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		fmt.Fprintln(t.out, tag.GetName())
	}

	return nil
}
