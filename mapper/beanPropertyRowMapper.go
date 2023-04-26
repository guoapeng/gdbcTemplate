package mapper

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"strings"

	interconv "github.com/mufti1/interconv/package"
)

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
