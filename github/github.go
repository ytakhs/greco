package github

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client contains github.Client
type Client struct {
	owner        string
	repo         string
	Repositories *github.RepositoriesService
}

func NewClient(owner string, repo string, token string) (*Client, error) {
	var gh *github.Client

	if token == "" {
		gh = github.NewClient(nil)
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		gh = github.NewClient(tc)
	}

	client := &Client{
		owner:        owner,
		repo:         repo,
		Repositories: gh.Repositories,
	}

	return client, nil
}

func (c *Client) Tags(per int, page int) ([]*github.RepositoryTag, error) {
	opt := &github.ListOptions{
		Page:    page,
		PerPage: per,
	}

	tags, _, err := c.Repositories.ListTags(context.Background(), c.owner, c.repo, opt)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) Compare(from string, to string) (*github.CommitsComparison, error) {
	comparison, _, err := c.Repositories.CompareCommits(context.Background(), c.owner, c.repo, from, to)
	if err != nil {
		return nil, err
	}

	return comparison, err
}
