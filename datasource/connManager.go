package datasource

import "database/sql"

type ConnManager interface {
	Open(driverName, dataSourceName string) (*sql.DB, error)
}

type connManager struct {
}

func (manager *connManager) Open(driverName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driverName, dataSourceName)
}

func NewConnManager() ConnManager {
	return &connManager{}
}
