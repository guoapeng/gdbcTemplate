package template_test

import (
	"database/sql"
	"gdbcTemplate/mocks"
	"gdbcTemplate/template"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestGdbcTemplateSuite(t *testing.T) {
	suite.Run(t, new(GdbcTemplateSuite))
}

type GdbcTemplateSuite struct {
	test *testing.T
	suite.Suite
	dataSource   *mocks.MockDataSource
	gdbcTemplate template.GdbcTemplate
}

func (suite *GdbcTemplateSuite) T() *testing.T {
	return suite.test
}

func (suite *GdbcTemplateSuite) SetT(t *testing.T) {
	suite.test = t
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.dataSource = mocks.NewMockDataSource(mockCtrl)
	suite.gdbcTemplate = template.NewWith(suite.dataSource)
}

func (suite *GdbcTemplateSuite) TestInsert() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()

	suite.dataSource.EXPECT().Open().Return(db, nil)
	mock.ExpectExec("xxx")
	suite.gdbcTemplate.Insert("xxx", "")
	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *GdbcTemplateSuite) TestQueryRow() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.dataSource.EXPECT().Open().Return(db, nil)
	mock.ExpectQuery("select .* from OUR_USERS where USER_NAME=?").WithArgs("eagle").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "one"))
	suite.gdbcTemplate.QueryRow("select * from OUR_USERS where USER_NAME=?", "eagle").Map(func(rows *sql.Row) interface{} { return "test" }).ToObject()

	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *GdbcTemplateSuite) TestQueryRows() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.dataSource.EXPECT().Open().Return(db, nil)
	mock.ExpectQuery("^select .* from TABLE_NAME where USER_NAME=").WithArgs("eagle").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "one").
			AddRow(2, "two"))
	suite.gdbcTemplate.QueryForArray("select * from TABLE_NAME where USER_NAME=?", "eagle").Map(func(rows *sql.Rows) interface{} { return "test" }).ToArray()
	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
