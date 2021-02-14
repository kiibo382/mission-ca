package service

import (
	"github.com/kiibo382/mission-ca/pkg/model"
)

type UserService struct {}

func (UserService) SetUsers(user *model.Users) error {
    _, err := DbEngine.Insert(user)
    if err!= nil{
         return  err
    }
    return nil
}

func (UserService) GetUsers() []model.Users {
    tests := make([]model.Users, 0)
    err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}

func (UserService) UpdateUsers(newUser *model.Users) error {
    _, err := DbEngine.Id(newUser.Id).Update(newUser)
    if err != nil {
        return err
    }
    return nil
}

func (UserService) DeleteUsers(id int) error {
    user := new(model.Users)
    _, err := DbEngine.Id(id).Delete(user)
    if err != nil{
        return err
    }
    return nil
}