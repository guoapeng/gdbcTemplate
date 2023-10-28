package transaction_test

import (
	"database/sql"
	"testing"

	"github.com/guoapeng/gdbcTemplate/mocks"
	"github.com/guoapeng/gdbcTemplate/template"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}

type TransactionSuite struct {
	test *testing.T
	suite.Suite
	dataSource   *mocks.DataSource
	gdbcTemplate template.GdbcTemplate
	openDbFn     *mocks.OpenDb
	sqlmock1     sqlmock.Sqlmock
	db           *sql.DB
}

func (suite *TransactionSuite) T() *testing.T {
	return suite.test
}

func (suite *TransactionSuite) SetT(t *testing.T) {
	suite.test = t
	suite.dataSource = &mocks.DataSource{}
	suite.openDbFn = mocks.NewOpenDb(t)
	suite.gdbcTemplate = template.NewWithDsAndOpenDbFn(suite.dataSource, suite.openDbFn.Execute)

	suite.dataSource.On("GetDriverName").Return("PGX")
	suite.dataSource.On("FormatDsn").Return("postgres://root:12345@127.0.0.1:5432/testdb", "postgres://root:******@127.0.0.1:5432/testdb")
}

func (suite *TransactionSuite) TestInsert() {
	db, mockObj, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	suite.db = db
	suite.openDbFn.On("Execute", mock.Anything, mock.Anything).Return(db, nil)

	mockObj.ExpectBegin()
	mockObj.ExpectPrepare("insert into test_table\\(id, name\\) values\\(100, 'test'\\)")
	mockObj.ExpectExec("insert into test_table\\(id, name\\) values\\(100, 'test'\\)").WillReturnResult(sqlmock.NewResult(100, 1))
	transaction, err := suite.gdbcTemplate.BeginTx()
	if err != nil {
		suite.Fail("transaction error: '%s'", err)
	}
	id, err := transaction.Insert("insert into test_table(id, name) values(100, 'test')", "")
	transaction.Commit()
	assert.Nil(suite.T(), err)
	lstId, _ := id.LastInsertId()
	assert.Equal(suite.T(), lstId, int64(100))
	if err := mockObj.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
	suite.openDbFn.AssertExpectations(suite.T())
}

func (suite *TransactionSuite) TearDownSuite() {
	suite.T().Log("close db")
	suite.db.Close()
}
