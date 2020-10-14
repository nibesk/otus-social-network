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
	ID        primitive.ObjectID `bson:"_id"`
	Text      string             `bson:"text"`
	Thread_id primitive.ObjectID `bson:"thread_id"`
	User_id   int                `bson:"from_user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

const messagesDbName = "messages"

func MessageCreate(message *Message) error {
	message.ID = primitive.NewObjectID()
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := storage.Mongo.Db.Collection(messagesDbName).InsertOne(ctx, message)

	return err
}

func MessageFindConversation(thread_id primitive.ObjectID) ([]*Message, error) {
	options := options.Find()
	options.SetSort(bson.D{{"created_at", 1}})
	filter := bson.D{{"thread_id", thread_id}}

	return MessageFilter(filter, options)
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
