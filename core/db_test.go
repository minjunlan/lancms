package lan

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	var db = new(DB)

	conn, err := db.Conn("test", "root", "root")

	row, err2 := conn.Query("select * from test")
	if err2 != nil {
		fmt.Printf("%v\n", err2)
	}
	fmt.Printf("%v\n", *db)
	//db.Close()
	var cols interface{}
	for row.Next() {
		row.Scan(cols)
	}

	if err != nil {
		t.Errorf("连接对象时出错:%s\n", err)
	}

}
