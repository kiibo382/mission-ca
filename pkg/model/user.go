package model

type User struct {
	Id    int64  `xorm:"pk autoincr int(64)" json:"id"`
	Name  string `xorm:"varchar(40)" json:"name"`
	Token string `xorm:"varchar(255)" json:"token"`
}