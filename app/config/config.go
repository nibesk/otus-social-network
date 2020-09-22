package config

import (
	"github.com/joho/godotenv"
	"log"
	"strings"
)

type Config struct {
	DB     DBConfig
	Server Server
	env    map[string]string
}

type DBConfig struct {
	Dialect      string
	MasterUrl    string
	ReplicasUrls []string
	Username     string
	Password     string
	Database     string
	Charset      string
}

type Server struct {
	HttpPort   string
	Env        string
	SessionKey string
	Host       string
}

func (s Server) IsDev() bool {
	return "dev" == s.Env
}

func InitConfig() *Config {
	return (&Config{}).initInner()
}

func (c *Config) initInner() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.env, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.DB = DBConfig{
		Dialect:      c.env["DB_DIALECT"],
		MasterUrl:    c.env["DB_MASTER_URL"],
		ReplicasUrls: strings.Split(c.env["DB_REPLICA_URLS"], ","),
		Username:     c.env["DB_USERNAME"],
		Password:     c.env["DB_PASSWORD"],
		Database:     c.env["DB_NAME"],
		Charset:      c.env["DB_CHARSET"],
	}

	c.Server = Server{
		HttpPort:   c.env["SERVER_HTTP_PORT"],
		Env:        c.env["ENVIRONMENT"],
		SessionKey: c.env["SESSION_KEY"],
		Host:       c.env["SERVER_HOST"],
	}

	return c
}
