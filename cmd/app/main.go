package main

import (
	"github.com/kiibo382/mission-ca/pkg/controller"
	"github.com/kiibo382/mission-ca/pkg/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.POST("/create", controller.UserAdd)
			v1.GET("/get", middleware.JwtMiddleware, controller.UserGet)
			v1.GET("/list", middleware.JwtMiddleware, controller.UserList)
			v1.PUT("/update", middleware.JwtMiddleware, controller.UserUpdate)
			v1.DELETE("/delete", middleware.JwtMiddleware, controller.UserDelete)
		}
	}
	engine.Run(":3000")
}