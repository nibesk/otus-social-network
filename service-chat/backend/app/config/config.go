package config

import (
	"github.com/joho/godotenv"
	"log"
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
	Env           string
	SessionKey    string
	Host          string
	Port          string
	WsCheckOrigin bool
}

func (s Server) IsDev() bool {
	return "dev" == s.Env
}

var Env *Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := &Config{}

	c.env, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.DB = DBConfig{
		Dialect:   c.env["DB_DIALECT"],
		MasterUrl: c.env["DB_URL"],
		Username:  c.env["DB_USERNAME"],
		Password:  c.env["DB_PASSWORD"],
		Database:  c.env["DB_NAME"],
	}

	wsCheckOrigin := ("1" == c.env["SERVER_WS_CHECK_ORIGIN"])

	c.Server = Server{
		Env:           c.env["ENVIRONMENT"],
		Host:          c.env["SERVER_HOST"],
		Port:          c.env["SERVER_PORT"],
		WsCheckOrigin: wsCheckOrigin,
	}

	Env = c
}
