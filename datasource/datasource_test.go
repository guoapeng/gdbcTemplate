package datasource_test

import (
	"gdbcTemplate/datasource"
	"gdbcTemplate/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatPostgresqlDsn(t *testing.T) {
	prop := &mocks.AppConfigProperties{}
	prop.On("Get", datasource.PORT).Return("5432")
	prop.On("Get", datasource.DRIVER_NAME).Return("PGX")
	prop.On("Get", datasource.USERNAME).Return("root")
	prop.On("Get", datasource.PASSWORD).Return("12345")
	prop.On("Get", datasource.NETWORK).Return("tcp")
	prop.On("Get", datasource.SERVER).Return("127.0.0.1")
	prop.On("Get", datasource.DATABASE).Return("testdb")
	prop.On("Get", datasource.CHARSET).Return("UTF8")
	ds := datasource.NewDataSource(prop)
	dsn, maskedDsb := ds.FormatDsn()
	assert.Equal(t, "postgres://root:12345@127.0.0.1:5432/testdb", dsn)
	assert.Equal(t, "postgres://root:******@127.0.0.1:5432/testdb", maskedDsb)
	assert.Equal(t, "PGX", ds.GetDriverName())
}

func TestFormatMysqlDsn(t *testing.T) {
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
	dsn, maskedDsb := ds.FormatDsn()
	assert.Equal(t, "eagle:123456@TCP(localhost:1234)/testdb?charset=utf8", dsn)
	assert.Equal(t, "eagle:******@TCP(localhost:1234)/testdb?charset=utf8", maskedDsb)
	assert.Equal(t, "mysql", ds.GetDriverName())
}
