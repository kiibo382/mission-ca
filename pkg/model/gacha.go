package model

type Gacha struct{
	Time int `json:"times"`
	Token string `json:"token"`
}
type Post struct{
	PostID uint `gorm:"primaryKey"`
	CharaID uint
	Chara string 
}
type Result struct {
	CharaID  uint `gorm: "primaryKey string"`
	Chara string `gorm: "string" json: "chara"`
}

type Character struct {
	ID uint `gorm:"unique;not null;PRIMARY_KEY;autoIncrement"`
  	Name string 	`gorm:"not null"`
	Percent string `gorm:"not null"`
}