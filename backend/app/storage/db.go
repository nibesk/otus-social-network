package storage

import (
	"database/sql"
	"fmt"
	"github.com/badThug/otus-social-network/app/config"
	"github.com/badThug/otus-social-network/app/storage/balancer"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DbConnection struct {
	dbConfig config.DBConfig
	cdb      *balancer.DB
}

func ConnectDb(dbConfig config.DBConfig) (*DbConnection, error) {
	dsnFormat := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=True"
	masterDSN := fmt.Sprintf(dsnFormat, dbConfig.Username, dbConfig.Password, dbConfig.MasterUrl, dbConfig.Database, dbConfig.Charset)

	replicaDSNs := make([]string, 0, len(dbConfig.ReplicasUrls))
	for _, replicaUrl := range dbConfig.ReplicasUrls {
		replicaDSNs = append(replicaDSNs, fmt.Sprintf(dsnFormat, dbConfig.Username, dbConfig.Password, replicaUrl, dbConfig.Database, dbConfig.Charset))
	}

	cdb, err := balancer.Open(dbConfig.Dialect, masterDSN, replicaDSNs)
	if err != nil {
		return nil, err
	}

	cdb.SetConnMaxLifetime(10 * time.Minute)
	cdb.SetMaxIdleConns(25)
	cdb.SetMaxOpenConns(25)

	return &DbConnection{dbConfig: dbConfig, cdb: cdb}, nil
}

func (c *DbConnection) GetDb() *sql.DB {
	return c.cdb.Master()
}

func (c *DbConnection) GetCDb() *balancer.DB {
	return c.cdb
}
