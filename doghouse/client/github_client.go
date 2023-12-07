package client

import (
	"context"

	"github.com/google/go-github/v57/github"

	"github.com/reviewdog/reviewdog/doghouse"
	"github.com/reviewdog/reviewdog/doghouse/server"
)

// GitHubClient is client which talks to GitHub directly instead of talking to
// doghouse server.
type GitHubClient struct {
	Client *github.Client
}

func (c *GitHubClient) Check(ctx context.Context, req *doghouse.CheckRequest, makeRequest bool) (*doghouse.CheckResponse, error) {
	return server.NewChecker(req, c.Client).Check(ctx, makeRequest)
}
