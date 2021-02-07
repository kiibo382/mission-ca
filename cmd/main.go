package main

import (
    "fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
    "github.com/davecgh/go-spew/spew"
    "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users", users)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

	var err error
	
	secret := "secret"
	
	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "hoge",
		"iss":   "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})
	
	//Dumpを吐く
	spew.Dump(token)
	
	tokenString, err := token.SignedString([]byte(secret))
	
	fmt.Println(len(tokenString))
	fmt.Println(tokenString)
	fmt.Println(err)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func users(c echo.Context) error {
	// TODO: DBからデータを取得しJSON形式にして返す
	return c.String(http.StatusOK, "Users")
}