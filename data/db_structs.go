package data

import (
	"database/sql"
	"log"
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
