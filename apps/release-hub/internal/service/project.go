package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/maodou24/release-hub/internal/model"
	req "github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
)

func CreateProject(ctx context.Context, projectReq req.ProjectReq) (*resp.Project, error) {
	var namespace *model.Namespace
	var err error
	if projectReq.NamespaceID != 0 {
		namespace, err = model.GetNamespace(ctx, projectReq.NamespaceID)
	} else {
		namespace, err = model.GetOrCreateNamespace(ctx, projectReq.NamespaceName)
	}
	if err != nil {
		return nil, fmt.Errorf("get namespace: %w", err)
	}

	projectName := strings.TrimSpace(projectReq.ProjectName)
	if projectName == "" {
		return nil, fmt.Errorf("projectName is required")
	}
	project := &model.Project{Name: projectName, NamespaceID: namespace.ID}
	deps, err := toModelDependencies(projectReq.Dependencies)
	if err != nil {
		return nil, err
	}
	if err := model.CreateProjectWithDependencies(ctx, project, deps); err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}
	return assembleProject(ctx, project, namespace)
}

func UpdateProject(ctx context.Context, projectID uint, projectReq req.ProjectReq) (*resp.Project, error) {
	project, err := model.GetProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if projectReq.ProjectID != 0 && projectReq.ProjectID != project.ID {
		return nil, fmt.Errorf("projectId does not match route")
	}
	if projectReq.NamespaceID != 0 && projectReq.NamespaceID != project.NamespaceID {
		return nil, fmt.Errorf("namespaceId does not match project")
	}
	if name := strings.TrimSpace(projectReq.ProjectName); name != "" {
		project.Name = name
	}
	deps, err := toModelDependencies(projectReq.Dependencies)
	if err != nil {
		return nil, err
	}
	if err := model.UpdateProjectWithDependencies(ctx, project, deps); err != nil {
		return nil, err
	}
	return assembleProject(ctx, project, nil)
}

func ListProjects(ctx context.Context, namespaceID uint) ([]resp.Project, error) {
	namespace, err := model.GetNamespace(ctx, namespaceID)
	if err != nil {
		return nil, err
	}
	projects, err := model.ListProjects(ctx, namespaceID)
	if err != nil {
		return nil, err
	}
	result := make([]resp.Project, 0, len(projects))
	for _, project := range projects {
		view, err := assembleProject(ctx, &project, namespace)
		if err != nil {
			return nil, err
		}
		result = append(result, *view)
	}
	return result, nil
}

func GetProject(ctx context.Context, projectID uint) (*resp.Project, error) {
	project, err := model.GetProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return assembleProject(ctx, project, nil)
}

func assembleProject(ctx context.Context, project *model.Project, namespace *model.Namespace) (*resp.Project, error) {
	if namespace == nil {
		var err error
		namespace, err = model.GetNamespace(ctx, project.NamespaceID)
		if err != nil {
			return nil, err
		}
	}
	releases, err := model.ListReleases(ctx, project.ID)
	if err != nil {
		return nil, err
	}
	views := make([]resp.Release, 0, len(releases))
	for _, release := range releases {
		dependencies, err := model.ListDependencies(ctx, release.ID)
		if err != nil {
			return nil, err
		}
		views = append(views, toReleaseResponse(release, dependencies))
	}
	return &resp.Project{
		ID:          project.ID,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		Name:        project.Name,
		NamespaceID: project.NamespaceID,
		Namespace:   toNamespaceResponse(*namespace),
		Releases:    views,
	}, nil
}

func DeleteProject(ctx context.Context, projectID uint) error {
	if _, err := model.GetProject(ctx, projectID); err != nil {
		return err
	}
	return model.DeleteProject(ctx, projectID)
}

func DeleteNamespace(ctx context.Context, namespaceID uint) error {
	if _, err := model.GetNamespace(ctx, namespaceID); err != nil {
		return err
	}
	return model.DeleteNamespace(ctx, namespaceID)
}

