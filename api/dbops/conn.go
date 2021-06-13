package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "admin:l5341500@tcp(database-2.cx0zufsjusrp.us-east-2.rds.amazonaws.com:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}
}
