package lan

import (
	"fmt"
	"testing"
)

func TestM(t *testing.T) {
	m, err := M("test")

	if err != nil {
		t.Errorf("错误值返回不为空：%#v", err)
	} else {
		fmt.Printf("正确:%q\n", m)
		fmt.Printf("正确:%X\n", Dblink)
	}

}

func TestSelect(t *testing.T) {
	test, err := M("test")
	if err != nil {
		t.Errorf("最后结果出错1:%#v", err)
	}
	data, err := test.Where("id > 8").Order("id desc").Limit("1,3").Select()
	if err != nil {
		t.Errorf("最后结果出错2:%#v", data)
	} else {
		fmt.Printf("结果为:%#v\n", data)
		fmt.Printf("列名为:%#v\n", test.cols)
		fmt.Printf("正确:%X\n", Dblink)
	}
}

func TestInsert(t *testing.T) {
	test, err := M("test")
	if err != nil {
		t.Errorf("最后结果出错1:%#v", err)
	}
	var data = make(map[string]string, 4)
	data["name"] = "链接测试2"
	data["age"] = "98"
	data["talk"] = `<div>我是"nainai"<p class='test'>我是<a href="http://test.baidu.com">是超链接</a></p><code>\n代表了什么</code></div>`
	data["tell"] = "0000000"
	rs, err := test.Data(&data).Insert()
	if err != nil {
		t.Errorf("最后结果出错2:%#v", err)
	} else {
		fmt.Printf("结果为:%#v\n", (*rs)[0]["effectlines"])
		fmt.Printf("正确:%X\n", Dblink)
	}
}

func TestUpdate(t *testing.T) {
	test, err := M("test")
	if err != nil {
		t.Errorf("最后结果出错1:%#v", err)
	}
	var data = make(map[string]string, 4)
	data["name"] = "链接测试3"
	data["age"] = "98"
	data["talk"] = `<div>我是"nainai"<p class='test'>我是<a href="http://test.baidu.com">是超链接</a></p><code>\n代表了什么</code></div>`
	data["tell"] = "0000000"
	rs, err := test.Where("id = 6").Data(&data).Update()
	if err != nil {
		t.Errorf("最后结果出错2:%#v", err)
	} else {
		fmt.Printf("结果为:%#v\n", (*rs)[0]["effectlines"])
		fmt.Printf("正确:%X\n", Dblink)
	}
}

func TestDelete(t *testing.T) {
	test, err := M("test")
	if err != nil {
		t.Errorf("最后结果出错1:%#v", err)
	}

	rs, err := test.Where("id = 28").Delete()
	if err != nil {
		t.Errorf("最后结果出错2:%#v", err)
	} else {
		fmt.Printf("结果为:%#v\n", (*rs)[0]["effectlines"])
		fmt.Printf("正确:%X\n", Dblink)
	}

}
