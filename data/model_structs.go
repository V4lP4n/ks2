package data

import (
	"fmt"
	"strconv"
)

// USER
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Init(id int) {

	queryString := fmt.Sprintf(`
	SELECT
	name
	FROM user
	WHERE id = %s`, strconv.Itoa(id))
	// queryString := fmt.Sprintf(`
	// SELECT
	// id,
	// bg_color
	// FROM settings
	// WHERE owner_id =  %s`, strconv.Itoa(id))

	rows := Select(queryString)

	for rows.Next() {

		err := rows.Scan(&u.Name)

		if err != nil {

			fmt.Println(err)
			fmt.Println("from getSettings on first scan")

		}

	}
	u.Id = id
	u.Password = "u dont need it"

}

// Settings
type Settings struct {
	Id      int
	Owner   *User
	BgColor string
}

func (s *Settings) Init(id int) {
	var u = User{}
	u.Init(id)
	queryString := fmt.Sprintf(`
	SELECT 
	id, 
	bg_color 
	FROM settings 
	WHERE owner_id =  %s`, strconv.Itoa(id))

	rows := Select(queryString)

	for rows.Next() {

		err := rows.Scan(&s.Id, &s.BgColor)
		if err != nil {

			fmt.Println(err)
			fmt.Println("from getSettings on first scan")

		}
	}
	s.Owner = &u

}
