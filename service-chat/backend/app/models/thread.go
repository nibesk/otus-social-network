package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"service-chat/app/storage"
	"time"
)

type Thread struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Initiator_user_id  int                `bson:"initiator_user_id"`
	Respondent_user_id int                `bson:"respondent_user_id"`
}

const threadDbName = "threads"

func ThreadInsure(Initiator_user_id, Respondent_user_id int) (*Thread, error) {
	filter := bson.D{
		primitive.E{
			Key: "$or",
			Value: []interface{}{
				bson.D{{"initiator_user_id", Initiator_user_id}, {"respondent_user_id", Respondent_user_id}},
				bson.D{{"initiator_user_id", Respondent_user_id}, {"respondent_user_id", Initiator_user_id}},
			},
		},
	}

	var t *Thread
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := storage.Mongo.Db.Collection(threadDbName).FindOne(ctx, filter).Decode(&t)
	switch true {
	case errors.Is(err, mongo.ErrNoDocuments):
		return ThreadCreate(Initiator_user_id, Respondent_user_id)
	case nil != err:
		return nil, err
	}

	return t, nil
}

func ThreadCreate(Initiator_user_id, Respondent_user_id int) (*Thread, error) {
	t := &Thread{
		ID:                 primitive.NewObjectID(),
		Initiator_user_id:  Initiator_user_id,
		Respondent_user_id: Respondent_user_id,
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := storage.Mongo.Db.Collection(threadDbName).InsertOne(ctx, t)

	return t, err
}
