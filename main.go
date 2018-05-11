package main

import (
	"os"
)

func main() {
	r := &Runner{
		out: os.Stdout,
		err: os.Stderr,
	}

	os.Exit(r.Run(os.Args))
}
