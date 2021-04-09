package model

type GormUser struct {
	Id    string `gorm:"primaryKey string"`
	Name  string `gorm:"string unique" json:"name"`
	Token string `gorm:"string"`
}