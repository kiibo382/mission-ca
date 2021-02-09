package db

import  (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *sql.DB

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	DbConnection, _ := sql.Open("mysql",dsn)
	defer DbConnection.Close()
	
	cmd := "INSER INTO users (name) VALUES (?)"
	_, err := DbConnection.Exec(cmd, "Mike")
	if err != nil {
		log.Fatalln(err)
	}


}