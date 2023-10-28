package template

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/guoapeng/gdbcTemplate/datasource"
	"github.com/guoapeng/gdbcTemplate/mapper"
	"github.com/guoapeng/gdbcTemplate/transaction"

	propsReader "github.com/guoapeng/props"
)

type GdbcTemplate interface {
	Insert(sqlstr string, args ...interface{}) (sql.Result, error)
	Update(sqlstr string, args ...interface{}) (sql.Result, error)
	Execute(sqlstr string, args ...interface{}) (sql.Result, error)
	QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor
	QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor
	BeginTx() (transaction.Transaction, error)
	CloseDB() error
}

type gdbcTemplate struct {
	dbM       datasource.DbManager
	fetchSize int
}

func (template *gdbcTemplate) BeginTx() (transaction.Transaction, error) {

	if db, err := template.dbM.GetDb(); err == nil {
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
	if db, err := template.dbM.GetDb(); err == nil {
		log.Println("update using sql: ", sqlstr, "\nwith arguments ", args)
		preparedStmt, updErr := db.Prepare(sqlstr)
		if updErr != nil {
			log.Printf("prepare sql statement failed, err:%v \n", err)
			return nil, err
		}
		result, updErr := preparedStmt.Exec(args...)
		if updErr != nil {
			log.Println("Encountering error when execting sql: ", updErr)
		}
		return result, updErr
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.dbM.GetDb(); err == nil {
		log.Println("Execute using sql: ", sqlstr, "\nwith arguments ", args)
		preparedStmt, err := db.Prepare(sqlstr)
		if err != nil {
			log.Printf("prepare sql statement failed, err:%v \n", err)
			return nil, err
		}
		result, err := preparedStmt.Exec(args...)
		if err != nil {
			log.Println("Encountering error when execting sql: ", sqlstr, err)
		}
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Insert(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.dbM.GetDb(); err == nil {
		log.Println("Insert using sql: ", sqlstr, "\nwith arguments", args)
		preparedStmt, err := db.Prepare(sqlstr)
		if err != nil {
			log.Printf("prepare sql statement failed, err:%v \n", err)
			return nil, err
		}
		result, err := preparedStmt.Exec(args...)
		if err != nil {
			log.Println("Encountering error when inserting: ", err)
		}
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor {
	return mapper.NewRowsConvertor(template.dbM, sqlstr, args)
}

func (template *gdbcTemplate) QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor {
	return mapper.NewRowConvertor(template.dbM, sqlstr, args)
}

func (template *gdbcTemplate) CloseDB() error {
	return template.dbM.Close()
}

func New(appConf propsReader.AppConfigProperties) GdbcTemplate {
	return &gdbcTemplate{dbM: datasource.NewDbManager(datasource.NewDataSource(appConf))}
}

func NewWith(ds datasource.DataSource) GdbcTemplate {
	return &gdbcTemplate{dbM: datasource.NewDbManager(ds)}
}

func NewWithDsAndOpenDbFn(ds datasource.DataSource, openDbFunc datasource.OpenDb) GdbcTemplate {
	return &gdbcTemplate{dbM: datasource.NewDbManagerWithOpenDbFn(ds, openDbFunc)}
}
