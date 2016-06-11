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
	var x1 = make([]interface{}, 4)

	for row.Next() {
		// cols, _ := row.Columns()

		err := row.Scan(&x1[0], &x1[1], &x1[2], &x1[3])
		if err != nil {
			fmt.Printf("cols:%#v\n", err)
		}

	}
	fmt.Printf("cols:%s\n", x1[3])
	if err != nil {
		t.Errorf("连接对象时出错:%s\n", err)
	}

}
