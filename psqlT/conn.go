package psqlT

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func SqlConn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgresql://ckex:1234@quote-postgres/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
