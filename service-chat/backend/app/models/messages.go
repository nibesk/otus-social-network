package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"service-chat/app/storage"
	"time"
)

type Message struct {
	Text         string    `bson:"text"`
	From_user_id int       `bson:"from_user_id"`
	To_user_id   int       `bson:"to_user_id"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

const messagesDbName = "messages"

func MessageCreate(message *Message) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := storage.Mongo.Db.Collection(messagesDbName).InsertOne(ctx, message)

	return err
}

func MessageFilter(filter interface{}, options *options.FindOptions) ([]*Message, error) {
	var messages []*Message

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := storage.Mongo.Db.Collection(messagesDbName).Find(ctx, filter, options)
	if err != nil {
		return messages, err
	}

	for cur.Next(ctx) {
		var m Message
		err := cur.Decode(&m)
		if err != nil {
			return messages, err
		}

		messages = append(messages, &m)
	}

	if err := cur.Err(); err != nil {
		return messages, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	return messages, nil
}

func MessageFindConversation(fromUserId, toUserId int) ([]*Message, error) {
	options := options.Find()
	options.SetSort(bson.D{{"created_at", 1}})

	filter := bson.D{
		primitive.E{
			Key: "$or",
			Value: []interface{}{
				bson.D{{"from_user_id", fromUserId}, {"to_user_id", toUserId}},
				bson.D{{"from_user_id", toUserId}, {"to_user_id", fromUserId}},
			},
		},
	}

	return MessageFilter(filter, options)
}
