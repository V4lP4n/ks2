package main

import (
	"fmt"
	"ks/model"
)

func main() {
	est1 := new(model.Estate)
	est1.Cost = 100
	est1.Id = 123

	vasya := new(model.Person)
	vasya.Name = "vasya"
	petya := new(model.Person)
	petya.Name = "petya"
	vasya.Id = 1
	vasya.Currency = 7
	petya.Id = 2
	petya.Currency = 120

	est1.Owner = vasya

	fmt.Println(vasya, petya, est1.Owner)

	// // err := est1.Trade(vasya, petya, 110)
	// // fmt.Println(err)
	// // fmt.Println(vasya, petya, est1)
	// loan := new(model.Loan)
	// fmt.Println(loan)
	// loan.Create(vasya, 1000, 365, 11005)
	// loan.AttachMoney(petya)
	// for true {

	// 	fmt.Println(loan, vasya.Currency, petya.Currency, petya.Debth)
	// 	for _, l := range petya.Loans {
	// 		l.Pay()
	// 	}
	// 	for _, d := range petya.Debth {
	// 		fmt.Println(d)
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }
	// var SQLITE db.SQL = db.SQL{"test.db", "sqlite3"}
	// var db = SQLITE.Open()
	// result, error := db.Exec(`
	// CREATE TABLE IF NOT EXISTS user (
	// 	id int PRIMARY KEY,
	// 	n_name varchar
	// 		)
	// `)

	// fmt.Println(result, error)
}
