package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/maodou24/release-hub/internal/model"
	"github.com/maodou24/release-hub/internal/pkg/gitprovider"
	req "github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/pkg/security"
)

func UpsertUser(ctx context.Context, userReq req.User, cipher *security.TokenCipher) (*resp.User, error) {
	token := strings.TrimSpace(userReq.AccessToken)
	if token == "" {
		token = strings.TrimSpace(userReq.Token)
	}
	if token == "" {
		return nil, fmt.Errorf("access token is required")
	}
	client := gitprovider.NewClient(ctx, userReq.GitProviderType, token)
	if client == nil {
		return nil, fmt.Errorf("unsupported git provider: %d", userReq.GitProviderType)
	}
	gitUser, err := client.Login(ctx)
	if err != nil {
		return nil, err
	}
	cipherText, err := cipher.Encrypt(token)
	if err != nil {
		return nil, fmt.Errorf("encrypt provider token: %w", err)
	}
	user := &model.User{
		PublicId: userReq.PublicId,
		Name: gitUser.Login,
		Login: gitUser.Login,
		Email: strings.TrimSpace(userReq.Email),
		TokenCipher: cipherText,
		GitProviderType: userReq.GitProviderType,
	}
	if user.Email == "" {
		user.Email = gitUser.Email
	}
	if err := model.UpsertUser(ctx, user); err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func GetUser(ctx context.Context, publicID string) (*resp.User, error) {
	user, err := model.GetUserByPublicID(ctx, publicID)
	if err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func toUserResponse(user *model.User) *resp.User {
	return &resp.User{
		PublicID: user.PublicId, Name: user.Name, Email: user.Email,
		Login: user.Login, GitProviderType: user.GitProviderType,
	}
}
