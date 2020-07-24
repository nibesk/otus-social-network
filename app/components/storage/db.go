package storage

import (
	"database/sql"
	"fmt"
	"github.com/badThug/otus-social-network/app/components/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Connection struct {
	config config.DBConfig
	Db     *sql.DB
}

func ConnectDatabase(DBConfig config.DBConfig) *Connection {
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

	connection := &Connection{config: DBConfig, Db: db}

	return connection
}
