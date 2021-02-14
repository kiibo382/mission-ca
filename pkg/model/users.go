package model

type Users struct {
    Id    int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
    Name string `xorm:"varchar(40)" json:"name" form:"name"`
    Token string `xorm:"varchar(40)" json:"token" form:"token"`
}