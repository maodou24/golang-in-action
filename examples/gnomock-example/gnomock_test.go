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

var c *gnomock.Container
var db *sql.DB

func DB() *sql.DB {
	return db
}

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
	sqlDb, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = sqlDb.Exec(`CREATE TABLE cars (
		brand VARCHAR(255),
		model VARCHAR(255),
		year INT
		);`)
	if err != nil {
		fmt.Println("create table", err)
		return
	}

	c = container
	db = sqlDb
	code := t.Run()

	_ = gnomock.Stop(container)
	os.Exit(code)
}

func TestGnomockPosgreSQL(t *testing.T) {
	tx, err := DB().Begin()
	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec(`INSERT INTO cars (brand, model, year) VALUES ('Ford', 'Mustang', 1964);`)
	if err != nil {
		fmt.Println("Prepare record", err)
		return
	}

	rows, err := tx.Query(`SELECT * FROM cars;`)
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

	if tx.Rollback() != nil {
		t.Fail()
	}

	rows, err = db.Query(`SELECT * FROM cars;`)
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

func TestGnomockPosgreSQL2(t *testing.T) {
	tx, err := DB().Begin()
	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec(`INSERT INTO cars (brand, model, year) VALUES ('BYD', 'qin', 2000);`)
	if err != nil {
		fmt.Println("Prepare record", err)
		return
	}

	rows, err := tx.Query(`SELECT * FROM cars;`)
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

	if tx.Rollback() != nil {
		t.Fail()
	}

	rows, err = db.Query(`SELECT * FROM cars;`)
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