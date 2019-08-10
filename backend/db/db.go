package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	datasource string
}

func NewDB(datasource string) *DB {
	return &DB{
		datasource: datasource,
	}
}

func (db *DB) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", db.datasource)
}
