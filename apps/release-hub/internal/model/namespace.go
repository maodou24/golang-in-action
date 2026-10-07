package model

import (
	"context"
	"strings"
	"time"

	"github.com/maodou24/release-hub/internal/database"
	"gorm.io/gorm"
)

type Namespace struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"column:name;not null;uniqueIndex" json:"name"`
}

func QueryNamespaces(ctx context.Context) ([]Namespace, error) {
	var result []Namespace
	err := database.GetDb().WithContext(ctx).Order("id DESC").Find(&result).Error
	return result, err
}

func CreateNamespace(ctx context.Context, name string) (*Namespace, error) {
	namespace := &Namespace{Name: strings.TrimSpace(name)}
	if namespace.Name == "" {
		return nil, gorm.ErrInvalidData
	}
	if err := database.GetDb().WithContext(ctx).Create(namespace).Error; err != nil {
		return nil, err
	}
	return namespace, nil
}

func GetNamespace(ctx context.Context, id uint) (*Namespace, error) {
	var namespace Namespace
	err := database.GetDb().WithContext(ctx).First(&namespace, id).Error
	return &namespace, err
}

func DeleteNamespace(ctx context.Context, id uint) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		var projects []Project
		if err := tx.Where("namespace_id = ?", id).Find(&projects).Error; err != nil {
			return err
		}
		for _, project := range projects {
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
		}
		if err := tx.Where("namespace_id = ?", id).Delete(&Project{}).Error; err != nil {
			return err
		}
		return tx.Delete(&Namespace{}, id).Error
	})
}
