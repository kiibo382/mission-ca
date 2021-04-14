package service

import (
	"errors"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/kiibo382/mission-ca/pkg/model"
)

var Db *gorm.DB

func init() {
	dsn := "root:rootpass@tcp(127.0.0.1:3306)/mysql_db?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	Db.AutoMigrate(&model.User{})
}