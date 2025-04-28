package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func OpenPostgreSQL(conn string) (*sql.DB, error) {
	sqlDb, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return sqlDb, nil
}
