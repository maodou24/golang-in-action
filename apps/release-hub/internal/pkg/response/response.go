package resp

import "time"

type Result struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
	Cost      string `json:"cost"`
	RequestId string `json:"request_id"`
}

type Namespace struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

type User struct {
	PublicID        string `json:"publicId"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Login           string `json:"login"`
	GitProviderType int    `json:"gitProviderType"`
}

type AuthenticatedUser struct {
	Authenticated bool `json:"authenticated"`
	User          User `json:"user"`
}

type Project struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	NamespaceID uint      `json:"namespaceId"`
	Namespace   Namespace `json:"namespace"`
	Releases    []Release `json:"releases"`
}

type Release struct {
	ID                uint             `json:"id"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
	ProjectID         uint             `json:"projectId"`
	Version           string           `json:"version"`
	BaselineID        *uint            `json:"baselineId,omitempty"`
	BranchName        string           `json:"branchName"`
	Status            string           `json:"status"`
	Message           string           `json:"message"`
	ManifestCommitSHA string           `json:"manifestCommitSha,omitempty"`
	Project           *ProjectSummary  `json:"project,omitempty"`
	Namespace         *Namespace       `json:"namespace,omitempty"`
	Dependencies      []Dependency     `json:"dependencies"`
	BranchResults     []BranchResult   `json:"branchResults,omitempty"`
}

type ProjectSummary struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	NamespaceID uint   `json:"namespaceId"`
}

type Dependency struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ProjectID  uint      `json:"projectId"`
	ReleaseID  uint      `json:"releaseId"`
	RepoURL    string    `json:"repoUrl"`
	BranchName string    `json:"branchName"`
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	Groups     string    `json:"groups"`
}

type BranchResult struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ReleaseID  uint      `json:"releaseId"`
	RepoURL    string    `json:"repoUrl"`
	BranchName string    `json:"branchName"`
	Result     string    `json:"result"`
	Message    string    `json:"message"`
}

func Success(data any) Result {
	return Result{Code: 0, Msg: "success", Data: data}
}

func Failure(code int, message string) Result {
	return Result{Code: code, Msg: message}
}
