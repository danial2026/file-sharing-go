package domain

import (
	"database/sql"
	"fmt"
    _ "github.com/Go-SQL-Driver/MySQL"
)

var (
	dbClient *sql.DB
)

func init() {
	var err error
	dbClient, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s",
		"root", "root", "127.0.0.1:3306", "users_db"))
	if err != nil {
		panic(err)
	}
}
