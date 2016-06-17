package lan

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB 定义数据库类，用来操作数据库
type DB struct {
	db                    *sql.DB
	dbname, dbuser, dbpwd string
}

// Result 返回的结果
type Result []map[string]interface{}

// Conn 连接数据库
func Conn(dbname, dbuser, dbpwd string) (*DB, error) {
	db, err := sql.Open("mysql", dbuser+":"+dbpwd+"@/"+dbname)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	d := new(DB)
	d.db = db
	return d, nil
}

//Query 数据库操作语句
func (d *DB) Query(s string, args ...interface{}) (*Result, error) {

	rows, err := d.db.Query(s)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var arr Result

	for rows.Next() {
		arr1 := make(map[string]interface{}, len(columns))
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for i, col := range values {
			arr1[columns[i]] = string(col)
		}
		arr = append(arr, arr1)
	}
	return &arr, nil
}

// Close 关闭数据库连接
func (d *DB) Close() {
	d.db.Close()
}

func init() {

}
