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
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

type Server struct {
	HttpPort string
}

func Init() *Config {
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
		Dialect:  c.env["DB_DIALECT"],
		Host:     c.env["DB_HOST"],
		Port:     c.env["DB_PORT"],
		Username: c.env["DB_USERNAME"],
		Password: c.env["DB_PASSWORD"],
		Name:     c.env["DB_NAME"],
		Charset:  c.env["DB_CHARSET"],
	}

	c.Server = Server{
		HttpPort: c.env["SERVER_HTTP_PORT"],
	}

	return c
}
