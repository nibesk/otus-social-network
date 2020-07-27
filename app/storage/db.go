package storage

import (
	"database/sql"
	"fmt"
	"github.com/badThug/otus-social-network/app/config"
	_ "github.com/go-sql-driver/mysql"
)

type DbConnection struct {
	dbConfig config.DBConfig
	db       *sql.DB
}

func CreateDbConnection(DBConfig config.DBConfig) *DbConnection {
	connection := &DbConnection{dbConfig: DBConfig}

	return connection
}

func (c *DbConnection) Connect() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		c.dbConfig.Username,
		c.dbConfig.Password,
		c.dbConfig.Host,
		c.dbConfig.Port,
		c.dbConfig.Name,
		c.dbConfig.Charset)

	db, err := sql.Open(c.dbConfig.Dialect, dataSourceName)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	c.db = db

	return nil
}

func (c *DbConnection) Close() error {
	return c.db.Close()
}

func (c *DbConnection) GetDb() *sql.DB {
	return c.db
}
