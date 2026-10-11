package service

import (
	"context"
	"strings"

	"github.com/maodou24/release-hub/internal/database"
	"github.com/maodou24/release-hub/internal/model"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"gorm.io/gorm"
)

func UpdateNamespace(ctx context.Context, id uint, name string) (*resp.Namespace, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, gorm.ErrInvalidData
	}
	namespace, err := model.GetNamespace(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := database.GetDb().WithContext(ctx).Model(namespace).Update("name", name).Error; err != nil {
		return nil, err
	}
	response := toNamespaceResponse(*namespace)
	return &response, nil
}

func GetOrCreateNamespace(ctx context.Context, name string) (*resp.Namespace, error) {
	namespace, err := model.GetOrCreateNamespace(ctx, strings.TrimSpace(name))
	if err != nil {
		return nil, err
	}
	response := toNamespaceResponse(*namespace)
	return &response, nil
}

func QueryNamespaces(ctx context.Context) ([]resp.Namespace, error) {
	namespaces, err := model.QueryNamespaces(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]resp.Namespace, 0, len(namespaces))
	for _, namespace := range namespaces {
		result = append(result, toNamespaceResponse(namespace))
	}
	return result, nil
}
