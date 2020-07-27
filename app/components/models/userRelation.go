package models

import "github.com/badThug/otus-social-network/app/components/storage"

type UserRelation struct {
	Relation_id    int
	User_id        int
	Friend_user_id int
	Created_at     string
	Updated_at     string
}

func UserRelationCreate(conn *storage.DbConnection, userId, friendUserId int) (*UserRelation, error) {
	db := conn.GetDb()
	insert, err := db.Prepare("INSERT INTO `user_relation` (user_id, friend_user_id) VALUES(?, ?)")
	if err != nil {
		return nil, err
	}

	result, err := insert.Exec(userId, friendUserId)
	if err != nil {
		return nil, err
	}

	relationId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	userRelation := &UserRelation{
		Relation_id:    int(relationId),
		User_id:        userId,
		Friend_user_id: friendUserId,
	}

	return userRelation, nil
}
