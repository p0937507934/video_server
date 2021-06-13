package dbops

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "admin:l5341500@tcp(database-2.cx0zufsjusrp.us-east-2.rds.amazonaws.com:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}

	return dbConn
}

func AddUserCredential(loginName, pwd string) error {

}
func GetUserCredentail(loginName string) (string, error) {

}

func main() {
	db := openConn()
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connecting is ok.")
}
