package lan

import "testing"

func TestConn(t *testing.T) {
	var db = new(DB)

	conn, err := db.Conn("test", "root", "root")

	if err != nil {
		t.Errorf("conn 函数出错：%#v\n", err)
	}

	if conn == nil {
		t.Errorf("conn 为空：%#v\n", conn)
	}
}
