package main

import (
	"fmt"
	"ks/api"
	"ks/data"
)

// SETTINGS

func main() {

	s := data.Settings{}
	s.Init(1)
	u := data.User{}
	u.Init(2)
	fmt.Println(u)

	fmt.Println(s.Id, s.BgColor, s.Owner, s.Owner.Name)
	api.RunServer()

}
