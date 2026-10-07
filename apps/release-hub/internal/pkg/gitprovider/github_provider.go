package gitprovider

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v69/github"
	"golang.org/x/oauth2"
)

type GithubProvider struct {
	client *github.Client
}

func newGithubProvider(ctx context.Context, token string) *GithubProvider {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return &GithubProvider{
		client: github.NewClient(tc),
	}
}

func (p *GithubProvider) Login(ctx context.Context) (*User, error) {
	user, _, err := p.client.Users.Get(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("verify github token failed: %w", err)
	}
	if user == nil || user.Login == nil || strings.TrimSpace(*user.Login) == "" {
		return nil, fmt.Errorf("github login is empty")
	}

	return &User{
		Login: *user.Login,
		Email: *user.Email,
	}, nil
}

func (p *GithubProvider) CreateRef(ctx context.Context, owner string, repo string, baseBranch, newBranch string) error {
	baseRef, _, err := p.client.Git.GetRef(ctx, owner, repo, "heads/"+baseBranch)
	if err != nil {
		return err
	}

	newRef := &github.Reference{
		Ref: github.Ptr("refs/heads/" + newBranch),
		Object: &github.GitObject{
			SHA: baseRef.Object.SHA,
		},
	}

	_, _, err = p.client.Git.CreateRef(ctx, owner, repo, newRef)
	if err != nil {
		return err
	}

	return nil
}

func (p *GithubProvider) GetRef(ctx context.Context, owner string, repo string, ref string) error {
	baseRef, _, err := p.client.Git.GetRef(ctx, owner, repo, ref)
	if err != nil {
		return err
	}

	if baseRef == nil {
		return fmt.Errorf("ref not found")
	}

	return nil
}

func (p *GithubProvider) DeleteRef(ctx context.Context, owner string, repo string, branch string) error {
	_, err := p.client.Git.DeleteRef(ctx, owner, repo, "heads/"+branch)
	if err != nil {
		return err
	}
	return nil
}
