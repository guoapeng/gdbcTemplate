package template

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"gdbcTemplate/datasource"
	"gdbcTemplate/mapper"
	"gdbcTemplate/transaction"
	propsReader "github.com/guoapeng/props"
)

type GdbcTemplate interface {
	Insert(sqlstr string, args ...interface{}) (sql.Result, error)
	Update(sqlstr string, args ...interface{}) (sql.Result, error)
	Execute(sqlstr string, args ...interface{}) (sql.Result, error)
	QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor
	QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor
	BeginTx() (transaction.Transaction, error)
}

type gdbcTemplate struct {
	datasource datasource.DataSource
	fetchSize  int
}

func (template *gdbcTemplate) BeginTx() (transaction.Transaction, error) {

	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		log.Println("begin transaction")
		ctx := context.Background()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Println("fail to start a transaction: ", err)
		}
		return transaction.New(tx), err
	} else {
		return nil, errors.New("fail to open db")
	}
}

func (template *gdbcTemplate) Update(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		log.Println("update using sql: ", sqlstr, "\nwith arguments ", args)
		result, updErr := db.Exec(sqlstr, args...)
		if updErr != nil {
			log.Println("Encountering error when execting sql: ", updErr)
		}
		return result, updErr
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		log.Println("Execute using sql: ", sqlstr, "\nwith arguments ", args)
		result, err := db.Exec(sqlstr, args...)
		if err != nil {
			log.Println("Encountering error when execting sql: ", sqlstr, err)
		}
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Insert(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		log.Println("Insert using sql: ", sqlstr, "\nwith arguments", args)
		result, err := db.Exec(sqlstr, args...)
		if err != nil {
			log.Println("Encountering error when inserting: ", err)
		}
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor {
	return mapper.NewRowsConvertor(template.datasource, sqlstr, args)
}

func (template *gdbcTemplate) QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor {
	return mapper.NewRowConvertor(template.datasource, sqlstr, args)
}

func New(appConf propsReader.AppConfigProperties) GdbcTemplate {
	return &gdbcTemplate{datasource: datasource.NewDataSource(appConf, datasource.NewConnManager())}
}

func NewWith(ds datasource.DataSource) GdbcTemplate {
	return &gdbcTemplate{datasource: ds}
}
