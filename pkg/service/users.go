package service

import (
	"github.com/kiibo382/mission-ca/pkg/model"
)

type userService struct {}

func (userService) SetBook(book *model.Users) error {
    _, err := DbEngine.Insert(book)
    if err!= nil{
         return  err
    }
    return nil
}


func (userService) GetBookList() []model.Users {
    tests := make([]model.Users, 0)
    err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}

func (userService) UpdateUsers(newBook *model.Users) error {
    _, err := DbEngine.Id(newBook.Id).Update(newBook)
    if err != nil {
        return err
    }
    return nil
}

func (userService) DeleteBook(id int) error {
    book := new(model.Users)
    _, err := DbEngine.Id(id).Delete(book)
    if err != nil{
        return err
    }
    return nil
}