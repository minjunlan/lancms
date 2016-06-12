package lan

import "errors"

//Model 类
type Model struct {
	tableName string   //表名称
	cols      []string //列名称
}

//M 析构函数
func M(tableName string) (*Model, error) {
	if tableName == "" {
		return nil, errors.New("表名为空")
	}
	model := new(Model)
	model.tableName = tableName
	return model, nil
}
