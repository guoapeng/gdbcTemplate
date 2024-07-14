package mapper

import (
	"database/sql"
	"strings"

	"github.com/guoapeng/gdbcTemplate/datasource"
	"go.uber.org/zap"
)

type RowsMapper func(rows *sql.Rows) interface{}

type RowsConvertor interface {
	Map(mapper RowsMapper) RowsConvertor
	MapTo(example interface{}) RowsConvertor
	ToArray() []interface{}
}

func NewRowsConvertor(dataSource datasource.DbManager, sqlstr string, args []interface{}) RowsConvertor {
	return &rowsConvertor{ds: dataSource, sqlstr: sqlstr, args: args}
}

type rowsConvertor struct {
	args       []interface{}
	sqlstr     string
	ds         datasource.DbManager
	rowsMapper RowsMapper
}

func (rowsCon *rowsConvertor) Map(rowMapper RowsMapper) RowsConvertor {
	rowsCon.rowsMapper = rowMapper
	return rowsCon
}

func (rowsCon *rowsConvertor) MapTo(example interface{}) RowsConvertor {
	rowsCon.rowsMapper = NewBeanPropertyRowMapper(example).RowsMapper
	return rowsCon
}

func (rowsCon *rowsConvertor) ToArray() []interface{} {
	if db, err := rowsCon.ds.GetDb(); err == nil {
		zap.S().Debugf("query using sql: %s \nwith arguments %v ", rowsCon.sqlstr, rowsCon.args)
		
		dataRows, err := db.Query(rowsCon.sqlstr, rowsCon.args...)

		if err != nil {
			zap.S().Errorf("query failed, err:%v \n", err)
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