func UpdateProjectVersion(ctx context.Context, releaseID uint, releaseReq req.ReleaseReq) error {
	release, err := model.GetRelease(ctx, releaseID)
	if err != nil {
		return err
	}
	if releaseReq.ID != 0 && releaseReq.ID != release.ID {
		return fmt.Errorf("releaseId does not match route")
	}
	if releaseReq.ProjectID != 0 && releaseReq.ProjectID != release.ProjectID {
		return fmt.Errorf("projectId does not match release")
	}
	deps, err := toModelDependencies(releaseReq.Dependencies)
	if err != nil {
		return err
	}
	if releaseReq.ReleaseName != "" {
		if err := validatePathPart(releaseReq.ReleaseName, "releaseName"); err != nil || releaseReq.ReleaseName == "dev" || releaseReq.ReleaseName == "default" {
			return fmt.Errorf("invalid published version")
		}
		release.Version = releaseReq.ReleaseName
		release.BranchName = releaseReq.ReleaseName
	}
	if releaseReq.BaselineID != 0 {
		baseline, err := model.GetRelease(ctx, releaseReq.BaselineID)
		if err != nil {
			return fmt.Errorf("baseline release not found: %w", err)
		}
		if baseline.ProjectID != release.ProjectID {
			return fmt.Errorf("baseline release belongs to a different project")
		}
		release.BaselineID = &baseline.ID
	} else {
		release.BaselineID = nil
	}
	release.Status = "updated"
	return model.SaveRelease(ctx, release, deps)
}

func GetProjectVersion(ctx context.Context, releaseID uint) (*resp.Release, error) {
	release, err := model.GetRelease(ctx, releaseID)
	if err != nil {
		return nil, err
	}
	project, err := model.GetProject(ctx, release.ProjectID)
	if err != nil {
		return nil, err
	}
	namespace, err := model.GetNamespace(ctx, project.NamespaceID)
	if err != nil {
		return nil, err
	}
	deps, err := model.ListDependencies(ctx, release.ID)
	if err != nil {
		return nil, err
	}
	branchResults, err := model.ListReleaseBranchResults(ctx, release.ID)
	if err != nil {
		return nil, err
	}
	projectSummary := &resp.ProjectSummary{ID: project.ID, Name: project.Name, NamespaceID: project.NamespaceID}
	response := toReleaseResponse(*release, deps)
	response.Project = projectSummary
	namespaceResponse := toNamespaceResponse(*namespace)
	response.Namespace = &namespaceResponse
	response.BranchResults = make([]resp.BranchResult, 0, len(branchResults))
	for _, result := range branchResults {
		response.BranchResults = append(response.BranchResults, toBranchResultResponse(result))
	}
	return &response, nil
}

func ListProjectVersions(ctx context.Context, projectID uint) ([]resp.Release, error) {
	if _, err := model.GetProject(ctx, projectID); err != nil {
		return nil, err
	}
	releases, err := model.ListReleases(ctx, projectID)
	if err != nil {
		return nil, err
	}
	result := make([]resp.Release, 0, len(releases))
	for _, release := range releases {
		deps, err := model.ListDependencies(ctx, release.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, toReleaseResponse(release, deps))
	}
	return result, nil
}

func DeleteProjectVersion(ctx context.Context, releaseID uint) error {
	release, err := model.GetRelease(ctx, releaseID)
	if err != nil {
		return err
	}
	if release.Version == "dev" || release.Version == "default" {
		return fmt.Errorf("default project configuration cannot be deleted as a release")
	}
	releases, err := model.ListReleases(ctx, release.ProjectID)
	if err != nil {
		return err
	}
	for _, child := range releases {
		if child.BaselineID != nil && *child.BaselineID == release.ID {
			return fmt.Errorf("release is used as a baseline by release %d", child.ID)
		}
	}
	return model.DeleteRelease(ctx, releaseID)
}

