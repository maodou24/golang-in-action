package model

import (
	"time"

	"gorm.io/gorm"
)

type Dependency struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ProjectID  uint           `gorm:"column:project_id;not null;index" json:"projectId"`
	ReleaseID  uint           `gorm:"column:release_id;not null;index" json:"releaseId"`
	RemoteURL  string         `gorm:"column:remote_url;not null" json:"repoUrl"`
	BranchName string         `gorm:"column:branch_name;not null" json:"branchName"`
	Path       string         `gorm:"column:path" json:"path"`
	Name       string         `gorm:"column:name" json:"name"`
	Groups     string         `gorm:"column:groups" json:"groups"`
}
