package command

import (
	"flag"
	"fmt"
	"io"
	"os"

	client "github.com/jit-y/greco/github"
)

var tagsUsage = `
Usage: greco tags [options] <owner> <repo>
`

type Tags struct {
	name     string
	out, err io.Writer
}

func NewTagsCommand(name string, out, err io.Writer) (*Tags, error) {
	return &Tags{
		name: name,
		out:  out,
		err:  err,
	}, nil
}

func (c *Tags) Synopsis() string {
	return "Tag"
}

func (c *Tags) Help() string {
	return tagsUsage
}

func (c *Tags) Run(args []string) int {
	var (
		token string
		page  int
		per   int
	)

	flags := flag.NewFlagSet(c.name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(flags.Output(), c.Help())
		flags.PrintDefaults()
	}

	flags.StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "")
	flags.StringVar(&token, "t", os.Getenv("GITHUB_TOKEN"), "")
	flags.IntVar(&page, "page", 1, "")
	flags.IntVar(&per, "per", 10, "")
	flags.IntVar(&per, "p", 10, "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	cmdArgs := flags.Args()
	if len(cmdArgs) < 2 {
		fmt.Fprintln(flags.Output(), c.Help())
		flags.PrintDefaults()
		return 1
	}

	owner := cmdArgs[0]
	repo := cmdArgs[1]

	github, err := client.NewClient(owner, repo, token)

	if err != nil {
		fmt.Fprintf(c.err, "%s\n", err)
		return 1
	}

	tags, err := github.Tags(per, page)
	if err != nil {
		fmt.Fprintf(c.err, "%s\n", err)
		return 1
	}

	for _, tag := range tags {
		fmt.Fprintln(c.out, tag.GetName())
	}

	return 0
}
