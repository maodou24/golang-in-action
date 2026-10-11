package request

type ProjectReq struct {
	NamespaceID   uint            `json:"namespaceId,omitempty"`
	NamespaceName string          `json:"namespaceName,omitempty"`
	ProjectID     uint            `json:"projectId,omitempty"`
	ProjectName   string          `json:"projectName,omitempty"`
	Dependencies  []DependencyReq `json:"dependencies"`
}

type NamespaceReq struct {
	Name string `json:"name"`
}

type DependencyReq struct {
	ID         uint     `json:"id,omitempty"`
	ProjectID  uint     `json:"projectId,omitempty"`
	ReleaseID  uint     `json:"releaseId,omitempty"`
	RepoURL    string   `json:"repoUrl"`
	Path       string   `json:"path"`
	Name       string   `json:"name"`
	BranchName string   `json:"branchName"`
	Groups     []string `json:"groups,omitempty"`
}

type ReleaseReq struct {
	ID                uint            `json:"id,omitempty"`
	ProjectID         uint            `json:"projectId,omitempty"`
	ReleaseName       string          `json:"releaseName"`
	BaselineID        uint            `json:"baselineId,omitempty"`
	Dependencies      []DependencyReq `json:"dependencies"`
	CommitSummary     string          `json:"commitSummary"`
	CommitDescription string          `json:"commitDescription"`
}
