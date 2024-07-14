package template_test

import (
	"database/sql"
	"testing"

	"github.com/guoapeng/gdbcTemplate/mocks"
	"github.com/guoapeng/gdbcTemplate/template"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestGdbcTemplateSuite(t *testing.T) {
	suite.Run(t, new(GdbcTemplateSuite))
}

type GdbcTemplateSuite struct {
	test *testing.T
	suite.Suite
	dataSource   *mocks.DataSource
	gdbcTemplate template.GdbcTemplate
	openDbFn     *mocks.OpenDb
}

func (suite *GdbcTemplateSuite) T() *testing.T {
	return suite.test
}

func (suite *GdbcTemplateSuite) SetT(t *testing.T) {
	suite.test = t
	suite.dataSource = &mocks.DataSource{}
	suite.openDbFn = mocks.NewOpenDb(t)
	suite.gdbcTemplate = template.NewWithDsAndOpenDbFn(suite.dataSource, suite.openDbFn.Execute)
}

func (suite *GdbcTemplateSuite) TestInsert() {
	db, mockobj, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.openDbFn.On("Execute", mock.Anything, mock.Anything).Return(db, nil)
	suite.dataSource.On("GetDriverName").Return("PGX")
	suite.dataSource.On("FormatDsn").Return("postgres://root:12345@127.0.0.1:5432/testdb", "postgres://root:******@127.0.0.1:5432/testdb")
	mockobj.ExpectExec("xxx")
	suite.gdbcTemplate.Insert("xxx", "")
	if err := mockobj.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *GdbcTemplateSuite) TestQueryRow() {
	db, mockobj, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.openDbFn.On("Execute", mock.Anything, mock.Anything).Return(db, nil)
	suite.dataSource.On("FormatDsn").Return("postgres://root:12345@127.0.0.1:5432/testdb", "postgres://root:******@127.0.0.1:5432/testdb")
	suite.dataSource.On("GetDriverName").Return("PGX")
	mockobj.ExpectQuery("select .* from OUR_USERS where USER_NAME=?").WithArgs("eagle").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "one"))
	suite.gdbcTemplate.QueryRow("select * from OUR_USERS where USER_NAME=?", "eagle").Map(func(rows *sql.Row) interface{} { return "test" }).ToObject()

	if err := mockobj.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *GdbcTemplateSuite) TestQueryRows() {
	db, mockobj, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.openDbFn.On("Execute", mock.Anything, mock.Anything).Return(db, nil)
	suite.dataSource.On("FormatDsn").Return("postgres://root:12345@127.0.0.1:5432/testdb", "postgres://root:******@127.0.0.1:5432/testdb")
	suite.dataSource.On("GetDriverName").Return("PGX")
	mockobj.ExpectQuery("^select .* from TABLE_NAME where USER_NAME=").WithArgs("eagle").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "one").
			AddRow(2, "two"))
	suite.gdbcTemplate.QueryForArray("select * from TABLE_NAME where USER_NAME=?", "eagle").Map(func(rows *sql.Rows) interface{} { return "test" }).ToArray()
	if err := mockobj.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
