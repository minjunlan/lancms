package lan

import (
	"fmt"
	"testing"
)

// func TestConn(t *testing.T) {
// 	var db = new(DB)

// 	conn, err := db.Conn("test", "root", "root")

// 	if err != nil {
// 		t.Errorf("conn 函数出错：%#v\n", err)
// 	}

// 	if conn == nil {
// 		t.Errorf("conn 为空：%#v\n", conn)
// 	}
// }

func TestQuery(t *testing.T) {

	conn, _ := Conn("test", "root", "root")

	data, err := conn.Query(`insert into test(name,age,talk) values("王五",27,"他是他各大")`)

	if err != nil {
		t.Errorf("返回错:%#v\n", err)
	}

	if data != nil {
		fmt.Printf("数据是:%#v\n", data)
	}
}
