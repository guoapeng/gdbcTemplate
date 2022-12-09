package mapper

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/guoapeng/gdbcTemplate/datasource"
	interconv "github.com/mufti1/interconv/package"
)

type RowsMapper func(rows *sql.Rows) interface{}

type RowMapper func(row *sql.Row) interface{}

type BeanPropertyRowMapper interface {
	RowMapper(row *sql.Row) interface{}
	RowsMapper(rows *sql.Rows) interface{}
}

type beanPropertyRowMapper struct {
	Typ reflect.Type
}

func NewBeanPropertyRowMapper(example interface{}) BeanPropertyRowMapper {
	return &beanPropertyRowMapper{Typ: reflect.TypeOf(example)}
}

func ColumnMapRowMapper(rows *sql.Rows) interface{} {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}
	_ = rows.Scan(cache...)
	item := make(map[string]interface{})
	for i, data := range cache {
		item[strings.ToUpper(columns[i])] = *data.(*interface{}) //取实际类型
	}
	return item
}

func (mapper *beanPropertyRowMapper) RowMapper(row *sql.Row) interface{} {
	out := reflect.New(mapper.Typ).Elem().Interface()
	ss := reflect.ValueOf(out).Elem()
	columnLength := ss.Len()
	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}
	_ = row.Scan(cache...)

	item := make(map[string]interface{})
	for i, data := range cache {
		item[strings.ToUpper(ss.Type().Field(i).Name)] = *data.(*interface{})
	}
	for i := 0; i < ss.NumField(); i++ {
		val, _ := interconv.ParseString(item[strings.ToUpper(ss.Type().Field(i).Tag.Get("sql"))])
		name := ss.Type().Field(i).Name
		log.Printf("tag:%s, tag value:%s, filed name:%s", ss.Type().Field(i).Tag.Get("sql"), val, name)
		switch ss.Field(i).Kind() {
		case reflect.String:
			ss.FieldByName(name).SetString(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.Atoi(val)
			//  fmt.Println("i:", i, name)
			if err != nil {
				log.Printf("can't not atoi:%v", val)
				continue
			}
			ss.FieldByName(name).SetInt(int64(i))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.Atoi(val)
			//  fmt.Println("i:", i, name)
			if err != nil {
				log.Printf("can't not atoi:%v", val)
				continue
			}
			ss.FieldByName(name).SetUint(uint64(i))
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Printf("can't not ParseFloat:%v", val)
				continue
			}
			ss.FieldByName(name).SetFloat(f)
		default:
			log.Printf("unknown type:%+v", ss.Field(i).Kind())
		}
	}
	return out
}

func (mapper *beanPropertyRowMapper) RowsMapper(rows *sql.Rows) interface{} {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index := range cache { //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	_ = rows.Scan(cache...)
	out := reflect.New(mapper.Typ).Interface()
	ss := reflect.ValueOf(out).Elem()
	item := make(map[string]interface{})
	for i, data := range cache {
		item[strings.ToUpper(columns[i])] = *data.(*interface{})
	}
	for i := 0; i < ss.NumField(); i++ {
		val, err := interconv.ParseString(item[strings.ToUpper(ss.Type().Field(i).Tag.Get("sql"))])
		log.Println(err)
		name := ss.Type().Field(i).Name
		log.Printf("tag:%s, tag value:%s, field name:%s", ss.Type().Field(i).Tag.Get("sql"), val, name)
		switch ss.Field(i).Kind() {
		case reflect.String:
			ss.FieldByName(name).SetString(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("can't not atoi:%v", val)
				continue
			}
			ss.FieldByName(name).SetInt(int64(i))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("can't not atoi:%v", val)
				continue
			}
			ss.FieldByName(name).SetUint(uint64(i))
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Printf("can't not ParseFloat:%v", val)
				continue
			}
			ss.FieldByName(name).SetFloat(f)
		default:
			log.Printf("unknown type:%+v", ss.Field(i).Kind())
		}
	}
	return out
}

type RowConvertor interface {
	Map(mapper RowMapper) RowConvertor
	MapTo(example interface{}) RowConvertor
	ToObject() interface{}
}

func NewRowConvertor(datasource datasource.DataSource, sqlstr string, args []interface{}) RowConvertor {
	return &rowConvertor{ds: datasource, sqlstr: sqlstr, args: args}
}

type rowConvertor struct {
	args   []interface{}
	sqlstr string
	ds     datasource.DataSource
	mapper RowMapper
}

func (conv *rowConvertor) Map(mapper RowMapper) RowConvertor {
	conv.mapper = mapper
	return conv
}

func (conv *rowConvertor) MapTo(example interface{}) RowConvertor {
	conv.mapper = NewBeanPropertyRowMapper(example).RowMapper
	return conv
}

func (conv *rowConvertor) ToObject() interface{} {
	if db, err := conv.ds.Open(); err == nil {
		defer db.Close()
		log.Println("query using sql: ", conv.sqlstr, " \nwith arguments", conv.args)
		datarow := db.QueryRow(conv.sqlstr, conv.args...)
		if datarow.Err() != nil {
			log.Println("Encountering query error: ", datarow.Err())
		}
		if conv.mapper != nil {
			return conv.mapper(datarow)
		} else {
			return nil
		}
	} else {
		return nil
	}
}

type RowsConvertor interface {
	Map(mapper RowsMapper) RowsConvertor
	MapTo(example interface{}) RowsConvertor
	ToArray() []interface{}
}

func NewRowsConvertor(dataSource datasource.DataSource, sqlstr string, args []interface{}) RowsConvertor {
	return &rowsConvertor{ds: dataSource, sqlstr: sqlstr, args: args}
}

type rowsConvertor struct {
	args       []interface{}
	sqlstr     string
	ds         datasource.DataSource
	rowsMapper RowsMapper
}

func (rowsCon *rowsConvertor) Map(mapper RowsMapper) RowsConvertor {
	rowsCon.rowsMapper = mapper
	return rowsCon
}

func (rowsCon *rowsConvertor) MapTo(example interface{}) RowsConvertor {
	rowsCon.rowsMapper = NewBeanPropertyRowMapper(example).RowsMapper
	return rowsCon
}

func (rowsCon *rowsConvertor) ToArray() []interface{} {
	if db, err := rowsCon.ds.Open(); err == nil {
		defer db.Close()
		log.Println("query using sql: ", rowsCon.sqlstr)
		dataRows, err := db.Query(rowsCon.sqlstr, rowsCon.args...)
		if err != nil {
			log.Printf("scan failed, err:%v \n", err)
			return nil
		}
		defer dataRows.Close()
		items := []interface{}{}
		if rowsCon.rowsMapper == nil {
			for dataRows.Next() {
				item := ColumnMapRowMapper(dataRows)
				items = append(items, item)
			}
		} else {
			for dataRows.Next() {
				item := rowsCon.rowsMapper(dataRows)
				items = append(items, item)
			}
		}
		return items
	} else {
		return nil
	}

}
