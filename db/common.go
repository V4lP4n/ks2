package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQL struct {
	ConString string
	Driver    string
}

func (s *SQL) Open() *sql.DB {
	db, err := sql.Open(s.Driver, s.ConString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
