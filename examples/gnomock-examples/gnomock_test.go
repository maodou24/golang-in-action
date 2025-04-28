package gnomockexamples

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

func TestMain(t *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, err := gnomock.Start(p,
		gnomock.WithContext(ctx),
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
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`CREATE TABLE cars (
		brand VARCHAR(255),
		model VARCHAR(255),
		year INT
		);`)
	if err != nil {
		fmt.Println("create table", err)
		return
	}

	code := t.Run()

	_ = gnomock.Stop(container)
	os.Exit(code)
}

func TestGnomockPosgreSQL(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	container, err := gnomock.Start(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`CREATE TABLE cars (
  brand VARCHAR(255),
  model VARCHAR(255),
  year INT
);`)
	if err != nil {
		fmt.Println("create table", err)
		return
	}

	_, err = db.Exec(`INSERT INTO cars (brand, model, year)
VALUES ('Ford', 'Mustang', 1964);`)
	if err != nil {
		fmt.Println("Prepare record", err)
		return
	}

	rows, err := db.Query(`SELECT * FROM cars;`)
	if err != nil {
		fmt.Println("query", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var brand, model string
		var year int
		err = rows.Scan(&brand, &model, &year)
		fmt.Println(err, brand, model, year)
	}
}
