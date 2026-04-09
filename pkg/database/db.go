package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Interface chung
type Database interface {
	Connect() error
	Close() error
	Ping() error
}

// Postgres implementation
type PostgresDB struct {
	ConnStr string
	DB      *sql.DB
}

func (p *PostgresDB) Connect() error {
	db, err := sql.Open("postgres", p.ConnStr)
	if err != nil {
		return err
	}
	p.DB = db
	return p.Ping()
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

func (p *PostgresDB) Ping() error {
	return p.DB.Ping()
}
