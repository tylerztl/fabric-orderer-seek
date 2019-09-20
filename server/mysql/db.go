package mysql

import (
	"database/sql"
	"fabric-orderer-seek/server/helpers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbInfo = helpers.GetAppConf().Conf.DB

func init() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbInfo.User, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Name))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func CloseDB() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
