package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client contains github.Client
type Client struct {
	owner  string
	repo   string
	GitHub *github.Client
}

type ClientListOptions struct {
	ListOptions *github.ListOptions
}

func NewClient(owner string, repo string, token string) (*Client, error) {
	var gh *github.Client

	if token == "" {
		fmt.Println("without authentication...")
		gh = github.NewClient(nil)
	} else {
		fmt.Println("with authentication...")
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		gh = github.NewClient(tc)
	}

	client := &Client{
		owner:  owner,
		repo:   repo,
		GitHub: gh,
	}

	return client, nil
}

func NewListOptions(page int, per_page int) (*ClientListOptions, error) {
	listOptions := &github.ListOptions{
		Page:    page,
		PerPage: per_page,
	}
	return &ClientListOptions{
		ListOptions: listOptions,
	}, nil
}
