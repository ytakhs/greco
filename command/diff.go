package command

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/go-github/github"
	client "github.com/jit-y/greco/github"
)

type Diff struct {
	name     string
	out, err io.Writer
}

var diffUsage = `
Usage: greco diff [options] <owner> <repo> <from> <to>
`

func NewDiffCommand(name string, out, err io.Writer) (*Diff, error) {
	return &Diff{
		name: name,
		out:  out,
		err:  err,
	}, nil
}

// Synopsis returns a description of diff command.
func (c *Diff) Synopsis() string {
	return "Tag"
}

func (c *Diff) Help() string {
	return diffUsage
}

func (c *Diff) Run(args []string) int {
	var (
		token    string
		onlyName bool
	)

	flags := flag.NewFlagSet(c.name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(flags.Output(), c.Help())
		flags.PrintDefaults()
	}

	flags.StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "")
	flags.StringVar(&token, "t", os.Getenv("GITHUB_TOKEN"), "")
	flags.BoolVar(&onlyName, "only-name", false, "")

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

	gh, err := client.NewClient(owner, repo, token)

	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	comparison, err := gh.Compare(from, to)
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	files := comparison.Files

	var fn func(file github.CommitFile) string

	if onlyName {
		fn = func(file github.CommitFile) string { return file.GetFilename() }
	} else {
		fn = func(file github.CommitFile) string { return file.GetPatch() }
	}

	output(&c.out, &files, fn)

	return 0
}

func output(out *io.Writer, files *[]github.CommitFile, fn func(file github.CommitFile) string) {
	for _, file := range *files {
		fmt.Fprintln(*out, fn(file))
	}
}
