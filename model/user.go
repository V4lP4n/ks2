package model

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (p *User) Init(id int) {
	queryString := fmt.Sprintf(`
	SELECT
	name,
	from user
	where id= %s`, strconv.Itoa(p.Id))

	// rows := db.Select(queryString)

	// for rows.Next() {

	// 	err := rows.Scan(&p.Username)
	// 	if err != nil {

	// 		fmt.Println(err)
	// 		fmt.Println("from getSettings on first scan")

	// 	}
	// }
	fmt.Println(queryString)

}
