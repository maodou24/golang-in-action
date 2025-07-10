package testcases

import (
	"context"
	"fmt"
	"github.com/maodou24/gorm-example/models"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

func TestMain(t *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, err := gnomock.Start(p,
		gnomock.WithContext(ctx),
		gnomock.WithContainerName("test"),
		gnomock.WithContainerReuse(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(), "gnomock", "gnomick", "mydb",
	)

	gdb, err := gorm.Open(gormpg.Open(connStr))
	if err != nil {
		panic(err)
	}

	db = gdb

	tables := []any{
		&models.User{},
	}
	if err := db.AutoMigrate(tables...); err != nil {
		panic(err)
	}

	code := t.Run()

	_ = gnomock.Stop(container)
	os.Exit(code)
}
