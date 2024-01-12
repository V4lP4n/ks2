package model

type Settings struct {
	Id      int
	Owner   *User
	BgColor string
}

func (Settings) Init(UID int) {

}
