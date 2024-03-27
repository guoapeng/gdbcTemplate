package mapper

import (
	"database/sql"

	"go.uber.org/zap"
)

func NewTxRowConvertor(tx *sql.Tx, sqlstr string, args []interface{}) RowConvertor {
	return &txRowConvertor{tx: tx, sqlstr: sqlstr, args: args}
}

type txRowConvertor struct {
	args   []interface{}
	sqlstr string
	tx     *sql.Tx
	mapper RowMapper
}

func (conv *txRowConvertor) Map(rowMapper RowMapper) RowConvertor {
	conv.mapper = rowMapper
	return conv
}

func (conv *txRowConvertor) MapTo(example interface{}) RowConvertor {
	conv.mapper = NewBeanPropertyRowMapper(example).RowMapper
	return conv
}

func (conv *txRowConvertor) ToObject() interface{} {
	zap.S().Debug("query using sql: ", conv.sqlstr, " \nwith arguments", conv.args)
	preparedStmt, err := conv.tx.Prepare(conv.sqlstr)
	if err != nil {
		zap.S().Error("prepare sql statement failed, err:%v \n", err)
		return nil
	}
	datarow := preparedStmt.QueryRow(conv.args...)
	if datarow.Err() != nil {
		zap.S().Error("Encountering query error: ", datarow.Err())
		return nil
	}
	if conv.mapper != nil {
		return conv.mapper(datarow)
	} else {
		return nil
	}
}
