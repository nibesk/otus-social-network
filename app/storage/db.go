package storage

import (
	"database/sql"
	"fmt"
	"github.com/badThug/otus-social-network/app/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DbConnection struct {
	dbConfig config.DBConfig
	db       *sql.DB
}

func CreateDbConnection(dbConfig config.DBConfig) (*DbConnection, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset)

	// Initialise a new c pool
	db, err := sql.Open(dbConfig.Dialect, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)

	return &DbConnection{dbConfig, db}, nil
}

func (c *DbConnection) GetDb() *sql.DB {
	return c.db
}
