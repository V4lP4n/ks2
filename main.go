package main

import (
	"fmt"
	"ks/model"
	"time"
)

func main() {
	est1 := new(model.Estate)
	est1.Cost = 100
	est1.Id = 123

	vasya := new(model.Company)
	vasya.Title = "vasya"
	petya := new(model.Company)
	petya.Title = "petya"
	vasya.Id = 1
	vasya.Currency = 7
	petya.Id = 2
	petya.Currency = 120

	est1.Owner = vasya

	fmt.Println(vasya, petya, est1.Owner)

	err := est1.Trade(vasya, petya, 110)
	fmt.Println(err)
	fmt.Println(vasya, petya, est1.Owner)
	loan := new(model.Loan)
	fmt.Println(loan)
	loan.Create(vasya, 1000, 10, 20)
	loan.AttachMoney(petya)
	for true {

		fmt.Println(loan, vasya.Currency, petya.Currency)
		loan.Pay()

		time.Sleep(1 * time.Second)
	}
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
