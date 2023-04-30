package datasource_test

import (
	"gdbcTemplate/datasource"
	"gdbcTemplate/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
)

func TestGetDB(t *testing.T) {
	prop := &mocks.AppConfigProperties{}
	prop.On("Get", datasource.PORT).Return("1234")
	prop.On("Get", datasource.DRIVER_NAME).Return("mysql")
	prop.On("Get", datasource.USERNAME).Return("eagle")
	prop.On("Get", datasource.PASSWORD).Return("123456")
	prop.On("Get", datasource.NETWORK).Return("TCP")
	prop.On("Get", datasource.SERVER).Return("localhost")
	prop.On("Get", datasource.DATABASE).Return("testdb")
	prop.On("Get", datasource.CHARSET).Return("utf8")
	ds := datasource.NewDataSource(prop)
	openDbFn := mocks.NewOpenDb(t)

	db, mocker, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("mock error: '%s'", err)
	}
	openDbFn.On("Execute", mock.Anything, mock.Anything).Return(db, nil)
	dbM := datasource.NewDbManagerWithOpenDbFn(ds, openDbFn.Execute)
	dbM.GetDb()
	if err := mocker.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
