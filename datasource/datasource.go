package datasource

import (
	"fmt"
	"strconv"

	propsReader "github.com/guoapeng/props"
)

const (
	DRIVER_NAME = "DRIVER_NAME"
	USERNAME    = "USERNAME"
	PASSWORD    = "PASSWORD"
	NETWORK     = "NETWORK"
	SERVER      = "SERVER"
	PORT        = "PORT"
	DATABASE    = "DATABASE"
	CHARSET     = "CHARSET"
)

type dataSource struct {
	DriverName string
	UserName   string
	Password   string
	Network    string
	Server     string
	Port       int
	DataBase   string
	Charset    string
}

type DataSource interface {
	FormatDsn() (string, string)
	GetDriverName() string
}

func (ds *dataSource) GetDriverName() string {
	return ds.DriverName
}
func (ds *dataSource) FormatDsn() (string, string) {

	if ds.DriverName == "mysql" {
		dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s", ds.UserName, ds.Password, ds.Network, ds.Server,
			ds.Port, ds.DataBase, ds.Charset)

		maskedDsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s", ds.UserName, "******", ds.Network, ds.Server,
			ds.Port, ds.DataBase, ds.Charset)
		return dsn, maskedDsn
	} else {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", ds.UserName, ds.Password, ds.Server,
			ds.Port, ds.DataBase)

		maskedDsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", ds.UserName, "******", ds.Server,
			ds.Port, ds.DataBase)

		return dsn, maskedDsn
	}
}

func NewDataSource(appConf propsReader.AppConfigProperties) DataSource {
	port, errp := strconv.Atoi(appConf.Get(PORT))
	if errp == nil {
		return &dataSource{DriverName: appConf.Get(DRIVER_NAME),
			UserName: appConf.Get(USERNAME),
			Password: appConf.Get(PASSWORD),
			Network:  appConf.Get(NETWORK),
			Server:   appConf.Get(SERVER),
			Port:     port,
			DataBase: appConf.Get(DATABASE),
			Charset:  appConf.Get(CHARSET),
		}
	} else {
		panic("failed to create data source due to invalid port number")
	}
}
