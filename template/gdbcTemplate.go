package template

import (
	"database/sql"
	"errors"

	"github.com/guoapeng/gdbcTemplate/datasource"
	"github.com/guoapeng/gdbcTemplate/mapper"
	propsReader "github.com/guoapeng/props"
)

type GdbcTemplate interface {
	Insert(sqlstr string, args ...interface{}) (sql.Result, error)
	Update(sqlstr string, args ...interface{}) (sql.Result, error)
	Execute(sqlstr string, args ...interface{}) (sql.Result, error)
	QueryForArray(sqlstr string, args ...interface{}) mapper.RowsConvertor
	QueryRow(sqlstr string, args ...interface{}) mapper.RowConvertor
}

type gdbcTemplate struct {
	datasource datasource.DataSource
	fetchSize  int
}

func (template *gdbcTemplate) Update(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		result, updErr := db.Exec(sqlstr, args...)
		return result, updErr
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Execute(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		result, err := db.Exec(sqlstr, args...)
		return result, err
	} else {
		return nil, errors.New("failed to open db")
	}
}

func (template *gdbcTemplate) Insert(sqlstr string, args ...interface{}) (sql.Result, error) {
	if db, err := template.datasource.Open(); err == nil {
		defer db.Close()
		result, err := db.Exec(sqlstr, args...)
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
