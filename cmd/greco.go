package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greco",
		Short: "foo",
		Long:  "foo",
	}

	cmd.AddCommand(
		newTagsCmd(os.Stdout, os.Stderr),
		newBrowseCmd(os.Stdout, os.Stderr),
		newDiffCmd(os.Stdout, os.Stderr),
		newSearchCmd(os.Stdout, os.Stderr),
	)

	return cmd
}
