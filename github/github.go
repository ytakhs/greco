package github

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client contains github's RepositoriesService, owner name and repo name.
type Client struct {
	owner        string
	repo         string
	Repositories *github.RepositoriesService
	Search       *github.SearchService
}

// NewClient creates a new object which contains github's Repository object.
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
		Search:       gh.Search,
	}

	return client, nil
}

// Tags returns an array of github's RepositoryTag objects
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

// Compare returns github's CommitsComparison which is a result of comparison from and to.
func (c *Client) Compare(from string, to string) (*github.CommitsComparison, error) {
	comparison, _, err := c.Repositories.CompareCommits(context.Background(), c.owner, c.repo, from, to)
	if err != nil {
		return nil, err
	}

	return comparison, err
}

func (c *Client) SearchRepositories() (*github.RepositoriesSearchResult, error) {
	opts := &github.SearchOptions{Sort: "created", Order: "asc"}
	searchResult, _, err := c.Search.Repositories(context.Background(), c.repo, opts)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
