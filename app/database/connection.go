package database

import (
	"database/sql"
	"fmt"
	"github.com/badThug/otus-social-network/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Config config.DBConfig
	Db     *sql.DB
}

func Connect(DBConfig config.DBConfig) *Connection {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		DBConfig.Username,
		DBConfig.Password,
		DBConfig.Host,
		DBConfig.Port,
		DBConfig.Name,
		DBConfig.Charset)
	db, err := sql.Open(DBConfig.Dialect, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	connection := &Connection{Config: DBConfig, Db: db}

	return connection
}
