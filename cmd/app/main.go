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
			v1.GET("/list", controller.UserList)
			v1.PUT("/update", controller.UserUpdate)
			v1.DELETE("/delete", controller.UserDelete)
		}
	}
	engine.Run(":3000")
}