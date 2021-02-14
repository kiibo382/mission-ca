package service

import (
	"errors"
	"fmt"
	"log"
	"github.com/go-xorm/xorm"
	"github.com/kiibo382/mission-ca/pkg/model"
)

var DbEngine *xorm.Engine

func init()  {
    driverName := "mysql"
    DsName := "root:rootpass@@tcp([127.0.0.1]:3306)/mysql_db?charset=utf8"
    err := errors.New("")
    DbEngine, err = xorm.NewEngine(driverName,DsName)
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
    DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Users))
    fmt.Println("init data base ok")
}