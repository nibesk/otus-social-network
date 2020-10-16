package models

import (
	"database/sql"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"service-users/app/customErrors"
	"service-users/app/storage"
)

type UserRelation struct {
	Relation_id    int    `json:"relation_id"`
	User_id        int    `json:"user_id"`
	Friend_user_id int    `json:"friend_user_id"`
	Created_at     string `json:"-"`
	Updated_at     string `json:"-"`
}

func UserRelationCreate(db storage.Executable, userId, friendUserId int) (*UserRelation, error) {
	insert, err := db.Prepare("INSERT INTO `user_relation` (user_id, friend_user_id) VALUES(?, ?)")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result, err := insert.Exec(userId, friendUserId)
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errors.Wrap(&customErrors.TypedError{"User already have this relation"}, driverErr.Message)
		}

		return nil, err
	}

	relationId, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userRelation := &UserRelation{
		Relation_id:    int(relationId),
		User_id:        userId,
		Friend_user_id: friendUserId,
	}

	return userRelation, nil
}

func UserRelationDelete(db storage.Executable, userId, friendUserId int) error {
	insert, err := db.Prepare("DELETE FROM user_relation WHERE user_id = ? AND friend_user_id = ?")
	if err != nil {
		return errors.WithStack(err)
	}

	result, err := insert.Exec(userId, friendUserId)
	if err != nil {
		return errors.WithStack(err)
	}

	cnt, err := result.RowsAffected()
	if nil != err {
		return errors.WithStack(err)
	}

	if 0 == cnt {
		return sql.ErrNoRows
	}

	return nil
}

func UserRelationFindByUserId(db storage.Queryable, userId int) ([]*UserRelation, error) {
	query, err := db.Query("SELECT * FROM user_relation WHERE user_id = ? OR friend_user_id = ?", userId, userId)
	if err != nil {
		return nil, err
	}

	collection := []*UserRelation{}
	for query.Next() {
		userRelation := &UserRelation{}
		err := userRelationQueryScan(query.Scan, userRelation)
		if err != nil {
			return nil, err
		}
		collection = append(collection, userRelation)
	}

	return collection, nil
}

func userRelationQueryScan(scan func(dest ...interface{}) error, userRelation *UserRelation) error {
	err := scan(
		&userRelation.Relation_id,
		&userRelation.User_id,
		&userRelation.Friend_user_id,
		&userRelation.Created_at,
		&userRelation.Updated_at)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