func BuildManifest(ctx context.Context, releaseID uint, remoteType string) ([]byte, error) {
	if remoteType == "" {
		remoteType = "https"
	}
	if remoteType != "https" && remoteType != "ssh" {
		return nil, fmt.Errorf("invalid remoteType")
	}
	release, err := GetProjectVersion(ctx, releaseID)
	if err != nil {
		return nil, err
	}
	fetch := "https://github.com"
	if remoteType == "ssh" {
		fetch = "ssh://git@github.com"
	}
	baseline := ""
	if release.BaselineID != nil {
		baselineRelease, err := model.GetRelease(ctx, *release.BaselineID)
		if err != nil {
			return nil, fmt.Errorf("load baseline release: %w", err)
		}
		baseline = baselineRelease.Version
	} else if release.Version == "default" {
		baseline = "default"
	}
	manifest := manifestXML{
		Baseline: baseline, Remotes: []manifestRemote{{Name: "origin", Fetch: fetch}},
		Projects: make([]manifestProject, 0, len(release.Dependencies)),
	}
	for _, dep := range release.Dependencies {
		name := dep.Name
		if name == "" {
			name = dep.RepoURL
		}
		projectPath := dep.Path
		if projectPath == "" {
			projectPath = path.Base(strings.TrimSuffix(dep.RepoURL, ".git"))
		}
		manifest.Projects = append(manifest.Projects, manifestProject{
			Path: projectPath, Name: name, Revision: dep.BranchName,
			Remote: "origin", Groups: dep.Groups,
		})
	}
	data, err := xml.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(append([]byte(xml.Header), data...), '\n'), nil
}

type manifestXML struct {
	XMLName  xml.Name          `xml:"manifest"`
	Baseline string            `xml:"x-baseline,attr,omitempty"`
	Remotes  []manifestRemote  `xml:"remote"`
	Projects []manifestProject `xml:"project"`
}

type manifestRemote struct {
	Name  string `xml:"name,attr"`
	Fetch string `xml:"fetch,attr"`
}

type manifestProject struct {
	Path     string `xml:"path,attr"`
	Name     string `xml:"name,attr"`
	Revision string `xml:"revision,attr"`
	Remote   string `xml:"remote,attr"`
	Groups   string `xml:"groups,attr,omitempty"`
}

func toModelDependencies(requests []req.DependencyReq) ([]model.Dependency, error) {
	deps := make([]model.Dependency, 0, len(requests))
	for _, item := range requests {
		remoteURL := strings.TrimSpace(item.RepoURL)
		branch := strings.TrimSpace(item.BranchName)
		if remoteURL == "" || branch == "" {
			return nil, fmt.Errorf("each dependency requires repoUrl and branchName")
		}
		groups := append([]string(nil), item.Groups...)
		sort.Strings(groups)
		deps = append(deps, model.Dependency{
			RemoteURL: remoteURL, BranchName: branch, Path: strings.TrimSpace(item.Path),
			Name: strings.TrimSpace(item.Name), Groups: strings.Join(groups, ","),
		})
	}
	return deps, nil
}

func validatePathPart(value, name string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." || path.Base(value) != value {
		return fmt.Errorf("invalid %s", name)
	}
	return nil
}

func toNamespaceResponse(namespace model.Namespace) resp.Namespace {
	return resp.Namespace{
		ID: namespace.ID, CreatedAt: namespace.CreatedAt,
		UpdatedAt: namespace.UpdatedAt, Name: namespace.Name,
	}
}

func toDependencyResponse(dependency model.Dependency) resp.Dependency {
	return resp.Dependency{
		ID: dependency.ID, CreatedAt: dependency.CreatedAt, UpdatedAt: dependency.UpdatedAt,
		ProjectID: dependency.ProjectID, ReleaseID: dependency.ReleaseID,
		RepoURL: dependency.RemoteURL, BranchName: dependency.BranchName,
		Path: dependency.Path, Name: dependency.Name, Groups: dependency.Groups,
	}
}

func toReleaseResponse(release model.Release, dependencies []model.Dependency) resp.Release {
	response := resp.Release{
		ID: release.ID, CreatedAt: release.CreatedAt, UpdatedAt: release.UpdatedAt,
		ProjectID: release.ProjectID, Version: release.Version, BaselineID: release.BaselineID,
		BranchName: release.BranchName, Status: release.Status, Message: release.Message,
		ManifestCommitSHA: release.ManifestCommitSHA,
		Dependencies:      make([]resp.Dependency, 0, len(dependencies)),
	}
	for _, dependency := range dependencies {
		response.Dependencies = append(response.Dependencies, toDependencyResponse(dependency))
	}
	return response
}

func toBranchResultResponse(result model.ReleaseBranchResult) resp.BranchResult {
	return resp.BranchResult{
		ID: result.ID, CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt,
		ReleaseID: result.ReleaseID, RepoURL: result.RepoURL,
		BranchName: result.BranchName, Result: result.Result, Message: result.Message,
	}
}
