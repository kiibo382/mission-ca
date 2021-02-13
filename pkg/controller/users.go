package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/kiibo382/mission-ca/pkg/model"
    // "github.com/kiibo382/mission-ca/pkg/service"
    "strconv"
)

func UsersAdd(c *gin.Context) {
     user := model.Users{}
     err := c.Bind(&user)
     if err != nil{
         c.String(http.StatusBadRequest, "Bad request")
         return
     }
    userService :=service.userService{}
    err = userService.SetBook(&user)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

func UsersList(c *gin.Context){
    bookService :=service.userService{}
    BookLists := bookService.GetBookList()
    c.JSONP(http.StatusOK, gin.H{
        "message": "ok",
        "data": BookLists,
    })
}

func UsersUpdate(c *gin.Context){
    book := model.Book{}
    err := c.Bind(&book)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    bookService :=service.userService{}
    err = bookService.UpdateBook(&book)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

func UsersDelete(c *gin.Context){
    id := c.PostForm("id")
    intId, err := strconv.ParseInt(id, 10, 0)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    bookService :=service.userService{}
    err = bookService.DeleteBook(int(intId))
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}