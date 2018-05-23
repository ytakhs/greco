package cmd

import (
	"errors"
	"fmt"
	"io"

	client "github.com/jit-y/greco/github"
	"github.com/spf13/cobra"
)

type search struct {
	out, err io.Writer
	token    string
	repo     string
	client   *client.Client
}

func newSearchCmd(out, err io.Writer) *cobra.Command {
	s := search{
		out: out,
		err: err,
	}

	cmd := &cobra.Command{
		Use:   "search [options] <repo name>",
		Short: "search repository",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("command `search` requires <repo>")
			}

			repo := args[0]

			cl, err := client.NewClient("", repo, s.token)
			if err != nil {
				return err
			}

			s.client = cl

			if err = s.run(args); err != nil {
				return err
			}

			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&s.token, "token", "t", "", "github token")

	return cmd
}

func (s *search) run(args []string) error {
	searchResult, err := s.client.SearchRepositories()
	if err != nil {
		return err
	}

	for _, repository := range searchResult.Repositories {
		fmt.Fprintf(s.out, "%s %s\n", repository.GetOwner().GetLogin(), repository.GetName())
	}

	return nil
}
