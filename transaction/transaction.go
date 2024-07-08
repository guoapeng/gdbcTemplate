package transaction

import (
	"database/sql"

	"github.com/guoapeng/gdbcTemplate/mapper"
	"go.uber.org/zap"
)

type Transaction interface {
	Insert(sqlstr string, args ...interface{}) (sql.Result, error)
	Update(sqlstr string, args ...interface{}) (sql.Result, error)
	Execute(sqlstr string, args ...interface{}) (sql.Result, error)
	QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor
	QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor
	Commit() error
	Rollback() error
}

type transaction struct {
	tran *sql.Tx
}

func (tx *transaction) Update(sqlstr string, args ...interface{}) (sql.Result, error) {
	zap.S().Debug("update using sql: ", sqlstr, "\nwith arguments ", args)
	preparedStmt, err := tx.tran.Prepare(sqlstr)
	if err != nil {
		zap.S().Errorf("prepare sql statement failed, err:%v \n", err)
		return nil, err
	}
	result, updErr := preparedStmt.Exec(args...)
	if updErr != nil {
		zap.S().Error("Encountering error when execting sql: ", updErr)
		return result, updErr
	}
	return result, updErr
}

func (tx *transaction) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	zap.S().Debug("Execute using sql: ", sqlstr, "\nwith arguments ", args)
	preparedStmt, err := tx.tran.Prepare(sqlstr)
	if err != nil {
		zap.S().Errorf("prepare sql statement failed, err:%v \n", err)
		return nil, err
	}
	result, err := preparedStmt.Exec(args...)
	if err != nil {
		zap.S().Error("Encountering error when execting sql: ", sqlstr, err)
		return result, err
	}
	return result, err
}

func (tx *transaction) Insert(sqlstr string, args ...interface{}) (sql.Result, error) {
	zap.S().Debug("Insert using sql: ", sqlstr, "\nwith arguments", args)
	preparedStmt, err := tx.tran.Prepare(sqlstr)
	if err != nil {
		zap.S().Errorf("prepare sql statement failed, err:%v \n", err)
		return nil, err
	}
	result, err := preparedStmt.Exec(args...)
	if err != nil {
		zap.S().Error("Encountering error when inserting: ", err)
		return result, err
	}
	return result, err
}

func (tx *transaction) QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor {
	return mapper.NewTxRowsConvertor(tx.tran, sqlstr, args)
}

func (tx *transaction) QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor {
	return mapper.NewTxRowConvertor(tx.tran, sqlstr, args)
}

func (tx *transaction) Commit() error {
	return tx.tran.Commit()
}

func (tx *transaction) Rollback() error {
	return tx.tran.Rollback()
}

func New(tx *sql.Tx) Transaction {
	return &transaction{tran: tx}
}
