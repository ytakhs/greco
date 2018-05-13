package command

import (
	"flag"
	"fmt"
	"io"
	"os"

	client "github.com/jit-y/greco/github"
	"github.com/pkg/browser"
)

type Browse struct {
	name     string
	out, err io.Writer
}

var browseUsage = `
Usage: greco browse [options] <owner> <repo> <from> <to>
`

func NewBrowseCommand(name string, out, err io.Writer) (*Browse, error) {
	return &Browse{
		name: name,
		out:  out,
		err:  err,
	}, nil
}

func (c *Browse) Synopsis() string {
	return "Browse"
}

func (c *Browse) Help() string {
	return browseUsage
}

func (c *Browse) Run(args []string) int {
	var token string

	flags := flag.NewFlagSet(c.name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(flags.Output(), c.Help())
		flags.PrintDefaults()
	}

	flags.StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "")
	flags.StringVar(&token, "t", os.Getenv("GITHUB_TOKEN"), "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	cmdArgs := flags.Args()
	if len(cmdArgs) < 4 {
		fmt.Fprintln(flags.Output(), c.Help())
		flags.PrintDefaults()
		return 1
	}

	owner := cmdArgs[0]
	repo := cmdArgs[1]
	from := cmdArgs[2]
	to := cmdArgs[3]

	github, err := client.NewClient(owner, repo, token)

	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	comparison, err := github.Compare(from, to)
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	err = browser.OpenURL(*comparison.HTMLURL)
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	return 0
}
