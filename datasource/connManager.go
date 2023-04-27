package datasource

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type ConnManager interface {
	GetDb() (*sql.DB, error)
}

type connManager struct {
	ds DataSource
	db *sql.DB
}

func (manager *connManager) GetDb() (*sql.DB, error) {
	if manager.db != nil {
		return manager.db, nil
	} else {
		if db, err := manager.open(); err == nil {
			manager.db = db
			return manager.db, nil
		} else {
			return nil, errors.New("failed to open db")
		}
	}
}

func (manager *connManager) open() (*sql.DB, error) {

	dsn, maskedDsn := manager.ds.FormatDsn()

	log.Printf("connecting to %s\n", maskedDsn)

	db, err := sql.Open(manager.ds.GetDriverName(), dsn)
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

func NewConnManager(ds DataSource) ConnManager {
	return &connManager{ds: ds}
}
