package lan

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 定义数据库类，用来操作数据库
type DB struct {
	sql.DB
	dbname, dbuser, dbpwd string
}

func (d *DB) Conn(dbname, dbuser, dbpwd string) (*sql.DB, error) {
	link := dbuser + ":" + dbpwd + "@/" + dbname
	db, err := sql.Open("mysql", link)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init() {
	fmt.Print("chu shi hau")
}
