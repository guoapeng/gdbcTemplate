package template

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guoapeng/gdbcTemplate/datasource"
	"github.com/guoapeng/gdbcTemplate/mapper"
	"github.com/guoapeng/props"
	"log"
)

type GdbcTemplate interface {
	Insert(sql string, args ...interface{}) (sql.Result, error)
	Update(sql string, args ...interface{}) (sql.Result, error)
	Execute(sql string, args ...interface{}) (sql.Result, error)
	QueryForArray(sql string, args ...interface{}) mapper.RowsConvertor
	QueryRow(sql string, args ...interface{}) mapper.RowConvertor
}

type gdbcTemplate struct {
	datasource datasource.DataSource
	fetchSize  int
}

func (template *gdbcTemplate) Update(sql string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		result, updErr := db.Exec(sql, args...)
		return result, updErr
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		result, err := db.Exec(sqlstr, args...)
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}



func (template *gdbcTemplate) Insert(sql string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		result, err := db.Exec(sql, args...)
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) QueryForArray(sql string, args ...interface{}) mapper.RowsConvertor {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		datarows, err := db.Query(sql, args...)
		if err != nil {
			log.Fatal(err)
			log.Printf("scan failed, err:%v \n", err)
		}
		return mapper.NewRowsConvertor(datarows)
	} else {
		return mapper.NewRowsConvertor(nil)
	}
}

func (template *gdbcTemplate) QueryRow(sql string, args ...interface{}) (mapper.RowConvertor) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		datarow := db.QueryRow(sql, args...)
		return mapper.NewRowConvertor(datarow)
	} else {
		return mapper.NewRowConvertor(nil)
	}
}

func New(appName string) GdbcTemplate {
	if appCinfig, err := propsReader.New(appName); err == nil {
		return &gdbcTemplate{datasource: datasource.NewDataSource(appCinfig)}
	} else {
		panic("failed to create gdbcTemplate")
		return nil
	}
}

func NewWith(ds datasource.DataSource) GdbcTemplate {
	return &gdbcTemplate{datasource: ds}
}
