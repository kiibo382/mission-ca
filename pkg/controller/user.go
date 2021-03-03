package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kiibo382/mission-ca/pkg/model"
	"github.com/kiibo382/mission-ca/pkg/service"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func UserAdd(c *gin.Context) {
	user := model.User{}
	// token := uuid.New().String()
	fmt.Println(c)
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.SetUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserList(c *gin.Context) {
	userService := service.UserService{}
	UserLists := userService.GetUserList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    UserLists,
	})
}

func UserUpdate(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.UpdateUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.DeleteUser(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
