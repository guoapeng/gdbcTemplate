package datasource

import (
	"database/sql"
	"fmt"
	"github.com/guoapeng/props"
	"log"
	"strconv"
	"time"
)

const (
	USERNAME = "USERNAME"
	PASSWORD = "PASSWORD"
	NETWORK  = "NETWORK"
	SERVER   = "SERVER"
	PORT     = "PORT"
	DATABASE = "DATABASE"
	CHARSET  = "CHARSET"
)

type DataSource interface {
	Open() (*sql.DB, error)
}

type dataSource struct {
	UserName string
	Password string
	Network  string
	Server   string
	Port     int
	DataBase string
	Charset  string
}


func (ds *dataSource) Open() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s", ds.UserName, ds.Password, ds.Network, ds.Server,
		ds.Port, ds.DataBase, ds.Charset)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		return nil, err
	} else {
		db.SetConnMaxLifetime(100 * time.Second)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(16)
		return db, nil
	}
}

func NewDataSource(appConf propsReader.AppConfigProperties) DataSource {
	port, errp := strconv.Atoi(appConf.Get(PORT))
	if errp == nil {
		return &dataSource{UserName: appConf.Get(USERNAME), Password: appConf.Get(PASSWORD),
			Network: appConf.Get(NETWORK), Server: appConf.Get(SERVER), Port: port, DataBase: appConf.Get(DATABASE), Charset: appConf.Get(CHARSET)}
	} else {
		panic("failed to create data source")
	}
}


