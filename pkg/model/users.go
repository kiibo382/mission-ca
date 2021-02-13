package model

type Users struct {
    Id    int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
    Name string `xorm:"varchar(40)" json:"title" form:"title"`
    Token string `xorm:"varchar(40)" json:"content" form:"content"`
}