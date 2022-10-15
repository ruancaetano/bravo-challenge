package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager() *DBManager {
	return &DBManager{
		DB: &sql.DB{},
	}
}

func (m *DBManager) Open(databaseFile string) error {
	db, err := sql.Open("sqlite3", databaseFile)
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
