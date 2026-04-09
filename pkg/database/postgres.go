package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	ConnStr string
	DB      *sql.DB
}

func NewPostgres(connStr string) *Postgres {
	return &Postgres{ConnStr: connStr}
}

func (p *Postgres) Connect() error {
	db, err := sql.Open("postgres", p.ConnStr)
	if err != nil {
		return err
	}
	p.DB = db
	return p.DB.Ping()
}

func (p *Postgres) Close() error {
	if p.DB != nil {
		return p.DB.Close()
	}
	return nil
}
