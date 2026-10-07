package service

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/maodou24/release-hub/internal/model"
	"github.com/maodou24/release-hub/internal/pkg/gitprovider"
	"github.com/maodou24/release-hub/internal/pkg/request"
	"github.com/maodou24/release-hub/internal/pkg/security"
)

func CreateRelease(ctx context.Context, publicUserID string, releaseReq request.ReleaseReq, cipher *security.TokenCipher) (*model.Release, error) {
	version := strings.TrimSpace(releaseReq.ReleaseName)
	if err := validatePathPart(version, "releaseName"); err != nil || version == "dev" || version == "default" {
		return nil, fmt.Errorf("invalid releaseName")
	}
	project, err := model.GetProject(ctx, releaseReq.ProjectID)
	if err != nil {
		return nil, err
	}

	baselineReleaseID := releaseReq.BaselineID
	if baselineReleaseID == 0 {
		defaultRelease, err := model.GetReleaseByProjectVersion(ctx, project.ID, "default")
		if err != nil {
			return nil, fmt.Errorf("load default release: %w", err)
		}
		baselineReleaseID = defaultRelease.ID
	}
	baselineRelease, err := model.GetRelease(ctx, baselineReleaseID)
	if err != nil {
		return nil, fmt.Errorf("load baseline release: %w", err)
	}
	if baselineRelease.ProjectID != project.ID {
		return nil, fmt.Errorf("baseline release belongs to a different project")
	}
	baselineDependencies, err := model.ListDependencies(ctx, baselineRelease.ID)
	if err != nil {
		return nil, fmt.Errorf("load baseline dependencies: %w", err)
	}
	dependencies := make([]request.DependencyReq, 0, len(baselineDependencies))
	for _, dep := range baselineDependencies {
		dependencies = append(dependencies, request.DependencyReq{
			ProjectID: project.ID, ReleaseID: baselineRelease.ID,
			RepoURL: dep.RemoteURL, Path: dep.Path, Name: dep.Name,
			BranchName: dep.BranchName, Groups: strings.Split(dep.Groups, ","),
		})
	}
	if len(releaseReq.Dependencies) > 0 {
		dependencies = releaseReq.Dependencies
	}
	if len(dependencies) == 0 {
		return nil, fmt.Errorf("dependencies cannot be empty")
	}
	deps, err := toModelDependencies(dependencies)
	if err != nil {
		return nil, err
	}
	release := &model.Release{
		ProjectID: project.ID, Version: version, BaselineID: &baselineRelease.ID,
		BranchName: version, Status: "created", Message: "release configuration saved",
	}
	user, err := model.GetUserByPublicID(ctx, publicUserID)
	if err != nil {
		return nil, fmt.Errorf("load release user: %w", err)
	}
	token, err := cipher.Decrypt(user.TokenCipher)
	if err != nil {
		return nil, fmt.Errorf("decrypt provider token: %w", err)
	}
	client := gitprovider.NewClient(ctx, user.GitProviderType, token)
	if client == nil {
		return nil, fmt.Errorf("unsupported git provider: %d", user.GitProviderType)
	}
	if err := model.SaveRelease(ctx, release, deps); err != nil {
		return nil, fmt.Errorf("save release: %w", err)
	}
	results := make([]model.ReleaseBranchResult, 0, len(deps))
	anyFailed := false
	for _, dep := range deps {
		owner, repo, parseErr := parseRepository(dep.RemoteURL)
		if parseErr != nil {
			results = append(results, model.ReleaseBranchResult{RepoURL: dep.RemoteURL, BranchName: dep.BranchName, Result: "failed", Message: parseErr.Error()})
			anyFailed = true
			continue
		}
		baseBranch := findBaselineBranch(baselineDependencies, dep)
		if baseBranch == "" {
			results = append(results, model.ReleaseBranchResult{RepoURL: dep.RemoteURL, BranchName: version, Result: "failed", Message: "baseline dependency not found"})
			anyFailed = true
			continue
		}
		if createErr := client.CreateRef(ctx, owner, repo, baseBranch, dep.BranchName); createErr != nil {
			result, message := "failed", createErr.Error()
			if strings.Contains(strings.ToLower(message), "already exists") {
				result, message = "skipped", "branch already exists"
			} else {
				anyFailed = true
			}
			results = append(results, model.ReleaseBranchResult{RepoURL: dep.RemoteURL, BranchName: dep.BranchName, Result: result, Message: message})
			continue
		}
		results = append(results, model.ReleaseBranchResult{RepoURL: dep.RemoteURL, BranchName: dep.BranchName, Result: "success", Message: "created"})
	}
	status, message := "success", "release created"
	if anyFailed {
		status, message = "partial_success", "release saved; one or more remote branches failed"
	}
	if err := model.SaveReleaseBranchResults(ctx, release.ID, results); err != nil {
		return nil, fmt.Errorf("save release branch results: %w", err)
	}
	if err := model.UpdateReleaseResult(ctx, release.ID, status, message, "", ""); err != nil {
		return nil, fmt.Errorf("save release status: %w", err)
	}
	release.Status, release.Message = status, message
	return release, nil
}

func findBaselineBranch(baseline []model.Dependency, target model.Dependency) string {
	for _, dependency := range baseline {
		if dependency.RemoteURL == target.RemoteURL && dependency.Path == target.Path {
			return dependency.BranchName
		}
	}
	return ""
}

func parseRepository(raw string) (string, string, error) {
	parsed, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", "", fmt.Errorf("invalid repository URL %q", raw)
	}
	parts := strings.Split(strings.Trim(path.Clean(parsed.Path), "/"), "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("repository URL must identify owner/repo: %s", raw)
	}
	return parts[0], strings.TrimSuffix(parts[1], ".git"), nil
}
