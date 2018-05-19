package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"

	client "github.com/jit-y/greco/github"
	"github.com/pkg/browser"
)

type browse struct {
	out, err io.Writer
	token    string
	client   *client.Client
	from     string
	to       string
}

func newBrowseCmd(out, err io.Writer) *cobra.Command {
	b := browse{
		out: out,
		err: err,
	}

	cmd := &cobra.Command{
		Use:     "browse [flags] <owner name> <repo name> <from version> <to version>",
		Aliases: []string{"open"},
		Short:   "open page",
		Long:    "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 4 {
				return errors.New("command `tags` requires <owner> <repo>")
			}

			owner := args[0]
			repo := args[1]
			b.from = args[2]
			b.to = args[3]

			client, err := client.NewClient(owner, repo, b.token)
			if err != nil {
				return err
			}
			b.client = client

			if err := b.run(args); err != nil {
				return err
			}

			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&b.token, "token", "t", "", "github token")

	return cmd
}

func (b *browse) run(args []string) error {
	comparison, err := b.client.Compare(b.from, b.to)
	if err != nil {
		return err
	}

	err = browser.OpenURL(*comparison.HTMLURL)
	if err != nil {
		return err
	}

	return nil
}
