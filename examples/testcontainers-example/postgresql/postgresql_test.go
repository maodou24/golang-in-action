package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Stub struct {
	DB    *sql.DB
	Clean func()
}

func SetUp(t *testing.T) *Stub {
	ctx := context.Background()

	template, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("template_db"),
		postgres.WithUsername("template"),
		postgres.WithPassword("template"),
	)
	if err != nil {
		log.Fatalf("run postgres container error: %v", err)
	}

	connStr, err := template.ConnectionString(ctx, "dbname=template_db")
	if err != nil {
		log.Fatalf("get connection string error: %v", err)
	}

	connStr = connStr + "&sslmode=disable"
	db, err := OpenPostgreSQL(connStr)
	if err != nil {
		log.Fatalf("open postgresql error: %v", err)
	}
	defer db.Close()
	time.Sleep(1 * time.Second)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users 
		(
		    id   SERIAL PRIMARY KEY, 
		    name TEXT NOT NULL UNIQUE
		)
	`)
	if err != nil {
		log.Fatalf("init template error: %v", err)
	}
	db.Close()

	if err := template.Snapshot(context.Background(), postgres.WithSnapshotName("snapshot-user")); err != nil {
		panic(err)
	}

	testDB, err := postgres.Run(context.Background(),
		"snapshot-user",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithStartupTimeout(30*time.Second)),
		// testcontainers.WithEnv(map[string]string{
		// 	"POSTGRES_TEMPLATE":         "template_db",
		// 	"POSTGRES_HOST_AUTH_METHOD": "trust",
		// }),
	)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := testDB.ConnectionString(context.Background(), "dbname=template_db")
	if err != nil {
		panic(err)
	}

	conn = conn + "&sslmode=disable"
	db, err = OpenPostgreSQL(conn)
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	return &Stub{
		DB: db,
		Clean: func() {
			testcontainers.CleanupContainer(t, testDB)
		},
	}
}

func TestPostgreSQL(t *testing.T) {
	s := SetUp(t)
	defer s.Clean()

	stmt, err := s.DB.Prepare("INSERT INTO users(name) VALUES ($1)")
	if err != nil {
		log.Println(err)
		return
	}

	res, err := stmt.Exec("test")
	assert.NoError(t, err)

	affect, err := res.RowsAffected()
	assert.NoError(t, err)

	fmt.Println("rows affect:", affect)
}
