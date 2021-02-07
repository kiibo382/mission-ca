package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("../../web/index.html")
    data := "Hello Variables!"
    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{"data": data})
    })
    router.Run()
}