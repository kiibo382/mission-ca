package model

type User struct {
	Id    string `xorm:"pk varchar(255)"`
	Name  string `xorm:"varchar(255)" json:"name"`
	Token string `xorm:"varchar(512)"`
}
