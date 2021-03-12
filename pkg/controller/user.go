package controller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kiibo382/mission-ca/pkg/model"
	"github.com/kiibo382/mission-ca/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	jwt "github.com/dgrijalva/jwt-go"
)

func UserAdd(c *gin.Context) {
	user := model.User{}
	user.Id = uuid.New().String()
	err := c.Bind(&user)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = user.Id
	claims["name"] = user.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	user.Token = tokenString

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	userService := service.UserService{}
	err = userService.SetUser(&user)
	if err != nil {
		c.String(http.StatusConflict, "Sorry, username alredy exists. Please change username.")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
}

func UserList(c *gin.Context) {
	userService := service.UserService{}
	UserLists := userService.GetUserList()
	c.JSONP(http.StatusOK, gin.H{
		"data": UserLists,
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
