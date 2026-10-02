package database

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maodou/go-leaning/standard-library/database/model"
)

func TestDatabasePostgreSQL(t *testing.T) {
	dsn := "host=localhost user=postgres password=maodou24 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	u := model.User{
		Name:    "maodou2",
		Address: "every where",
		Age:     18,
	}
	if err := model.Create(db, &u); err != nil {
		t.Fatalf("create user failed, err=%v", err)
	}

	queryUser, err := model.QueryById(db, u.ID)
	if err != nil {
		t.Fatalf("query user failed, err=%v", err)
	}

	queryUser.Name = "maodou"
	if err := model.UpdateById(db, &queryUser); err != nil {
		t.Fatalf("update user failed, err=%v", err)
	}

	if err := model.DeleteById(db, queryUser.ID); err != nil {
		t.Fatalf("delete user failed, err=%v", err)
	}
}
