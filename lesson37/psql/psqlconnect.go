package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nurmuxammad"
	dbname   = "new"
	password = 1234
)

func ConnectDb() (*sql.DB, error) {
	con := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%d sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
