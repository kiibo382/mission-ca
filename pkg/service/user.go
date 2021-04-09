package service

import (
	"fmt"

	"github.com/kiibo382/mission-ca/pkg/model"
)

type GormUserService struct{}

type GormUserDataStruct struct{
	Name  string `json:"name"`
}

func (GormUserService) GormSetUser(gormUser *model.GormUser) error {
	_, err := db.Create(gormUser)
	if err != nil {
		return err
	}
	return nil
}

func (GormUserService) GormGetUserData(gormUser *model.GormUser) GormUserDataStruct {
	GormUserData := GormUserDataStruct{}
	err := DbEngine.Where("token = ?", gormUser.Token).Cols("name").Find(&gormUser)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return GormUserData
}

func (GormUserService) GormGetUserList() []model.GormUser {
	gormUserList := make([]model.GormUser, 0)
	err := DbEngine.Distinct("name").Limit(10, 0).Find(&gormUserList)
	if err != nil {
		panic(err)
	}
	return gormUserList
}

func (GormUserService) GormUpdateUser(gormNewUser *model.GormUser) error {
	_, err := DbEngine.Id(gormNewUser.Id).Update(gormNewUser)
	if err != nil {
		return err
	}
	return nil
}

func (GormUserService) GormDeleteUser(id int) error {
	gormUser := new(model.GormUser)
	_, err := DbEngine.Id(id).Delete(gormUser)
	if err != nil {
		return err
	}
	return nil
}
