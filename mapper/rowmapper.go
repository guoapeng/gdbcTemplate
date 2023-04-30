package mapper

import (
	"database/sql"
	"log"

	"github.com/guoapeng/gdbcTemplate/datasource"
)

type RowMapper func(row *sql.Row) interface{}

type RowConvertor interface {
	Map(mapper RowMapper) RowConvertor
	MapTo(example interface{}) RowConvertor
	ToObject() interface{}
}

func NewRowConvertor(datasource datasource.DbManager, sqlstr string, args []interface{}) RowConvertor {
	return &rowConvertor{dbM: datasource, sqlstr: sqlstr, args: args}
}

type rowConvertor struct {
	args   []interface{}
	sqlstr string
	dbM    datasource.DbManager
	mapper RowMapper
}

func (conv *rowConvertor) Map(rowMapper RowMapper) RowConvertor {
	conv.mapper = rowMapper
	return conv
}

func (conv *rowConvertor) MapTo(example interface{}) RowConvertor {
	conv.mapper = NewBeanPropertyRowMapper(example).RowMapper
	return conv
}

func (conv *rowConvertor) ToObject() interface{} {
	if db, err := conv.dbM.GetDb(); err == nil {
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
