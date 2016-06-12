package lan

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 定义数据库类，用来操作数据库
type DB struct {
	db                    *sql.DB
	dbname, dbuser, dbpwd string
	data                  []map[string]interface{}
}

// Result 返回的结果
type Result struct {
}

//连接数据库
func (d *DB) Conn(dbname, dbuser, dbpwd string) (*DB, error) {
	db, err := sql.Open("mysql", dbuser+":"+dbpwd+"@/"+dbname)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	d.db = db
	return d, nil
}

//获取结果
func (d *DB) Result() *DB {
	fmt.Println("resuilt")
	return d
}

func (d *DB) Insert(sql string, args ...interface{}) *Result {

}
func (d *DB) Select(sql string, args ...interface{}) *Result {

}
func (d *DB) Delete(sql string, args ...interface{}) *Result {

}
func (d *DB) Update(sql string, args ...interface{}) *Result {

}
func init() {
	fmt.Print("chu shi hau")
}
