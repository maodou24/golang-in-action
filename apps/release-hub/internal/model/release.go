package model

import (
	"context"
	"time"

	"github.com/maodou24/release-hub/internal/database"
	"gorm.io/gorm"
)

type Release struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	ProjectID         uint           `gorm:"column:project_id;not null;uniqueIndex:idx_releases_project_version;index" json:"projectId"`
	Version           string         `gorm:"column:version;not null;uniqueIndex:idx_releases_project_version" json:"version"`
	BaselineID        *uint          `gorm:"column:baseline_id;index" json:"baselineId,omitempty"`
	BranchName        string         `gorm:"column:branch_name" json:"branchName"`
	Status            string         `gorm:"column:status;not null;default:created" json:"status"`
	Message           string         `gorm:"column:message" json:"message"`
	ManifestCommitSHA string         `gorm:"column:manifest_commit_sha" json:"manifestCommitSha,omitempty"`
	BranchResults     string         `gorm:"column:branch_results;type:text" json:"-"`
}

type ReleaseBranchResult struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ReleaseID  uint           `gorm:"column:release_id;not null;index" json:"releaseId"`
	RepoURL    string         `gorm:"column:repo_url;not null" json:"repoUrl"`
	BranchName string         `gorm:"column:branch_name;not null" json:"branchName"`
	Result     string         `gorm:"column:result;not null" json:"result"`
	Message    string         `gorm:"column:message" json:"message"`
}

func UpdateReleaseResult(ctx context.Context, releaseID uint, status, message, commitSHA, branchResults string) error {
	return database.GetDb().WithContext(ctx).Model(&Release{}).Where("id = ?", releaseID).Updates(map[string]any{
		"status": status, "message": message,
		"manifest_commit_sha": commitSHA, "branch_results": branchResults,
	}).Error
}

func SaveReleaseBranchResults(ctx context.Context, releaseID uint, results []ReleaseBranchResult) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("release_id = ?", releaseID).Delete(&ReleaseBranchResult{}).Error; err != nil {
			return err
		}
		for i := range results {
			results[i].ReleaseID = releaseID
		}
		if len(results) == 0 {
			return nil
		}
		return tx.Create(&results).Error
	})
}

func ListReleaseBranchResults(ctx context.Context, releaseID uint) ([]ReleaseBranchResult, error) {
	var results []ReleaseBranchResult
	err := database.GetDb().WithContext(ctx).Where("release_id = ?", releaseID).Order("id ASC").Find(&results).Error
	return results, err
}
