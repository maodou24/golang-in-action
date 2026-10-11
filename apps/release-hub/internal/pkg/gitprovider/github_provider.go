package gitprovider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

	var u User
	if user.Login != nil {
		u.Login = *user.Login
	}
	if user.Email != nil {
		u.Email = *user.Email
	}

	return &u, nil
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

// CommitManifest creates the publishing branch if absent and commits one manifest
// without force-updating refs or overwriting unrelated repository files.
func (p *GithubProvider) CommitManifest(ctx context.Context, owner, repo, baseBranch, branch, filePath, message string, content []byte) (string, error) {
	if _, _, err := p.client.Git.GetRef(ctx, owner, repo, "heads/"+branch); err != nil {
		if !isNotFound(err) {
			return "", fmt.Errorf("get manifest branch: %w", err)
		}
		if baseBranch == "" {
			repository, _, err := p.client.Repositories.Get(ctx, owner, repo)
			if err != nil {
				return "", fmt.Errorf("get manifest repository: %w", err)
			}
			baseBranch = repository.GetDefaultBranch()
		}
		if baseBranch == "" {
			return "", fmt.Errorf("manifest repository has no default branch")
		}
		if err := p.CreateRef(ctx, owner, repo, baseBranch, branch); err != nil {
			// A concurrent release may have created the branch in the meantime.
			if _, _, lookupErr := p.client.Git.GetRef(ctx, owner, repo, "heads/"+branch); lookupErr != nil {
				return "", fmt.Errorf("create manifest branch: %w", err)
			}
		}
	}

	file, _, _, err := p.client.Repositories.GetContents(ctx, owner, repo, filePath, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil && !isNotFound(err) {
		return "", fmt.Errorf("get existing manifest: %w", err)
	}
	options := &github.RepositoryContentFileOptions{
		Message: github.Ptr(message), Content: content, Branch: github.Ptr(branch),
	}
	if err == nil {
		if file == nil || file.GetSHA() == "" {
			return "", fmt.Errorf("manifest path is not a file: %s", filePath)
		}
		options.SHA = file.SHA
	}
	result, _, err := p.client.Repositories.CreateFile(ctx, owner, repo, filePath, options)
	if err != nil {
		return "", fmt.Errorf("commit manifest: %w", err)
	}
	if result == nil || result.Commit.GetSHA() == "" {
		return "", fmt.Errorf("manifest commit returned no SHA")
	}
	return result.Commit.GetSHA(), nil
}

func isNotFound(err error) bool {
	var response *github.ErrorResponse
	return errors.As(err, &response) && response.Response != nil && response.Response.StatusCode == http.StatusNotFound
}
