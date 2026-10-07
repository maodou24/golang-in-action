package service

import (
	"context"

	"github.com/maodou24/release-hub/internal/model"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
)

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
