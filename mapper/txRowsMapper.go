package mapper

import (
	"database/sql"

	"go.uber.org/zap"
)

type TxRowsConvertor interface {
	Map(mapper RowsMapper) RowsConvertor
	MapTo(example interface{}) RowsConvertor
	ToArray() []interface{}
}

func NewTxRowsConvertor(tx *sql.Tx, sqlstr string, args []interface{}) RowsConvertor {
	return &txRowsConvertor{tx: tx, sqlstr: sqlstr, args: args}
}

type txRowsConvertor struct {
	args       []interface{}
	sqlstr     string
	tx         *sql.Tx
	rowsMapper RowsMapper
}

func (rowsCon *txRowsConvertor) Map(rowMapper RowsMapper) RowsConvertor {
	rowsCon.rowsMapper = rowMapper
	return rowsCon
}

func (rowsCon *txRowsConvertor) MapTo(example interface{}) RowsConvertor {
	rowsCon.rowsMapper = NewBeanPropertyRowMapper(example).RowsMapper
	return rowsCon
}

func (rowsCon *txRowsConvertor) ToArray() []interface{} {
	zap.S().Debug("query using sql: ", rowsCon.sqlstr, "\nwith arguments ", rowsCon.args)
	preparedStmt, err := rowsCon.tx.Prepare(rowsCon.sqlstr)
	if err != nil {
		zap.S().Error("Prepare statement failed, err:%v \n", err)
		return nil
	}
	dataRows, err := preparedStmt.Query(rowsCon.args...)
	if err != nil {
		zap.S().Errorf("Query failed, err:%v \n", err)
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
}
