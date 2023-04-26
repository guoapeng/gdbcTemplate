package mapper

import (
	"database/sql"
	"log"

	"gdbcTemplate/datasource"
)

type RowMapper func(row *sql.Row) interface{}

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

func (conv *rowConvertor) Map(rowMapper RowMapper) RowConvertor {
	conv.mapper = rowMapper
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
