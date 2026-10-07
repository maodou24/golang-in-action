package model

import (
	"context"
	"time"

	"github.com/maodou24/release-hub/internal/database"
	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"column:name;not null;uniqueIndex:idx_projects_namespace_name" json:"name"`
	NamespaceID uint           `gorm:"column:namespace_id;not null;uniqueIndex:idx_projects_namespace_name;index" json:"namespaceId"`
}

func CreateProject(ctx context.Context, p *Project) error {
	db := database.GetDb().WithContext(ctx)
	return db.Create(p).Error
}

func GetProject(ctx context.Context, id uint) (*Project, error) {
	var project Project
	err := database.GetDb().WithContext(ctx).First(&project, id).Error
	return &project, err
}

func ListProjects(ctx context.Context, namespaceID uint) ([]Project, error) {
	var projects []Project
	err := database.GetDb().WithContext(ctx).Where("namespace_id = ?", namespaceID).Order("name ASC").Find(&projects).Error
	return projects, err
}

func DeleteProject(ctx context.Context, id uint) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		var project Project
		if err := tx.First(&project, id).Error; err != nil {
			return err
		}
		var releases []Release
		if err := tx.Where("project_id = ?", project.ID).Find(&releases).Error; err != nil {
			return err
		}
		for _, release := range releases {
			if err := tx.Where("release_id = ?", release.ID).Delete(&ReleaseBranchResult{}).Error; err != nil {
				return err
			}
			if err := tx.Where("release_id = ?", release.ID).Delete(&Dependency{}).Error; err != nil {
				return err
			}
		}
		if err := tx.Where("project_id = ?", project.ID).Delete(&Dependency{}).Error; err != nil {
			return err
		}
		if err := tx.Where("project_id = ?", project.ID).Delete(&Release{}).Error; err != nil {
			return err
		}
		return tx.Delete(&project).Error
	})
}

func UpdateDefaultDependencies(ctx context.Context, projectID uint, deps []Dependency) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		var release Release
		if err := tx.Where("project_id = ? AND version = ?", projectID, "default").FirstOrCreate(&release, Release{ProjectID: projectID, Version: "default", BranchName: "default", Status: "configured"}).Error; err != nil {
			return err
		}
		if err := tx.Where("release_id = ?", release.ID).Delete(&Dependency{}).Error; err != nil {
			return err
		}
		for i := range deps {
			deps[i].ProjectID = projectID
			deps[i].ReleaseID = release.ID
		}
		if len(deps) == 0 {
			return nil
		}
		return tx.Create(&deps).Error
	})
}

func UpdateProjectWithDependencies(ctx context.Context, project *Project, deps []Dependency) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(project).Update("name", project.Name).Error; err != nil {
			return err
		}
		var release Release
		if err := tx.Where("project_id = ? AND version = ?", project.ID, "default").FirstOrCreate(&release, Release{ProjectID: project.ID, Version: "default", BranchName: "default", Status: "configured"}).Error; err != nil {
			return err
		}
		if err := tx.Where("release_id = ?", release.ID).Delete(&Dependency{}).Error; err != nil {
			return err
		}
		for i := range deps {
			deps[i].ProjectID = project.ID
			deps[i].ReleaseID = release.ID
		}
		if len(deps) == 0 {
			return nil
		}
		return tx.Create(&deps).Error
	})
}

func CreateProjectWithDependencies(ctx context.Context, project *Project, deps []Dependency) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(project).Error; err != nil {
			return err
		}
		defaultRelease := &Release{ProjectID: project.ID, Version: "default", BranchName: "default", Status: "configured"}
		if err := tx.Create(defaultRelease).Error; err != nil {
			return err
		}
		for i := range deps {
			deps[i].ProjectID = project.ID
			deps[i].ReleaseID = defaultRelease.ID
		}
		if len(deps) == 0 {
			return nil
		}
		return tx.Create(&deps).Error
	})
}

func ListDependencies(ctx context.Context, releaseID uint) ([]Dependency, error) {
	var deps []Dependency
	err := database.GetDb().WithContext(ctx).Where("release_id = ?", releaseID).Order("id ASC").Find(&deps).Error
	return deps, err
}

func ListReleases(ctx context.Context, projectID uint) ([]Release, error) {
	var releases []Release
	err := database.GetDb().WithContext(ctx).Where("project_id = ?", projectID).Order("id ASC").Find(&releases).Error
	return releases, err
}

func GetRelease(ctx context.Context, id uint) (*Release, error) {
	var release Release
	err := database.GetDb().WithContext(ctx).First(&release, id).Error
	return &release, err
}

func SaveRelease(ctx context.Context, release *Release, deps []Dependency) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		if release.ID == 0 {
			if err := tx.Create(release).Error; err != nil {
				return err
			}
		} else if err := tx.Model(release).Updates(map[string]any{
			"baseline_id": release.BaselineID, "version": release.Version,
			"branch_name": release.BranchName, "status": release.Status, "message": release.Message,
		}).Error; err != nil {
			return err
		}
		if err := tx.Where("release_id = ?", release.ID).Delete(&Dependency{}).Error; err != nil {
			return err
		}
		for i := range deps {
			deps[i].ProjectID = release.ProjectID
			deps[i].ReleaseID = release.ID
		}
		if len(deps) == 0 {
			return nil
		}
		return tx.Create(&deps).Error
	})
}

func DeleteRelease(ctx context.Context, id uint) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		var release Release
		if err := tx.First(&release, id).Error; err != nil {
			return err
		}
		if err := tx.Where("release_id = ?", release.ID).Delete(&Dependency{}).Error; err != nil {
			return err
		}
		if err := tx.Where("release_id = ?", release.ID).Delete(&ReleaseBranchResult{}).Error; err != nil {
			return err
		}
		return tx.Delete(&release).Error
	})
}

func GetReleaseByProjectVersion(ctx context.Context, projectID uint, version string) (*Release, error) {
	var release Release
	err := database.GetDb().WithContext(ctx).Where("project_id = ? AND version = ?", projectID, version).First(&release).Error
	return &release, err
}
