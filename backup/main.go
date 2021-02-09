package backup

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
	"os"
)

func main() {

	// Connect database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	ctrl := NewController(db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", ctrl.GetHello)

	e.GET("/users", ctrl.GetUsers)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// メソッドをまとめる構造体
type Controller struct {
	db *sqlx.DB
}

// Controllerのコンストラクタ
func NewController(db *sqlx.DB) Controller {
	ctrl := new(Controller)
	ctrl.db = db
	return *ctrl
}

// Hello, Worldを返す
func (ctrl Controller) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type UserDTO struct {
	Name  string `db:"name" json:"name"`
}

// ユーザー情報を返す
func (ctrl Controller) GetUser(c echo.Context) error {
	var usersDto UserDTO
	query := `SELECT name FROM users WHERE token=` + token
	err := ctrl.db.Select(&usersDto, query)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, usersDto)
}

func (ctrl Controller) CreateUser(c echo.Context) error {
	var usersDto []UserDTO
	query := `SELECT name FROM users WHERE token=`
	err := ctrl.db.Select(&usersDto, query)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, usersDto)
}


func tmp() int {
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
	return tokenString
}