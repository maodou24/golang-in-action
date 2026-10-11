package model

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/maodou24/release-hub/internal/database"
	"gorm.io/gorm"
)

var ErrNamespaceNotEmpty = errors.New("namespace contains projects")

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

func GetOrCreateNamespace(ctx context.Context, name string) (*Namespace, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, gorm.ErrInvalidData
	}
	var namespace Namespace
	db := database.GetDb().WithContext(ctx)
	err := db.Where("name = ?", name).First(&namespace).Error
	if err == nil {
		return &namespace, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	created, err := CreateNamespace(ctx, name)
	if err == nil {
		return created, nil
	}
	// Another request may have created the same uniquely named namespace concurrently.
	if lookupErr := db.Where("name = ?", name).First(&namespace).Error; lookupErr == nil {
		return &namespace, nil
	}
	return nil, err
}

func GetNamespace(ctx context.Context, id uint) (*Namespace, error) {
	var namespace Namespace
	err := database.GetDb().WithContext(ctx).First(&namespace, id).Error
	return &namespace, err
}

func DeleteNamespace(ctx context.Context, id uint) error {
	db := database.GetDb().WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		var projectCount int64
		if err := tx.Model(&Project{}).Where("namespace_id = ?", id).Count(&projectCount).Error; err != nil {
			return err
		}
		if projectCount != 0 {
			return ErrNamespaceNotEmpty
		}
		return tx.Delete(&Namespace{}, id).Error
	})
}
