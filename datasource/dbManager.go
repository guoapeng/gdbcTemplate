package datasource

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type OpenDb func(driverName, dataSourceName string) (*sql.DB, error)

type DbManager interface {
	GetDb() (*sql.DB, error)
	Close() error
}

type dbManager struct {
	ds         DataSource
	db         *sql.DB
	openDbFunc OpenDb
}

func (manager *dbManager) GetDb() (*sql.DB, error) {
	if manager.db != nil {
		return manager.db, nil
	} else {
		if db, err := manager.openDb(); err == nil {
			manager.db = db
			return manager.db, nil
		} else {
			return nil, errors.New("failed to open db")
		}
	}
}

func (manager *dbManager) openDb() (*sql.DB, error) {

	dsn, maskedDsn := manager.ds.FormatDsn()

	log.Printf("connecting to %s\n", maskedDsn)
	db, err := manager.openDbFunc(manager.ds.GetDriverName(), dsn)
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

func (manager *dbManager) Close() error {
	if manager.db != nil {
		log.Println("Closing db")
		return manager.db.Close()
	} else {
		log.Println("db is not opened, need not to close")
		return nil
	}

}

func NewDbManager(ds DataSource) DbManager {
	return &dbManager{ds: ds, openDbFunc: sql.Open}
}

func NewDbManagerWithOpenDbFn(ds DataSource, openDbFunc OpenDb) DbManager {
	return &dbManager{ds: ds, openDbFunc: openDbFunc}
}
