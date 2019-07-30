package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	datasource string
}

func NewDb(datasource string) *Db {
	return &Db{
		datasource: datasource,
	}
}

func (db *Db) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", db.datasource)
}
