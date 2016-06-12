package lan

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	m, err := M("test")

	if err != nil {
		t.Errorf("错误值返回不为空：%#v", err)
	} else {
		fmt.Printf("正确:%#v", m)
	}

}
