package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/kiibo382/pkg/controller"
    "github.com/kiibo382/pkg/middleware"
)

func main(){
    engine := gin.Default()
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    // CRUD 書籍
    bookEngine := engine.Group("/users")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/add", controller.UsersAdd)
            v1.GET("/list", controller.UsersList)
            v1.PUT("/update", controller.UsersUpdate)
            v1.DELETE("/delete", controller.UsersDelete)
        }
    }
    engine.Run(":3000")
}