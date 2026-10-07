package main

import (
	"log"
	"strings"

	"github.com/maodou24/release-hub/config"
	"github.com/maodou24/release-hub/internal/controller"
	"github.com/maodou24/release-hub/internal/database"
	"github.com/maodou24/release-hub/internal/model"
	"github.com/maodou24/release-hub/internal/pkg/security"
	"github.com/maodou24/release-hub/internal/router"
)

func main() {
	cfg := config.Load()
	if strings.TrimSpace(cfg.TokenEncryptionKey) == "" || strings.TrimSpace(cfg.SessionSecret) == "" {
		log.Fatal("TOKEN_ENCRYPTION_KEY and SESSION_SECRET must be configured")
	}

	if err := database.Init(cfg.DataPath); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	if err := database.GetDb().AutoMigrate(
		&model.User{},
		&model.Namespace{},
		&model.Project{},
		&model.Release{},
		&model.Dependency{},
		&model.ReleaseBranchResult{},
	); err != nil {
		log.Fatalf("auto migrate err: %v", err)
	}
	controller.ConfigureAuth(security.NewTokenCipher(cfg.TokenEncryptionKey), cfg.SessionSecret)

	e := router.New(cfg.SessionSecret)
	if err := e.Run(cfg.HTTPAddr); err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}
}
