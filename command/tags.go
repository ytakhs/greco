package command

import (
	"context"
	"flag"
	"fmt"
	"os"

	client "github.com/jit-y/greco/github"
)

var usage = `
Usage: greco tags [options] <owner> <repo>
`

type Tags struct {
	name string
}

func (c *Tags) Synopsis() string {
	return "Tag"
}

func (c *Tags) Help() string {
	return usage
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
	flags.IntVar(&per, "per", 50, "")
	flags.IntVar(&per, "p", 50, "")

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
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	fmt.Println(page)
	fmt.Println(per)
	opt, err := client.NewListOptions(page, per)
	if err != nil {
		return 1
	}

	tags, _, err := github.GitHub.Repositories.ListTags(context.Background(), owner, repo, opt.ListOptions)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	for _, tag := range tags {
		fmt.Println(tag.GetName())
	}

	return 0
}
