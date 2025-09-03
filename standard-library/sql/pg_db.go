package sql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var pgDB *sql.DB

func init() {
	dsn := "host=localhost user=postgres password=maodou24 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	pgDB = db
}
