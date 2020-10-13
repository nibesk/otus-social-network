package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"service-chat/app/config"
	"time"
)

type Conn struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mongo *Conn

func ConnectDb() *Conn {
	cred := options.Credential{
		Username: config.Env.DB.Username,
		Password: config.Env.DB.Password,
	}
	options := options.Client().ApplyURI(fmt.Sprintf("%s://%s", config.Env.DB.Dialect, config.Env.DB.Url)).SetAuth(cred)
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Mongo = &Conn{Client: client, Db: client.Database(config.Env.DB.Database)}

	return Mongo
}
