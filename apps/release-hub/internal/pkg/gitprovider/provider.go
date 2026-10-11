package gitprovider

import (
	"context"

	"github.com/maodou24/release-hub/internal/pkg/base"
)

type Client interface {
	Login(ctx context.Context) (*User, error)
	// git
	DeleteRef(ctx context.Context, owner string, repo string, ref string) error
	GetRef(ctx context.Context, owner string, repo string, ref string) error
	CreateRef(ctx context.Context, owner string, repo string, oldref, newref string) error
	CommitManifest(ctx context.Context, owner, repo, baseBranch, branch, filePath, message string, content []byte) (string, error)
}

type User struct {
	Login string
	Email string
}

func NewClient(ctx context.Context, provider int, token string) Client {
	switch provider {
	case base.GitProviderGithub:
		return newGithubProvider(ctx, token)
	default:
		return nil
	}
}
