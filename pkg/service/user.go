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
	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService) GetUserData(user *model.User) UserDataStruct {
	UserData := UserDataStruct{}
	err := DbEngine.Where("token = ?", user.Token).Cols("name").Find(&user)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return UserData
}

func (UserService) GetUserList() []model.User {
	userList := make([]model.User, 0)
	err := DbEngine.Distinct("name").Limit(10, 0).Find(&userList)
	if err != nil {
		panic(err)
	}
	return userList
}

func (UserService) UpdateUser(newUser *model.User) error {
	_, err := DbEngine.Id(newUser.Id).Update(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (UserService) DeleteUser(id int) error {
	user := new(model.User)
	_, err := DbEngine.Id(id).Delete(user)
	if err != nil {
		return err
	}
	return nil
}
