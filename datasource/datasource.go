package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

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

type DataSource interface {
	Open() (*sql.DB, error)
}

type dataSource struct {
	DriverName  string
	UserName    string
	Password    string
	Network     string
	Server      string
	Port        int
	DataBase    string
	Charset     string
	connManager ConnManager
}

func (ds *dataSource) Open() (*sql.DB, error) {

	dsn, maskedDsn := ds.format()

	log.Printf("connecting to %s\n", maskedDsn)

	db, err := ds.connManager.Open(ds.DriverName, dsn)
	if err != nil {
		log.Printf("Open database failed,err:%v\n", err)
		return nil, err
	} else {
		log.Printf("connected to the database successfully!\n")
		db.SetConnMaxLifetime(100 * time.Second)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(16)
		return db, nil
	}
}

func (ds *dataSource) format() (string, string) {

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

func NewDataSource(appConf propsReader.AppConfigProperties, connManager ConnManager) DataSource {
	port, errp := strconv.Atoi(appConf.Get(PORT))
	if errp == nil {
		return &dataSource{DriverName: appConf.Get(DRIVER_NAME),
			UserName:    appConf.Get(USERNAME),
			Password:    appConf.Get(PASSWORD),
			Network:     appConf.Get(NETWORK),
			Server:      appConf.Get(SERVER),
			Port:        port,
			DataBase:    appConf.Get(DATABASE),
			Charset:     appConf.Get(CHARSET),
			connManager: connManager,
		}
	} else {
		panic("failed to create data source due to invalid port number")
	}
}
