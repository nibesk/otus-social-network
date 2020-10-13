package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB       DBConfig
	Server   Server
	Services Services
	env      map[string]string
}

type DBConfig struct {
	Dialect  string
	Url      string
	Username string
	Password string
	Database string
}

type Server struct {
	Env           string
	SessionKey    string
	Host          string
	Port          string
	WsCheckOrigin bool
}

type Services struct {
	UsersUrl string
}

func (s Server) IsDev() bool {
	return "dev" == s.Env
}

var Env *Config

func InitConfig() {
	var envFileName string
	var err error

	if _, err = os.Stat(".env"); os.IsNotExist(err) {
		envFileName = ".env.example"
	} else {
		envFileName = ".env"
	}

	c := &Config{}

	c.env, err = godotenv.Read(envFileName)
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}

	c.DB = DBConfig{
		Dialect:  c.getVar("DB_DIALECT"),
		Url:      c.getVar("DB_URL"),
		Username: c.getVar("DB_USERNAME"),
		Password: c.getVar("DB_PASSWORD"),
		Database: c.getVar("DB_NAME"),
	}

	wsCheckOrigin := ("1" == c.getVar("SERVER_WS_CHECK_ORIGIN"))

	c.Server = Server{
		Env:           c.getVar("ENVIRONMENT"),
		Host:          c.getVar("SERVER_HOST"),
		Port:          c.getVar("SERVER_PORT"),
		WsCheckOrigin: wsCheckOrigin,
	}

	c.Services = Services{
		UsersUrl: c.getVar("SERVICE_USERS_URL"),
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
