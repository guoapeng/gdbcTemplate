package transaction

import (
	"database/sql"
	"log"

	"github.com/guoapeng/gdbcTemplate/mapper"
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
	log.Println("update using sql: ", sqlstr, "\nwith arguments ", args)
	result, updErr := tx.tran.Exec(sqlstr, args...)
	if updErr != nil {
		log.Println("Encountering error when execting sql: ", updErr)
	}
	return result, updErr
}

func (tx *transaction) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	log.Println("Execute using sql: ", sqlstr, "\nwith arguments ", args)
	result, err := tx.tran.Exec(sqlstr, args...)
	if err != nil {
		log.Println("Encountering error when execting sql: ", sqlstr, err)
	}
	return result, err
}

func (tx *transaction) Insert(sqlstr string, args ...interface{}) (sql.Result, error) {
	log.Println("Insert using sql: ", sqlstr, "\nwith arguments", args)
	result, err := tx.tran.Exec(sqlstr, args...)
	if err != nil {
		log.Println("Encountering error when inserting: ", err)
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
