package sqlite

import (
	"database/sql"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager() *DBManager {
	return &DBManager{
		DB: &sql.DB{},
	}
}

func (m *DBManager) Open(driverName string, dsn string) error {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return err
	}

	m.DB = db

	return nil
}

func (m *DBManager) Close() error {
	err := m.DB.Close()
	if err != nil {
		return err
	}

	return nil
}
