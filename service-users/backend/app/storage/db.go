package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"service-users/app/config"
	"service-users/app/storage/balancer"
	"time"
)

type DbConnection struct {
	dbConfig config.DBConfig
	cdb      *balancer.DB
}

func (c *DbConnection) GetDb() *sql.DB {
	return c.cdb.Master()
}

func (c *DbConnection) GetCDb() *balancer.DB {
	return c.cdb
}

type Queryable interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Executable interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var conn *DbConnection

func ConnectDb() error {
	dbConfig := config.Env.DB
	dsnFormat := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=True"
	masterDSN := fmt.Sprintf(dsnFormat, dbConfig.Username, dbConfig.Password, dbConfig.MasterUrl, dbConfig.Database, dbConfig.Charset)

	replicaDSNs := make([]string, 0, len(dbConfig.ReplicasUrls))
	for _, replicaUrl := range dbConfig.ReplicasUrls {
		replicaDSNs = append(replicaDSNs, fmt.Sprintf(dsnFormat, dbConfig.Username, dbConfig.Password, replicaUrl, dbConfig.Database, dbConfig.Charset))
	}

	cdb, err := balancer.Open(dbConfig.Dialect, masterDSN, replicaDSNs)
	if err != nil {
		return err
	}

	cdb.SetConnMaxLifetime(10 * time.Minute)
	cdb.SetMaxIdleConns(25)
	cdb.SetMaxOpenConns(25)

	conn = &DbConnection{dbConfig: dbConfig, cdb: cdb}

	return nil
}

func GetDb() *sql.DB {
	return conn.GetDb()
}

func GetCDb() *balancer.DB {
	return conn.GetCDb()
}
