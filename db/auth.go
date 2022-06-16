package db

import (
	"errors"
	"fmt"
	"strconv"
)

func FindUser(uname, pwd string) bool {
	var SQLITE SQL = SQL{"test.db", "sqlite3"}
	var db = SQLITE.Open()
	defer db.Close()

	result, err := db.Query(`
	Select id from user where name = "` + uname + `"
	and password = "` + pwd + `"`)
	if err != nil {
		fmt.Println(err)
	}

	return result.Next()

}
func GetPwd(uname string) (string, error) {
	var SQLITE SQL = SQL{"test.db", "sqlite3"}
	var db = SQLITE.Open()
	defer db.Close()

	result, err := db.Query(`
	Select password from user where name = "` + uname + `"`)
	if err != nil {
		fmt.Println(err)
	}
	if result.Next() {
		var pwd string
		err = result.Scan(&pwd)
		return pwd, err
	} else {
		return "", errors.New("no such user")
	}
}

func GetId(uname string) (int, error) {
	var SQLITE SQL = SQL{"test.db", "sqlite3"}
	var db = SQLITE.Open()
	defer db.Close()

	result, err := db.Query(`
	Select id from user where name = "` + uname + `"`)
	if err != nil {
		fmt.Println(err)
	}
	if result.Next() {
		var pwd string
		err = result.Scan(&pwd)
		if err != nil {
			return -1, err
		}
		res, err := strconv.Atoi(pwd)
		return res, err
	} else {
		return -1, errors.New("no such user")
	}
}
