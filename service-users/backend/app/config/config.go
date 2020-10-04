package config

import (
	"github.com/joho/godotenv"
	"log"
	"strings"
)

var Env *Config

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
	Env        string
	SessionKey string
	Host       string
	Port       string
}

func (s Server) IsDev() bool {
	return "dev" == s.Env
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	c := &Config{}

	c.env, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	replicasUrlsEnv := c.env["DB_REPLICA_URLS"]
	var replicasUrlsList []string
	if "" != replicasUrlsEnv {
		replicasUrlsList = strings.Split(replicasUrlsEnv, ",")
	}

	c.DB = DBConfig{
		Dialect:      c.env["DB_DIALECT"],
		MasterUrl:    c.env["DB_MASTER_URL"],
		ReplicasUrls: replicasUrlsList,
		Username:     c.env["DB_USERNAME"],
		Password:     c.env["DB_PASSWORD"],
		Database:     c.env["DB_NAME"],
		Charset:      c.env["DB_CHARSET"],
	}

	c.Server = Server{
		Env:        c.env["ENVIRONMENT"],
		SessionKey: c.env["SESSION_KEY"],
		Host:       c.env["SERVER_HOST"],
		Port:       c.env["SERVER_PORT"],
	}

	Env = c
}
