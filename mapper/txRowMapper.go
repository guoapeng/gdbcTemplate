package mapper

import (
	"database/sql"
	"log"
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
	log.Println("query using sql: ", conv.sqlstr, " \nwith arguments", conv.args)
	datarow := conv.tx.QueryRow(conv.sqlstr, conv.args...)
	if datarow.Err() != nil {
		log.Println("Encountering query error: ", datarow.Err())
	}
	if conv.mapper != nil {
		return conv.mapper(datarow)
	} else {
		return nil
	}
}