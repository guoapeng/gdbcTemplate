package datasource_test

import (
	"testing"

	"gdbcTemplate/datasource"
	"gdbcTemplate/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestDataSourceSuite(t *testing.T) {
	suite.Run(t, new(DataSourceSuite))
}

type DataSourceSuite struct {
	suite.Suite
	test        *testing.T
	appConf     *mocks.MockAppConfigProperties
	connManager *mocks.MockConnManager
	datasource  datasource.DataSource
}

func (suite *DataSourceSuite) T() *testing.T {
	return suite.test
}

func (suite *DataSourceSuite) SetT(t *testing.T) {
	suite.test = t
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.connManager = mocks.NewMockConnManager(mockCtrl)
	suite.appConf = mocks.NewMockAppConfigProperties(mockCtrl)

}

func (suite *DataSourceSuite) TestOpenPostgreDb() {

	suite.appConf.EXPECT().Get("DRIVER_NAME").Return("pgx")
	suite.appConf.EXPECT().Get("USERNAME").Return("eagle")
	suite.appConf.EXPECT().Get("PASSWORD").Return("123456")
	suite.appConf.EXPECT().Get("NETWORK").Return("TCP")
	suite.appConf.EXPECT().Get("SERVER").Return("localhost")
	suite.appConf.EXPECT().Get("PORT").Return("1234")
	suite.appConf.EXPECT().Get("DATABASE").Return("test")
	suite.appConf.EXPECT().Get("CHARSET").Return("utf8")
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.connManager.EXPECT().Open("pgx", "postgres://eagle:123456@localhost:1234/test").Return(db, nil)
	suite.datasource = datasource.NewDataSource(suite.appConf, suite.connManager)
	suite.datasource.Open()
	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *DataSourceSuite) TestOpenMySqlDb() {

	suite.appConf.EXPECT().Get("DRIVER_NAME").Return("mysql")
	suite.appConf.EXPECT().Get("USERNAME").Return("eagle")
	suite.appConf.EXPECT().Get("PASSWORD").Return("123456")
	suite.appConf.EXPECT().Get("NETWORK").Return("TCP")
	suite.appConf.EXPECT().Get("SERVER").Return("localhost")
	suite.appConf.EXPECT().Get("PORT").Return("1234")
	suite.appConf.EXPECT().Get("DATABASE").Return("test")
	suite.appConf.EXPECT().Get("CHARSET").Return("utf8")
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.Fail("mock error: '%s'", err)
	}
	defer db.Close()
	suite.connManager.EXPECT().Open("mysql", "eagle:123456@TCP(localhost:1234)/test?charset=utf8").Return(db, nil)
	suite.datasource = datasource.NewDataSource(suite.appConf, suite.connManager)
	suite.datasource.Open()
	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
