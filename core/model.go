package lan

import (
	"errors"
	"fmt"
	"strconv"
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
	k1        []rune   //要插入修改的关键字
	v1        []rune   //要插入修改的值
}

var Dblink *DB
var dbname, dbuser, dbpwd string

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
		Dblink, _ = Conn(dbname, dbuser, dbpwd)
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

//Insert 插入函数
func (m *Model) Insert() (*Result, error) {
	sql := m.getSql(insertSQL)

	_, err := Dblink.Query(sql)
	if err != nil {
		return nil, err
	}
	r, _ := Dblink.Query("SELECT ROW_COUNT() as effectlines")
	if i, _ := strconv.Atoi((*r)[0]["effectlines"].(string)); i <= 0 {
		return nil, errors.New("插入失败")
	}

	m.rows = r
	return m.rows, nil
}

//Update 插入函数
func (m *Model) Update() (*Result, error) {
	sql := m.getSql(updateSQL)

	_, err := Dblink.Query(sql)
	if err != nil {
		return nil, err
	}
	r, _ := Dblink.Query("SELECT ROW_COUNT() as effectlines")
	if i, _ := strconv.Atoi((*r)[0]["effectlines"].(string)); i <= 0 {
		return nil, errors.New("更新失败")
	}

	m.rows = r
	m.where = ""
	return m.rows, nil
}

//delete 删除函数
func (m *Model) Delete() (*Result, error) {
	sql := m.getSql(deleteSQL)

	_, err := Dblink.Query(sql)
	if err != nil {
		return nil, err
	}
	r, _ := Dblink.Query("SELECT ROW_COUNT() as effectlines")
	if i, _ := strconv.Atoi((*r)[0]["effectlines"].(string)); i <= 0 {
		return nil, errors.New("删除失败")
	}

	m.rows = r
	m.where = ""
	return m.rows, nil
}

func (m *Model) Data(data *map[string]string) *Model {
	d := *data
	var key, value string

	for k, v := range d {
		key += k + `,`
		v = changeChart(v)
		value += `"` + v + `",`
	}
	m.k1 = []rune(key)
	m.v1 = []rune(value)

	return m
}

func changeChart(src string) string {
	var s string
	s = strings.Replace(src, `\`, `\\`, -1)
	s = strings.Replace(s, `"`, `\"`, -1)
	return s
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

	switch pre {
	case selectSQL:
		fmt.Printf("sql2:%#v\n", pre)
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
		m.stat = `(` + string(m.k1[:len(m.k1)-1]) + `) values(` + string(m.v1[:len(m.v1)-1]) + `)`
		m.stat = "Insert into " + m.tableName + m.stat
	case updateSQL:
		var s string

		karr := strings.Split(string(m.k1), `,`)
		varr := strings.Split(string(m.v1), `,`)

		for i := 0; i < len(karr)-1; i++ {
			if karr[i] != "," {
				s += karr[i] + `=` + varr[i] + `, `
			}
		}
		if m.where != "" {
			m.stat = "Update " + m.tableName + " set " + strings.TrimRight(s, `, `) + ` where ` + m.where
		} else {
			m.stat = "Update " + m.tableName + " set " + strings.TrimRight(s, `, `)
		}
	case deleteSQL:
		if m.where != "" {
			m.stat = "Delete from " + m.tableName + ` where ` + m.where
		} else {
			m.stat = "Delete from " + m.tableName
		}

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

func SetDNS(name, user, pwd string) {
	dbname = name
	dbuser = user
	dbpwd = pwd
}

// func init() {
// 	if dbname == "" {
// 		data := make([]byte, 100)
// 		f, err := os.Open("../conf/db.conf")
// 		if err != nil {
// 			fmt.Errorf("读文件出错:%#v", err)
// 		}
// 		count, err := f.Read(data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		s := strings.Split(string(data[:count]), "\r\n")
// 		for i := 0; i < len(s); i++ {
// 			t := strings.Split(s[i], ":")
// 			if t[0] == "dbname" {
// 				dbname = t[1]
// 			}
// 			if t[0] == "dbuser" {
// 				dbuser = t[1]
// 			}
// 			if t[0] == "dbpwd" {
// 				dbpwd = t[1]
// 			}
// 		}

// 	}
// }
