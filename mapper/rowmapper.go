package mapper

import (
	"database/sql"

	"github.com/guoapeng/gdbcTemplate/datasource"
	"go.uber.org/zap"
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
		zap.S().Debug("query using sql: ", conv.sqlstr, " \nwith arguments", conv.args)
		preparedStmt, err := db.Prepare(conv.sqlstr)
		if err != nil {
			zap.S().Errorf("prepare sql statement failed, err:%v \n", err)
			return nil
		}
		datarow := preparedStmt.QueryRow(conv.args...)
		if datarow.Err() != nil {
			zap.S().Error("Encountering query error: ", datarow.Err())
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
