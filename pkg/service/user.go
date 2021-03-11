package service

import (
	"github.com/kiibo382/mission-ca/pkg/model"
)

type UserService struct{}

func (UserService) SetUser(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService) GetUserList() []model.User {
	tests := make([]model.User, 0)
	err := DbEngine.Distinct("id", "name").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
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
