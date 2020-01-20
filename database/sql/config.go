package sql

import (
	"database/sql"
	"fmt"
	sqlite3 "github.com/mattn/go-sqlite3"

	"github.com/1token/email-services/database"
)

type SQLite3 struct {
	// File to
	File string `yaml:"file"`
}

func (s *SQLite3) Open() (database.DatabaseX, error) {
	conn, err := s.open()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *SQLite3) open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s.File)
	if err != nil {
		sqlErr, ok := err.(sqlite3.Error)
		if !ok {
			return nil, err
		}
		return nil, sqlErr.ExtendedCode
	}
	if s.File == ":memory:" {
		// sqlite3 uses file locks to coordinate concurrent access. In memory
		// doesn't support this, so limit the number of connections to 1.
		db.SetMaxOpenConns(1)
	}

	return db, nil
}

// Postgres options for creating an SQL db.
type Postgres struct {
	Database string
	User     string
	Password string
	Host     string
	Port     uint16
}

func (p *Postgres) Open() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", p.User, p.Password, p.Host, p.Port, p.Database)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
