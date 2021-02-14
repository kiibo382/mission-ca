package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/kiibo382/mission-ca/pkg/model"
    "github.com/kiibo382/mission-ca/pkg/service"
    "strconv"
)

func UsersAdd(c *gin.Context) {
     user := model.Users{}
     err := c.ShouldBindJSON(&user)
     if err != nil{
         c.String(http.StatusBadRequest, "Bad request")
         return
     }
    userService := service.UserService{}
    err = userService.SetUsers(&user)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

func UsersGet(c *gin.Context){
    userService := service.UserService{}
    userdata := userService.GetUsers()
    c.JSONP(http.StatusOK, gin.H{
        "message": "ok",
        "data": userdata,
    })
}

func UsersUpdate(c *gin.Context){
    user := model.Users{}
    err := c.Bind(&user)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    bookService := service.UserService{}
    err = bookService.UpdateUsers(&user)
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
    userService := service.UserService{}
    err = userService.DeleteUsers(int(intId))
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}