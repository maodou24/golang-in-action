package database

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(path string) error {
	if path == "" {
		return fmt.Errorf("sqlite database path must not be empty")
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create sqlite database directory: %w", err)
	}

	sqliteDB, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open sqlite database %q: %w", path, err)
	}

	db = sqliteDB
	return nil
}

func GetDb() *gorm.DB {
	return db
}
