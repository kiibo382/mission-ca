package service

import (
	"fmt"

	"github.com/kiibo382/mission-ca/pkg/model"
)

type UserService struct{}

type UserDataStruct struct{
	Name  string `json:"name"`
}

func (UserService) SetUser(user *model.User) error {
	result := Db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserService) GetUserData(user *model.User) UserDataStruct {
	UserData := UserDataStruct{}
	result := Db.Where("token = ?", user.Token).Select("name").Find(user)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return UserData
}

func (UserService) GetUserList() []model.User {
	userList := make([]model.User, 0)
	result := Db.Limit(10).Select("name").Find(&userList)
	if result.Error != nil {
		panic(result.Error)
	}
	return userList
}

func (UserService) UpdateUser(gormNewUser *model.User) error {
	result := Db.Model(gormNewUser).Updates(gormNewUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserService) DeleteUser(id int) error {
	gormUser := new(model.User)
	result := Db.Delete(gormUser, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
