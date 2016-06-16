package lan

import (
	"fmt"
	"testing"
)

// func TestM(t *testing.T) {
// 	m, err := M("test")

// 	if err != nil {
// 		t.Errorf("错误值返回不为空：%#v", err)
// 	} else {
// 		fmt.Printf("正确:%#v", m)
// 	}

// }

// func TestSelect(t *testing.T) {
// 	test, err := M("test")
// 	if err != nil {
// 		t.Errorf("最后结果出错1:%#v", err)
// 	}
// 	data, err := test.Where("id = 8").Order("id desc").Limit("1,3").Select()
// 	if err != nil {
// 		t.Errorf("最后结果出错2:%#v", data)
// 	} else {
// 		fmt.Printf("结果为:%#v\n", data)
// 		fmt.Printf("列名为:%#v\n", test.cols)
// 		fmt.Printf("主键为:%#v\n", test.cols)
// 	}
// }

func TestInsert(t *testing.T) {
	test, err := M("test")
	if err != nil {
		t.Errorf("最后结果出错1:%#v", err)
	}
	var data = make(map[string]string, 4)
	data["name"] = "兰州"
	data["age"] = "45"
	data["talk"] = "我是他们的爷爷"
	data["tell"] = "18575"
	rs, err := test.Data(&data).Insert()
	if err != nil {
		t.Errorf("最后结果出错2:%#v", err)
	} else {
		fmt.Printf("结果为:%#v\n", rs)
	}
}
