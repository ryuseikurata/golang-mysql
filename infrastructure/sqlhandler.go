package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "docker:docker@tcp(0.0.0.0:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(sqlHandler)
	sqlHandler.conn = conn
	return sqlHandler
}