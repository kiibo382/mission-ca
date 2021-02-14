package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/kiibo382/mission-ca/pkg/controller"
    "github.com/kiibo382/mission-ca/pkg/middleware"
)

func main(){
    router := gin.Default()
    // Middleware
    router.Use(middleware.RecordUaAndTime)
    // Routing
    usersRouter := router.Group("/users")
    {
        v1 := usersRouter.Group("/v1")
        {
            v1.POST("/create", controller.UsersAdd)
            v1.GET("/get", controller.UsersGet)
            v1.PUT("/update", controller.UsersUpdate)
            v1.DELETE("/delete", controller.UsersDelete)
        }
    }

    // This handler will match /user/john but will not match /user/ or /user
    // router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

    // Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
    router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath() == "/user/:name/*action" // true
	// })

	router.Run(":8080")
}