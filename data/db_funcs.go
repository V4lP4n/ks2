package data

import (
	"database/sql"
	"fmt"
)

func Select(req string) *sql.Rows {
	db := DBSETTINGS.Open()
	defer db.Close()
	rows, err := db.Query(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("from sql.Select")
		return rows
	}

	return rows
}
