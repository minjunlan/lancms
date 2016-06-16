package lan

import (
	"errors"
	"fmt"
	"strings"
)

//Model 类
type Model struct {
	tableName string   //表名称
	cols      []string //列名称
	rows      *Result  //行数据
	priKey    string   //主键
	stat      string   //sql语句
	where     string   //where条件
	group     string   //group by
	having    string   //having
	order     string   //order by
	limit     string   //limit 2,3 limit off, n
}

var Dblink *DB

type sqltype int

const (
	selectSQL sqltype = iota
	insertSQL
	updateSQL
	deleteSQL
)

//M 析构函数
func M(tableName string) (*Model, error) {
	if tableName == "" {
		return nil, errors.New("表名为空")
	}
	if Dblink == nil {
		Dblink, _ = Conn("test", "root", "root")
		if Dblink == nil {
			return nil, errors.New("连接数据库错误")
		}
	}
	cols, err := Dblink.Query("show columns from " + tableName)
	if err != nil {
		return nil, errors.New("连接表时出错，或这个表不存在!")
	}

	model := new(Model)
	model.tableName = tableName
	model.cols = model.getCols(cols)
	return model, nil
}

//Select 查询函数
func (m *Model) Select() (*Result, error) {
	sql := m.getSql(selectSQL)
	r, err := Dblink.Query(sql)
	if err != nil {
		return nil, err
	}
	m.rows = r

	m.stat = ""
	m.where = ""
	m.group = ""
	m.having = ""
	m.limit = ""
	return m.rows, nil
}

//Insert 查询函数
func (m *Model) Insert() (*Result, error) {
	sql := m.getSql(insertSQL)

	_, err := Dblink.Query(sql)
	if err != nil {
		return nil, err
	}
	r, err := Dblink.Query("SELECT ROW_COUNT()")
	var rs Result
	rs[0]["affectLine"] = r
	fmt.Printf("影响行数为:%#v", r)
	m.rows = &rs
	return m.rows, nil
}

func (m *Model) Data(data *map[string]string) *Model {
	d := *data
	var key, value string

	for k, v := range d {
		key += `'` + k + `',`
		value += `'` + v + `',`
	}
	k1 := []rune(key)
	v1 := []rune(value)
	m.stat = `(` + string(k1[:len(k1)-1]) + `) values(` + string(v1[:len(v1)-1]) + `)`
	return m
}

//Where条件语句
func (m *Model) Where(sql string) *Model {
	if sql == "" {
		return m
	}
	m.where = sql
	return m
}

//Group条件语句
func (m *Model) Group(sql string) *Model {
	if sql == "" {
		return m
	}
	m.group = " group by " + sql
	return m
}

//Having条件语句
func (m *Model) Having(sql string) *Model {
	if sql == "" {
		return m
	}
	m.having = " having " + sql
	return m
}

//Order条件语句
func (m *Model) Order(sql string) *Model {
	if sql == "" {
		return m
	}
	m.order = " order by " + sql
	return m
}

//limit
func (m *Model) Limit(sql string) *Model {
	if sql == "" {
		m.limit = " limit 1 "
		return m
	}
	m.limit = " limit " + sql
	return m
}

//getSql得到sql语句
func (m *Model) getSql(pre sqltype) string {
	fmt.Printf("sql:%#v\n", pre)
	
		switch pre {
		case selectSQL:
			if m.where != "" {
				m.stat = `Select * from ` + m.tableName + " where " + m.where
			} else {
				m.stat = `Select * from ` + m.tableName
			}

			if m.group != "" {
				m.stat += m.group
			}
			if m.having != "" {
				m.stat += m.having
			}
			if m.order != "" {
				m.stat += m.order
			}
			if m.limit != "" {
				m.stat += m.limit
			}
		case insertSQL:

			m.stat = "Insert into " + m.tableName + m.stat
			m.stat = strings.Replace(m.stat, `\`, "", 1)

		}
	

	return m.stat
}

func (m *Model) getCols(cols *Result) []string {
	var s []string
	var comns = *cols
	for i := 0; i < len(comns); i++ {
		for k, v := range comns[i] {
			if k == "Field" {
				s = append(s, v.(string))
			}
			if k == "Key" && v == "PRI" {
				m.priKey = comns[i]["Field"].(string)
			}
		}
	}
	return s
}
