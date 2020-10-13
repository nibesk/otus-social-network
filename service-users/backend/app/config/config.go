package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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
	var envFileName string
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		envFileName = ".env.example"
	} else {
		envFileName = ".env"
	}

	c := &Config{}

	var err error
	c.env, err = godotenv.Read(envFileName)
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}

	replicasUrlsEnv := c.getVar("DB_REPLICA_URLS")
	var replicasUrlsList []string
	if "" != replicasUrlsEnv {
		replicasUrlsList = strings.Split(replicasUrlsEnv, ",")
	}

	c.DB = DBConfig{
		Dialect:      c.getVar("DB_DIALECT"),
		MasterUrl:    c.getVar("DB_MASTER_URL"),
		ReplicasUrls: replicasUrlsList,
		Username:     c.getVar("DB_USERNAME"),
		Password:     c.getVar("DB_PASSWORD"),
		Database:     c.getVar("DB_NAME"),
		Charset:      c.getVar("DB_CHARSET"),
	}

	c.Server = Server{
		Env:        c.getVar("ENVIRONMENT"),
		SessionKey: c.getVar("SESSION_KEY"),
		Host:       c.getVar("SERVER_HOST"),
		Port:       c.getVar("SERVER_PORT"),
	}

	Env = c
}

func (c *Config) getVar(key string) string {
	var val string

	val = os.Getenv(key)
	if "" == val {
		val = c.env[key]
	}

	return val
}
