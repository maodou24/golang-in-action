package service

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/maodou24/release-hub/internal/database"
	"github.com/maodou24/release-hub/internal/model"
	"gorm.io/gorm"
)

func TestNamespaceManagement(t *testing.T) {
	ctx := context.Background()
	if err := database.Init(filepath.Join(t.TempDir(), "namespace.db")); err != nil {
		t.Fatal(err)
	}
	db := database.GetDb()
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	if err := db.AutoMigrate(&model.Namespace{}, &model.Project{}); err != nil {
		t.Fatal(err)
	}
	namespace, err := GetOrCreateNamespace(ctx, " platform ")
	if err != nil {
		t.Fatal(err)
	}
	if namespace.Name != "platform" {
		t.Fatalf("name = %q", namespace.Name)
	}
	existing, err := GetOrCreateNamespace(ctx, "platform")
	if err != nil || existing.ID != namespace.ID {
		t.Fatalf("expected existing namespace, got %v, %v", existing, err)
	}
	project := &model.Project{Name: "app", NamespaceID: namespace.ID}
	if err := model.CreateProject(ctx, project); err != nil {
		t.Fatal(err)
	}
	renamed, err := UpdateNamespace(ctx, namespace.ID, " mobile ")
	if err != nil {
		t.Fatal(err)
	}
	if renamed.ID != namespace.ID || renamed.Name != "mobile" {
		t.Fatalf("unexpected rename response: %+v", renamed)
	}
	stored, err := model.GetProject(ctx, project.ID)
	if err != nil || stored.NamespaceID != namespace.ID {
		t.Fatalf("project binding changed: %v, %v", stored, err)
	}
	if err := DeleteNamespace(ctx, namespace.ID); !errors.Is(err, model.ErrNamespaceNotEmpty) {
		t.Fatalf("expected nonempty error, got %v", err)
	}
	if _, err := UpdateNamespace(ctx, namespace.ID, "   "); !errors.Is(err, gorm.ErrInvalidData) {
		t.Fatalf("expected empty name error, got %v", err)
	}
	if _, err := GetOrCreateNamespace(ctx, "other"); err != nil {
		t.Fatal(err)
	}
	if _, err := UpdateNamespace(ctx, namespace.ID, "other"); err == nil {
		t.Fatal("expected duplicate name to be rejected")
	}
	listed, err := QueryNamespaces(ctx)
	if err != nil || len(listed) != 2 {
		t.Fatalf("unexpected list: %v, %v", listed, err)
	}
	if err := db.Delete(project).Error; err != nil {
		t.Fatal(err)
	}
	if err := DeleteNamespace(ctx, namespace.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := model.GetNamespace(ctx, namespace.ID); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("namespace was not deleted: %v", err)
	}
	if _, err := UpdateNamespace(ctx, namespace.ID, "deleted"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected missing namespace error, got %v", err)
	}
}